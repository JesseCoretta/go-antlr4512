// Code generated from RFC4512.g4 by ANTLR 4.13.1. DO NOT EDIT.

package rfc4512antlr // RFC4512
import "github.com/antlr4-go/antlr/v4"

// BaseRFC4512Listener is a complete listener for a parse tree produced by RFC4512Parser.
type BaseRFC4512Listener struct{}

var _ RFC4512Listener = &BaseRFC4512Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRFC4512Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRFC4512Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRFC4512Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRFC4512Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseRFC4512Listener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseRFC4512Listener) ExitParse(ctx *ParseContext) {}

// EnterDefinitions is called when production definitions is entered.
func (s *BaseRFC4512Listener) EnterDefinitions(ctx *DefinitionsContext) {}

// ExitDefinitions is called when production definitions is exited.
func (s *BaseRFC4512Listener) ExitDefinitions(ctx *DefinitionsContext) {}

// EnterObjectClassDescriptions is called when production objectClassDescriptions is entered.
func (s *BaseRFC4512Listener) EnterObjectClassDescriptions(ctx *ObjectClassDescriptionsContext) {}

// ExitObjectClassDescriptions is called when production objectClassDescriptions is exited.
func (s *BaseRFC4512Listener) ExitObjectClassDescriptions(ctx *ObjectClassDescriptionsContext) {}

// EnterObjectClassLabel is called when production objectClassLabel is entered.
func (s *BaseRFC4512Listener) EnterObjectClassLabel(ctx *ObjectClassLabelContext) {}

// ExitObjectClassLabel is called when production objectClassLabel is exited.
func (s *BaseRFC4512Listener) ExitObjectClassLabel(ctx *ObjectClassLabelContext) {}

// EnterObjectClassDescription is called when production objectClassDescription is entered.
func (s *BaseRFC4512Listener) EnterObjectClassDescription(ctx *ObjectClassDescriptionContext) {}

// ExitObjectClassDescription is called when production objectClassDescription is exited.
func (s *BaseRFC4512Listener) ExitObjectClassDescription(ctx *ObjectClassDescriptionContext) {}

// EnterOCKind is called when production oCKind is entered.
func (s *BaseRFC4512Listener) EnterOCKind(ctx *OCKindContext) {}

// ExitOCKind is called when production oCKind is exited.
func (s *BaseRFC4512Listener) ExitOCKind(ctx *OCKindContext) {}

// EnterAbstractKind is called when production abstractKind is entered.
func (s *BaseRFC4512Listener) EnterAbstractKind(ctx *AbstractKindContext) {}

// ExitAbstractKind is called when production abstractKind is exited.
func (s *BaseRFC4512Listener) ExitAbstractKind(ctx *AbstractKindContext) {}

// EnterStructuralKind is called when production structuralKind is entered.
func (s *BaseRFC4512Listener) EnterStructuralKind(ctx *StructuralKindContext) {}

// ExitStructuralKind is called when production structuralKind is exited.
func (s *BaseRFC4512Listener) ExitStructuralKind(ctx *StructuralKindContext) {}

// EnterAuxiliaryKind is called when production auxiliaryKind is entered.
func (s *BaseRFC4512Listener) EnterAuxiliaryKind(ctx *AuxiliaryKindContext) {}

// ExitAuxiliaryKind is called when production auxiliaryKind is exited.
func (s *BaseRFC4512Listener) ExitAuxiliaryKind(ctx *AuxiliaryKindContext) {}

// EnterOCSuperClasses is called when production oCSuperClasses is entered.
func (s *BaseRFC4512Listener) EnterOCSuperClasses(ctx *OCSuperClassesContext) {}

// ExitOCSuperClasses is called when production oCSuperClasses is exited.
func (s *BaseRFC4512Listener) ExitOCSuperClasses(ctx *OCSuperClassesContext) {}

// EnterAttributeTypeDescriptions is called when production attributeTypeDescriptions is entered.
func (s *BaseRFC4512Listener) EnterAttributeTypeDescriptions(ctx *AttributeTypeDescriptionsContext) {}

// ExitAttributeTypeDescriptions is called when production attributeTypeDescriptions is exited.
func (s *BaseRFC4512Listener) ExitAttributeTypeDescriptions(ctx *AttributeTypeDescriptionsContext) {}

