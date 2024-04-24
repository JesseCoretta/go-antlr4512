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

/*
MatchingRuleUses is an instance of slices of [MatchingRuleUse] instances.
*/
type MatchingRuleUses []MatchingRuleUse

/*
ParseMatchingRuleUse processes raw into an instance of [MatchingRuleUse]. Note
that it is unusual for this function to be used and exists mainly for the sake
of completeness.

A [MatchingRuleUse] instance is simply a manifest of attributeType instances
which utilize a given [MatchingRule] for EQUALITY, SUBSTR and ORDERING fields.

Generally, these [MatchingRuleUse] instances are created -- by the DSA -- based
on equivalent [MatchingRule] instances, and only manifest when the number of
"applied" [AttributeType] users is greater than zero (0) for the relevant
[MatchingRule].

In short, one generally need only load ldapSyntax, matchingRule and attributeType
definitions to initialize or extend matchingRuleUse instances.
*/
func ParseMatchingRuleUse(raw string) (mu MatchingRuleUse, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}

	err = mu.process(i.P.MatchingRuleUseDescription())

	return

}

func (r *MatchingRuleUse) process(ctx IMatchingRuleUseDescriptionContext) (err error) {
        if err = parenContext(ctx.OpenParen(), ctx.CloseParen()); err != nil {
                return                                                  
        }

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *NumericOIDContext:
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

	if len(r.OID) == 0 {
		err = errorf("Missing numeric OID for %T", r)
	} else if len(r.Applies) == 0 {
		err = errorf("Missing applied OIDs for %T", r)
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
                        if err = parenContext(y.OpenParen(), y.CloseParen()); err != nil {
                                return
                        }
			n, d, err = oIDsContext(y.(*OIDsContext))
		}

		var applies []string
		if err == nil {
			applies = append(applies, n...)
			applies = append(applies, d...)
		}

		if len(applies) == 0 {
			if err != nil {
				err = errorf("No applied types found in %T: %v", ctx, err)
			} else {
				err = errorf("No applied types found in %T", ctx)
			}
		} else {
			r.Applies = applies
		}
	}

	return
}
