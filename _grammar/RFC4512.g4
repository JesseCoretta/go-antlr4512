/*
RFC4512 - LEXER & PARSER GRAMMAR

MIT LICENSE

Copyright (c) 2024 Jesse Coretta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

// --------------------------------------------------------------------

ABOUT THIS FILE

This file, RFC4512.g4, contains textual definitions of various descriptions
and constructs found within RFC4512.

SECTION 2.5

Section 2.5 constructs are implemented for attributeDescription and options
support for parameters such as "lang-de", "binary" and many others.

SECTION 4

Section 4 is implemented for complete schema parsing at multiple levels.
The following grammar rules (and all requisite subcomponents) are available:

    - objectClassDescription(s)
    - attributeTypeDescription(s)
    - matchingRuleDescription(s)
    - matchingRuleUseDescription(s)
    - lDAPSyntaxDescription(s)
    - dITContentRuleDescription(s)
    - dITStructureRuleDescription(s)
    - nameFormDescription(s)

An entire schema file may be parsed using the 'parse' top-level rule.  As
an alternative, individual subcomponents such as definitions or even portions
of definitions (e.g.: "NAME") may be parsed in an ad hoc manner.

While this grammar file was written with strict RFC4512 compliance in mind, this
grammar field was written to be extremely forgiving in terms of the presence of
NEWLINE characters.  For instance:

   ( numericOIDOrMacro
	FIELD ...
	FIELD ...
	FIELD ... )

... is the same as ...

   ( numericOID FIELD ... FIELD ... FIELD ... )

Additionally, Bash-style line comments (# ...) are permitted.  For instance:

   ## this is <definition> ...
   ( numericOID		## {iso(1) identified-organization(3) ...}
   	## FIELD comment here
   	FIELD ...	## or over here
   	FIELD ...	## etc etc
   	FIELD ... )	## other comment

This allows administrators more flexibility in terms of commenting on
individual definitions with respect to their components.  Good schema
annotation makes for fewer clerical errors.

The only downside of this grammar implementation are the inherently
deep cyclomatics that are present in the ANTLR-generated codebase.

Adopters may wish to severely modify this grammar file -- for example,
to enforce strict formatting rules without newlines -- and regenerate
the codebase in the preferred language if cyclomatic complexity is
a problem.
*/

// BEGIN GRAMMAR ------------------------------------------------------

grammar RFC4512;

// Parse is used from a top-level (file) context, parsing any number
// of definitions.
fileparse	: NEWLINE? LineComment? definitions EOF ;

// Definitions describes a sequence of any type of definition and,
// optionally, comments.
definitions	: (LineComment
		  |lDAPSyntaxDescriptions
		  |matchingRuleDescriptions
		  |attributeTypeDescriptions
		  |matchingRuleUseDescriptions
		  |objectClassDescriptions
		  |dITContentRuleDescriptions
		  |nameFormDescriptions
		  |dITStructureRuleDescriptions
		  )+
		;

// objectClassDescriptions is a sequence of one (1) or more
// objectClassDescription instances separated by newlines.
objectClassDescriptions	: (objectClassLabel objectClassDescription
				SP* (NEWLINE|EOF))+
			;
objectClassLabel	: OCLabel definitionLabelDelim			  ;

// objectClassDescription conforms to section 4.1.1 of RFC4512.
objectClassDescription          : openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (SP+ 'NAME' SP+ definitionName SP* LineComment? NEWLINE?)?
                                        (SP+ 'DESC' SP+ definitionDescription SP* LineComment? NEWLINE?)?
                                        (SP+ definitionObsolete SP* LineComment? NEWLINE?)?
					(SP+ 'SUP' SP+ oCSuperClasses SP* LineComment? NEWLINE?)?
                                        (SP+ oCKind SP* LineComment? NEWLINE?)?
					(SP+ 'MUST' SP+ definitionMust SP* LineComment? NEWLINE?)?
					(SP+ 'MAY' SP+ definitionMay SP* LineComment? NEWLINE?)?
                                        definitionExtensions? SP* LineComment? NEWLINE?
						SP* LineComment? closeParen LineComment?
                                ;

