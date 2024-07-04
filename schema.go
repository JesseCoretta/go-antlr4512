package antlr4512

/*
Schema contains slices of RFC 4512 Section 4.1 contexts.
*/
type Schema struct {
	DN string
	OM Macros
	LS LDAPSyntaxes
	MR MatchingRules
	MU MatchingRuleUses
	AT AttributeTypes
	OC ObjectClasses
	DC DITContentRules
	NF NameForms
	DS DITStructureRules
}

/*
NewSchema initializes and returns a new empty instance of *Schema.
*/
func NewSchema() Schema {
	return Schema{
		LS: make(LDAPSyntaxes, 0),
		OM: make(Macros, 0),
		MR: make(MatchingRules, 0),
		MU: make(MatchingRuleUses, 0),
		AT: make(AttributeTypes, 0),
		OC: make(ObjectClasses, 0),
		DC: make(DITContentRules, 0),
		NF: make(NameForms, 0),
		DS: make(DITStructureRules, 0),
	}
}

/*
Macros contains a manifest of name to oid/name pairs, each describing
a resolution path for oid/name into a distinct OID reference.
*/
type Macros map[string]string

/*
IsZero returns a Boolean value indicative of receiver nilness.
*/
func (r *Schema) IsZero() bool {
	return r == nil
}

/*
UpdateMatchingRuleUses updates the underlying [MatchingRuleUses]
instance within the receiver instance.
*/
func (r *Schema) UpdateMatchingRuleUses() {
	if len(r.MU) == 0 {
		r.MU = make(MatchingRuleUses, 0)
	}

	for i := 0; i < len(r.MR); i++ {
		var (
			mr MatchingRule    = r.MR[i]
			mu MatchingRuleUse = MatchingRuleUse{
				OID:     mr.OID,
				Name:    mr.Name,
				Desc:    mr.Desc,
				Applies: make([]string, 0),
			}
		)

		for j := 0; j < len(r.AT); j++ {
			var at AttributeType = r.AT[j]
			for _, amr := range []string{
				at.Equality,
				at.Substring,
				at.Ordering,
			} {
				if eq(amr, mr.OID) || strInSlice(amr, mr.Name) {
					if !strInSlice(amr, mu.Applies) {
						mu.Applies = append(mu.Applies, at.OID)
					}
				}
			}
		}

		if len(mu.Applies) > 0 {
			var found bool
			for j := 0; j < len(r.MU); j++ {
				if found = r.MU[j].OID == mu.OID; found {
					for a := 0; a < len(mu.Applies); a++ {
						o := mu.Applies[a]
						if !strInSlice(o, r.MU[j].Applies) {
							r.MU[j].Applies = append(r.MU[j].Applies, o)
						}
					}
					break
				}
			}

			if !found {
				r.MU = append(r.MU, mu)
			}
		}
	}
}

/*
ParseRaw returns an error following an attempt to parse the input raw
bytes into the receiver instance. Both the [Schema.ParseFile] method
and the [Schema.ParseDirectory] method are dependents of this method.

The raw input value must be non-nil, and must contain valid schema
definitions with optional bash comments.
*/
func (r *Schema) ParseRaw(raw []byte) (err error) {
	if r.IsZero() {
		err = errorf("%T instance is not initialized", r)
		return
	} else if len(raw) == 0 {
		err = errorf("No raw content to parse")
		return
	}

        var i Instance                                                  
        if i, err = ParseInstance(string(raw)); err != nil {        
                return                                                  
        }                                                               
                                                                        
        fp := i.P.Fileparse()                                           
                                                                        
        defs := fp.Definitions()                                        
        if defs == nil {                                                
                err = errorf("No definitions found in %T", defs)        
                return                                                  
        }                                                               
                                                                        
        r.processObjectIdentifiers(defs.AllObjectIdentifier())          
                                                                        
        if dn := fp.SchemaDN(); dn != nil && len(r.DN) == 0 {           
                r.DN = trimS(trimL(dn.GetText(), `dn:`))                
        }                                                               
                                                                        
        if err = r.processSchemaDefinitions(defs); err == nil {         
                r.UpdateMatchingRuleUses()                              
        }                                                               
                                                                        
        return
}