// EnterAttributeTypeLabel is called when production attributeTypeLabel is entered.
func (s *BaseRFC4512Listener) EnterAttributeTypeLabel(ctx *AttributeTypeLabelContext) {}

// ExitAttributeTypeLabel is called when production attributeTypeLabel is exited.
func (s *BaseRFC4512Listener) ExitAttributeTypeLabel(ctx *AttributeTypeLabelContext) {}

// EnterAttributeTypeDescription is called when production attributeTypeDescription is entered.
func (s *BaseRFC4512Listener) EnterAttributeTypeDescription(ctx *AttributeTypeDescriptionContext) {}

// ExitAttributeTypeDescription is called when production attributeTypeDescription is exited.
func (s *BaseRFC4512Listener) ExitAttributeTypeDescription(ctx *AttributeTypeDescriptionContext) {}

// EnterMinimumUpperBounds is called when production minimumUpperBounds is entered.
func (s *BaseRFC4512Listener) EnterMinimumUpperBounds(ctx *MinimumUpperBoundsContext) {}

// ExitMinimumUpperBounds is called when production minimumUpperBounds is exited.
func (s *BaseRFC4512Listener) ExitMinimumUpperBounds(ctx *MinimumUpperBoundsContext) {}

// EnterATEquality is called when production aTEquality is entered.
func (s *BaseRFC4512Listener) EnterATEquality(ctx *ATEqualityContext) {}

// ExitATEquality is called when production aTEquality is exited.
func (s *BaseRFC4512Listener) ExitATEquality(ctx *ATEqualityContext) {}

// EnterATOrdering is called when production aTOrdering is entered.
func (s *BaseRFC4512Listener) EnterATOrdering(ctx *ATOrderingContext) {}

// ExitATOrdering is called when production aTOrdering is exited.
func (s *BaseRFC4512Listener) ExitATOrdering(ctx *ATOrderingContext) {}

// EnterATSubstring is called when production aTSubstring is entered.
func (s *BaseRFC4512Listener) EnterATSubstring(ctx *ATSubstringContext) {}

// ExitATSubstring is called when production aTSubstring is exited.
func (s *BaseRFC4512Listener) ExitATSubstring(ctx *ATSubstringContext) {}

// EnterATSuperType is called when production aTSuperType is entered.
func (s *BaseRFC4512Listener) EnterATSuperType(ctx *ATSuperTypeContext) {}

// ExitATSuperType is called when production aTSuperType is exited.
func (s *BaseRFC4512Listener) ExitATSuperType(ctx *ATSuperTypeContext) {}

// EnterATUsage is called when production aTUsage is entered.
func (s *BaseRFC4512Listener) EnterATUsage(ctx *ATUsageContext) {}

// ExitATUsage is called when production aTUsage is exited.
func (s *BaseRFC4512Listener) ExitATUsage(ctx *ATUsageContext) {}

// EnterATSingleValue is called when production aTSingleValue is entered.
func (s *BaseRFC4512Listener) EnterATSingleValue(ctx *ATSingleValueContext) {}

// ExitATSingleValue is called when production aTSingleValue is exited.
func (s *BaseRFC4512Listener) ExitATSingleValue(ctx *ATSingleValueContext) {}

// EnterATCollective is called when production aTCollective is entered.
func (s *BaseRFC4512Listener) EnterATCollective(ctx *ATCollectiveContext) {}

// ExitATCollective is called when production aTCollective is exited.
func (s *BaseRFC4512Listener) ExitATCollective(ctx *ATCollectiveContext) {}

// EnterATNoUserModification is called when production aTNoUserModification is entered.
func (s *BaseRFC4512Listener) EnterATNoUserModification(ctx *ATNoUserModificationContext) {}

// ExitATNoUserModification is called when production aTNoUserModification is exited.
func (s *BaseRFC4512Listener) ExitATNoUserModification(ctx *ATNoUserModificationContext) {}

// EnterMatchingRuleDescriptions is called when production matchingRuleDescriptions is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleDescriptions(ctx *MatchingRuleDescriptionsContext) {}

// ExitMatchingRuleDescriptions is called when production matchingRuleDescriptions is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleDescriptions(ctx *MatchingRuleDescriptionsContext) {}

