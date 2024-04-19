// Code generated from RFC4512.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr4512 // RFC4512
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

// EnterFileparse is called when production fileparse is entered.
func (s *BaseRFC4512Listener) EnterFileparse(ctx *FileparseContext) {}

// ExitFileparse is called when production fileparse is exited.
func (s *BaseRFC4512Listener) ExitFileparse(ctx *FileparseContext) {}

// EnterDefinitions is called when production definitions is entered.
func (s *BaseRFC4512Listener) EnterDefinitions(ctx *DefinitionsContext) {}

// ExitDefinitions is called when production definitions is exited.
func (s *BaseRFC4512Listener) ExitDefinitions(ctx *DefinitionsContext) {}

// EnterObjectClassDescriptions is called when production objectClassDescriptions is entered.
func (s *BaseRFC4512Listener) EnterObjectClassDescriptions(ctx *ObjectClassDescriptionsContext) {}

// ExitObjectClassDescriptions is called when production objectClassDescriptions is exited.
func (s *BaseRFC4512Listener) ExitObjectClassDescriptions(ctx *ObjectClassDescriptionsContext) {}

// EnterAttributeTypeDescriptions is called when production attributeTypeDescriptions is entered.
func (s *BaseRFC4512Listener) EnterAttributeTypeDescriptions(ctx *AttributeTypeDescriptionsContext) {}

// ExitAttributeTypeDescriptions is called when production attributeTypeDescriptions is exited.
func (s *BaseRFC4512Listener) ExitAttributeTypeDescriptions(ctx *AttributeTypeDescriptionsContext) {}

// EnterMatchingRuleDescriptions is called when production matchingRuleDescriptions is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleDescriptions(ctx *MatchingRuleDescriptionsContext) {}

// ExitMatchingRuleDescriptions is called when production matchingRuleDescriptions is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleDescriptions(ctx *MatchingRuleDescriptionsContext) {}

// EnterMatchingRuleUseDescriptions is called when production matchingRuleUseDescriptions is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleUseDescriptions(ctx *MatchingRuleUseDescriptionsContext) {
}

// ExitMatchingRuleUseDescriptions is called when production matchingRuleUseDescriptions is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleUseDescriptions(ctx *MatchingRuleUseDescriptionsContext) {
}

// EnterLDAPSyntaxDescriptions is called when production lDAPSyntaxDescriptions is entered.
func (s *BaseRFC4512Listener) EnterLDAPSyntaxDescriptions(ctx *LDAPSyntaxDescriptionsContext) {}

// ExitLDAPSyntaxDescriptions is called when production lDAPSyntaxDescriptions is exited.
func (s *BaseRFC4512Listener) ExitLDAPSyntaxDescriptions(ctx *LDAPSyntaxDescriptionsContext) {}

// EnterDITStructureRuleDescriptions is called when production dITStructureRuleDescriptions is entered.
func (s *BaseRFC4512Listener) EnterDITStructureRuleDescriptions(ctx *DITStructureRuleDescriptionsContext) {
}

// ExitDITStructureRuleDescriptions is called when production dITStructureRuleDescriptions is exited.
func (s *BaseRFC4512Listener) ExitDITStructureRuleDescriptions(ctx *DITStructureRuleDescriptionsContext) {
}

// EnterNameFormDescriptions is called when production nameFormDescriptions is entered.
func (s *BaseRFC4512Listener) EnterNameFormDescriptions(ctx *NameFormDescriptionsContext) {}

// ExitNameFormDescriptions is called when production nameFormDescriptions is exited.
func (s *BaseRFC4512Listener) ExitNameFormDescriptions(ctx *NameFormDescriptionsContext) {}

// EnterDITContentRuleDescriptions is called when production dITContentRuleDescriptions is entered.
func (s *BaseRFC4512Listener) EnterDITContentRuleDescriptions(ctx *DITContentRuleDescriptionsContext) {
}

