# go-antlr4512

[![ANTLR4](https://img.shields.io/badge/antlr-blue?label=%E2%92%B6&labelColor=gray&color=magenta&cacheSeconds=86400)](https://www.antlr.org/) [![RFC4512](https://img.shields.io/badge/RFC-4512-blue)](https://datatracker.ietf.org/doc/html/rfc4512) [![Reference](https://pkg.go.dev/badge/github.com/JesseCoretta/go-antlr4512.svg)](https://pkg.go.dev/github.com/JesseCoretta/go-antlr4512) [![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-antlr4512/blob/main/LICENSE) [![Issues](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-antlr4512/issues) [![Experimental](https://img.shields.io/badge/experimental-blue?logoColor=blue&label=%F0%9F%A7%AA%20%F0%9F%94%AC&labelColor=blue&color=gray)](https://github.com/JesseCoretta/JesseCoretta/blob/main/EXPERIMENTAL.md)

Package antlr4512 implements [Section 2.5](https://www.rfc-editor.org/rfc/rfc4512.html#section-2.5) and [Section 4 of RFC 4512](https://tools.ietf.org/html/rfc4512#section-4) by way of an unmodified ANTLR4-generated target.

## ANTLR

See the `_grammar/RFC4512.g4` file within this repository for ANTLR grammar. Using this file, the target was auto-generated using the following `antlr4` command:

```
$ antlr4 -no-listener -Dlanguage=Go -package antlr4512 RFC4512.g4
```

For details on ANTLR, click the above ANTLR badge. This page was generated using **antlr4-go** version 4.13.1.

If alterations are desired -- such as the introduction of bonafide listeners or visitors -- modify the above grammar file and/or regenerate code as needed. Remember it is NOT recommended for auto-generated code to be modified in any way. Code enhancements, convenience wrappers and other extensions devised by the adopter(s) should ONLY be made in separate files sharing the same package namespace.

## Features

  - Compliant
    - [RFC 4512](https://tools.ietf.org/html/rfc4512) syntax implementation honors relevant Augmented Backus-Naur Form (ABNF) productions
  - Compatible
    - 389DS and OpenLDAP style schema formats supported
    - Auto-resolving OpenLDAP `objectidentifier` macro support
    - Unresolvable macros accepted silently, allowing user remediation during the post-parsing stage
  - Flexible -- all of the following are supported:
    - Variable description labels 
    - Newline and hanging indents 
    - Line-terminating Bash comments
    - Optional clause ordering
  - Convenient
    - File and directory parsing functions
      - `ParseFile`
      - `ParseDirectory`
    - Helper functions for those with minimal ANTLR experience:
      - `ParseLDAPSyntax`
      - `ParseMatchingRule`
      - `ParseMatchingRuleUse`
      - `ParseAttributeType`
      - `ParseObjectClass`
      - `ParseDITContentRule`
      - `ParseNameForm`
      - `ParseDITStructureRule`
    - Extra methods
      - `Schema.UpdateMatchingRuleUses`, for simple auto-generation of `matchingRuleUse` instances
  - Simple, easily understood type instances
    - Struct field values are confined to primitive and basic map instances
    - No unnecessary pointers, no embedding, no private fields, no type assertion
    - Just parse your instance and go nuts
    - Great for templating and external processing

## Cyclomatics Notice

Please note that due to the complexity and flexibility of the grammar design (and some questionable [`antlr4-go`](https://pkg.go.dev/github.com/antlr4-go/antlr) design decisions) the ANTLR-generated codebase for the [RFC 4512](https://tools.ietf.org/html/rfc4512) grammar has some pretty high cyclomatic factors, with the highest factor at just above 100.

The upside of this is that a schema instance is usually only spawned once per runtime, and not in rapid succession through multiple threads. A schema is a source of truth, and conceptually it should be a _singular_ entity in a process, model or state machine.

Thus the performance of this package -- while still adequately fast -- was really the lowest priority in terms of design goals.

## Ordering of Fields

In addition to this package's extreme flexibility in terms of newlines and whitespace indenting, this package also allows flexibility in terms of field ordering within a given definition. This only applies to contiguous sequences of optional fields, such as `NAME`, `DESC` and others.

For example, consider the following `nameForm` instance derived from [Section 2.7.1 of draft-coretta-oiddir-schema](https://datatracker.ietf.org/doc/html/draft-coretta-oiddir-schema#section-2.7.1):

```
 nameForm ( 1.3.6.1.4.1.56521.101.2.7.1
	NAME 'nRootArcForm'
	DESC 'root arc name form for a number form RDN'
	OC rootArc
	MUST n )
```

Observe the `NameFormDescription` ABNF production per [Section 4.1.7.2 of RFC 4512](https://www.rfc-editor.org/rfc/rfc4512.html#section-4.1.7.2):

```
 NameFormDescription = LPAREN WSP
	numericoid                 ; object identifier
	[ SP "NAME" SP qdescrs ]   ; short names (descriptors)
	[ SP "DESC" SP qdstring ]  ; description
	[ SP "OBSOLETE" ]          ; not active
	SP "OC" SP oid             ; structural object class
	SP "MUST" SP oids          ; attribute types
	[ SP "MAY" SP oids ]       ; attribute types
	extensions WSP RPAREN      ; extensions
```

Both `OC` and `MUST` are required clauses, while `NAME` and `DESC` are optional (as denoted by the encapsulating square brackets). If defined, these optional fields can be reversed order-wise (e.g.: `DESC` before `NAME`), but their placement must always precede the subsequent required `OC` and `MUST` clauses as defined in the above [RFC 4512](https://www.rfc-editor.org/rfc/rfc4512.html) excerpt.

This rule applies to any other definition type that contains a mixture of optional and required clauses, but is moot if said optional fields were not specified at all.  Required fields themselves CANNOT be ordered in any manner other than that specified within the above `NameFormDescription` ABNF production, et al.

In general, it is always strongly advised to honor the ordering of fields as specified in [RFC 4512](https://www.rfc-editor.org/rfc/rfc4512.html).  This flexibility is primarily intended for parsing definitions within RFCs that have long since been accepted despite being technically poorly formed.

## Macro Support

The OpenLDAP `objectidentifier` parameter is supported during the parsing of schema **files** and **directories** only.  This allows succinct aliases -- called macros -- to be used in place of cumbersome and error-prone numeric OIDs alone.

For instance, the [RFC 2307](https://datatracker.ietf.org/doc/html/rfc2307) alias macro `nisSchema` can be declared as follows in the relevant schema files:

```
objectidentifier nisSchema 1.3.6.1.1.1
```

When used within an attributeType in place of the numeric OID:

```
   ( nisSchema.1.2
      NAME 'gecos'
      DESC 'The GECOS field; the common name'
      EQUALITY caseIgnoreIA5Match
      SUBSTRINGS caseIgnoreIA5SubstringsMatch
      SYNTAX 'IA5String'
      SINGLE-VALUE )
```

... the macro properly resolves as follows:

```
   ( 1.3.6.1.1.1.1.2
      NAME 'gecos'
      DESC 'The GECOS field; the common name'
      EQUALITY caseIgnoreIA5Match
      SUBSTRINGS caseIgnoreIA5SubstringsMatch
      SYNTAX 'IA5String'
      SINGLE-VALUE )
```

When using `objectidentifier` macros, a reference to a macro must only be made _after_ the needed macro has been properly defined. Failure to follow this requirement will result in botched OID resolutions during the parsing process and the return of some very unhelpful error messages.

Note that `objectidentifier` macros can only be used for the following definition types:

  - `attributeType`
  - `objectClass`
  - `nameForm`
  - `dITContentRule`

`objectidentifier` macros also support nesting. These are normally referred to as _nested_ or _sub_ macros. For example, one could establish a `nisAttrs` sub macro like so:

```
objectidentifier nisAttrs nisSchema:1
```

... which implies the ultimately resolved numeric OID of "1.3.6.1.1.1.1", which is also known by the macro name "nisAttrs".

Use of nesting requires dependency macros are defined prior to being referenced by other subordinate macros.

If using macros as shown in [RFC 2307](https://datatracker.ietf.org/doc/html/rfc2307) for definitions such as `ldapSyntax` or `matchingRule`, OR if using macros in "piecemeal" definitions -- in other words, definitions NOT parsed from a file or directory -- macros are NOT resolvable.  Rather, the definition will be released without an error but lacking a literal OID. The definition will, however, bear a Macro field value which the user may resolve (or not) using their own means after-the-fact.  In this case, users MAY simply edit the returned instances directly, or they may simply transfer the non-zero contents elsewhere for further handling.

Regardless of the form of macro OID resolution used, users may specify either a dot (".") or colon (":") for the macro name/oid delimiter. In other words, either `nisSchema:0` or `nisSchema.0` would be acceptable.


## Minimal Verification

The parser provided by this package is only capable of the most basic verification checks. For example, beyond the basic required contexts of [RFC 4512 Section 4](https://www.rfc-editor.org/rfc/rfc4512.html#section-4) -- such as a `nameForm`'s `OC` and `MUST` fields -- this package does not perform any in-depth analysis upon the respective values.

The intent of this package is mainly focused on easy tokenization. Very few checks are performed upon individual tokens. To offer some examples, this means:

  - No validity of `matchingRule` combinations found within `attributeType` definitions is confirmed
  - An OID erroneously shared by multiple definitions is not mitigated and will not raise an error
  - References to dependent definitions (e.g.: `name` and `cn`) are not verified presence-wise
  - Bogus OIDs will be produced when bogus macros are used; macros themselves are unverified

The only restrictions imposed during the parsing process come from ANTLR itself; for instance, missing, unexpected or invalid token contents. Token _placement_ and content is defined based upon a reasonable implementation of [RFC 4512](https://www.rfc-editor.org/rfc/rfc4512.html). Beyond this, all bets are off.

If it is desirable to enhance the capabilities of a parser solution, the user is advised to consider manual implementation of the necessary checks and balances, or use an external solution such as [`go-schemax`](https://github.com/JesseCoretta/go-schemax).

## File and Definition Parsing Order

Although this package does not perform in-depth verification on individual references, it is important that files -- and more importantly, their contents -- are parsed in a manner that honors reference dependencies, such as those used for super types, super classes and super rules.

The reason for this is that parsed definitions are stored in the order in which they were originally parsed within the slice-based fields of a Schema definition.  If and when this instance is processed using external processes, such as those affored by the [`go-schemax`](https://github.com/JesseCoretta/go-schemax) package, this will almost certainly result in errors.

Therefore users are strongly advised to verify and ensure the proper ordering of the following:

  - "`.schema`" files
  - Definitions within "`.schema`" files

In either context, the ordering scheme must always favor fundamental elements prior to any and all definitions which reference such elements.

