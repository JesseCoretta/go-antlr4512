package antlr4512

//import "fmt"

/*
LDAPSyntax implements RFC 4512 Section 4.1.5.
*/
type LDAPSyntax struct {
	OID        string
	Macro	   []string
	Desc       string
	Extensions map[string][]string
}

/*
LDAPSyntaxes is an instance of slices of [LDAPSyntax] instances.
*/
type LDAPSyntaxes []LDAPSyntax

func newLDAPSyntax() LDAPSyntax {
	return LDAPSyntax{
		Extensions: make(map[string][]string, 0),
	}
}

/*
ParseLDAPSyntax processes raw into an instance of [LDAPSyntax],
which is returned alongside an error.
*/
func ParseLDAPSyntax(raw string) (ls LDAPSyntax, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}
	ls = newLDAPSyntax()

	err = ls.process(i.P.LDAPSyntaxDescription())

	return

}

func (r *LDAPSyntax) process(ctx ILDAPSyntaxDescriptionContext) (err error) {
	if err = parenContext(ctx.OpenParen(), ctx.CloseParen()); err != nil {
		return
	}

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *NumericOIDOrMacroContext:
			r.OID, r.Macro, err = numOIDContext(tv)

		case *ExtensionsContext,
			*DescriptionContext:
			err = r.setMisc(tv)

		default:
			err = isErrImpl(`ldapSyntax`, r.OID, tv)
		}
	}

	if len(r.OID) == 0 && len(r.Macro) == 0 {
		err = errorf("Neither OID nor Macro literals found in %T", r)
	}

	return
}

func (r *LDAPSyntax) setMisc(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *DescriptionContext:
		r.Desc, err = descContext(tv)
	case *ExtensionsContext:
		r.Extensions, err = extContext(tv)
	default:
		err = errorf("Unknown miscellaneous context '%T'", ctx)
	}

	return
}

