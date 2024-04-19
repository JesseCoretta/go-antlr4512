package antlr4512

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

var (
	hasPfx  func(string, string) bool           = strings.HasPrefix
	isAlpha func(rune) bool                     = unicode.IsLetter
	isDigit func(rune) bool                     = unicode.IsDigit
	atoi    func(string) (int, error)           = strconv.Atoi
	trimS   func(string) string                 = strings.TrimSpace
	trim    func(string, string) string         = strings.Trim
	trimL   func(string, string) string         = strings.TrimLeft
	trimR   func(string, string) string         = strings.TrimRight
	repAll  func(string, string, string) string = strings.ReplaceAll
	hasSfx  func(string, string) bool           = strings.HasSuffix
	eq      func(string, string) bool           = strings.EqualFold
	printf  func(string, ...any) (int, error)   = fmt.Printf
)

func isAlnum(x rune) bool {
	return isAlpha(x) || isDigit(x)
}

/*
readFile uses os.ReadFile to attempt to read f into raw, which is
returned alongside an error.

Only files with a suffix of ".schema" will be eligible for processing.
*/
func readFile(f string) (raw []byte, err error) {
	if hasSfx(f, `.schema`) {
		raw, err = os.ReadFile(f)
	}

	return
}

func (r *Schema) processObjectIdentifiers(ctx []IObjectIdentifierContext) (err error) {
	for i := 0; i < len(ctx) && err == nil; i++ {
		oioids := ctx[i].AllObjectIdentifierOIDOrName()
		for j := 0; j < len(oioids); j++ {
			txt := oioids[j].Descriptor()
			val := oioids[j].NumericOIDOrMacro()
			if noid := val.NumericOID(); noid != nil {
				r.OM[txt.GetText()] = noid.GetText()
			} else if mc := val.Macro(); mc != nil {
				if descr := mc.Descriptor(); descr != nil {
					resl, found := r.OM[descr.GetText()]
					if !found {
						err = errorf("Could not resolve macro '%s'",
							descr.GetText())
						break
					}
					r.OM[txt.GetText()] = resl +
						repAll(mc.MacroSuffix().GetText(), `:`, `.`)
				}
			}
		}
	}

	return
}

/*
readDirectory recurses all files and folders specified at 'dir',
returning parsed schema bytes (content) alongside an error.

Only files with an extension of ".schema" will be parsed, but all
subdirectories will be traversed in search of these files.

File and directory naming schemes MUST guarantee the appropriate
ordering of any and all sub types, sub rules and sub classes which
would rely on the presence of dependency definitions (e.g.: 'cn'
cannot exist without 'name').
*/
func readDirectory(dir string) (content []byte, err error) {
	// remove any number of trailing
	// slashes from dir.
	dir = trimR(dir, `/`)

	// interim content storage
	var _content []byte

	// recurse dir path
	err = filepath.Walk(dir, func(p string, d fs.FileInfo, err error) error {
		if !d.IsDir() {
			var subcnt []byte
			// to avoid accidental splicing of two files
			// with regards to a missing newline (and
			// thus avoiding an error), insert a newline
			// char (rune(10)) if not already present at
			// the end of the byte sequence.
			if subcnt, err = readFile(p); err == nil && len(subcnt) > 0 {
				if string(subcnt[len(subcnt)-1]) != string(rune(10)) {
					_content = append(_content, []byte(string(rune(10)))...)
				}
				_content = append(_content, subcnt...)
			}
		}

		return err
	})

	if len(_content) > 0 && err == nil {
		content = _content
	}

	return
}

/*
errorf wraps errors.New and returns a non-nil instance of error
based upon a non-nil/non-zero msg input value with optional args.
*/
func errorf(msg any, x ...any) error {
	switch tv := msg.(type) {
	case string:
		if len(tv) > 0 {
			return errors.New(fmt.Sprintf(tv, x...))
		}
	case error:
		if tv != nil {
			return errors.New(fmt.Sprintf(tv.Error(), x...))
		}
	}

	return nil
}