// EnterMatchingRuleLabel is called when production matchingRuleLabel is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleLabel(ctx *MatchingRuleLabelContext) {}

// ExitMatchingRuleLabel is called when production matchingRuleLabel is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleLabel(ctx *MatchingRuleLabelContext) {}

// EnterMatchingRuleDescription is called when production matchingRuleDescription is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleDescription(ctx *MatchingRuleDescriptionContext) {}

// ExitMatchingRuleDescription is called when production matchingRuleDescription is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleDescription(ctx *MatchingRuleDescriptionContext) {}

// EnterMatchingRuleUseDescriptions is called when production matchingRuleUseDescriptions is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleUseDescriptions(ctx *MatchingRuleUseDescriptionsContext) {
}

// ExitMatchingRuleUseDescriptions is called when production matchingRuleUseDescriptions is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleUseDescriptions(ctx *MatchingRuleUseDescriptionsContext) {
}

// EnterMatchingRuleUseLabel is called when production matchingRuleUseLabel is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleUseLabel(ctx *MatchingRuleUseLabelContext) {}

// ExitMatchingRuleUseLabel is called when production matchingRuleUseLabel is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleUseLabel(ctx *MatchingRuleUseLabelContext) {}

// EnterMatchingRuleUseDescription is called when production matchingRuleUseDescription is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleUseDescription(ctx *MatchingRuleUseDescriptionContext) {
}

// ExitMatchingRuleUseDescription is called when production matchingRuleUseDescription is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleUseDescription(ctx *MatchingRuleUseDescriptionContext) {
}

// EnterMRUApplies is called when production mRUApplies is entered.
func (s *BaseRFC4512Listener) EnterMRUApplies(ctx *MRUAppliesContext) {}

// ExitMRUApplies is called when production mRUApplies is exited.
func (s *BaseRFC4512Listener) ExitMRUApplies(ctx *MRUAppliesContext) {}

// EnterLDAPSyntaxDescriptions is called when production lDAPSyntaxDescriptions is entered.
func (s *BaseRFC4512Listener) EnterLDAPSyntaxDescriptions(ctx *LDAPSyntaxDescriptionsContext) {}

// ExitLDAPSyntaxDescriptions is called when production lDAPSyntaxDescriptions is exited.
func (s *BaseRFC4512Listener) ExitLDAPSyntaxDescriptions(ctx *LDAPSyntaxDescriptionsContext) {}

// EnterLDAPSyntaxLabel is called when production lDAPSyntaxLabel is entered.
func (s *BaseRFC4512Listener) EnterLDAPSyntaxLabel(ctx *LDAPSyntaxLabelContext) {}

// ExitLDAPSyntaxLabel is called when production lDAPSyntaxLabel is exited.
func (s *BaseRFC4512Listener) ExitLDAPSyntaxLabel(ctx *LDAPSyntaxLabelContext) {}

// EnterLDAPSyntaxDescription is called when production lDAPSyntaxDescription is entered.
func (s *BaseRFC4512Listener) EnterLDAPSyntaxDescription(ctx *LDAPSyntaxDescriptionContext) {}

// ExitLDAPSyntaxDescription is called when production lDAPSyntaxDescription is exited.
func (s *BaseRFC4512Listener) ExitLDAPSyntaxDescription(ctx *LDAPSyntaxDescriptionContext) {}

// EnterDITContentRuleDescriptions is called when production dITContentRuleDescriptions is entered.
func (s *BaseRFC4512Listener) EnterDITContentRuleDescriptions(ctx *DITContentRuleDescriptionsContext) {
}

// ExitDITContentRuleDescriptions is called when production dITContentRuleDescriptions is exited.
func (s *BaseRFC4512Listener) ExitDITContentRuleDescriptions(ctx *DITContentRuleDescriptionsContext) {
}

// EnterDITContentRuleLabel is called when production dITContentRuleLabel is entered.
func (s *BaseRFC4512Listener) EnterDITContentRuleLabel(ctx *DITContentRuleLabelContext) {}

// ExitDITContentRuleLabel is called when production dITContentRuleLabel is exited.
func (s *BaseRFC4512Listener) ExitDITContentRuleLabel(ctx *DITContentRuleLabelContext) {}

