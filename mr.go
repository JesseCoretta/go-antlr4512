package antlr4512

/*
MatchingRule implements RFC 4512 Section 4.1.3.
*/
type MatchingRule struct {
	OID        string
	Macro	   []string
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
        if err = parenContext(ctx.OpenParen(), ctx.CloseParen()); err != nil {
                return                                                  
        }

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *NumericOIDOrMacroContext:
			r.OID, r.Macro, err = numOIDContext(tv)

		case *NameContext,
			*ExtensionsContext,
			*DescriptionContext:
			err = r.setMisc(tv)

		case *ObsolescenceContext:
			r.Obsolete = true

		case *SyntaxContext:
			r.Syntax, err = syntaxContext(tv)

		default:
			err = isErrImpl(`matchingRule`, r.OID, tv)
		}
	}

	if len(r.OID) == 0 && len(r.Macro) == 0 {
                err = errorf("Neither OID nor Macro literals found in %T", r)
	} else if len(r.Syntax) == 0 {
		err = errorf("No numeric OID found within %T.Syntax", r)
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