/*
Instance contain tokenized content following a successful parse
of any "thing" in its complete form.

The P field contains an instance of *RFC4512Parser. The L field
contains an instance of *RFC4512Lexer.
*/
type Instance struct {
	P    *RFC4512Parser
	L    *RFC4512Lexer
	lErr *ErrorListener
	pErr *ErrorListener
}

/*
ParseInstance returns a new instance of Instance, alongside an error,
following an attempt to parse raw. The stdErr variadic bool allows for
the ability to toggle stdErr error dispatching.  The default is to
handle errors in the standard idiomatic way (recommended).
*/
func ParseInstance(raw string, stdErr ...bool) (r Instance, err error) {
	// don't waste time on a zero string ...
	if len(raw) == 0 {
		err = errorf("Zero-length input string; impossible parsing request")
		return
	}

	var useStdErr bool
	if len(stdErr) > 0 {
		useStdErr = stdErr[0]
	}

	// Prepare lexer error listener
	r.lErr = &ErrorListener{}

	// create a new input stream using the
	// raw input and initialize lexer using
	// the stream.
	r.L = NewRFC4512Lexer(antlr.NewInputStream(raw))

	if !useStdErr {
		r.L.RemoveErrorListeners()
		r.L.AddErrorListener(r.lErr)
	}

	// Prepare parser error listener
	r.pErr = &ErrorListener{}

	// Prepare the parser using a new CommonTokenStream
	// as input, define as p, and verify for nilness.
	if r.P = NewRFC4512Parser(antlr.NewCommonTokenStream(r.L, 0)); r.P == nil {
		err = errorf("Unknown ANTLR error; received a nil %T instance", r.P)
		return
	}

	if !useStdErr {
		r.P.RemoveErrorListeners()
		r.P.AddErrorListener(r.pErr)
	}

	// Declare our intent to build a parse tree
	r.P.BuildParseTrees = true

	err = r.Error()

	return
}

