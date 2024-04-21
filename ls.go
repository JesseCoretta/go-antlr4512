package antlr4512

/*
LDAPSyntax implements RFC 4512 Section 4.1.5.
*/
type LDAPSyntax struct {
	OID        string
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
	if !checkParen(ctx.OpenParen()) || !checkParen(ctx.CloseParen()) {
		err = errorf("Unbalanced parenthetical state for %T", r)
		return
	}

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *NumericOIDContext,
			*OpenParenContext,
			*CloseParenContext:
			err = r.setCritical(tv)

		case *ExtensionsContext,
			*DescriptionContext:
			err = r.setMisc(tv)

		default:
			err = isErrImpl(`ldapSyntax`, r.OID, tv)
		}
	}

	if len(r.OID) == 0 {
		err = errorf("No OID literal for %T", r)
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

func (r *LDAPSyntax) setCritical(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *NumericOIDContext:
		r.OID, _, err = numOIDContext(tv)
	case *OpenParenContext, *CloseParenContext:
		err = parenContext(tv)
	default:
		err = errorf("Unknown critical context '%T'", ctx)
	}

	return
}