// oCKind implements the "kind" of object class. This concept is covered
// in Section 2.4.
oCKind				: (structuralKind|auxiliaryKind|abstractKind)	;
abstractKind			: 'ABSTRACT'					;
structuralKind			: 'STRUCTURAL'					;
auxiliaryKind			: 'AUXILIARY'					;

// oCSuperClasses defines one (1) or more superior object classes of the
// bearing object class.
oCSuperClasses			: (oID|oIDs)					;

// attributeTypeDescriptions is a sequence of one (1) or more
// attributeTypeDescription instances separated by newlines.
attributeTypeDescriptions: (attributeTypeLabel attributeTypeDescription
				SP* (NEWLINE|EOF))+
			 ;
attributeTypeLabel	 : ATLabel definitionLabelDelim			  ;

// attributeTypeDescription conforms to section 4.1.2 of RFC4512.
attributeTypeDescription        : openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (SP+ 'NAME' SP+ definitionName SP* LineComment? NEWLINE?)?
                                        (SP+ 'DESC' SP+ definitionDescription SP* LineComment? NEWLINE?)?
                                        (SP+ definitionObsolete SP* LineComment? NEWLINE?)?
                                        (SP+ 'SUP' SP+ aTSuperType SP* LineComment? NEWLINE?)?
                                        (SP+ 'EQUALITY' SP+ aTEquality SP* LineComment? NEWLINE?)?
                                        (SP+ 'ORDERING' SP+ aTOrdering SP* LineComment? NEWLINE?)?
                                        (SP+ 'SUBSTR''ING'? SP+ aTSubstring SP* LineComment? NEWLINE?)?
                                        (SP+ 'SYNTAX' SP+ definitionSyntax minimumUpperBounds? SP* LineComment? NEWLINE?)?
                                        (SP+ (aTSingleValue|aTCollective) SP* LineComment? NEWLINE?)?
                                        (SP+ aTNoUserModification SP* LineComment? NEWLINE?)?
                                        (SP+ 'USAGE' SP+ aTUsage SP* LineComment? NEWLINE?)?
                                        definitionExtensions? SP* LineComment? NEWLINE?
						SP* LineComment? closeParen LineComment?
                                ;

// minimumUpperBounds defines the maximum length of the value governed
// by the given SYNTAX.
minimumUpperBounds		: MinUpperBounds			;

// aTEquality defines the equality matchingRule of the attributeType.
aTEquality			: oID					;

// aTOrdering defines the ordering matchingRule of the attributeType.
aTOrdering			: oID					;

// aTSubstring defines the substring matchingRule of the attributeType.
aTSubstring			: oID					;

// aTSuperType defines the super type of the attributeType.
aTSuperType			: oID					;

// aTUsage defines the usage of the attributeType.  See Usage lexer rule.
aTUsage				: userApplication
				| directoryOperation
				| distributedOperation
				| dSAOperation
				;
userApplication			: 'userApplication'			;
directoryOperation		: 'directoryOperation'			;
distributedOperation		: 'distributedOperation'		;
dSAOperation			: 'dSAOperation'			;

// aTSingleValue, if present, indicates the attributeType is single-valued.
// Single-valued attributes CANNOT be collective.
aTSingleValue			: 'SINGLE-VALUE'			;

// aTCollective, if present, indicates the attributeType is collective.
// Collective attributes CANNOT be single-valued.
aTCollective			: 'COLLECTIVE'				;

// aTNoUserModification, if present, indicates values of this attributeType
// cannot be set nor modified by users.
aTNoUserModification		: 'NO-USER-MODIFICATION'		;

// matchingRuleDescriptions is a sequence of one (1) or more
// matchingRuleDescription instances separated by newlines.
matchingRuleDescriptions: (matchingRuleLabel matchingRuleDescription
				SP* (NEWLINE|EOF))+
			;