// syntaxError overrides ANTLR's default
// (and most horrible) error design.
type syntaxError struct {
	line, column int
	msg          string
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func (r *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException) {
	r.Errors = append(r.Errors, &syntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (r *ErrorListener) clear() {
	r.Errors = []error{}
}

func (r syntaxError) Error() string {
	return r.msg
}

/*
Instance returns an error instance based on any parser and/or lexer
errors that transpire during processing of the tokenized value.

Once accessed, error content is cleared from the receiver.
*/
func (r Instance) Error() (err error) {
	var errtxt string

	// Process all lexer errors
	if len(r.lErr.Errors) > 0 {
		errtxt += "Lexer: "
		for i := 0; i < len(r.lErr.Errors); i++ {
			errtxt += r.lErr.Errors[i].Error()
			if i < len(r.lErr.Errors) {
				errtxt += "; "
			}
		}
		r.lErr.clear()
	}

	// Process all parser errors
	if len(r.pErr.Errors) > 0 {
		if len(errtxt) > 0 {
			errtxt += `; `
		}
		errtxt += "Parser: "
		for i := 0; i < len(r.pErr.Errors); i++ {
			errtxt += r.pErr.Errors[i].Error()
			if i < len(r.pErr.Errors) {
				errtxt += "; "
			}
		}
		r.pErr.clear()
	}

	if len(errtxt) > 0 {
		err = errorf(errtxt)
	}

	return
}

// nameContext extracts zero (0) or more names from the ctx instance,
// and returns them via names ([]string).  Note that presence of a name
// value applies to all definitions EXCEPT LDAPSyntax.
func nameContext(ctx INamesContext) (names []string, err error) {
	if ctx == nil {
		err = errorf("%T instance is nil", ctx)
		return
	}

	names = make([]string, 0)
	_names := ctx.AllQuotedDescriptor()
	for i := 0; i < len(_names); i++ {
		if sym := _names[i].GetText(); len(sym) > 0 {
			names = append(names, trim(sym, `''`))
		}
	}

	// Determine if parentheticals were needed, and
	// if so, that both L & R were specified.  Note
	// that parentheticals are required if more than
	// one name is present, though a single name MAY
	// be similarly encapsulated, but this is optional.
	switch l := len(names); l {
	case 0:
		err = errorf("No names parsed from %T", ctx)
	default:
		var ct int
		if ctx.OpenParen() != nil {
			ct++
		}
		if ctx.CloseParen() != nil {
			ct++
		}

		if ct%2 != 0 {
			// odd number of parentheticals is bad in any scenario.
			err = errorf("Missing open or close parenthesis in %T", ctx)
		} else if ct != 2 && l > 1 {
			// A multi-valued NAME MUST be parenthetical.
			err = errorf("Multi-valued NAME must be parenthetical")
		}
		// A single-valued NAME MAY be parenthetical.
	}

	return
}

/*
extContext extracts the X-<NAME> extension key/value pair from the
ctx instance, and returns it via e.  This applies to all definition
types.
*/
func extContext(ctx IExtensionsContext) (e map[string][]string, err error) {
	if ctx != nil {
		exts := ctx.AllExtension()
		e = make(map[string][]string, 0)
		for i := 0; i < len(exts); i++ {
			if en := exts[i].XString(); en != nil {
				qdesc := exts[i].AllExtensionValue()
				if ev := extVContext(qdesc); len(ev) > 0 {
					e[trimS(en.GetText())] = ev
				}
			}
		}
	}

	return
}

/*
extVContext extracts the extension values from the ctx instance, and
returns them via extv ([]string).
*/
func extVContext(ctx []IExtensionValueContext) (extv []string) {
	if ctx != nil {
		for i := 0; i < len(ctx); i++ {
			var ev string
			if qs := ctx[i].QuotedString(); qs != nil {
				ev = qs.GetText()
			} else if qd := ctx[i].QuotedDescriptor(); qd != nil {
				ev = qd.GetText()
			}
			if ev = trimS(trim(ev, `''`)); len(ev) > 0 {
				extv = append(extv, ev)
			}
		}
	}

	return
}

/*
descContext extracts descriptive quoted-string or quoted-descriptor
text from the ctx instance, returning it via desc.
*/
func descContext(ctx *DescriptionContext) (desc string, err error) {
	if ctx != nil {
		if qs := ctx.QuotedString(); qs != nil {
			desc = trimS(trim(qs.GetText(), `''`))
		} else if qd := ctx.QuotedDescriptor(); qd != nil {
			desc = trimS(trim(qd.GetText(), `''`))
		}
	}

	return
}

/*
numOIDContext extracts the numeric OID from the definition, returning
it via noid.  This applies to all definitions EXCEPT dITStructureRule.
*/
func numOIDContext(ctx any) (noid string, macro []string, err error) {
	switch tv := ctx.(type) {
	case *NumericOIDContext:
		if tv != nil {
			if noid = tv.GetText(); hasPfx(noid, `<missing`) {
				err = errorf("Missing numeric OID for definition")
			}
		} else {
			err = errorf("Nil %T instance", tv)
		}
	case *NumericOIDOrMacroContext:
		if tv != nil {
			if n := tv.NumericOID(); n != nil {
				if noid = n.GetText(); hasPfx(noid, `<missing`) {
					err = errorf("Missing numeric OID for definition")
				}
			} else if mc := tv.Macro(); mc != nil {
				var ok bool
				if macro, ok = macroContext(mc); !ok {
					err = errorf("Invalid macro value(s): %v", macro)
				}
			}
		} else {
			err = errorf("Nil %T instance", tv)
		}
	}

	return
}

/*
parenContext returns an error based on a lack of opening and/or closing
parenthetical encapsulation.  This applies to any multi-valued attribute
sequence value, such as MUST ( cn $ sn $ o ) and names ( 'cn' 'commonName' ).
*/
func parenContext(ctx any) (err error) {
	if ctx != nil {
		switch tv := ctx.(type) {
		case *OpenParenContext:
			if hasPfx(tv.GetText(), `<missing`) {
				err = errorf("Missing open parenthesis for definition '%T'", tv)
			}
		case *CloseParenContext:
			if hasPfx(tv.GetText(), `<missing`) {
				err = errorf("Missing close parenthesis for definition '%T'", tv)
			}
		}
	}

	return
}

/*
isErrImpl returns an error following an attempt to assert
ctx to *antlr.ErrorNodeImpl. If asserted, determine if the
underlying instance describes an erroneous token.
*/
func isErrImpl(typ, id string, ctx any) (err error) {
	if ern, ok := ctx.(*antlr.ErrorNodeImpl); ok {
		// Don't count stray newlines or
		// whitespaces as problem tokens
		txt := trimS(repAll(ern.GetText(), string(rune(10)), ``))
		if len(txt) > 0 {
			err = errorf("Unexpected or bogus token '%s' in %s '%s'",
				txt, typ, id)
		}
	}

	return
}

/*
descrNumOIDContext returns the definition descriptor OR numeric OID
within x and identifies which source was positive via noid.  In no
scenario do both sources apply.
*/
func descrNumOIDContext(ctx *OIDContext) (x string, noid bool, err error) {
	if ctx != nil {
		if _x := ctx.Descriptor(); _x != nil {
			x = _x.GetText()
		} else if _y := ctx.NumericOID(); _y != nil {
			x = _y.GetText()
			noid = true
		}

		if len(x) == 0 {
			err = errorf("No %T parsed", ctx)
		}
	}

	return
}

/*
Use of single-valued OIDs applies to attributeType, dITStructureRule
and nameForm definitions.
*/
func oIDContext(ctx *OIDContext) (n, d []string, err error) {
	var (
		x    string
		noid bool
	)

	if x, noid, err = descrNumOIDContext(ctx); err != nil {
		return
	}

	if x = trimS(trim(x, `''`)); len(x) > 0 {
		if noid {
			n = append(n, x)
		} else {
			d = append(d, x)
		}
	}

	return
}

/*
Use of multi-valued OIDs applies to objectClass, dITContentRule,
nameForm, and matchingRuleUse definitions.
*/
func oIDsContext(ctx *OIDsContext) (n, d []string, err error) {
	ch := ctx.AllOID()

	for i := 0; i < len(ch); i++ {
		var _n, _d []string
		assert, ok := ch[i].(*OIDContext)
		if !ok {
			err = errorf("%T type assertion failed for %T", ch[i], assert)
			break
		}

		if _n, _d, err = oIDContext(assert); err != nil {
			break
		}

		if len(_n) > 0 {
			n = append(n, _n...)
		}

		if len(_d) > 0 {
			d = append(d, _d...)
		}
	}

	return
}

/*
syntaxContext extracts the syntax numeric OID from the ctx instance,
and returns it through soid.  This applies to attributeType and
matchingRule definitions.
*/
func syntaxContext(ctx *SyntaxContext) (soid string, err error) {
	if ctx != nil {
		_soid := ctx.NumericOID()
		if pfx := _soid.GetText(); hasPfx(pfx, `<missing`) {
			err = errorf("Missing numeric OID for definition syntax")
		} else {
			soid = _soid.GetText()
		}
	}

	return
}

/*
macroContext extracts the macro data from the definition, returning
slices of macro values (name,oid_suffix) alongside a Boolean value.
The macro slice values are used to resolve and acquire the complete
numeric OID.  The Boolean value indicates that a complete macro was
acquired.
*/
func macroContext(ctx IMacroContext) (macro []string, ok bool) {
	var m, s string

	if m = ctx.Descriptor().GetText(); hasPfx(m, `<missing`) {
		return
	} else if s = ctx.MacroSuffix().GetText(); hasPfx(s, `<missing`) {
		return
	}

	if ok = len(m) > 0 && len(s) > 0; ok {
		if hasPfx(s, `:`) || hasPfx(s, `.`) {
			s = s[1:]
		}
		macro = []string{m, s}
	}
	return
}

func strInSlice(str string, sl []string) bool {
	if len(str) == 0 {
		return false
	}

	for i := 0; i < len(sl); i++ {
		if eq(sl[i], str) {
			return true
		}
	}

	return false
}
