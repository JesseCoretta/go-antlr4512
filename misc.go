package rfc4512antlr

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Error option types.
//
// TODO: make configurable.
//var (
//	antlrUseStderr bool // not recommended!
//	lxErr,
//	pErr *errorListener
//)

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

func (c *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &syntaxError{
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
	if len(r.lErr.Errors) > 0 {
		errtxt += "Lexer: " + r.lErr.Errors[0].Error()
		r.lErr.clear()
	}

	if len(r.pErr.Errors) > 0 {
		if len(errtxt) > 0 {
			errtxt += `; `
		}
		errtxt += "Parser: " + r.pErr.Errors[0].Error()
		r.pErr.clear()
	}

	if len(errtxt) > 0 {
		err = errorf(errtxt)
	}

	return
}
