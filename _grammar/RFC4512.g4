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

/*
fileparse is used from a top-level (file) context,
parsing any number of definitions.
*/
fileparse	: NEWLINE? LineComment? schemaDN? definitions EOF ;

definitions	: (LineComment
		  |objectIdentifier
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

objectClassDescriptions	    : (objectClassLabel objectClassDescription SP* (NEWLINE|EOF))+	     ;
attributeTypeDescriptions   : (attributeTypeLabel attributeTypeDescription SP* (NEWLINE|EOF))+	     ;
matchingRuleDescriptions    : (matchingRuleLabel matchingRuleDescription SP* (NEWLINE|EOF))+	     ;
matchingRuleUseDescriptions : (matchingRuleUseLabel matchingRuleUseDescription SP* (NEWLINE|EOF))+   ;
lDAPSyntaxDescriptions	    : (lDAPSyntaxLabel lDAPSyntaxDescription SP* (NEWLINE|EOF))+ 	     ;
dITStructureRuleDescriptions: (dITStructureRuleLabel dITStructureRuleDescription SP* (NEWLINE|EOF))+ ;
nameFormDescriptions	    : (nameFormLabel nameFormDescription SP* (NEWLINE|EOF))+		     ;
dITContentRuleDescriptions  : (dITContentRuleLabel dITContentRuleDescription SP* (NEWLINE|EOF))+     ;
objectIdentifier	    : ('objectidentifier' SP+ objectIdentifierOIDOrName SP* (NEWLINE|EOF))+  ;

objectClassLabel	    	: OCLabel definitionLabelDelim				;
matchingRuleLabel	    	: MRLabel definitionLabelDelim				;
matchingRuleUseLabel	    	: MULabel definitionLabelDelim				;
lDAPSyntaxLabel		    	: LSLabel definitionLabelDelim				;
dITContentRuleLabel	    	: DCLabel definitionLabelDelim				;
dITStructureRuleLabel	    	: DSLabel definitionLabelDelim				;
attributeTypeLabel	    	: ATLabel definitionLabelDelim				;
nameFormLabel		    	: NFLabel definitionLabelDelim				;
definitionLabelDelim	    	: ((SP*':'SP*|SP*'='SP*)|SP+) 				;

kind				: SP+ (structuralKind|auxiliaryKind|abstractKind)	;
abstractKind			: 'ABSTRACT'						;
structuralKind			: 'STRUCTURAL'						;
auxiliaryKind			: 'AUXILIARY'						;

applies				: SP+ 'APPLIES' SP+ (oID|oIDs)				;
must				: SP+ 'MUST' SP+ (oID|oIDs)				;
aux				: SP+ 'AUX'  SP+ (oID|oIDs)				;
may				: SP+ 'MAY'  SP+ (oID|oIDs)				;
not				: SP+ 'NOT'  SP+ (oID|oIDs)				;

form				: SP+ 'FORM'     SP+ oID				;
ordering			: SP+ 'ORDERING' SP+ oID				;
equality			: SP+ 'EQUALITY' SP+ oID				;
structuralClass			: SP+ 'OC'       SP+ oID				;
superType			: SP+ 'SUP'      SP+ oID				;
superClasses			: SP+ 'SUP'      SP+ (oID|oIDs)				;
superRules			: SP+ 'SUP'      SP+ (structureRule|structureRules)	;
substr				: SP+ 'SUBSTR'('ING''S'?)? SP+ oID			;

name				: SP+ 'NAME' SP+ names					;
usage				: SP+ 'USAGE' SP+ usages				;
description			: SP+ 'DESC' SP+ (quotedString|quotedDescriptor)	;
syntax				: SP+ 'SYNTAX' SP+ (numericOID|quotedDescriptor)	;
singleValue			: SP+ 'SINGLE-VALUE'					;
immutability			: SP+ 'NO-USER-MODIFICATION'				;
collective			: SP+ 'COLLECTIVE'					;
obsolescence			: SP+ 'OBSOLETE'					;
minUpperBounds			: MUB							;
structureRule			: number						;

usages				: userApplications
				| directoryOperation
				| distributedOperation
				| dSAOperation
				;

quotedDescriptor		: QDescr						;
quotedString			: QString						;
schemaDN			: DistinguishedName					;
objectIdentifierOIDOrName	: descriptor SP+ numericOIDOrMacro			;
attributeDescription		: oID attrOptions?					;
openParen			: OpenParenthesis					;
closeParen			: CloseParenthesis					;
numericOIDOrMacro		: (macro|numericOID)					;
macro				: descriptor macroSuffix				;
numericOID			: NumOID						;
macroSuffix			: MacroSuffix						;
number				: Digit							;

// attrOptions is a sequence attribute option, per section 2.5
attrOptions			: ( ';' descriptor )+					;

userApplications		: 'userApplications'					;
directoryOperation		: 'directoryOperation'					;
distributedOperation		: 'distributedOperation'				;
dSAOperation			: 'dSAOperation'					;

structureRules			: openParen SP* LineComment? NEWLINE? SP* LineComment?
					structureRule SP* LineComment? NEWLINE?
					( SP+ structureRule )* SP* LineComment?
					NEWLINE? SP* closeParen SP* LineComment?
				;

oID		    		: (descriptor|numericOID)				;
oIDs				: openParen SP* LineComment? NEWLINE? SP* oID SP*
					LineComment? NEWLINE? ( ODelim SP* oID )*
					SP* LineComment? NEWLINE? SP* LineComment?
					closeParen SP* LineComment?
				;

extensionValue			: (quotedString|quotedDescriptor)			;
extensions			: extension+						;
extension			: NEWLINE? XString (extensionValue|                       
                        	        openParen SP* LineComment? NEWLINE? SP*
					LineComment? extensionValue SP* LineComment?
					NEWLINE? ( SP+ extensionValue )* SP*
					LineComment? NEWLINE? SP* closeParen SP*
					LineComment? )
				;

names				: (quotedDescriptor|
					openParen SP* LineComment? NEWLINE? SP*
					LineComment? quotedDescriptor SP* LineComment? NEWLINE?
					( SP+ quotedDescriptor )* SP* LineComment? NEWLINE?
					SP* closeParen SP* LineComment? )
				;

// descriptor is a "name", such as that held by an attributeType or
// matchingRule.  The name must conform to: ALPHA ('-'? (ALPHA|DIGIT)+)*
// as shown in the 'descr' ABNF production of RFC 4512.
//
// The alternative "special keywords" are used as a workaround to migitate
// ANTLR's tendency to interpolate tokens rather than accepting them as
// coincidental literal values.  One example of this is the OCLabel, which
// is the literal string 'objectClass' that is used in more than one context
// in any given schema.
descriptor			: Descr
				| NFLabel
				| LSLabel
				| MRLabel
				| MULabel
				| ATLabel
				| OCLabel
				| DCLabel
				| DSLabel
				;

objectClassDescription          : openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (( name
					 | description
					 | obsolescence
					 | superClasses
					 | kind
					 | must
					 | may
					 ) SP* LineComment? NEWLINE?)*
                                        extensions? closeParen LineComment?
                                ;

attributeTypeDescription        : openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (( name
                                         | description
                                         | obsolescence
                                         | superType
                                         | equality
                                         | ordering
                                         | substr
                                         | syntax minUpperBounds?
                                         | (singleValue|collective)
                                         | immutability
                                         | usage
					 ) SP* LineComment? NEWLINE?)*
                                        extensions? closeParen LineComment?
				;

matchingRuleDescription		: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (( name
					 | description
					 | obsolescence
					 ) SP* LineComment? NEWLINE?)*
					syntax SP* LineComment? NEWLINE?
                                        extensions? closeParen LineComment?
				;

matchingRuleUseDescription	: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (( name
					 | description
					 | obsolescence
					 ) SP* LineComment? NEWLINE?)*
					applies SP* LineComment? NEWLINE?
                                        extensions? closeParen LineComment?
				;

lDAPSyntaxDescription		: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
					description? SP* LineComment? NEWLINE?
                                        extensions? NEWLINE? closeParen LineComment?
				;

dITContentRuleDescription	: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (( name
					 | description
					 | obsolescence
					 | aux
					 | must
					 | may
					 | not
					 ) SP* LineComment? NEWLINE?)*
                                        extensions? closeParen LineComment?
				;

dITStructureRuleDescription	: openParen SP* LineComment? NEWLINE?
                                        SP* structureRule SP* LineComment? NEWLINE?
                                        (( name
					 | description
					 | obsolescence
					 ) SP* LineComment? NEWLINE?)*
					(form SP* LineComment? NEWLINE?)
					(superRules SP* LineComment? NEWLINE?)?
                                        extensions? closeParen LineComment?
                                ;

nameFormDescription		: openParen SP* LineComment? NEWLINE?
                                        SP* numericOIDOrMacro SP* LineComment? NEWLINE?
                                        (( name
					 | description
					 | obsolescence
					 ) SP* LineComment? NEWLINE?)*
					(structuralClass SP* LineComment? NEWLINE?)
					(must SP* LineComment? NEWLINE?)
					(may SP* LineComment? NEWLINE?)?
                                        extensions? closeParen LineComment?
				;

////////////////////////////////////////
// LEXERS

// RULES ///////////////////////////////

OpenParenthesis		: '(' SP*					;
CloseParenthesis	: SP* ')'					;
ODelim			: NEWLINE? SP* '$' SP* NEWLINE?			;
MacroSuffix		: ( ':' | '.' ) Digit ( '.' Digit )*		;
NumOID			: ('0'|'1'|'2') ( '.' Digit )+  		;
XString			: SP+ 'X-' ALPHA (('-'|'_')? ALPHA+)* SP+	;
DistinguishedName	: 'dn:' SP* Descr '=' Descr SP* NEWLINE		; 
QDescr			: '\'' Descr '\''				;
QString			: '\'' (Descr|DSTR+) '\''			;
MUB			: '{' Digit '}'					;
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

fragment QQ		: '\u005c' '\u0027'				; // ESC \27
fragment QS		: '\u005c' '\u005c'				; // ESC 5[C|c]

fragment UTF0		: ( '\u0080' .. '\u00bf' )			; // UTF 0 production

NEWLINE			: '\n'+						;

LineComment		: '#' ~[\r\n]* [\r\n]? -> skip			;