// ExitDITContentRuleDescriptions is called when production dITContentRuleDescriptions is exited.
func (s *BaseRFC4512Listener) ExitDITContentRuleDescriptions(ctx *DITContentRuleDescriptionsContext) {
}

// EnterObjectIdentifier is called when production objectIdentifier is entered.
func (s *BaseRFC4512Listener) EnterObjectIdentifier(ctx *ObjectIdentifierContext) {}

// ExitObjectIdentifier is called when production objectIdentifier is exited.
func (s *BaseRFC4512Listener) ExitObjectIdentifier(ctx *ObjectIdentifierContext) {}

// EnterObjectClassLabel is called when production objectClassLabel is entered.
func (s *BaseRFC4512Listener) EnterObjectClassLabel(ctx *ObjectClassLabelContext) {}

// ExitObjectClassLabel is called when production objectClassLabel is exited.
func (s *BaseRFC4512Listener) ExitObjectClassLabel(ctx *ObjectClassLabelContext) {}

// EnterMatchingRuleLabel is called when production matchingRuleLabel is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleLabel(ctx *MatchingRuleLabelContext) {}

// ExitMatchingRuleLabel is called when production matchingRuleLabel is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleLabel(ctx *MatchingRuleLabelContext) {}

// EnterMatchingRuleUseLabel is called when production matchingRuleUseLabel is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleUseLabel(ctx *MatchingRuleUseLabelContext) {}

// ExitMatchingRuleUseLabel is called when production matchingRuleUseLabel is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleUseLabel(ctx *MatchingRuleUseLabelContext) {}

// EnterLDAPSyntaxLabel is called when production lDAPSyntaxLabel is entered.
func (s *BaseRFC4512Listener) EnterLDAPSyntaxLabel(ctx *LDAPSyntaxLabelContext) {}

// ExitLDAPSyntaxLabel is called when production lDAPSyntaxLabel is exited.
func (s *BaseRFC4512Listener) ExitLDAPSyntaxLabel(ctx *LDAPSyntaxLabelContext) {}

// EnterDITContentRuleLabel is called when production dITContentRuleLabel is entered.
func (s *BaseRFC4512Listener) EnterDITContentRuleLabel(ctx *DITContentRuleLabelContext) {}

// ExitDITContentRuleLabel is called when production dITContentRuleLabel is exited.
func (s *BaseRFC4512Listener) ExitDITContentRuleLabel(ctx *DITContentRuleLabelContext) {}

// EnterDITStructureRuleLabel is called when production dITStructureRuleLabel is entered.
func (s *BaseRFC4512Listener) EnterDITStructureRuleLabel(ctx *DITStructureRuleLabelContext) {}

// ExitDITStructureRuleLabel is called when production dITStructureRuleLabel is exited.
func (s *BaseRFC4512Listener) ExitDITStructureRuleLabel(ctx *DITStructureRuleLabelContext) {}

// EnterAttributeTypeLabel is called when production attributeTypeLabel is entered.
func (s *BaseRFC4512Listener) EnterAttributeTypeLabel(ctx *AttributeTypeLabelContext) {}

// ExitAttributeTypeLabel is called when production attributeTypeLabel is exited.
func (s *BaseRFC4512Listener) ExitAttributeTypeLabel(ctx *AttributeTypeLabelContext) {}

// EnterNameFormLabel is called when production nameFormLabel is entered.
func (s *BaseRFC4512Listener) EnterNameFormLabel(ctx *NameFormLabelContext) {}

// ExitNameFormLabel is called when production nameFormLabel is exited.
func (s *BaseRFC4512Listener) ExitNameFormLabel(ctx *NameFormLabelContext) {}

// EnterDefinitionLabelDelim is called when production definitionLabelDelim is entered.
func (s *BaseRFC4512Listener) EnterDefinitionLabelDelim(ctx *DefinitionLabelDelimContext) {}