/*
ParseFile returns an error following an attempt to parse file. Only
files ending in ".schema" will be considered, however submission of
non-qualifying files shall not produce an error.
*/
func (r *Schema) ParseFile(file string) (err error) {
	if r.IsZero() {
		err = errorf("%T instance is not initialized", r)
		return
	}

	var content []byte
	if content, err = readFile(file); err != nil {
		return
	}

	err = r.ParseRaw(content)

	return
}

/*
ParseDirectory returns an error following an attempt to parse the
directory found at dir. Sub-directories encountered are traversed
indefinitely.  Files encountered will only be read if their name
ends in ".schema", at which point their contents are read into
bytes, processed using ANTLR and written to the receiver instance.
*/
func (r *Schema) ParseDirectory(dir string) (err error) {
	if r.IsZero() {
		err = errorf("%T instance is not initialized", r)
		return
	}

	var content []byte
	if content, err = readDirectory(dir); err != nil {
		return
	}

	err = r.ParseRaw(content)

	return
}

/*
processSchemaDefinitions returns an error instance following an
attempt to parse all schema definitions found within the defs
IDefinitionsContext instance.
*/
func (r *Schema) processSchemaDefinitions(defs IDefinitionsContext) (err error) {
	ct := defs.GetChildCount()
	for i := 0; i < ct && err == nil; i++ {
		switch tv := defs.GetChild(i).(type) {
		case *LDAPSyntaxDescriptionsContext:
			err = r.processLSs(tv)
		case *MatchingRuleDescriptionsContext:
			err = r.processMRs(tv)
		case *MatchingRuleUseDescriptionsContext:
			err = r.processMUs(tv)
		case *AttributeTypeDescriptionsContext:
			err = r.processATs(tv)
		case *ObjectClassDescriptionsContext:
			err = r.processOCs(tv)
		case *DITContentRuleDescriptionsContext:
			err = r.processDCs(tv)
		case *NameFormDescriptionsContext:
			err = r.processNFs(tv)
		case *DITStructureRuleDescriptionsContext:
			err = r.processDSs(tv)
		}
	}

	return
}

/*
processATs returns an error instance following an attempt to parse all AttributeType
definitions found within the input at *AttributeTypeDescriptionsContext.
*/
func (r *Schema) processATs(x *AttributeTypeDescriptionsContext) (err error) {
	_z := x.AllAttributeTypeDescription()
	if len(r.AT) == 0 {
		r.AT = make(AttributeTypes, 0)
	}

	for j := 0; j < len(_z); j++ {
		var def AttributeType
		if err = def.process(_z[j]); err != nil {
			break
		}

		if len(def.Macro) == 2 && len(def.OID) == 0 {
			resl, found := r.OM[def.Macro[0]]
			if !found {
				err = errorf("Unresolvable macro for attributeType")
				break
			}
			def.OID = resl + def.Macro[1]
		} else if len(def.Macro) < 2 && len(def.OID) == 0 {
			err = errorf("Incomplete or missing macro for OID resolution")
			break
		}

		r.AT = append(r.AT, def)
	}

	return
}

/*
processLSs returns an error instance following an attempt to parse all LDAPSyntax
definitions found within the input at *LDAPSyntaxDescriptionsContext.
*/
func (r *Schema) processLSs(x *LDAPSyntaxDescriptionsContext) (err error) {
	_z := x.AllLDAPSyntaxDescription()
	if len(r.LS) == 0 {
		r.LS = make(LDAPSyntaxes, 0)
	}

	for j := 0; j < len(_z); j++ {
		var def LDAPSyntax
		if err = def.process(_z[j]); err != nil {
			break
		}

		r.LS = append(r.LS, def)
	}

	return
}

/*
processMRs returns an error instance following an attempt to parse all MatchingRule
definitions found within the input at *MatchingRuleDescriptionsContext.
*/
func (r *Schema) processMRs(x *MatchingRuleDescriptionsContext) (err error) {
	_z := x.AllMatchingRuleDescription()
	if len(r.MR) == 0 {
		r.MR = make(MatchingRules, 0)
	}

	for j := 0; j < len(_z); j++ {
		var def MatchingRule
		if err = def.process(_z[j]); err != nil {
			break
		}
		r.MR = append(r.MR, def)
	}

	return
}

