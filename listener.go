// Code generated from RFC4512.g4 by ANTLR 4.13.1. DO NOT EDIT.

package rfc4512antlr // RFC4512
import "github.com/antlr4-go/antlr/v4"

// RFC4512Listener is a complete listener for a parse tree produced by RFC4512Parser.
type RFC4512Listener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterDefinitions is called when entering the definitions production.
	EnterDefinitions(c *DefinitionsContext)

	// EnterObjectClassDescriptions is called when entering the objectClassDescriptions production.
	EnterObjectClassDescriptions(c *ObjectClassDescriptionsContext)

	// EnterObjectClassDescription is called when entering the objectClassDescription production.
	EnterObjectClassDescription(c *ObjectClassDescriptionContext)

	// EnterOCKind is called when entering the oCKind production.
	EnterOCKind(c *OCKindContext)

	// EnterAbstractKind is called when entering the abstractKind production.
	EnterAbstractKind(c *AbstractKindContext)

	// EnterStructuralKind is called when entering the structuralKind production.
	EnterStructuralKind(c *StructuralKindContext)

	// EnterAuxiliaryKind is called when entering the auxiliaryKind production.
	EnterAuxiliaryKind(c *AuxiliaryKindContext)

	// EnterOCSuperClasses is called when entering the oCSuperClasses production.
	EnterOCSuperClasses(c *OCSuperClassesContext)

	// EnterAttributeTypeDescriptions is called when entering the attributeTypeDescriptions production.
	EnterAttributeTypeDescriptions(c *AttributeTypeDescriptionsContext)

	// EnterAttributeTypeDescription is called when entering the attributeTypeDescription production.
	EnterAttributeTypeDescription(c *AttributeTypeDescriptionContext)

	// EnterMinimumUpperBounds is called when entering the minimumUpperBounds production.
	EnterMinimumUpperBounds(c *MinimumUpperBoundsContext)

	// EnterATEquality is called when entering the aTEquality production.
	EnterATEquality(c *ATEqualityContext)

	// EnterATOrdering is called when entering the aTOrdering production.
	EnterATOrdering(c *ATOrderingContext)

	// EnterATSubstring is called when entering the aTSubstring production.
	EnterATSubstring(c *ATSubstringContext)

	// EnterATSuperType is called when entering the aTSuperType production.
	EnterATSuperType(c *ATSuperTypeContext)

	// EnterATUsage is called when entering the aTUsage production.
	EnterATUsage(c *ATUsageContext)

	// EnterATSingleValue is called when entering the aTSingleValue production.
	EnterATSingleValue(c *ATSingleValueContext)

	// EnterATCollective is called when entering the aTCollective production.
	EnterATCollective(c *ATCollectiveContext)

	// EnterATNoUserModification is called when entering the aTNoUserModification production.
	EnterATNoUserModification(c *ATNoUserModificationContext)

	// EnterMatchingRuleDescriptions is called when entering the matchingRuleDescriptions production.
	EnterMatchingRuleDescriptions(c *MatchingRuleDescriptionsContext)

	// EnterMatchingRuleDescription is called when entering the matchingRuleDescription production.
	EnterMatchingRuleDescription(c *MatchingRuleDescriptionContext)

	// EnterMatchingRuleUseDescriptions is called when entering the matchingRuleUseDescriptions production.
	EnterMatchingRuleUseDescriptions(c *MatchingRuleUseDescriptionsContext)

	// EnterMatchingRuleUseDescription is called when entering the matchingRuleUseDescription production.
	EnterMatchingRuleUseDescription(c *MatchingRuleUseDescriptionContext)

	// EnterMRUApplies is called when entering the mRUApplies production.
	EnterMRUApplies(c *MRUAppliesContext)

	// EnterLDAPSyntaxDescriptions is called when entering the lDAPSyntaxDescriptions production.
	EnterLDAPSyntaxDescriptions(c *LDAPSyntaxDescriptionsContext)

	// EnterLDAPSyntaxDescription is called when entering the lDAPSyntaxDescription production.
	EnterLDAPSyntaxDescription(c *LDAPSyntaxDescriptionContext)

	// EnterDITContentRuleDescriptions is called when entering the dITContentRuleDescriptions production.
	EnterDITContentRuleDescriptions(c *DITContentRuleDescriptionsContext)

	// EnterDITContentRuleDescription is called when entering the dITContentRuleDescription production.
	EnterDITContentRuleDescription(c *DITContentRuleDescriptionContext)

	// EnterDCRNot is called when entering the dCRNot production.
	EnterDCRNot(c *DCRNotContext)

	// EnterDCRAux is called when entering the dCRAux production.
	EnterDCRAux(c *DCRAuxContext)

	// EnterDITStructureRuleDescriptions is called when entering the dITStructureRuleDescriptions production.
	EnterDITStructureRuleDescriptions(c *DITStructureRuleDescriptionsContext)

	// EnterDITStructureRuleDescription is called when entering the dITStructureRuleDescription production.
	EnterDITStructureRuleDescription(c *DITStructureRuleDescriptionContext)

	// EnterStructureRules is called when entering the structureRules production.
	EnterStructureRules(c *StructureRulesContext)

	// EnterDSRSuperRules is called when entering the dSRSuperRules production.
	EnterDSRSuperRules(c *DSRSuperRulesContext)

	// EnterStructureRule is called when entering the structureRule production.
	EnterStructureRule(c *StructureRuleContext)

	// EnterDSRForm is called when entering the dSRForm production.
	EnterDSRForm(c *DSRFormContext)

	// EnterNameFormDescriptions is called when entering the nameFormDescriptions production.
	EnterNameFormDescriptions(c *NameFormDescriptionsContext)

	// EnterNameFormDescription is called when entering the nameFormDescription production.
	EnterNameFormDescription(c *NameFormDescriptionContext)

	// EnterNFStructuralOC is called when entering the nFStructuralOC production.
	EnterNFStructuralOC(c *NFStructuralOCContext)

	// EnterDefinitionObsolete is called when entering the definitionObsolete production.
	EnterDefinitionObsolete(c *DefinitionObsoleteContext)

	// EnterDefinitionSyntax is called when entering the definitionSyntax production.
	EnterDefinitionSyntax(c *DefinitionSyntaxContext)

	// EnterDefinitionMust is called when entering the definitionMust production.
	EnterDefinitionMust(c *DefinitionMustContext)

	// EnterDefinitionMay is called when entering the definitionMay production.
	EnterDefinitionMay(c *DefinitionMayContext)

	// EnterDefinitionDescription is called when entering the definitionDescription production.
	EnterDefinitionDescription(c *DefinitionDescriptionContext)

	// EnterDefinitionName is called when entering the definitionName production.
	EnterDefinitionName(c *DefinitionNameContext)

	// EnterDefinitionExtensions is called when entering the definitionExtensions production.
	EnterDefinitionExtensions(c *DefinitionExtensionsContext)

	// EnterDefinitionExtension is called when entering the definitionExtension production.
	EnterDefinitionExtension(c *DefinitionExtensionContext)

	// EnterOIDs is called when entering the oIDs production.
	EnterOIDs(c *OIDsContext)

	// EnterAttributeDescription is called when entering the attributeDescription production.
	EnterAttributeDescription(c *AttributeDescriptionContext)

	// EnterOID is called when entering the oID production.
	EnterOID(c *OIDContext)

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

	// EnterDescriptor is called when entering the descriptor production.
	EnterDescriptor(c *DescriptorContext)

	// EnterExtensionValue is called when entering the extensionValue production.
	EnterExtensionValue(c *ExtensionValueContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitDefinitions is called when exiting the definitions production.
	ExitDefinitions(c *DefinitionsContext)

	// ExitObjectClassDescriptions is called when exiting the objectClassDescriptions production.
	ExitObjectClassDescriptions(c *ObjectClassDescriptionsContext)

	// ExitObjectClassDescription is called when exiting the objectClassDescription production.
	ExitObjectClassDescription(c *ObjectClassDescriptionContext)

	// ExitOCKind is called when exiting the oCKind production.
	ExitOCKind(c *OCKindContext)

	// ExitAbstractKind is called when exiting the abstractKind production.
	ExitAbstractKind(c *AbstractKindContext)

	// ExitStructuralKind is called when exiting the structuralKind production.
	ExitStructuralKind(c *StructuralKindContext)

	// ExitAuxiliaryKind is called when exiting the auxiliaryKind production.
	ExitAuxiliaryKind(c *AuxiliaryKindContext)

	// ExitOCSuperClasses is called when exiting the oCSuperClasses production.
	ExitOCSuperClasses(c *OCSuperClassesContext)

	// ExitAttributeTypeDescriptions is called when exiting the attributeTypeDescriptions production.
	ExitAttributeTypeDescriptions(c *AttributeTypeDescriptionsContext)

	// ExitAttributeTypeDescription is called when exiting the attributeTypeDescription production.
	ExitAttributeTypeDescription(c *AttributeTypeDescriptionContext)

	// ExitMinimumUpperBounds is called when exiting the minimumUpperBounds production.
	ExitMinimumUpperBounds(c *MinimumUpperBoundsContext)

	// ExitATEquality is called when exiting the aTEquality production.
	ExitATEquality(c *ATEqualityContext)

	// ExitATOrdering is called when exiting the aTOrdering production.
	ExitATOrdering(c *ATOrderingContext)

	// ExitATSubstring is called when exiting the aTSubstring production.
	ExitATSubstring(c *ATSubstringContext)

	// ExitATSuperType is called when exiting the aTSuperType production.
	ExitATSuperType(c *ATSuperTypeContext)

	// ExitATUsage is called when exiting the aTUsage production.
	ExitATUsage(c *ATUsageContext)

	// ExitATSingleValue is called when exiting the aTSingleValue production.
	ExitATSingleValue(c *ATSingleValueContext)

	// ExitATCollective is called when exiting the aTCollective production.
	ExitATCollective(c *ATCollectiveContext)

	// ExitATNoUserModification is called when exiting the aTNoUserModification production.
	ExitATNoUserModification(c *ATNoUserModificationContext)

	// ExitMatchingRuleDescriptions is called when exiting the matchingRuleDescriptions production.
	ExitMatchingRuleDescriptions(c *MatchingRuleDescriptionsContext)

	// ExitMatchingRuleDescription is called when exiting the matchingRuleDescription production.
	ExitMatchingRuleDescription(c *MatchingRuleDescriptionContext)

	// ExitMatchingRuleUseDescriptions is called when exiting the matchingRuleUseDescriptions production.
	ExitMatchingRuleUseDescriptions(c *MatchingRuleUseDescriptionsContext)

	// ExitMatchingRuleUseDescription is called when exiting the matchingRuleUseDescription production.
	ExitMatchingRuleUseDescription(c *MatchingRuleUseDescriptionContext)

	// ExitMRUApplies is called when exiting the mRUApplies production.
	ExitMRUApplies(c *MRUAppliesContext)

	// ExitLDAPSyntaxDescriptions is called when exiting the lDAPSyntaxDescriptions production.
	ExitLDAPSyntaxDescriptions(c *LDAPSyntaxDescriptionsContext)

	// ExitLDAPSyntaxDescription is called when exiting the lDAPSyntaxDescription production.
	ExitLDAPSyntaxDescription(c *LDAPSyntaxDescriptionContext)

	// ExitDITContentRuleDescriptions is called when exiting the dITContentRuleDescriptions production.
	ExitDITContentRuleDescriptions(c *DITContentRuleDescriptionsContext)

	// ExitDITContentRuleDescription is called when exiting the dITContentRuleDescription production.
	ExitDITContentRuleDescription(c *DITContentRuleDescriptionContext)

	// ExitDCRNot is called when exiting the dCRNot production.
	ExitDCRNot(c *DCRNotContext)

	// ExitDCRAux is called when exiting the dCRAux production.
	ExitDCRAux(c *DCRAuxContext)

	// ExitDITStructureRuleDescriptions is called when exiting the dITStructureRuleDescriptions production.
	ExitDITStructureRuleDescriptions(c *DITStructureRuleDescriptionsContext)

	// ExitDITStructureRuleDescription is called when exiting the dITStructureRuleDescription production.
	ExitDITStructureRuleDescription(c *DITStructureRuleDescriptionContext)

	// ExitStructureRules is called when exiting the structureRules production.
	ExitStructureRules(c *StructureRulesContext)

	// ExitDSRSuperRules is called when exiting the dSRSuperRules production.
	ExitDSRSuperRules(c *DSRSuperRulesContext)

	// ExitStructureRule is called when exiting the structureRule production.
	ExitStructureRule(c *StructureRuleContext)

	// ExitDSRForm is called when exiting the dSRForm production.
	ExitDSRForm(c *DSRFormContext)

	// ExitNameFormDescriptions is called when exiting the nameFormDescriptions production.
	ExitNameFormDescriptions(c *NameFormDescriptionsContext)

	// ExitNameFormDescription is called when exiting the nameFormDescription production.
	ExitNameFormDescription(c *NameFormDescriptionContext)

	// ExitNFStructuralOC is called when exiting the nFStructuralOC production.
	ExitNFStructuralOC(c *NFStructuralOCContext)

	// ExitDefinitionObsolete is called when exiting the definitionObsolete production.
	ExitDefinitionObsolete(c *DefinitionObsoleteContext)

	// ExitDefinitionSyntax is called when exiting the definitionSyntax production.
	ExitDefinitionSyntax(c *DefinitionSyntaxContext)

	// ExitDefinitionMust is called when exiting the definitionMust production.
	ExitDefinitionMust(c *DefinitionMustContext)

	// ExitDefinitionMay is called when exiting the definitionMay production.
	ExitDefinitionMay(c *DefinitionMayContext)

	// ExitDefinitionDescription is called when exiting the definitionDescription production.
	ExitDefinitionDescription(c *DefinitionDescriptionContext)

	// ExitDefinitionName is called when exiting the definitionName production.
	ExitDefinitionName(c *DefinitionNameContext)

	// ExitDefinitionExtensions is called when exiting the definitionExtensions production.
	ExitDefinitionExtensions(c *DefinitionExtensionsContext)

	// ExitDefinitionExtension is called when exiting the definitionExtension production.
	ExitDefinitionExtension(c *DefinitionExtensionContext)

	// ExitOIDs is called when exiting the oIDs production.
	ExitOIDs(c *OIDsContext)

	// ExitAttributeDescription is called when exiting the attributeDescription production.
	ExitAttributeDescription(c *AttributeDescriptionContext)

	// ExitOID is called when exiting the oID production.
	ExitOID(c *OIDContext)

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

	// ExitDescriptor is called when exiting the descriptor production.
	ExitDescriptor(c *DescriptorContext)

	// ExitExtensionValue is called when exiting the extensionValue production.
	ExitExtensionValue(c *ExtensionValueContext)
}