// EnterDITContentRuleDescription is called when production dITContentRuleDescription is entered.
func (s *BaseRFC4512Listener) EnterDITContentRuleDescription(ctx *DITContentRuleDescriptionContext) {}

// ExitDITContentRuleDescription is called when production dITContentRuleDescription is exited.
func (s *BaseRFC4512Listener) ExitDITContentRuleDescription(ctx *DITContentRuleDescriptionContext) {}

// EnterDCRNot is called when production dCRNot is entered.
func (s *BaseRFC4512Listener) EnterDCRNot(ctx *DCRNotContext) {}

// ExitDCRNot is called when production dCRNot is exited.
func (s *BaseRFC4512Listener) ExitDCRNot(ctx *DCRNotContext) {}

// EnterDCRAux is called when production dCRAux is entered.
func (s *BaseRFC4512Listener) EnterDCRAux(ctx *DCRAuxContext) {}

// ExitDCRAux is called when production dCRAux is exited.
func (s *BaseRFC4512Listener) ExitDCRAux(ctx *DCRAuxContext) {}

// EnterDITStructureRuleDescriptions is called when production dITStructureRuleDescriptions is entered.
func (s *BaseRFC4512Listener) EnterDITStructureRuleDescriptions(ctx *DITStructureRuleDescriptionsContext) {
}

// ExitDITStructureRuleDescriptions is called when production dITStructureRuleDescriptions is exited.
func (s *BaseRFC4512Listener) ExitDITStructureRuleDescriptions(ctx *DITStructureRuleDescriptionsContext) {
}

// EnterDITStructureRuleLabel is called when production dITStructureRuleLabel is entered.
func (s *BaseRFC4512Listener) EnterDITStructureRuleLabel(ctx *DITStructureRuleLabelContext) {}

// ExitDITStructureRuleLabel is called when production dITStructureRuleLabel is exited.
func (s *BaseRFC4512Listener) ExitDITStructureRuleLabel(ctx *DITStructureRuleLabelContext) {}

// EnterDITStructureRuleDescription is called when production dITStructureRuleDescription is entered.
func (s *BaseRFC4512Listener) EnterDITStructureRuleDescription(ctx *DITStructureRuleDescriptionContext) {
}

// ExitDITStructureRuleDescription is called when production dITStructureRuleDescription is exited.
func (s *BaseRFC4512Listener) ExitDITStructureRuleDescription(ctx *DITStructureRuleDescriptionContext) {
}

// EnterStructureRules is called when production structureRules is entered.
func (s *BaseRFC4512Listener) EnterStructureRules(ctx *StructureRulesContext) {}

// ExitStructureRules is called when production structureRules is exited.
func (s *BaseRFC4512Listener) ExitStructureRules(ctx *StructureRulesContext) {}

// EnterDSRSuperRules is called when production dSRSuperRules is entered.
func (s *BaseRFC4512Listener) EnterDSRSuperRules(ctx *DSRSuperRulesContext) {}

// ExitDSRSuperRules is called when production dSRSuperRules is exited.
func (s *BaseRFC4512Listener) ExitDSRSuperRules(ctx *DSRSuperRulesContext) {}

// EnterStructureRule is called when production structureRule is entered.
func (s *BaseRFC4512Listener) EnterStructureRule(ctx *StructureRuleContext) {}

// ExitStructureRule is called when production structureRule is exited.
func (s *BaseRFC4512Listener) ExitStructureRule(ctx *StructureRuleContext) {}

// EnterDSRForm is called when production dSRForm is entered.
func (s *BaseRFC4512Listener) EnterDSRForm(ctx *DSRFormContext) {}

// ExitDSRForm is called when production dSRForm is exited.
func (s *BaseRFC4512Listener) ExitDSRForm(ctx *DSRFormContext) {}

// EnterNameFormDescriptions is called when production nameFormDescriptions is entered.
func (s *BaseRFC4512Listener) EnterNameFormDescriptions(ctx *NameFormDescriptionsContext) {}

// ExitNameFormDescriptions is called when production nameFormDescriptions is exited.
func (s *BaseRFC4512Listener) ExitNameFormDescriptions(ctx *NameFormDescriptionsContext) {}

