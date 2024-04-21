package antlr4512

/*
MatchingRule implements RFC 4512 Section 4.1.3.
*/
type MatchingRule struct {
	OID        string
	Name       []string
	Desc       string
	Obsolete   bool
	Syntax     string
	Extensions map[string][]string
}

/*
MatchingRules is an instance of slices of [MatchingRule] instances.
*/
type MatchingRules []MatchingRule

/*
ParseMatchingRule processes raw into an instance of [MatchingRule],
which is returned alongside an error.
*/
func ParseMatchingRule(raw string) (mr MatchingRule, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}

	err = mr.process(i.P.MatchingRuleDescription())

	return

}

func (r *MatchingRule) process(ctx IMatchingRuleDescriptionContext) (err error) {
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

		case *NameContext,
			*ExtensionsContext,
			*DescriptionContext:
			err = r.setMisc(tv)

		case *ObsolescenceContext:
			r.Obsolete = true

		case *SyntaxContext:
			if o := tv.NumericOID(); o != nil {
				r.Syntax = o.GetText()
			} else if p := tv.QuotedDescriptor(); p != nil {
				r.Syntax = trim(p.GetText(), `''`)
			}

		default:
			err = isErrImpl(`matchingRule`, r.OID, tv)
		}
	}

	if len(r.OID) == 0 {
		err = errorf("No OID literal for %T", r)
	} else if len(r.Syntax) == 0 {
		err = errorf("No numeric OID for syntax found within %T", r)
	}

	return
}

func (r *MatchingRule) setMisc(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *NameContext:
		r.Name, err = nameContext(tv.Names())
	case *DescriptionContext:
		r.Desc, err = descContext(tv)
	case *ExtensionsContext:
		r.Extensions, err = extContext(tv)
	default:
		err = errorf("Unknown miscellaneous context '%T'", ctx)
	}

	return
}

func (r *MatchingRule) setCritical(ctx any) (err error) {
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
