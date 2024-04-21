package antlr4512

/*
DITStructureRule implements RFC 4512 Section 4.1.1.
*/
type DITStructureRule struct {
	ID         string
	Name       []string
	Desc       string
	Obsolete   bool
	Form       string
	SuperRules []string
	Extensions map[string][]string
}

/*
DITStructureRules is an instance of slices of [DITStructureRule] instances.
*/
type DITStructureRules []DITStructureRule

/*
ParseDITStructureRule processes raw into an instance of [DITStructureRule],
which is returned alongside an error.
*/
func ParseDITStructureRule(raw string) (oc DITStructureRule, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}

	err = oc.process(i.P.DITStructureRuleDescription())

	return

}

func (r *DITStructureRule) process(ctx IDITStructureRuleDescriptionContext) (err error) {
	if !checkParen(ctx.OpenParen()) || !checkParen(ctx.CloseParen()) {
		err = errorf("Unbalanced parenthetical state for %T", r)
		return
	}

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *StructureRuleContext:
			err = r.setCritical(tv)

		case *NameContext,
			*ExtensionsContext,
			*DescriptionContext:
			err = r.setMisc(tv)

		case *ObsolescenceContext:
			r.Obsolete = true

		case *FormContext:
			if o := tv.OID(); o != nil {
				r.Form = o.GetText()
			}

		case *SuperRulesContext:
			err = r.superRulesContext(tv)

		default:
			err = isErrImpl(`dITStructureRule`, r.ID, tv)
		}
	}

	if len(r.Form) == 0 {
		err = errorf("No nameForm reference set for %T", r)
	}

	return
}

func (r *DITStructureRule) setMisc(ctx any) (err error) {

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

func (r *DITStructureRule) setCritical(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *StructureRuleContext:
		r.ID = tv.GetText()
	default:
		err = errorf("Unknown critical context '%T'", ctx)
	}

	return
}

func (r *DITStructureRule) superRulesContext(ctx *SuperRulesContext) (err error) {
	if ctx != nil {
		var sids []string
		if x := ctx.StructureRule(); x != nil {
			var sid string
			sid, err = ruleIDContext(x.(*StructureRuleContext))
			if err == nil {
				sids = append(sids, sid)
			}

		} else if y := ctx.StructureRules(); y != nil {
			sids, err = ruleIDsContext(y.(*StructureRulesContext))
		}
		if len(sids) > 0 {
			r.SuperRules = sids
		}
	}

	return
}

// Use of single-valued OIDs applies to attributeType, dITStructureRule
// and nameForm definitions.
func ruleIDContext(ctx *StructureRuleContext) (id string, err error) {

	dig := ctx.Number()
	if dig == nil {
		err = errorf("%T.%T is nil", ctx, dig)
		return
	}

	id = trimS(trim(dig.GetText(), `''`))

	return
}

// Use of multi-valued IDs applies solely to subordinate dITStructureRules
// in reference to their respective superior rules.
func ruleIDsContext(ctx *StructureRulesContext) (sids []string, err error) {
	ch := ctx.AllStructureRule()

	for i := 0; i < len(ch) && err == nil; i++ {
		assert, ok := ch[i].(*StructureRuleContext)
		if !ok {
			err = errorf("%T type assertion failed for %T", ch[i], assert)
			break
		}

		var sid string
		if sid, err = ruleIDContext(assert); err != nil {
			break
		}
		sids = append(sids, sid)
	}

	if len(sids) == 0 {
		err = errorf("No rule IDs parsed")
	}

	return
}
