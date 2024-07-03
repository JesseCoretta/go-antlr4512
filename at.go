package antlr4512

/*
AttributeType implements RFC 4512 Section 4.1.2.
*/
type AttributeType struct {
	OID        string
	Macro      []string
	Name       []string
	Desc       string
	Obsolete   bool
	SuperType  string
	Equality   string
	Ordering   string
	Substring  string
	Syntax     string
	Single     bool
	Collective bool
	Immutable  bool
	Usage      string
	MUB        uint
	Extensions map[string][]string
}

/*
AttributeTypes is an instance of slices of [AttributeType] instances.
*/
type AttributeTypes []AttributeType

/*
ParseAttributeType processes raw into an instance of [AttributeType],
which is returned alongside an error.
*/
func ParseAttributeType(raw string) (at AttributeType, err error) {
	var i Instance
	if i, err = ParseInstance(raw); err != nil {
		return
	}

	err = at.process(i.P.AttributeTypeDescription())

	return

}

func (r *AttributeType) process(ctx IAttributeTypeDescriptionContext) (err error) {
	if !checkParen(ctx.OpenParen()) || !checkParen(ctx.CloseParen()) {
		err = errorf("Unbalanced parenthetical state for %T", r)
		return
	}

	for k, ct := 0, ctx.GetChildCount(); k < ct && err == nil; k++ {
		switch tv := ctx.GetChild(k).(type) {
		case *NumericOIDOrMacroContext:
			r.OID, r.Macro, err = numOIDContext(tv)

		case *UsageContext,
			*NameContext,
			*ExtensionsContext,
			*DescriptionContext:
			err = r.setMisc(tv)

		case *CollectiveContext,
			*SingleValueContext,
			*ObsolescenceContext,
			*ImmutabilityContext:
			r.setBoolean(tv)

		case *SyntaxContext,
			*MinUpperBoundsContext:
			err = r.setSyntaxContexts(tv)

		case *EqualityContext,
			*OrderingContext,
			*SubstrContext:
			err = r.setMatchingRuleContexts(tv)

		case *SuperTypeContext:
			if o := tv.OID(); o != nil {
				r.SuperType = o.GetText()
			}

		default:
			err = isErrImpl(`attributeType`, r.OID, tv)
		}
	}

	if len(r.OID) == 0 && len(r.Macro) == 0 {
                err = errorf("Neither OID nor Macro literals found in %T", r)
	}

	return
}

func (r *AttributeType) setMisc(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *NameContext:
		r.Name, err = nameContext(tv.Names())
	case *UsageContext:
		r.Usage = trimL(tv.GetText(),` USAGE `)
	case *DescriptionContext:
		r.Desc, err = descContext(tv)
	case *ExtensionsContext:
		r.Extensions, err = extContext(tv)
	default:
		err = errorf("Unknown miscellaneous context '%T'", ctx)
	}

	return
}

func (r *AttributeType) setMatchingRuleContexts(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *EqualityContext:
		err = r.setMatchingRuleContext(1, tv.OID())
	case *OrderingContext:
		err = r.setMatchingRuleContext(2, tv.OID())
	case *SubstrContext:
		err = r.setMatchingRuleContext(3, tv.OID())
	default:
		err = errorf("Unknown matchingRule context %T", ctx)
	}

	return
}

func (r *AttributeType) setMatchingRuleContext(typ int, ctx IOIDContext) (err error) {
	var n, d []string
	if n, d, err = oIDContext(ctx.(*OIDContext)); err != nil {
		return
	}

	var mr string
	if len(n) > 0 {
		mr = n[0]
	} else if len(d) > 0 {
		mr = d[0]
	} else {
		err = errorf("No %T value found", ctx)
		return
	}

	switch typ {
	case 1:
		r.Equality = mr
	case 2:
		r.Ordering = mr
	case 3:
		r.Substring = mr
	default:
		err = errorf("Invalid matchingRule type '%d'", typ)
	}

	return
}

func (r *AttributeType) setSyntaxContexts(ctx any) (err error) {

	switch tv := ctx.(type) {
	case *SyntaxContext:
		var syn string
		if syn, err = syntaxContext(tv); err != nil {
			break
		}

		r.Syntax = syn

	case *MinUpperBoundsContext:
		var m int
		m, err = atoi(trimS(trim(tv.GetText(), `{}`)))
		if err != nil || m < 0 {
			break
		}
		r.MUB = uint(m)
	default:
		err = errorf("Unknown attribute syntax context %T", ctx)
	}

	return
}

func (r *AttributeType) setBoolean(x any) {

	switch x.(type) {
	case *SingleValueContext:
		r.Single = true
	case *ImmutabilityContext:
		r.Immutable = true
	case *ObsolescenceContext:
		r.Obsolete = true
	case *CollectiveContext:
		r.Collective = true
	}
}
