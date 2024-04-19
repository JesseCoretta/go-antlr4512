package antlr4512

/*
DITContentRule implements RFC 4512 Section 4.1.6.
*/
type DITContentRule struct {
	OID        string
	Macro      []string
	Name       []string
	Desc       string
	Obsolete   bool
	Aux        []string
	Must       []string
	May        []string
	Not        []string
	Extensions map[string][]string
}

/*
DITContentRules is an instance of slices of [DITContentRule] instances.
*/
type DITContentRules []DITContentRule

/*
ParseDITContentRule processes raw into an instance of [DITContentRule],
which is returned alongside an error.
*/
func ParseDITContentRule(raw string) (oc DITContentRule, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}

	err = oc.process(i.P.DITContentRuleDescription())

	return

}

func (r *DITContentRule) process(ctx IDITContentRuleDescriptionContext) (err error) {

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *NumericOIDOrMacroContext,
			*OpenParenContext,
			*CloseParenContext:
			err = r.setCritical(tv)

		case *NameContext,
			*ExtensionsContext,
			*DescriptionContext:
			err = r.setMisc(tv)

		case *ObsolescenceContext:
			r.Obsolete = true

		case *AuxContext:
			err = r.auxContext(tv)

		case *MustContext:
			err = r.mustContext(tv)

		case *MayContext:
			err = r.mayContext(tv)

		case *NotContext:
			err = r.notContext(tv)

		default:
			err = isErrImpl(`dITContentRule`, r.OID, tv)
		}
	}

	if len(r.OID) == 0 && len(r.Macro) == 0 {
		err = errorf("No macro or OID literal for %T", r)
	}

	return
}

func (r *DITContentRule) setMisc(ctx any) (err error) {

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

func (r *DITContentRule) setCritical(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *NumericOIDOrMacroContext:
		r.OID, r.Macro, err = numOIDContext(tv)
	case *OpenParenContext, *CloseParenContext:
		err = parenContext(tv)
	default:
		err = errorf("Unknown critical context '%T'", ctx)
	}

	return
}

func (r *DITContentRule) mustContext(ctx *MustContext) (err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*OIDsContext))
		}

		var must []string
		if err == nil {
			must = append(must, n...)
			must = append(must, d...)
		}

		if len(must) == 0 {
			if err != nil {
				err = errorf("No required types found in %T: %v", ctx, err)
			} else {
				err = errorf("No required types found in %T", ctx)
			}
		} else {
			r.Must = must
		}
	}

	return
}

func (r *DITContentRule) mayContext(ctx *MayContext) (err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*OIDsContext))
		}

		var may []string
		if err == nil {
			may = append(may, n...)
			may = append(may, d...)
		}

		if len(may) == 0 {
			if err != nil {
				err = errorf("No permitted types found in %T: %v", ctx, err)
			} else {
				err = errorf("No permitted types found in %T", ctx)
			}
		} else {
			r.May = may
		}
	}

	return
}

func (r *DITContentRule) notContext(ctx *NotContext) (err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*OIDsContext))
		}

		var not []string
		if err == nil {
			not = append(not, n...)
			not = append(not, d...)
		}

		if len(not) == 0 {
			if err != nil {
				err = errorf("No disallowed types found in %T: %v", ctx, err)
			} else {
				err = errorf("No disallowed types found in %T", ctx)
			}
		} else {
			r.Not = not
		}
	}

	return
}

func (r *DITContentRule) auxContext(ctx *AuxContext) (err error) {
	if ctx != nil {
		var n, d []string

		if x := ctx.OID(); x != nil {
			n, d, err = oIDContext(x.(*OIDContext))
		} else if y := ctx.OIDs(); y != nil {
			n, d, err = oIDsContext(y.(*OIDsContext))
		}

		var aux []string
		if err == nil {
			aux = append(aux, n...)
			aux = append(aux, d...)
		}

		if len(aux) == 0 {
			if err != nil {
				err = errorf("No auxiliary classes found in %T: %v", ctx, err)
			} else {
				err = errorf("No auxiliary classes types found in %T", ctx)
			}
		} else {
			r.Aux = aux
		}
	}

	return
}
