// Code generated from RFC4512.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr4512 // RFC4512
import "github.com/antlr4-go/antlr/v4"

// RFC4512Listener is a complete listener for a parse tree produced by RFC4512Parser.
type RFC4512Listener interface {
	antlr.ParseTreeListener

	// EnterFileparse is called when entering the fileparse production.
	EnterFileparse(c *FileparseContext)

	// EnterDefinitions is called when entering the definitions production.
	EnterDefinitions(c *DefinitionsContext)

	// EnterObjectClassDescriptions is called when entering the objectClassDescriptions production.
	EnterObjectClassDescriptions(c *ObjectClassDescriptionsContext)

	// EnterAttributeTypeDescriptions is called when entering the attributeTypeDescriptions production.
	EnterAttributeTypeDescriptions(c *AttributeTypeDescriptionsContext)

	// EnterMatchingRuleDescriptions is called when entering the matchingRuleDescriptions production.
	EnterMatchingRuleDescriptions(c *MatchingRuleDescriptionsContext)

	// EnterMatchingRuleUseDescriptions is called when entering the matchingRuleUseDescriptions production.
	EnterMatchingRuleUseDescriptions(c *MatchingRuleUseDescriptionsContext)

	// EnterLDAPSyntaxDescriptions is called when entering the lDAPSyntaxDescriptions production.
	EnterLDAPSyntaxDescriptions(c *LDAPSyntaxDescriptionsContext)

	// EnterDITStructureRuleDescriptions is called when entering the dITStructureRuleDescriptions production.
	EnterDITStructureRuleDescriptions(c *DITStructureRuleDescriptionsContext)

	// EnterNameFormDescriptions is called when entering the nameFormDescriptions production.
	EnterNameFormDescriptions(c *NameFormDescriptionsContext)

	// EnterDITContentRuleDescriptions is called when entering the dITContentRuleDescriptions production.
	EnterDITContentRuleDescriptions(c *DITContentRuleDescriptionsContext)

	// EnterObjectIdentifier is called when entering the objectIdentifier production.
	EnterObjectIdentifier(c *ObjectIdentifierContext)

	// EnterObjectClassLabel is called when entering the objectClassLabel production.
	EnterObjectClassLabel(c *ObjectClassLabelContext)

	// EnterMatchingRuleLabel is called when entering the matchingRuleLabel production.
	EnterMatchingRuleLabel(c *MatchingRuleLabelContext)

	// EnterMatchingRuleUseLabel is called when entering the matchingRuleUseLabel production.
	EnterMatchingRuleUseLabel(c *MatchingRuleUseLabelContext)

	// EnterLDAPSyntaxLabel is called when entering the lDAPSyntaxLabel production.
	EnterLDAPSyntaxLabel(c *LDAPSyntaxLabelContext)

	// EnterDITContentRuleLabel is called when entering the dITContentRuleLabel production.
	EnterDITContentRuleLabel(c *DITContentRuleLabelContext)

	// EnterDITStructureRuleLabel is called when entering the dITStructureRuleLabel production.
	EnterDITStructureRuleLabel(c *DITStructureRuleLabelContext)

	// EnterAttributeTypeLabel is called when entering the attributeTypeLabel production.
	EnterAttributeTypeLabel(c *AttributeTypeLabelContext)

	// EnterNameFormLabel is called when entering the nameFormLabel production.
	EnterNameFormLabel(c *NameFormLabelContext)

	// EnterDefinitionLabelDelim is called when entering the definitionLabelDelim production.
	EnterDefinitionLabelDelim(c *DefinitionLabelDelimContext)

	// EnterKind is called when entering the kind production.
	EnterKind(c *KindContext)

	// EnterAbstractKind is called when entering the abstractKind production.
	EnterAbstractKind(c *AbstractKindContext)

	// EnterStructuralKind is called when entering the structuralKind production.
	EnterStructuralKind(c *StructuralKindContext)

	// EnterAuxiliaryKind is called when entering the auxiliaryKind production.
	EnterAuxiliaryKind(c *AuxiliaryKindContext)

	// EnterApplies is called when entering the applies production.
	EnterApplies(c *AppliesContext)

	// EnterMust is called when entering the must production.
	EnterMust(c *MustContext)

	// EnterAux is called when entering the aux production.
	EnterAux(c *AuxContext)

	// EnterMay is called when entering the may production.
	EnterMay(c *MayContext)

	// EnterNot is called when entering the not production.
	EnterNot(c *NotContext)

	// EnterForm is called when entering the form production.
	EnterForm(c *FormContext)

	// EnterOrdering is called when entering the ordering production.
	EnterOrdering(c *OrderingContext)

	// EnterEquality is called when entering the equality production.
	EnterEquality(c *EqualityContext)

	// EnterStructuralClass is called when entering the structuralClass production.
	EnterStructuralClass(c *StructuralClassContext)

	// EnterSuperType is called when entering the superType production.
	EnterSuperType(c *SuperTypeContext)

	// EnterSuperClasses is called when entering the superClasses production.
	EnterSuperClasses(c *SuperClassesContext)

	// EnterSuperRules is called when entering the superRules production.
	EnterSuperRules(c *SuperRulesContext)

	// EnterSubstr is called when entering the substr production.
	EnterSubstr(c *SubstrContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterUsage is called when entering the usage production.
	EnterUsage(c *UsageContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterSingleValue is called when entering the singleValue production.
	EnterSingleValue(c *SingleValueContext)

	// EnterImmutability is called when entering the immutability production.
	EnterImmutability(c *ImmutabilityContext)

	// EnterCollective is called when entering the collective production.
	EnterCollective(c *CollectiveContext)

	// EnterObsolescence is called when entering the obsolescence production.
	EnterObsolescence(c *ObsolescenceContext)

	// EnterMinUpperBounds is called when entering the minUpperBounds production.
	EnterMinUpperBounds(c *MinUpperBoundsContext)

	// EnterStructureRule is called when entering the structureRule production.
	EnterStructureRule(c *StructureRuleContext)

	// EnterUsages is called when entering the usages production.
	EnterUsages(c *UsagesContext)

	// EnterQuotedDescriptor is called when entering the quotedDescriptor production.
	EnterQuotedDescriptor(c *QuotedDescriptorContext)

	// EnterQuotedString is called when entering the quotedString production.
	EnterQuotedString(c *QuotedStringContext)

	// EnterSchemaDN is called when entering the schemaDN production.
	EnterSchemaDN(c *SchemaDNContext)

	// EnterObjectIdentifierOIDOrName is called when entering the objectIdentifierOIDOrName production.
	EnterObjectIdentifierOIDOrName(c *ObjectIdentifierOIDOrNameContext)

	// EnterAttributeDescription is called when entering the attributeDescription production.
	EnterAttributeDescription(c *AttributeDescriptionContext)

	// EnterOpenParen is called when entering the openParen production.
	EnterOpenParen(c *OpenParenContext)

	// EnterCloseParen is called when entering the closeParen production.
	EnterCloseParen(c *CloseParenContext)

	// EnterNumericOIDOrMacro is called when entering the numericOIDOrMacro production.
	EnterNumericOIDOrMacro(c *NumericOIDOrMacroContext)

	// EnterMacro is called when entering the macro production.
	EnterMacro(c *MacroContext)

	// EnterNumericOID is called when entering the numericOID production.
	EnterNumericOID(c *NumericOIDContext)

	// EnterMacroSuffix is called when entering the macroSuffix production.
	EnterMacroSuffix(c *MacroSuffixContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterAttrOptions is called when entering the attrOptions production.
	EnterAttrOptions(c *AttrOptionsContext)

	// EnterUserApplications is called when entering the userApplications production.
	EnterUserApplications(c *UserApplicationsContext)

	// EnterDirectoryOperation is called when entering the directoryOperation production.
	EnterDirectoryOperation(c *DirectoryOperationContext)

	// EnterDistributedOperation is called when entering the distributedOperation production.
	EnterDistributedOperation(c *DistributedOperationContext)

	// EnterDSAOperation is called when entering the dSAOperation production.
	EnterDSAOperation(c *DSAOperationContext)

	// EnterStructureRules is called when entering the structureRules production.
	EnterStructureRules(c *StructureRulesContext)

	// EnterOID is called when entering the oID production.
	EnterOID(c *OIDContext)

	// EnterOIDs is called when entering the oIDs production.
	EnterOIDs(c *OIDsContext)

	// EnterExtensionValue is called when entering the extensionValue production.
	EnterExtensionValue(c *ExtensionValueContext)

	// EnterExtensions is called when entering the extensions production.
	EnterExtensions(c *ExtensionsContext)

	// EnterExtension is called when entering the extension production.
	EnterExtension(c *ExtensionContext)

	// EnterNames is called when entering the names production.
	EnterNames(c *NamesContext)

	// EnterDescriptor is called when entering the descriptor production.
	EnterDescriptor(c *DescriptorContext)

	// EnterObjectClassDescription is called when entering the objectClassDescription production.
	EnterObjectClassDescription(c *ObjectClassDescriptionContext)

	// EnterAttributeTypeDescription is called when entering the attributeTypeDescription production.
	EnterAttributeTypeDescription(c *AttributeTypeDescriptionContext)

	// EnterMatchingRuleDescription is called when entering the matchingRuleDescription production.
	EnterMatchingRuleDescription(c *MatchingRuleDescriptionContext)

	// EnterMatchingRuleUseDescription is called when entering the matchingRuleUseDescription production.
	EnterMatchingRuleUseDescription(c *MatchingRuleUseDescriptionContext)

	// EnterLDAPSyntaxDescription is called when entering the lDAPSyntaxDescription production.
	EnterLDAPSyntaxDescription(c *LDAPSyntaxDescriptionContext)

	// EnterDITContentRuleDescription is called when entering the dITContentRuleDescription production.
	EnterDITContentRuleDescription(c *DITContentRuleDescriptionContext)

	// EnterDITStructureRuleDescription is called when entering the dITStructureRuleDescription production.
	EnterDITStructureRuleDescription(c *DITStructureRuleDescriptionContext)

	// EnterNameFormDescription is called when entering the nameFormDescription production.
	EnterNameFormDescription(c *NameFormDescriptionContext)

	// ExitFileparse is called when exiting the fileparse production.
	ExitFileparse(c *FileparseContext)

	// ExitDefinitions is called when exiting the definitions production.
	ExitDefinitions(c *DefinitionsContext)

	// ExitObjectClassDescriptions is called when exiting the objectClassDescriptions production.
	ExitObjectClassDescriptions(c *ObjectClassDescriptionsContext)

	// ExitAttributeTypeDescriptions is called when exiting the attributeTypeDescriptions production.
	ExitAttributeTypeDescriptions(c *AttributeTypeDescriptionsContext)

	// ExitMatchingRuleDescriptions is called when exiting the matchingRuleDescriptions production.
	ExitMatchingRuleDescriptions(c *MatchingRuleDescriptionsContext)

	// ExitMatchingRuleUseDescriptions is called when exiting the matchingRuleUseDescriptions production.
	ExitMatchingRuleUseDescriptions(c *MatchingRuleUseDescriptionsContext)

	// ExitLDAPSyntaxDescriptions is called when exiting the lDAPSyntaxDescriptions production.
	ExitLDAPSyntaxDescriptions(c *LDAPSyntaxDescriptionsContext)

	// ExitDITStructureRuleDescriptions is called when exiting the dITStructureRuleDescriptions production.
	ExitDITStructureRuleDescriptions(c *DITStructureRuleDescriptionsContext)

	// ExitNameFormDescriptions is called when exiting the nameFormDescriptions production.
	ExitNameFormDescriptions(c *NameFormDescriptionsContext)

	// ExitDITContentRuleDescriptions is called when exiting the dITContentRuleDescriptions production.
	ExitDITContentRuleDescriptions(c *DITContentRuleDescriptionsContext)

	// ExitObjectIdentifier is called when exiting the objectIdentifier production.
	ExitObjectIdentifier(c *ObjectIdentifierContext)

	// ExitObjectClassLabel is called when exiting the objectClassLabel production.
	ExitObjectClassLabel(c *ObjectClassLabelContext)

	// ExitMatchingRuleLabel is called when exiting the matchingRuleLabel production.
	ExitMatchingRuleLabel(c *MatchingRuleLabelContext)

	// ExitMatchingRuleUseLabel is called when exiting the matchingRuleUseLabel production.
	ExitMatchingRuleUseLabel(c *MatchingRuleUseLabelContext)

	// ExitLDAPSyntaxLabel is called when exiting the lDAPSyntaxLabel production.
	ExitLDAPSyntaxLabel(c *LDAPSyntaxLabelContext)

	// ExitDITContentRuleLabel is called when exiting the dITContentRuleLabel production.
	ExitDITContentRuleLabel(c *DITContentRuleLabelContext)

	// ExitDITStructureRuleLabel is called when exiting the dITStructureRuleLabel production.
	ExitDITStructureRuleLabel(c *DITStructureRuleLabelContext)

	// ExitAttributeTypeLabel is called when exiting the attributeTypeLabel production.
	ExitAttributeTypeLabel(c *AttributeTypeLabelContext)

	// ExitNameFormLabel is called when exiting the nameFormLabel production.
	ExitNameFormLabel(c *NameFormLabelContext)

	// ExitDefinitionLabelDelim is called when exiting the definitionLabelDelim production.
	ExitDefinitionLabelDelim(c *DefinitionLabelDelimContext)

	// ExitKind is called when exiting the kind production.
	ExitKind(c *KindContext)

	// ExitAbstractKind is called when exiting the abstractKind production.
	ExitAbstractKind(c *AbstractKindContext)

	// ExitStructuralKind is called when exiting the structuralKind production.
	ExitStructuralKind(c *StructuralKindContext)

	// ExitAuxiliaryKind is called when exiting the auxiliaryKind production.
	ExitAuxiliaryKind(c *AuxiliaryKindContext)

	// ExitApplies is called when exiting the applies production.
	ExitApplies(c *AppliesContext)

	// ExitMust is called when exiting the must production.
	ExitMust(c *MustContext)

	// ExitAux is called when exiting the aux production.
	ExitAux(c *AuxContext)

	// ExitMay is called when exiting the may production.
	ExitMay(c *MayContext)

	// ExitNot is called when exiting the not production.
	ExitNot(c *NotContext)

	// ExitForm is called when exiting the form production.
	ExitForm(c *FormContext)

	// ExitOrdering is called when exiting the ordering production.
	ExitOrdering(c *OrderingContext)

	// ExitEquality is called when exiting the equality production.
	ExitEquality(c *EqualityContext)

	// ExitStructuralClass is called when exiting the structuralClass production.
	ExitStructuralClass(c *StructuralClassContext)

	// ExitSuperType is called when exiting the superType production.
	ExitSuperType(c *SuperTypeContext)

	// ExitSuperClasses is called when exiting the superClasses production.
	ExitSuperClasses(c *SuperClassesContext)

	// ExitSuperRules is called when exiting the superRules production.
	ExitSuperRules(c *SuperRulesContext)

	// ExitSubstr is called when exiting the substr production.
	ExitSubstr(c *SubstrContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitUsage is called when exiting the usage production.
	ExitUsage(c *UsageContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitSingleValue is called when exiting the singleValue production.
	ExitSingleValue(c *SingleValueContext)

	// ExitImmutability is called when exiting the immutability production.
	ExitImmutability(c *ImmutabilityContext)

	// ExitCollective is called when exiting the collective production.
	ExitCollective(c *CollectiveContext)

	// ExitObsolescence is called when exiting the obsolescence production.
	ExitObsolescence(c *ObsolescenceContext)

	// ExitMinUpperBounds is called when exiting the minUpperBounds production.
	ExitMinUpperBounds(c *MinUpperBoundsContext)

	// ExitStructureRule is called when exiting the structureRule production.
	ExitStructureRule(c *StructureRuleContext)

	// ExitUsages is called when exiting the usages production.
	ExitUsages(c *UsagesContext)

	// ExitQuotedDescriptor is called when exiting the quotedDescriptor production.
	ExitQuotedDescriptor(c *QuotedDescriptorContext)

	// ExitQuotedString is called when exiting the quotedString production.
	ExitQuotedString(c *QuotedStringContext)

	// ExitSchemaDN is called when exiting the schemaDN production.
	ExitSchemaDN(c *SchemaDNContext)

	// ExitObjectIdentifierOIDOrName is called when exiting the objectIdentifierOIDOrName production.
	ExitObjectIdentifierOIDOrName(c *ObjectIdentifierOIDOrNameContext)

	// ExitAttributeDescription is called when exiting the attributeDescription production.
	ExitAttributeDescription(c *AttributeDescriptionContext)

	// ExitOpenParen is called when exiting the openParen production.
	ExitOpenParen(c *OpenParenContext)

	// ExitCloseParen is called when exiting the closeParen production.
	ExitCloseParen(c *CloseParenContext)

	// ExitNumericOIDOrMacro is called when exiting the numericOIDOrMacro production.
	ExitNumericOIDOrMacro(c *NumericOIDOrMacroContext)

	// ExitMacro is called when exiting the macro production.
	ExitMacro(c *MacroContext)

	// ExitNumericOID is called when exiting the numericOID production.
	ExitNumericOID(c *NumericOIDContext)

	// ExitMacroSuffix is called when exiting the macroSuffix production.
	ExitMacroSuffix(c *MacroSuffixContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitAttrOptions is called when exiting the attrOptions production.
	ExitAttrOptions(c *AttrOptionsContext)

	// ExitUserApplications is called when exiting the userApplications production.
	ExitUserApplications(c *UserApplicationsContext)

	// ExitDirectoryOperation is called when exiting the directoryOperation production.
	ExitDirectoryOperation(c *DirectoryOperationContext)

	// ExitDistributedOperation is called when exiting the distributedOperation production.
	ExitDistributedOperation(c *DistributedOperationContext)

	// ExitDSAOperation is called when exiting the dSAOperation production.
	ExitDSAOperation(c *DSAOperationContext)

	// ExitStructureRules is called when exiting the structureRules production.
	ExitStructureRules(c *StructureRulesContext)

	// ExitOID is called when exiting the oID production.
	ExitOID(c *OIDContext)

	// ExitOIDs is called when exiting the oIDs production.
	ExitOIDs(c *OIDsContext)

	// ExitExtensionValue is called when exiting the extensionValue production.
	ExitExtensionValue(c *ExtensionValueContext)

	// ExitExtensions is called when exiting the extensions production.
	ExitExtensions(c *ExtensionsContext)

	// ExitExtension is called when exiting the extension production.
	ExitExtension(c *ExtensionContext)

	// ExitNames is called when exiting the names production.
	ExitNames(c *NamesContext)

	// ExitDescriptor is called when exiting the descriptor production.
	ExitDescriptor(c *DescriptorContext)

	// ExitObjectClassDescription is called when exiting the objectClassDescription production.
	ExitObjectClassDescription(c *ObjectClassDescriptionContext)

	// ExitAttributeTypeDescription is called when exiting the attributeTypeDescription production.
	ExitAttributeTypeDescription(c *AttributeTypeDescriptionContext)

	// ExitMatchingRuleDescription is called when exiting the matchingRuleDescription production.
	ExitMatchingRuleDescription(c *MatchingRuleDescriptionContext)

	// ExitMatchingRuleUseDescription is called when exiting the matchingRuleUseDescription production.
	ExitMatchingRuleUseDescription(c *MatchingRuleUseDescriptionContext)

	// ExitLDAPSyntaxDescription is called when exiting the lDAPSyntaxDescription production.
	ExitLDAPSyntaxDescription(c *LDAPSyntaxDescriptionContext)

	// ExitDITContentRuleDescription is called when exiting the dITContentRuleDescription production.
	ExitDITContentRuleDescription(c *DITContentRuleDescriptionContext)

	// ExitDITStructureRuleDescription is called when exiting the dITStructureRuleDescription production.
	ExitDITStructureRuleDescription(c *DITStructureRuleDescriptionContext)

	// ExitNameFormDescription is called when exiting the nameFormDescription production.
	ExitNameFormDescription(c *NameFormDescriptionContext)
}