// EnterNameFormLabel is called when production nameFormLabel is entered.
func (s *BaseRFC4512Listener) EnterNameFormLabel(ctx *NameFormLabelContext) {}

// ExitNameFormLabel is called when production nameFormLabel is exited.
func (s *BaseRFC4512Listener) ExitNameFormLabel(ctx *NameFormLabelContext) {}

// EnterNameFormDescription is called when production nameFormDescription is entered.
func (s *BaseRFC4512Listener) EnterNameFormDescription(ctx *NameFormDescriptionContext) {}

// ExitNameFormDescription is called when production nameFormDescription is exited.
func (s *BaseRFC4512Listener) ExitNameFormDescription(ctx *NameFormDescriptionContext) {}

// EnterNFStructuralOC is called when production nFStructuralOC is entered.
func (s *BaseRFC4512Listener) EnterNFStructuralOC(ctx *NFStructuralOCContext) {}

// ExitNFStructuralOC is called when production nFStructuralOC is exited.
func (s *BaseRFC4512Listener) ExitNFStructuralOC(ctx *NFStructuralOCContext) {}

// EnterDefinitionObsolete is called when production definitionObsolete is entered.
func (s *BaseRFC4512Listener) EnterDefinitionObsolete(ctx *DefinitionObsoleteContext) {}

// ExitDefinitionObsolete is called when production definitionObsolete is exited.
func (s *BaseRFC4512Listener) ExitDefinitionObsolete(ctx *DefinitionObsoleteContext) {}

// EnterDefinitionSyntax is called when production definitionSyntax is entered.
func (s *BaseRFC4512Listener) EnterDefinitionSyntax(ctx *DefinitionSyntaxContext) {}

// ExitDefinitionSyntax is called when production definitionSyntax is exited.
func (s *BaseRFC4512Listener) ExitDefinitionSyntax(ctx *DefinitionSyntaxContext) {}

// EnterDefinitionMust is called when production definitionMust is entered.
func (s *BaseRFC4512Listener) EnterDefinitionMust(ctx *DefinitionMustContext) {}

// ExitDefinitionMust is called when production definitionMust is exited.
func (s *BaseRFC4512Listener) ExitDefinitionMust(ctx *DefinitionMustContext) {}

// EnterDefinitionMay is called when production definitionMay is entered.
func (s *BaseRFC4512Listener) EnterDefinitionMay(ctx *DefinitionMayContext) {}

// ExitDefinitionMay is called when production definitionMay is exited.
func (s *BaseRFC4512Listener) ExitDefinitionMay(ctx *DefinitionMayContext) {}

// EnterDefinitionDescription is called when production definitionDescription is entered.
func (s *BaseRFC4512Listener) EnterDefinitionDescription(ctx *DefinitionDescriptionContext) {}

// ExitDefinitionDescription is called when production definitionDescription is exited.
func (s *BaseRFC4512Listener) ExitDefinitionDescription(ctx *DefinitionDescriptionContext) {}

// EnterDefinitionName is called when production definitionName is entered.
func (s *BaseRFC4512Listener) EnterDefinitionName(ctx *DefinitionNameContext) {}

// ExitDefinitionName is called when production definitionName is exited.
func (s *BaseRFC4512Listener) ExitDefinitionName(ctx *DefinitionNameContext) {}

// EnterDefinitionExtensions is called when production definitionExtensions is entered.
func (s *BaseRFC4512Listener) EnterDefinitionExtensions(ctx *DefinitionExtensionsContext) {}

// ExitDefinitionExtensions is called when production definitionExtensions is exited.
func (s *BaseRFC4512Listener) ExitDefinitionExtensions(ctx *DefinitionExtensionsContext) {}

// EnterDefinitionExtension is called when production definitionExtension is entered.
func (s *BaseRFC4512Listener) EnterDefinitionExtension(ctx *DefinitionExtensionContext) {}

// ExitDefinitionExtension is called when production definitionExtension is exited.
func (s *BaseRFC4512Listener) ExitDefinitionExtension(ctx *DefinitionExtensionContext) {}

// EnterDefinitionLabelDelim is called when production definitionLabelDelim is entered.
func (s *BaseRFC4512Listener) EnterDefinitionLabelDelim(ctx *DefinitionLabelDelimContext) {}