// ExitDefinitionLabelDelim is called when production definitionLabelDelim is exited.
func (s *BaseRFC4512Listener) ExitDefinitionLabelDelim(ctx *DefinitionLabelDelimContext) {}

// EnterKind is called when production kind is entered.
func (s *BaseRFC4512Listener) EnterKind(ctx *KindContext) {}

// ExitKind is called when production kind is exited.
func (s *BaseRFC4512Listener) ExitKind(ctx *KindContext) {}

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

// EnterApplies is called when production applies is entered.
func (s *BaseRFC4512Listener) EnterApplies(ctx *AppliesContext) {}

// ExitApplies is called when production applies is exited.
func (s *BaseRFC4512Listener) ExitApplies(ctx *AppliesContext) {}

// EnterMust is called when production must is entered.
func (s *BaseRFC4512Listener) EnterMust(ctx *MustContext) {}

// ExitMust is called when production must is exited.
func (s *BaseRFC4512Listener) ExitMust(ctx *MustContext) {}

// EnterAux is called when production aux is entered.
func (s *BaseRFC4512Listener) EnterAux(ctx *AuxContext) {}

// ExitAux is called when production aux is exited.
func (s *BaseRFC4512Listener) ExitAux(ctx *AuxContext) {}

// EnterMay is called when production may is entered.
func (s *BaseRFC4512Listener) EnterMay(ctx *MayContext) {}

// ExitMay is called when production may is exited.
func (s *BaseRFC4512Listener) ExitMay(ctx *MayContext) {}