matchingRuleLabel	: MRLabel definitionLabelDelim			  ;

// matchingRuleDescription conforms to section 4.1.3 of RFC4512.
matchingRuleDescription		: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (SP+ 'NAME' SP+ definitionName SP* LineComment? NEWLINE?)?
                                        (SP+ 'DESC' SP+ definitionDescription SP* LineComment? NEWLINE?)?
                                        (SP+ definitionObsolete SP* LineComment? NEWLINE?)?
                                        (SP+ 'SYNTAX' SP+ definitionSyntax SP* LineComment? NEWLINE?)?
                                        definitionExtensions? SP* LineComment? NEWLINE?
						SP* LineComment? closeParen LineComment?
				;

// matchingRuleUseDescriptions is a sequence of one (1) or more
// matchingRuleUseDescription instances separated by newlines.
matchingRuleUseDescriptions: (matchingRuleUseLabel matchingRuleUseDescription
				SP* (NEWLINE|EOF))+
			   ;
matchingRuleUseLabel	   : MULabel definitionLabelDelim			  ;

// matchingRuleUseDescription conforms to section 4.1.4 of RFC4512.
// Note that its numericOID is NOT unique, but rather is identical to
// the numericOID held by the foundation matchingRule instance upon
// upon which a matchingRuleUse is based.
matchingRuleUseDescription	: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (SP+ 'NAME' SP+ definitionName SP* LineComment? NEWLINE?)?
                                        (SP+ 'DESC' SP+ definitionDescription SP* LineComment? NEWLINE?)?
                                        (SP+ definitionObsolete SP* LineComment? NEWLINE?)?
					(SP+ 'APPLIES' SP+ mRUApplies SP* LineComment? NEWLINE?)?
                                        definitionExtensions? SP* LineComment? NEWLINE?
						SP* LineComment? closeParen LineComment?
				;

// mRUApplies contains attributeType OID instances which identify users of the
// matchingRule in question.
mRUApplies			: (oID|oIDs)					;

// lDAPSyntaxDescriptions is a sequence of one (1) or more
// lDAPSyntaxDescription instances separated by newlines.
lDAPSyntaxDescriptions	: (lDAPSyntaxLabel lDAPSyntaxDescription SP* (NEWLINE|EOF))+ ;
lDAPSyntaxLabel		: LSLabel definitionLabelDelim				;

// lDAPSyntaxDescription conforms to section 4.1.5 of RFC4512.
lDAPSyntaxDescription		: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (SP+ 'DESC' SP+ definitionDescription SP* LineComment? NEWLINE?)?
                                        definitionExtensions? SP* LineComment? NEWLINE?
						SP* LineComment? closeParen LineComment?
				;

// dITContentRuleDescriptions is a sequence of one (1) or more
// dITContentRuleDescription instances separated by newlines.
dITContentRuleDescriptions: (dITContentRuleLabel dITContentRuleDescription
				 SP* (NEWLINE|EOF))+
			  ;
dITContentRuleLabel	  : DCLabel definitionLabelDelim		;

// dITContentRuleDescription conforms to section 4.1.6 of RFC4512.
// Note that its numericOID is NOT unique, but rather is identical
// to the numericOID held by the foundation (structural) objectClass
// instance upon which a dITContentRule is based.
dITContentRuleDescription	: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (SP+ 'NAME' SP+ definitionName SP* LineComment? NEWLINE?)?
					(SP+ 'DESC' SP+ definitionDescription SP* LineComment? NEWLINE?)?
                                        (SP+ definitionObsolete SP* LineComment? NEWLINE?)?
					(SP+ 'AUX' SP+ dCRAux SP* LineComment? NEWLINE?)?
					(SP+ 'MUST' SP+ definitionMust SP* LineComment? NEWLINE?)?
					(SP+ 'MAY' SP+ definitionMay SP* LineComment? NEWLINE?)?
					(SP+ 'NOT' SP+ dCRNot SP* LineComment? NEWLINE?)?
					definitionExtensions? SP* LineComment?
						NEWLINE? SP* LineComment? closeParen
				;

