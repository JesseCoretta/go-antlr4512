package antlr4512

/*
NameForm implements RFC 4512 Section 4.1.7.2.
*/
type NameForm struct {
	OID        string
	Macro      []string
	Name       []string
	Desc       string
	Obsolete   bool
	OC         string
	Must       []string
	May        []string
	Extensions map[string][]string
}

/*
NameForms is an instance of slices of [NameForm] instances.
*/
type NameForms []NameForm

/*
ParseNameForm processes raw into an instance of [NameForm],
which is returned alongside an error.
*/
func ParseNameForm(raw string) (nf NameForm, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}

	err = nf.process(i.P.NameFormDescription())

	return

}

func (r *NameForm) process(ctx INameFormDescriptionContext) (err error) {

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

		case *StructuralClassContext:
			if o := tv.OID(); o != nil {
				r.OC = o.GetText()
			}

		case *MustContext:
			err = r.mustContext(tv)

		case *MayContext:
			err = r.mayContext(tv)

		default:
			err = isErrImpl(`nameForm`, r.OID, tv)
		}
	}

	if len(r.OID) == 0 && len(r.Macro) == 0 {
		err = errorf("No macro or OID literal for %T", r)
	} else if len(r.OC) == 0 || len(r.Must) == 0 {
		err = errorf("No OC or MUST clauses satisfied for %T", r)
	}

	return
}

func (r *NameForm) setMisc(ctx any) (err error) {

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

func (r *NameForm) setCritical(ctx any) (err error) {

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

func (r *NameForm) mustContext(ctx *MustContext) (err error) {
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

func (r *NameForm) mayContext(ctx *MayContext) (err error) {
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
