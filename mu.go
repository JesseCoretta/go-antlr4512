package antlr4512

/*
MatchingRuleUse implements RFC 4512 Section 4.1.4.
*/
type MatchingRuleUse struct {
	OID        string
	Name       []string
	Desc       string
	Obsolete   bool
	Applies    []string
	Extensions map[string][]string
}

type MatchingRuleUses []MatchingRuleUse

func ParseMatchingRuleUse(raw string) (mu MatchingRuleUse, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}

	err = mu.process(i.P.MatchingRuleUseDescription())

	return

}

func (r *MatchingRuleUse) process(ctx IMatchingRuleUseDescriptionContext) (err error) {

	if ctx.OpenParen() == nil || ctx.CloseParen() == nil {
		err = errorf("Definition encapsulation error; missing parenthesis?")
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

		case *AppliesContext:
			err = r.appliesContext(tv)

		default:
			err = isErrImpl(`matchingRuleUse`, r.OID, tv)
		}
	}

	return
}

func (r *MatchingRuleUse) setMisc(ctx any) (err error) {

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

func (r *MatchingRuleUse) setCritical(ctx any) (err error) {

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

func (r *MatchingRuleUse) appliesContext(ctx *AppliesContext) (err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*OIDsContext))
		}

		var applies []string
		if err == nil {
			applies = append(applies, n...)
			applies = append(applies, d...)
		}

		if len(applies) == 0 {
			if err != nil {
				err = errorf("No required types found in %T: %v", ctx, err)
			} else {
				err = errorf("No required types found in %T", ctx)
			}
		} else {
			r.Applies = applies
		}
	}

	return
}