// ExitDefinitionLabelDelim is called when production definitionLabelDelim is exited.
func (s *BaseRFC4512Listener) ExitDefinitionLabelDelim(ctx *DefinitionLabelDelimContext) {}

// EnterOIDs is called when production oIDs is entered.
func (s *BaseRFC4512Listener) EnterOIDs(ctx *OIDsContext) {}

// ExitOIDs is called when production oIDs is exited.
func (s *BaseRFC4512Listener) ExitOIDs(ctx *OIDsContext) {}

// EnterAttributeDescription is called when production attributeDescription is entered.
func (s *BaseRFC4512Listener) EnterAttributeDescription(ctx *AttributeDescriptionContext) {}

// ExitAttributeDescription is called when production attributeDescription is exited.
func (s *BaseRFC4512Listener) ExitAttributeDescription(ctx *AttributeDescriptionContext) {}

// EnterOID is called when production oID is entered.
func (s *BaseRFC4512Listener) EnterOID(ctx *OIDContext) {}

// ExitOID is called when production oID is exited.
func (s *BaseRFC4512Listener) ExitOID(ctx *OIDContext) {}

// EnterOpenParen is called when production openParen is entered.
func (s *BaseRFC4512Listener) EnterOpenParen(ctx *OpenParenContext) {}

// ExitOpenParen is called when production openParen is exited.
func (s *BaseRFC4512Listener) ExitOpenParen(ctx *OpenParenContext) {}

// EnterCloseParen is called when production closeParen is entered.
func (s *BaseRFC4512Listener) EnterCloseParen(ctx *CloseParenContext) {}

// ExitCloseParen is called when production closeParen is exited.
func (s *BaseRFC4512Listener) ExitCloseParen(ctx *CloseParenContext) {}

// EnterNumericOIDOrMacro is called when production numericOIDOrMacro is entered.
func (s *BaseRFC4512Listener) EnterNumericOIDOrMacro(ctx *NumericOIDOrMacroContext) {}

// ExitNumericOIDOrMacro is called when production numericOIDOrMacro is exited.
func (s *BaseRFC4512Listener) ExitNumericOIDOrMacro(ctx *NumericOIDOrMacroContext) {}

// EnterMacro is called when production macro is entered.
func (s *BaseRFC4512Listener) EnterMacro(ctx *MacroContext) {}

// ExitMacro is called when production macro is exited.
func (s *BaseRFC4512Listener) ExitMacro(ctx *MacroContext) {}

// EnterNumericOID is called when production numericOID is entered.
func (s *BaseRFC4512Listener) EnterNumericOID(ctx *NumericOIDContext) {}

// ExitNumericOID is called when production numericOID is exited.
func (s *BaseRFC4512Listener) ExitNumericOID(ctx *NumericOIDContext) {}

// EnterMacroSuffix is called when production macroSuffix is entered.
func (s *BaseRFC4512Listener) EnterMacroSuffix(ctx *MacroSuffixContext) {}

// ExitMacroSuffix is called when production macroSuffix is exited.
func (s *BaseRFC4512Listener) ExitMacroSuffix(ctx *MacroSuffixContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseRFC4512Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseRFC4512Listener) ExitNumber(ctx *NumberContext) {}

// EnterAttrOptions is called when production attrOptions is entered.
func (s *BaseRFC4512Listener) EnterAttrOptions(ctx *AttrOptionsContext) {}

// ExitAttrOptions is called when production attrOptions is exited.
func (s *BaseRFC4512Listener) ExitAttrOptions(ctx *AttrOptionsContext) {}

// EnterDescriptor is called when production descriptor is entered.
func (s *BaseRFC4512Listener) EnterDescriptor(ctx *DescriptorContext) {}

// ExitDescriptor is called when production descriptor is exited.
func (s *BaseRFC4512Listener) ExitDescriptor(ctx *DescriptorContext) {}

// EnterExtensionValue is called when production extensionValue is entered.
func (s *BaseRFC4512Listener) EnterExtensionValue(ctx *ExtensionValueContext) {}

// ExitExtensionValue is called when production extensionValue is exited.
func (s *BaseRFC4512Listener) ExitExtensionValue(ctx *ExtensionValueContext) {}