// EnterNot is called when production not is entered.
func (s *BaseRFC4512Listener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production not is exited.
func (s *BaseRFC4512Listener) ExitNot(ctx *NotContext) {}

// EnterForm is called when production form is entered.
func (s *BaseRFC4512Listener) EnterForm(ctx *FormContext) {}

// ExitForm is called when production form is exited.
func (s *BaseRFC4512Listener) ExitForm(ctx *FormContext) {}

// EnterOrdering is called when production ordering is entered.
func (s *BaseRFC4512Listener) EnterOrdering(ctx *OrderingContext) {}

// ExitOrdering is called when production ordering is exited.
func (s *BaseRFC4512Listener) ExitOrdering(ctx *OrderingContext) {}

// EnterEquality is called when production equality is entered.
func (s *BaseRFC4512Listener) EnterEquality(ctx *EqualityContext) {}

// ExitEquality is called when production equality is exited.
func (s *BaseRFC4512Listener) ExitEquality(ctx *EqualityContext) {}

// EnterStructuralClass is called when production structuralClass is entered.
func (s *BaseRFC4512Listener) EnterStructuralClass(ctx *StructuralClassContext) {}

// ExitStructuralClass is called when production structuralClass is exited.
func (s *BaseRFC4512Listener) ExitStructuralClass(ctx *StructuralClassContext) {}

// EnterSuperType is called when production superType is entered.
func (s *BaseRFC4512Listener) EnterSuperType(ctx *SuperTypeContext) {}

// ExitSuperType is called when production superType is exited.
func (s *BaseRFC4512Listener) ExitSuperType(ctx *SuperTypeContext) {}

// EnterSuperClasses is called when production superClasses is entered.
func (s *BaseRFC4512Listener) EnterSuperClasses(ctx *SuperClassesContext) {}

// ExitSuperClasses is called when production superClasses is exited.
func (s *BaseRFC4512Listener) ExitSuperClasses(ctx *SuperClassesContext) {}

// EnterSuperRules is called when production superRules is entered.
func (s *BaseRFC4512Listener) EnterSuperRules(ctx *SuperRulesContext) {}

// ExitSuperRules is called when production superRules is exited.
func (s *BaseRFC4512Listener) ExitSuperRules(ctx *SuperRulesContext) {}

// EnterSubstr is called when production substr is entered.
func (s *BaseRFC4512Listener) EnterSubstr(ctx *SubstrContext) {}

// ExitSubstr is called when production substr is exited.
func (s *BaseRFC4512Listener) ExitSubstr(ctx *SubstrContext) {}

// EnterName is called when production name is entered.
func (s *BaseRFC4512Listener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseRFC4512Listener) ExitName(ctx *NameContext) {}

// EnterUsage is called when production usage is entered.
func (s *BaseRFC4512Listener) EnterUsage(ctx *UsageContext) {}

// ExitUsage is called when production usage is exited.
func (s *BaseRFC4512Listener) ExitUsage(ctx *UsageContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseRFC4512Listener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseRFC4512Listener) ExitDescription(ctx *DescriptionContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *BaseRFC4512Listener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BaseRFC4512Listener) ExitSyntax(ctx *SyntaxContext) {}

// EnterSingleValue is called when production singleValue is entered.
func (s *BaseRFC4512Listener) EnterSingleValue(ctx *SingleValueContext) {}

// ExitSingleValue is called when production singleValue is exited.
func (s *BaseRFC4512Listener) ExitSingleValue(ctx *SingleValueContext) {}

// EnterImmutability is called when production immutability is entered.
func (s *BaseRFC4512Listener) EnterImmutability(ctx *ImmutabilityContext) {}

// ExitImmutability is called when production immutability is exited.
func (s *BaseRFC4512Listener) ExitImmutability(ctx *ImmutabilityContext) {}

// EnterCollective is called when production collective is entered.
func (s *BaseRFC4512Listener) EnterCollective(ctx *CollectiveContext) {}

// ExitCollective is called when production collective is exited.
func (s *BaseRFC4512Listener) ExitCollective(ctx *CollectiveContext) {}

// EnterObsolescence is called when production obsolescence is entered.
func (s *BaseRFC4512Listener) EnterObsolescence(ctx *ObsolescenceContext) {}

// ExitObsolescence is called when production obsolescence is exited.
func (s *BaseRFC4512Listener) ExitObsolescence(ctx *ObsolescenceContext) {}

// EnterMinUpperBounds is called when production minUpperBounds is entered.
func (s *BaseRFC4512Listener) EnterMinUpperBounds(ctx *MinUpperBoundsContext) {}

// ExitMinUpperBounds is called when production minUpperBounds is exited.
func (s *BaseRFC4512Listener) ExitMinUpperBounds(ctx *MinUpperBoundsContext) {}

// EnterStructureRule is called when production structureRule is entered.
func (s *BaseRFC4512Listener) EnterStructureRule(ctx *StructureRuleContext) {}

// ExitStructureRule is called when production structureRule is exited.
func (s *BaseRFC4512Listener) ExitStructureRule(ctx *StructureRuleContext) {}

// EnterUsages is called when production usages is entered.
func (s *BaseRFC4512Listener) EnterUsages(ctx *UsagesContext) {}

// ExitUsages is called when production usages is exited.
func (s *BaseRFC4512Listener) ExitUsages(ctx *UsagesContext) {}

// EnterQuotedDescriptor is called when production quotedDescriptor is entered.
func (s *BaseRFC4512Listener) EnterQuotedDescriptor(ctx *QuotedDescriptorContext) {}

// ExitQuotedDescriptor is called when production quotedDescriptor is exited.
func (s *BaseRFC4512Listener) ExitQuotedDescriptor(ctx *QuotedDescriptorContext) {}

// EnterQuotedString is called when production quotedString is entered.
func (s *BaseRFC4512Listener) EnterQuotedString(ctx *QuotedStringContext) {}

// ExitQuotedString is called when production quotedString is exited.
func (s *BaseRFC4512Listener) ExitQuotedString(ctx *QuotedStringContext) {}

// EnterSchemaDN is called when production schemaDN is entered.
func (s *BaseRFC4512Listener) EnterSchemaDN(ctx *SchemaDNContext) {}

// ExitSchemaDN is called when production schemaDN is exited.
func (s *BaseRFC4512Listener) ExitSchemaDN(ctx *SchemaDNContext) {}

// EnterObjectIdentifierOIDOrName is called when production objectIdentifierOIDOrName is entered.
func (s *BaseRFC4512Listener) EnterObjectIdentifierOIDOrName(ctx *ObjectIdentifierOIDOrNameContext) {}

// ExitObjectIdentifierOIDOrName is called when production objectIdentifierOIDOrName is exited.
func (s *BaseRFC4512Listener) ExitObjectIdentifierOIDOrName(ctx *ObjectIdentifierOIDOrNameContext) {}

// EnterAttributeDescription is called when production attributeDescription is entered.
func (s *BaseRFC4512Listener) EnterAttributeDescription(ctx *AttributeDescriptionContext) {}

// ExitAttributeDescription is called when production attributeDescription is exited.
func (s *BaseRFC4512Listener) ExitAttributeDescription(ctx *AttributeDescriptionContext) {}

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

// EnterUserApplication is called when production userApplication is entered.
func (s *BaseRFC4512Listener) EnterUserApplication(ctx *UserApplicationContext) {}

// ExitUserApplication is called when production userApplication is exited.
func (s *BaseRFC4512Listener) ExitUserApplication(ctx *UserApplicationContext) {}

// EnterDirectoryOperation is called when production directoryOperation is entered.
func (s *BaseRFC4512Listener) EnterDirectoryOperation(ctx *DirectoryOperationContext) {}

// ExitDirectoryOperation is called when production directoryOperation is exited.
func (s *BaseRFC4512Listener) ExitDirectoryOperation(ctx *DirectoryOperationContext) {}

// EnterDistributedOperation is called when production distributedOperation is entered.
func (s *BaseRFC4512Listener) EnterDistributedOperation(ctx *DistributedOperationContext) {}

// ExitDistributedOperation is called when production distributedOperation is exited.
func (s *BaseRFC4512Listener) ExitDistributedOperation(ctx *DistributedOperationContext) {}

// EnterDSAOperation is called when production dSAOperation is entered.
func (s *BaseRFC4512Listener) EnterDSAOperation(ctx *DSAOperationContext) {}

// ExitDSAOperation is called when production dSAOperation is exited.
func (s *BaseRFC4512Listener) ExitDSAOperation(ctx *DSAOperationContext) {}

// EnterStructureRules is called when production structureRules is entered.
func (s *BaseRFC4512Listener) EnterStructureRules(ctx *StructureRulesContext) {}

// ExitStructureRules is called when production structureRules is exited.
func (s *BaseRFC4512Listener) ExitStructureRules(ctx *StructureRulesContext) {}

// EnterOID is called when production oID is entered.
func (s *BaseRFC4512Listener) EnterOID(ctx *OIDContext) {}

// ExitOID is called when production oID is exited.
func (s *BaseRFC4512Listener) ExitOID(ctx *OIDContext) {}

// EnterOIDs is called when production oIDs is entered.
func (s *BaseRFC4512Listener) EnterOIDs(ctx *OIDsContext) {}

// ExitOIDs is called when production oIDs is exited.
func (s *BaseRFC4512Listener) ExitOIDs(ctx *OIDsContext) {}

// EnterExtensionValue is called when production extensionValue is entered.
func (s *BaseRFC4512Listener) EnterExtensionValue(ctx *ExtensionValueContext) {}

// ExitExtensionValue is called when production extensionValue is exited.
func (s *BaseRFC4512Listener) ExitExtensionValue(ctx *ExtensionValueContext) {}

// EnterExtensions is called when production extensions is entered.
func (s *BaseRFC4512Listener) EnterExtensions(ctx *ExtensionsContext) {}

// ExitExtensions is called when production extensions is exited.
func (s *BaseRFC4512Listener) ExitExtensions(ctx *ExtensionsContext) {}

// EnterExtension is called when production extension is entered.
func (s *BaseRFC4512Listener) EnterExtension(ctx *ExtensionContext) {}

// ExitExtension is called when production extension is exited.
func (s *BaseRFC4512Listener) ExitExtension(ctx *ExtensionContext) {}

// EnterNames is called when production names is entered.
func (s *BaseRFC4512Listener) EnterNames(ctx *NamesContext) {}

// ExitNames is called when production names is exited.
func (s *BaseRFC4512Listener) ExitNames(ctx *NamesContext) {}

// EnterDescriptor is called when production descriptor is entered.
func (s *BaseRFC4512Listener) EnterDescriptor(ctx *DescriptorContext) {}

// ExitDescriptor is called when production descriptor is exited.
func (s *BaseRFC4512Listener) ExitDescriptor(ctx *DescriptorContext) {}

// EnterObjectClassDescription is called when production objectClassDescription is entered.
func (s *BaseRFC4512Listener) EnterObjectClassDescription(ctx *ObjectClassDescriptionContext) {}

// ExitObjectClassDescription is called when production objectClassDescription is exited.
func (s *BaseRFC4512Listener) ExitObjectClassDescription(ctx *ObjectClassDescriptionContext) {}

// EnterAttributeTypeDescription is called when production attributeTypeDescription is entered.
func (s *BaseRFC4512Listener) EnterAttributeTypeDescription(ctx *AttributeTypeDescriptionContext) {}

// ExitAttributeTypeDescription is called when production attributeTypeDescription is exited.
func (s *BaseRFC4512Listener) ExitAttributeTypeDescription(ctx *AttributeTypeDescriptionContext) {}

// EnterMatchingRuleDescription is called when production matchingRuleDescription is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleDescription(ctx *MatchingRuleDescriptionContext) {}

// ExitMatchingRuleDescription is called when production matchingRuleDescription is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleDescription(ctx *MatchingRuleDescriptionContext) {}

// EnterMatchingRuleUseDescription is called when production matchingRuleUseDescription is entered.
func (s *BaseRFC4512Listener) EnterMatchingRuleUseDescription(ctx *MatchingRuleUseDescriptionContext) {
}

// ExitMatchingRuleUseDescription is called when production matchingRuleUseDescription is exited.
func (s *BaseRFC4512Listener) ExitMatchingRuleUseDescription(ctx *MatchingRuleUseDescriptionContext) {
}

// EnterLDAPSyntaxDescription is called when production lDAPSyntaxDescription is entered.
func (s *BaseRFC4512Listener) EnterLDAPSyntaxDescription(ctx *LDAPSyntaxDescriptionContext) {}

// ExitLDAPSyntaxDescription is called when production lDAPSyntaxDescription is exited.
func (s *BaseRFC4512Listener) ExitLDAPSyntaxDescription(ctx *LDAPSyntaxDescriptionContext) {}

// EnterDITContentRuleDescription is called when production dITContentRuleDescription is entered.
func (s *BaseRFC4512Listener) EnterDITContentRuleDescription(ctx *DITContentRuleDescriptionContext) {}

// ExitDITContentRuleDescription is called when production dITContentRuleDescription is exited.
func (s *BaseRFC4512Listener) ExitDITContentRuleDescription(ctx *DITContentRuleDescriptionContext) {}

// EnterDITStructureRuleDescription is called when production dITStructureRuleDescription is entered.
func (s *BaseRFC4512Listener) EnterDITStructureRuleDescription(ctx *DITStructureRuleDescriptionContext) {
}

// ExitDITStructureRuleDescription is called when production dITStructureRuleDescription is exited.
func (s *BaseRFC4512Listener) ExitDITStructureRuleDescription(ctx *DITStructureRuleDescriptionContext) {
}

// EnterNameFormDescription is called when production nameFormDescription is entered.
func (s *BaseRFC4512Listener) EnterNameFormDescription(ctx *NameFormDescriptionContext) {}

// ExitNameFormDescription is called when production nameFormDescription is exited.
func (s *BaseRFC4512Listener) ExitNameFormDescription(ctx *NameFormDescriptionContext) {}
