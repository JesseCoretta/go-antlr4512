// Generated from RFC4512.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link RFC4512Parser}.
 */
public interface RFC4512Listener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#fileparse}.
	 * @param ctx the parse tree
	 */
	void enterFileparse(RFC4512Parser.FileparseContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#fileparse}.
	 * @param ctx the parse tree
	 */
	void exitFileparse(RFC4512Parser.FileparseContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#definitions}.
	 * @param ctx the parse tree
	 */
	void enterDefinitions(RFC4512Parser.DefinitionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#definitions}.
	 * @param ctx the parse tree
	 */
	void exitDefinitions(RFC4512Parser.DefinitionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#objectClassDescriptions}.
	 * @param ctx the parse tree
	 */
	void enterObjectClassDescriptions(RFC4512Parser.ObjectClassDescriptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#objectClassDescriptions}.
	 * @param ctx the parse tree
	 */
	void exitObjectClassDescriptions(RFC4512Parser.ObjectClassDescriptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#attributeTypeDescriptions}.
	 * @param ctx the parse tree
	 */
	void enterAttributeTypeDescriptions(RFC4512Parser.AttributeTypeDescriptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#attributeTypeDescriptions}.
	 * @param ctx the parse tree
	 */
	void exitAttributeTypeDescriptions(RFC4512Parser.AttributeTypeDescriptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#matchingRuleDescriptions}.
	 * @param ctx the parse tree
	 */
	void enterMatchingRuleDescriptions(RFC4512Parser.MatchingRuleDescriptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#matchingRuleDescriptions}.
	 * @param ctx the parse tree
	 */
	void exitMatchingRuleDescriptions(RFC4512Parser.MatchingRuleDescriptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#matchingRuleUseDescriptions}.
	 * @param ctx the parse tree
	 */
	void enterMatchingRuleUseDescriptions(RFC4512Parser.MatchingRuleUseDescriptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#matchingRuleUseDescriptions}.
	 * @param ctx the parse tree
	 */
	void exitMatchingRuleUseDescriptions(RFC4512Parser.MatchingRuleUseDescriptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#lDAPSyntaxDescriptions}.
	 * @param ctx the parse tree
	 */
	void enterLDAPSyntaxDescriptions(RFC4512Parser.LDAPSyntaxDescriptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#lDAPSyntaxDescriptions}.
	 * @param ctx the parse tree
	 */
	void exitLDAPSyntaxDescriptions(RFC4512Parser.LDAPSyntaxDescriptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#dITStructureRuleDescriptions}.
	 * @param ctx the parse tree
	 */
	void enterDITStructureRuleDescriptions(RFC4512Parser.DITStructureRuleDescriptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#dITStructureRuleDescriptions}.
	 * @param ctx the parse tree
	 */
	void exitDITStructureRuleDescriptions(RFC4512Parser.DITStructureRuleDescriptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#nameFormDescriptions}.
	 * @param ctx the parse tree
	 */
	void enterNameFormDescriptions(RFC4512Parser.NameFormDescriptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#nameFormDescriptions}.
	 * @param ctx the parse tree
	 */
	void exitNameFormDescriptions(RFC4512Parser.NameFormDescriptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#dITContentRuleDescriptions}.
	 * @param ctx the parse tree
	 */
	void enterDITContentRuleDescriptions(RFC4512Parser.DITContentRuleDescriptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#dITContentRuleDescriptions}.
	 * @param ctx the parse tree
	 */
	void exitDITContentRuleDescriptions(RFC4512Parser.DITContentRuleDescriptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#objectIdentifier}.
	 * @param ctx the parse tree
	 */
	void enterObjectIdentifier(RFC4512Parser.ObjectIdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#objectIdentifier}.
	 * @param ctx the parse tree
	 */
	void exitObjectIdentifier(RFC4512Parser.ObjectIdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#objectClassLabel}.
	 * @param ctx the parse tree
	 */
	void enterObjectClassLabel(RFC4512Parser.ObjectClassLabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#objectClassLabel}.
	 * @param ctx the parse tree
	 */
	void exitObjectClassLabel(RFC4512Parser.ObjectClassLabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#matchingRuleLabel}.
	 * @param ctx the parse tree
	 */
	void enterMatchingRuleLabel(RFC4512Parser.MatchingRuleLabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#matchingRuleLabel}.
	 * @param ctx the parse tree
	 */
	void exitMatchingRuleLabel(RFC4512Parser.MatchingRuleLabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#matchingRuleUseLabel}.
	 * @param ctx the parse tree
	 */
	void enterMatchingRuleUseLabel(RFC4512Parser.MatchingRuleUseLabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#matchingRuleUseLabel}.
	 * @param ctx the parse tree
	 */
	void exitMatchingRuleUseLabel(RFC4512Parser.MatchingRuleUseLabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#lDAPSyntaxLabel}.
	 * @param ctx the parse tree
	 */
	void enterLDAPSyntaxLabel(RFC4512Parser.LDAPSyntaxLabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#lDAPSyntaxLabel}.
	 * @param ctx the parse tree
	 */
	void exitLDAPSyntaxLabel(RFC4512Parser.LDAPSyntaxLabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#dITContentRuleLabel}.
	 * @param ctx the parse tree
	 */
	void enterDITContentRuleLabel(RFC4512Parser.DITContentRuleLabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#dITContentRuleLabel}.
	 * @param ctx the parse tree
	 */
	void exitDITContentRuleLabel(RFC4512Parser.DITContentRuleLabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#dITStructureRuleLabel}.
	 * @param ctx the parse tree
	 */
	void enterDITStructureRuleLabel(RFC4512Parser.DITStructureRuleLabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#dITStructureRuleLabel}.
	 * @param ctx the parse tree
	 */
	void exitDITStructureRuleLabel(RFC4512Parser.DITStructureRuleLabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#attributeTypeLabel}.
	 * @param ctx the parse tree
	 */
	void enterAttributeTypeLabel(RFC4512Parser.AttributeTypeLabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#attributeTypeLabel}.
	 * @param ctx the parse tree
	 */
	void exitAttributeTypeLabel(RFC4512Parser.AttributeTypeLabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#nameFormLabel}.
	 * @param ctx the parse tree
	 */
	void enterNameFormLabel(RFC4512Parser.NameFormLabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#nameFormLabel}.
	 * @param ctx the parse tree
	 */
	void exitNameFormLabel(RFC4512Parser.NameFormLabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#definitionLabelDelim}.
	 * @param ctx the parse tree
	 */
	void enterDefinitionLabelDelim(RFC4512Parser.DefinitionLabelDelimContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#definitionLabelDelim}.
	 * @param ctx the parse tree
	 */
	void exitDefinitionLabelDelim(RFC4512Parser.DefinitionLabelDelimContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#kind}.
	 * @param ctx the parse tree
	 */
	void enterKind(RFC4512Parser.KindContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#kind}.
	 * @param ctx the parse tree
	 */
	void exitKind(RFC4512Parser.KindContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#abstractKind}.
	 * @param ctx the parse tree
	 */
	void enterAbstractKind(RFC4512Parser.AbstractKindContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#abstractKind}.
	 * @param ctx the parse tree
	 */
	void exitAbstractKind(RFC4512Parser.AbstractKindContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#structuralKind}.
	 * @param ctx the parse tree
	 */
	void enterStructuralKind(RFC4512Parser.StructuralKindContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#structuralKind}.
	 * @param ctx the parse tree
	 */
	void exitStructuralKind(RFC4512Parser.StructuralKindContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#auxiliaryKind}.
	 * @param ctx the parse tree
	 */
	void enterAuxiliaryKind(RFC4512Parser.AuxiliaryKindContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#auxiliaryKind}.
	 * @param ctx the parse tree
	 */
	void exitAuxiliaryKind(RFC4512Parser.AuxiliaryKindContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#applies}.
	 * @param ctx the parse tree
	 */
	void enterApplies(RFC4512Parser.AppliesContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#applies}.
	 * @param ctx the parse tree
	 */
	void exitApplies(RFC4512Parser.AppliesContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#must}.
	 * @param ctx the parse tree
	 */
	void enterMust(RFC4512Parser.MustContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#must}.
	 * @param ctx the parse tree
	 */
	void exitMust(RFC4512Parser.MustContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#aux}.
	 * @param ctx the parse tree
	 */
	void enterAux(RFC4512Parser.AuxContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#aux}.
	 * @param ctx the parse tree
	 */
	void exitAux(RFC4512Parser.AuxContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#may}.
	 * @param ctx the parse tree
	 */
	void enterMay(RFC4512Parser.MayContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#may}.
	 * @param ctx the parse tree
	 */
	void exitMay(RFC4512Parser.MayContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#not}.
	 * @param ctx the parse tree
	 */
	void enterNot(RFC4512Parser.NotContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#not}.
	 * @param ctx the parse tree
	 */
	void exitNot(RFC4512Parser.NotContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#form}.
	 * @param ctx the parse tree
	 */
	void enterForm(RFC4512Parser.FormContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#form}.
	 * @param ctx the parse tree
	 */
	void exitForm(RFC4512Parser.FormContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#ordering}.
	 * @param ctx the parse tree
	 */
	void enterOrdering(RFC4512Parser.OrderingContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#ordering}.
	 * @param ctx the parse tree
	 */
	void exitOrdering(RFC4512Parser.OrderingContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#equality}.
	 * @param ctx the parse tree
	 */
	void enterEquality(RFC4512Parser.EqualityContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#equality}.
	 * @param ctx the parse tree
	 */
	void exitEquality(RFC4512Parser.EqualityContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#structuralClass}.
	 * @param ctx the parse tree
	 */
	void enterStructuralClass(RFC4512Parser.StructuralClassContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#structuralClass}.
	 * @param ctx the parse tree
	 */
	void exitStructuralClass(RFC4512Parser.StructuralClassContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#superType}.
	 * @param ctx the parse tree
	 */
	void enterSuperType(RFC4512Parser.SuperTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#superType}.
	 * @param ctx the parse tree
	 */
	void exitSuperType(RFC4512Parser.SuperTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#superClasses}.
	 * @param ctx the parse tree
	 */
	void enterSuperClasses(RFC4512Parser.SuperClassesContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#superClasses}.
	 * @param ctx the parse tree
	 */
	void exitSuperClasses(RFC4512Parser.SuperClassesContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#superRules}.
	 * @param ctx the parse tree
	 */
	void enterSuperRules(RFC4512Parser.SuperRulesContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#superRules}.
	 * @param ctx the parse tree
	 */
	void exitSuperRules(RFC4512Parser.SuperRulesContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#substr}.
	 * @param ctx the parse tree
	 */
	void enterSubstr(RFC4512Parser.SubstrContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#substr}.
	 * @param ctx the parse tree
	 */
	void exitSubstr(RFC4512Parser.SubstrContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#name}.
	 * @param ctx the parse tree
	 */
	void enterName(RFC4512Parser.NameContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#name}.
	 * @param ctx the parse tree
	 */
	void exitName(RFC4512Parser.NameContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#usage}.
	 * @param ctx the parse tree
	 */
	void enterUsage(RFC4512Parser.UsageContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#usage}.
	 * @param ctx the parse tree
	 */
	void exitUsage(RFC4512Parser.UsageContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#description}.
	 * @param ctx the parse tree
	 */
	void enterDescription(RFC4512Parser.DescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#description}.
	 * @param ctx the parse tree
	 */
	void exitDescription(RFC4512Parser.DescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#syntax}.
	 * @param ctx the parse tree
	 */
	void enterSyntax(RFC4512Parser.SyntaxContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#syntax}.
	 * @param ctx the parse tree
	 */
	void exitSyntax(RFC4512Parser.SyntaxContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#singleValue}.
	 * @param ctx the parse tree
	 */
	void enterSingleValue(RFC4512Parser.SingleValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#singleValue}.
	 * @param ctx the parse tree
	 */
	void exitSingleValue(RFC4512Parser.SingleValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#immutability}.
	 * @param ctx the parse tree
	 */
	void enterImmutability(RFC4512Parser.ImmutabilityContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#immutability}.
	 * @param ctx the parse tree
	 */
	void exitImmutability(RFC4512Parser.ImmutabilityContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#collective}.
	 * @param ctx the parse tree
	 */
	void enterCollective(RFC4512Parser.CollectiveContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#collective}.
	 * @param ctx the parse tree
	 */
	void exitCollective(RFC4512Parser.CollectiveContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#obsolescence}.
	 * @param ctx the parse tree
	 */
	void enterObsolescence(RFC4512Parser.ObsolescenceContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#obsolescence}.
	 * @param ctx the parse tree
	 */
	void exitObsolescence(RFC4512Parser.ObsolescenceContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#minUpperBounds}.
	 * @param ctx the parse tree
	 */
	void enterMinUpperBounds(RFC4512Parser.MinUpperBoundsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#minUpperBounds}.
	 * @param ctx the parse tree
	 */
	void exitMinUpperBounds(RFC4512Parser.MinUpperBoundsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#structureRule}.
	 * @param ctx the parse tree
	 */
	void enterStructureRule(RFC4512Parser.StructureRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#structureRule}.
	 * @param ctx the parse tree
	 */
	void exitStructureRule(RFC4512Parser.StructureRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#usages}.
	 * @param ctx the parse tree
	 */
	void enterUsages(RFC4512Parser.UsagesContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#usages}.
	 * @param ctx the parse tree
	 */
	void exitUsages(RFC4512Parser.UsagesContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#quotedDescriptor}.
	 * @param ctx the parse tree
	 */
	void enterQuotedDescriptor(RFC4512Parser.QuotedDescriptorContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#quotedDescriptor}.
	 * @param ctx the parse tree
	 */
	void exitQuotedDescriptor(RFC4512Parser.QuotedDescriptorContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#quotedString}.
	 * @param ctx the parse tree
	 */
	void enterQuotedString(RFC4512Parser.QuotedStringContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#quotedString}.
	 * @param ctx the parse tree
	 */
	void exitQuotedString(RFC4512Parser.QuotedStringContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#schemaDN}.
	 * @param ctx the parse tree
	 */
	void enterSchemaDN(RFC4512Parser.SchemaDNContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#schemaDN}.
	 * @param ctx the parse tree
	 */
	void exitSchemaDN(RFC4512Parser.SchemaDNContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#objectIdentifierOIDOrName}.
	 * @param ctx the parse tree
	 */
	void enterObjectIdentifierOIDOrName(RFC4512Parser.ObjectIdentifierOIDOrNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#objectIdentifierOIDOrName}.
	 * @param ctx the parse tree
	 */
	void exitObjectIdentifierOIDOrName(RFC4512Parser.ObjectIdentifierOIDOrNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#attributeDescription}.
	 * @param ctx the parse tree
	 */
	void enterAttributeDescription(RFC4512Parser.AttributeDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#attributeDescription}.
	 * @param ctx the parse tree
	 */
	void exitAttributeDescription(RFC4512Parser.AttributeDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#openParen}.
	 * @param ctx the parse tree
	 */
	void enterOpenParen(RFC4512Parser.OpenParenContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#openParen}.
	 * @param ctx the parse tree
	 */
	void exitOpenParen(RFC4512Parser.OpenParenContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#closeParen}.
	 * @param ctx the parse tree
	 */
	void enterCloseParen(RFC4512Parser.CloseParenContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#closeParen}.
	 * @param ctx the parse tree
	 */
	void exitCloseParen(RFC4512Parser.CloseParenContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#numericOIDOrMacro}.
	 * @param ctx the parse tree
	 */
	void enterNumericOIDOrMacro(RFC4512Parser.NumericOIDOrMacroContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#numericOIDOrMacro}.
	 * @param ctx the parse tree
	 */
	void exitNumericOIDOrMacro(RFC4512Parser.NumericOIDOrMacroContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#macro}.
	 * @param ctx the parse tree
	 */
	void enterMacro(RFC4512Parser.MacroContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#macro}.
	 * @param ctx the parse tree
	 */
	void exitMacro(RFC4512Parser.MacroContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#numericOID}.
	 * @param ctx the parse tree
	 */
	void enterNumericOID(RFC4512Parser.NumericOIDContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#numericOID}.
	 * @param ctx the parse tree
	 */
	void exitNumericOID(RFC4512Parser.NumericOIDContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#macroSuffix}.
	 * @param ctx the parse tree
	 */
	void enterMacroSuffix(RFC4512Parser.MacroSuffixContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#macroSuffix}.
	 * @param ctx the parse tree
	 */
	void exitMacroSuffix(RFC4512Parser.MacroSuffixContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#number}.
	 * @param ctx the parse tree
	 */
	void enterNumber(RFC4512Parser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#number}.
	 * @param ctx the parse tree
	 */
	void exitNumber(RFC4512Parser.NumberContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#attrOptions}.
	 * @param ctx the parse tree
	 */
	void enterAttrOptions(RFC4512Parser.AttrOptionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#attrOptions}.
	 * @param ctx the parse tree
	 */
	void exitAttrOptions(RFC4512Parser.AttrOptionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#userApplications}.
	 * @param ctx the parse tree
	 */
	void enterUserApplications(RFC4512Parser.UserApplicationsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#userApplications}.
	 * @param ctx the parse tree
	 */
	void exitUserApplications(RFC4512Parser.UserApplicationsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#directoryOperation}.
	 * @param ctx the parse tree
	 */
	void enterDirectoryOperation(RFC4512Parser.DirectoryOperationContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#directoryOperation}.
	 * @param ctx the parse tree
	 */
	void exitDirectoryOperation(RFC4512Parser.DirectoryOperationContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#distributedOperation}.
	 * @param ctx the parse tree
	 */
	void enterDistributedOperation(RFC4512Parser.DistributedOperationContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#distributedOperation}.
	 * @param ctx the parse tree
	 */
	void exitDistributedOperation(RFC4512Parser.DistributedOperationContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#dSAOperation}.
	 * @param ctx the parse tree
	 */
	void enterDSAOperation(RFC4512Parser.DSAOperationContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#dSAOperation}.
	 * @param ctx the parse tree
	 */
	void exitDSAOperation(RFC4512Parser.DSAOperationContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#structureRules}.
	 * @param ctx the parse tree
	 */
	void enterStructureRules(RFC4512Parser.StructureRulesContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#structureRules}.
	 * @param ctx the parse tree
	 */
	void exitStructureRules(RFC4512Parser.StructureRulesContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#oID}.
	 * @param ctx the parse tree
	 */
	void enterOID(RFC4512Parser.OIDContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#oID}.
	 * @param ctx the parse tree
	 */
	void exitOID(RFC4512Parser.OIDContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#oIDs}.
	 * @param ctx the parse tree
	 */
	void enterOIDs(RFC4512Parser.OIDsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#oIDs}.
	 * @param ctx the parse tree
	 */
	void exitOIDs(RFC4512Parser.OIDsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#extensionValue}.
	 * @param ctx the parse tree
	 */
	void enterExtensionValue(RFC4512Parser.ExtensionValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#extensionValue}.
	 * @param ctx the parse tree
	 */
	void exitExtensionValue(RFC4512Parser.ExtensionValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#extensions}.
	 * @param ctx the parse tree
	 */
	void enterExtensions(RFC4512Parser.ExtensionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#extensions}.
	 * @param ctx the parse tree
	 */
	void exitExtensions(RFC4512Parser.ExtensionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#extension}.
	 * @param ctx the parse tree
	 */
	void enterExtension(RFC4512Parser.ExtensionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#extension}.
	 * @param ctx the parse tree
	 */
	void exitExtension(RFC4512Parser.ExtensionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#names}.
	 * @param ctx the parse tree
	 */
	void enterNames(RFC4512Parser.NamesContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#names}.
	 * @param ctx the parse tree
	 */
	void exitNames(RFC4512Parser.NamesContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#descriptor}.
	 * @param ctx the parse tree
	 */
	void enterDescriptor(RFC4512Parser.DescriptorContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#descriptor}.
	 * @param ctx the parse tree
	 */
	void exitDescriptor(RFC4512Parser.DescriptorContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#objectClassDescription}.
	 * @param ctx the parse tree
	 */
	void enterObjectClassDescription(RFC4512Parser.ObjectClassDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#objectClassDescription}.
	 * @param ctx the parse tree
	 */
	void exitObjectClassDescription(RFC4512Parser.ObjectClassDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#attributeTypeDescription}.
	 * @param ctx the parse tree
	 */
	void enterAttributeTypeDescription(RFC4512Parser.AttributeTypeDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#attributeTypeDescription}.
	 * @param ctx the parse tree
	 */
	void exitAttributeTypeDescription(RFC4512Parser.AttributeTypeDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#matchingRuleDescription}.
	 * @param ctx the parse tree
	 */
	void enterMatchingRuleDescription(RFC4512Parser.MatchingRuleDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#matchingRuleDescription}.
	 * @param ctx the parse tree
	 */
	void exitMatchingRuleDescription(RFC4512Parser.MatchingRuleDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#matchingRuleUseDescription}.
	 * @param ctx the parse tree
	 */
	void enterMatchingRuleUseDescription(RFC4512Parser.MatchingRuleUseDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#matchingRuleUseDescription}.
	 * @param ctx the parse tree
	 */
	void exitMatchingRuleUseDescription(RFC4512Parser.MatchingRuleUseDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#lDAPSyntaxDescription}.
	 * @param ctx the parse tree
	 */
	void enterLDAPSyntaxDescription(RFC4512Parser.LDAPSyntaxDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#lDAPSyntaxDescription}.
	 * @param ctx the parse tree
	 */
	void exitLDAPSyntaxDescription(RFC4512Parser.LDAPSyntaxDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#dITContentRuleDescription}.
	 * @param ctx the parse tree
	 */
	void enterDITContentRuleDescription(RFC4512Parser.DITContentRuleDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#dITContentRuleDescription}.
	 * @param ctx the parse tree
	 */
	void exitDITContentRuleDescription(RFC4512Parser.DITContentRuleDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#dITStructureRuleDescription}.
	 * @param ctx the parse tree
	 */
	void enterDITStructureRuleDescription(RFC4512Parser.DITStructureRuleDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#dITStructureRuleDescription}.
	 * @param ctx the parse tree
	 */
	void exitDITStructureRuleDescription(RFC4512Parser.DITStructureRuleDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link RFC4512Parser#nameFormDescription}.
	 * @param ctx the parse tree
	 */
	void enterNameFormDescription(RFC4512Parser.NameFormDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link RFC4512Parser#nameFormDescription}.
	 * @param ctx the parse tree
	 */
	void exitNameFormDescription(RFC4512Parser.NameFormDescriptionContext ctx);
}