// dCRNot is a sequence of one (1) or more attributeType OIDs which the
// dITContentRule forbids being present for the applicable DIT entry.
dCRNot				: (oID|oIDs)				;

// dCRAux is a sequence of one (1) or more auxiliary objectClasses which
// the dITContentRule allows to be present upon qualifying DIT entries.
dCRAux				: (oID|oIDs)				;

// dITStructureRuleDescriptions is a sequence of one (1) or more
// dITStructureRuleDescription instances separated by newlines.
dITStructureRuleDescriptions: (dITStructureRuleLabel dITStructureRuleDescription
				SP* (NEWLINE|EOF))+
			    ;
dITStructureRuleLabel	    : DSLabel definitionLabelDelim		;

// dITStructureRuleDescription conforms to section 4.1.7.1 of RFC4512.
dITStructureRuleDescription	: openParen SP* LineComment? NEWLINE?
                                        SP* structureRule SP* LineComment? NEWLINE?
                                        (SP+ 'NAME' SP+ definitionName SP* LineComment? NEWLINE?)?
                                        (SP+ 'DESC' SP+ definitionDescription SP* LineComment? NEWLINE?)?
                                        (SP+ definitionObsolete SP* LineComment? NEWLINE?)?
					SP+ 'FORM' SP+ dSRForm SP* LineComment? NEWLINE?
					(SP+ 'SUP' SP+ dSRSuperRules SP* LineComment? NEWLINE?)?
                                        definitionExtensions? SP* LineComment? NEWLINE?
						SP* LineComment? closeParen LineComment?
                                ;

// structureRules is a sequence of one (1) or more structureRule instances
// enclosed within opening and closing parenthesis and delimited via space.
structureRules			: openParen SP* LineComment? NEWLINE? SP* LineComment?
					structureRule SP* LineComment? NEWLINE?
					( SP+ structureRule )* SP* LineComment?
					NEWLINE? SP* closeParen SP* LineComment?
				;

// dSRSuperRules defines one (1) or more structureRule instances to which
// the bearing dITStructureRule is subordinate.
dSRSuperRules			: (structureRule|structureRules)	;

// structuralRule defines the numeric rule number for the bearing
// dITStructureRule instance.  Instances of this type are also
// used in super rule contexts.
structureRule			: number				;

// dSRForm defines the OID of the associated nameForm instance of the
// bearing dITStructureRule.
dSRForm				: oID				;

// nameFormDescriptions is a sequence of one (1) or more
// nameFormDescription instances separated by newlines.
nameFormDescriptions	: (nameFormLabel nameFormDescription SP* (NEWLINE|EOF))+ ;
nameFormLabel		: NFLabel definitionLabelDelim				 ;

// nameFormDescription conforms to section 4.1.7.2
nameFormDescription		: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (SP+ 'NAME' SP+ definitionName SP* LineComment? NEWLINE?)?
                                        (SP+ 'DESC' SP+ definitionDescription SP* LineComment? NEWLINE?)?
                                        (SP+ definitionObsolete SP* LineComment? NEWLINE?)?
					SP+ 'OC' SP+ nFStructuralOC SP* LineComment? NEWLINE?
					SP+ 'MUST' SP+ definitionMust SP* LineComment? NEWLINE?
					(SP+ 'MAY' SP+ definitionMay SP* LineComment? NEWLINE?)?
                                        definitionExtensions? SP* LineComment? NEWLINE?
						SP* LineComment? closeParen LineComment?
                                ;

// nFStructuralOC defines the OID of the structural objectClass held by
// entries whose DNs are controlled by way of a dITStructureRule that
// bears a nameForm.
nFStructuralOC		: oID					;

// parser rules that are found in more than
// one schema definition type.
definitionObsolete	: 'OBSOLETE'					;