/*
processMUs returns an error instance following an attempt to parse all [MatchingRuleUse]
definitions found within the input at [MatchingRuleUseDescriptionsContext].
*/
func (r *Schema) processMUs(x *MatchingRuleUseDescriptionsContext) (err error) {
	_z := x.AllMatchingRuleUseDescription()
	if len(r.MU) == 0 {
		r.MU = make(MatchingRuleUses, 0)
	}

	for j := 0; j < len(_z); j++ {
		var def MatchingRuleUse
		if err = def.process(_z[j]); err != nil {
			break
		}
		r.MU = append(r.MU, def)
	}

	return
}

/*
processDCs returns an error instance following an attempt to parse all DITContentRule
definitions found within the input at *DITContentRuleDescriptionsContext.
*/
func (r *Schema) processDCs(x *DITContentRuleDescriptionsContext) (err error) {
	_z := x.AllDITContentRuleDescription()
	if len(r.DC) == 0 {
		r.DC = make(DITContentRules, 0)
	}
	for j := 0; j < len(_z); j++ {
		var def DITContentRule
		if err = def.process(_z[j]); err != nil {
			break
		}

		if len(def.Macro) == 2 && len(def.OID) == 0 {
			resl, found := r.OM[def.Macro[0]]
			if !found {
				err = errorf("Unresolvable macro for matchingRule")
				break
			}
			def.OID = resl + `.` + def.Macro[1]
		} else if len(def.Macro) < 2 && len(def.OID) == 0 {
			err = errorf("Incomplete or missing macro for OID resolution")
			break
		}

		r.DC = append(r.DC, def)
	}

	return
}

/*
processDSs returns an error instance following an attempt to parse all DITStructureRule
definitions found within the input at *DITStructureRuleDescriptionsContext.
*/
func (r *Schema) processDSs(x *DITStructureRuleDescriptionsContext) (err error) {
	_z := x.AllDITStructureRuleDescription()
	if len(r.DS) == 0 {
		r.DS = make(DITStructureRules, 0)
	}

	for j := 0; j < len(_z); j++ {
		var def DITStructureRule
		if err = def.process(_z[j]); err != nil {
			break
		}
		r.DS = append(r.DS, def)
	}

	return
}

/*
processOCs returns an error instance following an attempt to parse all ObjectClass
definitions found within the input at *ObjectClassDescriptionsContext.
*/
func (r *Schema) processOCs(x *ObjectClassDescriptionsContext) (err error) {
	_z := x.AllObjectClassDescription()
	if len(r.OC) == 0 {
		r.OC = make(ObjectClasses, 0)
	}

	for j := 0; j < len(_z); j++ {
		var def ObjectClass
		if err = def.process(_z[j]); err != nil {
			break
		}

		if len(def.Macro) == 2 && len(def.OID) == 0 {
			resl, found := r.OM[def.Macro[0]]
			if !found {
				err = errorf("Unresolvable macro for objectClass")
				break
			}
			def.OID = resl + `.` + def.Macro[1]
		} else if len(def.Macro) < 2 && len(def.OID) == 0 {
			err = errorf("Incomplete or missing macro for OID resolution")
			break
		}

		r.OC = append(r.OC, def)
	}

	return
}

/*
processNFs returns an error instance following an attempt to parse all NameForm
definitions found within the input at *NameFormDescriptionsContext.
*/
func (r *Schema) processNFs(x *NameFormDescriptionsContext) (err error) {
	_z := x.AllNameFormDescription()
	if len(r.NF) == 0 {
		r.NF = make(NameForms, 0)
	}

	for j := 0; j < len(_z); j++ {
		var def NameForm
		if err = def.process(_z[j]); err != nil {
			break
		}

		if len(def.Macro) == 2 && len(def.OID) == 0 {
			resl, found := r.OM[def.Macro[0]]
			if !found {
				err = errorf("Unresolvable macro for nameForm")
				break
			}
			def.OID = resl + `.` + def.Macro[1]
		} else if len(def.Macro) < 2 && len(def.OID) == 0 {
			err = errorf("Incomplete or missing macro for OID resolution")
			break
		}

		r.NF = append(r.NF, def)
	}

	return
}
