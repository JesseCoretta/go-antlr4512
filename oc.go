package antlr4512

/*
ObjectClass implements RFC 4512 Section 4.1.1.
*/
type ObjectClass struct {
	OID          string
	Macro        []string
	Name         []string
	Desc         string
	Obsolete     bool
	SuperClasses []string
	Kind         string
	Must         []string
	May          []string
	Extensions   map[string][]string
}

/*
ObjectClasses is an instance of slices of [ObjectClass] instances.
*/
type ObjectClasses []ObjectClass

/*
ParseObjectClass processes raw into an instance of [ObjectClass],
which is returned alongside an error.
*/
func ParseObjectClass(raw string) (oc ObjectClass, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}

	// Ensure what went in matches precisely what came out
	x := i.P.ObjectClassDescription()
	if err = oc.process(x); err == nil {
		ptext := x.GetText()
		if raw != ptext {
			err = errorf("Inconsistent %T parse results or bad content", oc)
		}
	}

	return

}

func (r *ObjectClass) process(ctx IObjectClassDescriptionContext) (err error) {
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

		case *KindContext:
			r.Kind = trimS(tv.GetText())

		case *SuperClassesContext:
			err = r.superClassesContext(tv)

		case *MustContext:
			err = r.mustContext(tv)

		case *MayContext:
			err = r.mayContext(tv)

		default:
			err = isErrImpl(`objectClass`, r.OID, tv)
		}
	}

	if len(r.Kind) == 0 {
		r.Kind = `STRUCTURAL`
	}

	if len(r.OID) == 0 && len(r.Macro) == 0 {
		err = errorf("Neither OID nor Macro literals found in %T", r)
	}

	return
}

func (r *ObjectClass) setMisc(ctx any) (err error) {
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

func (r *ObjectClass) setCritical(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *NumericOIDOrMacroContext:
		r.OID, r.Macro, err = numOIDContext(tv)
	default:
		err = errorf("Unknown critical context '%T'", ctx)
	}

	return
}

func (r *ObjectClass) superClassesContext(ctx *SuperClassesContext) (err error) {
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

		var sup []string
		if err == nil {
			sup = append(sup, n...)
			sup = append(sup, d...)
		}

		if len(sup) == 0 {
			if err != nil {
				err = errorf("No super classes found in %T: %v", ctx, err)
			} else {
				err = errorf("No super classes found in %T", ctx)
			}
		} else {
			r.SuperClasses = sup
		}
	}

	return
}

func (r *ObjectClass) mustContext(ctx *MustContext) (err error) {
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

func (r *ObjectClass) mayContext(ctx *MayContext) (err error) {
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