// definitionSyntax is used by attributeType and matchingRule definition
// instances.  It defines the numericOID of the LDAPSyntax associated
// with the definition.
definitionSyntax	: numericOID					;

// definitionMust is used by objectClass, dITContentRule and nameForm
// definitions. It is used to define the required attributeTypes for
// use in entries under the governance of the given definition.
definitionMust		: (oID|oIDs)					;

// definitionMay is used by objectClass, dITContentRule and nameForm
// definitions. It is used to define the allowed attributeTypes for
// use in entries under the governance of the given definition.
definitionMay		: (oID|oIDs)					;

// definitionDescription is optionally used in any definition type to
// further describe the nature and use of the definition.
definitionDescription	: (QuotedString|QuotedDescriptor)		;

// definitionName contains one (1) or more quoted descriptors, each being
// a declared name by which the bearing definition is known.
definitionName		: (QuotedDescriptor|
				openParen SP* LineComment? NEWLINE? SP*
				LineComment? QuotedDescriptor SP* LineComment? NEWLINE?
				( SP+ QuotedDescriptor )* SP* LineComment? NEWLINE?
				SP* closeParen SP* LineComment? )
			;

// definitionExtensions is the plural form of definitionExtension.
definitionExtensions	: definitionExtension+				;

// definitionExtension is a single RFC4512 definition "extension", as
// described in ABNF terms within the opening of Section 4.1.
definitionExtension	: NEWLINE? ExtensionName (extensionValue|                       
                                openParen SP* LineComment? NEWLINE? SP*
				LineComment? extensionValue SP* LineComment?
				NEWLINE? ( SP+ extensionValue )* SP*
				LineComment? NEWLINE? SP* closeParen SP*
				LineComment? )
                        ; 

definitionLabelDelim	: ((SP*':'SP*|SP*'='SP*)|SP+)			;

// oIDs is a sequence of one (1) or more oID instances enclosed within
// opening and closing parenthesis and delimited via ODelim.
oIDs			: openParen SP* LineComment? NEWLINE? SP* oID SP*
				LineComment? NEWLINE? ( ODelim SP* oID )*
				SP* LineComment? NEWLINE? SP* LineComment?
				closeParen SP* LineComment?
			;

// attributeDescription implements section 2.5.
attributeDescription	: oID attrOptions?				;

// oID is either a descriptor (e.g.: cn) OR a numeric OID (e.g.: 2.5.4.3).
oID		    	: descriptor
			| numericOID
			;

// openParen is used to mark the beginning of parenthetical encapsulation
// for certain expressions or values, as well as an entire definition
// itself.  All encapsulations MUST terminate at some point using the
// closeParen rule.
openParen		: OpenParenthesis				;

// closeParen is used to mark the completion of parenthetical encapsulation.
closeParen		: CloseParenthesis				;

// numericOIDOrMacro allows use of numeric OIDs or numeric OID macros.
// This is supported for all definition types except dITStructureRules.
numericOIDOrMacro	: macro
			| numericOID
			;

// macro offers a means to refer to numeric OIDs by friendlier names.
// A macro is a descriptor followed by a macroSuffix.  For an example
// of this concept in the wild, see the schema definitions in RFC 2307
// and also see OpenLDAP's 'objectIdentifier' configuration parameter.
macro			: descriptor macroSuffix			;

// numericOID is the numeric object identifier form, which consists of
// unsigned integers (mostly of any magnitude), each delimited by dots.
numericOID		: NumOID					;

// macroSuffix is the (partial) OID which a macro's descriptor references
macroSuffix		: MacroSuffix					;

number			: Digit						;

// attrOptions is a sequence attribute option, per section 2.5
attrOptions		: ( ';' descriptor )+				;

// descriptor is a "name", such as that held by an attributeType or
// matchingRule.  The name must conform to: ALPHA ('-'? (ALPHA|DIGIT)+)*
// as shown in the 'descr' ABNF production of RFC 4512.
//
// The alternative "special keywords" are used as a workaround to migitate
// ANTLR's tendency to interpolate tokens rather than accepting them as
// coincidental literal values.  One example of this is the OCLabel, which
// is the literal string 'objectClass' that is used in more than one context
// in any given schema.
descriptor		: Descr
			| NFLabel
			| LSLabel
			| MRLabel
			| MULabel
			| ATLabel
			| OCLabel
			| DCLabel
			| DSLabel
			;

// extensionValue allows a quotedString.  Note QuotedDescriptor is
// included only to mitigate the overlapping nature of the two forms
// when using ANTLR to avoid errors.
extensionValue		: (QuotedString|QuotedDescriptor)		;

////////////////////////////////////////
// LEXERS

// RULES ///////////////////////////////

OpenParenthesis		: '(' SP*					;
CloseParenthesis	: SP* ')'					;
ODelim			: NEWLINE? SP* '$' SP* NEWLINE?			;
MacroSuffix		: ( ':' | '.' ) Digit ( '.' Digit )*		;
NumOID			: ('0'|'1'|'2') ( '.' Digit )+  		;
ExtensionName		: SP+ 'X-' ALPHA (('-'|'_')? ALPHA+)* SP+	;
QuotedDescriptor	: '\'' Descr '\''				;
QuotedString		: '\'' (Descr|DSTR+) '\''			;
MinUpperBounds		: '{' Digit '}'					;
Digit			: [0-9]
		        | [1-9][0-9]+
                        ;
NFLabel			: 'name'[fF]'orm''s'?				;
OCLabel			: 'object'[cC]'lass''es'?			;
ATLabel			: 'attribute'[tT]'ype''s'?			;
LSLabel			: 'l'[dD][aA][pP][sS]'yntax''es'?		;
MRLabel			: 'matching'[rR]'ule''s'?			;
MULabel			: 'matching'[rR]'ule'[uU]'se''s'?		;
DSLabel			: 'd'[iI][tT][sS]'tructure'[rR]'ule''s'?	;
DCLabel			: 'd'[iI][tT][cC]'ontent'[rR]'ule''s'?		;
Descr			: ALPHA ( '-'? (ALPHA|[0-9])+ )*		;
SP			: [ \t]						;

// FRAGMENTS ///////////////////////////

fragment DSTR		: ( QS | QQ | QUTF8 )				;
fragment ALPHA		: ( 'a'..'z' | 'A'..'Z' )			;

fragment QUTF8		: ( QUTF1 | UTFMB )				;
fragment UTFMB		: ( UTF2 | UTF3 | UTF4 )			;

fragment UTF4		: '\u00f0' ( '\u0090' .. '\u00bf' ) UTF0 UTF0
			| ( '\u00f1' .. '\u00f3' ) UTF0 UTF0 UTF0
			| '\u00f4' ( '\u0080' .. '\u008f' ) UTF0 UTF0
			;

fragment UTF3		: '\u00e0' ( '\u00a0' .. '\u00bf' ) UTF0
			| ( '\u00e1' .. '\u00ec' ) UTF0 UTF0
			| '\u00ed' ( '\u0080' .. '\u009f' ) UTF0
			| ( '\u00ee' .. '\u00ef' ) UTF0 UTF0
			;

fragment UTF2		: ( '\u00c2' .. '\u00df' ) UTF0
			;

fragment QUTF1		: ( '\u0000' .. '\u0026' )
                        | ( '\u0028' .. '\u005b' )
                        | ( '\u005d' .. '\u007f' )
			;

fragment QQ		: '\u005c' '\u0032' '\u0037'			; // ESC \27
fragment QS		: '\u005c' '\u0035' ('\u0043'|'\u0063')		; // ESC 5[C|c]

fragment UTF0		: ( '\u0080' .. '\u00bf' )			; // UTF 0 production

NEWLINE			: '\n'+						;

LineComment		: '#' ~[\r\n]* [\r\n]? -> skip			;
