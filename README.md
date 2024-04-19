# go-rfc4512-antlr

[![ANTLR4](https://img.shields.io/badge/antlr-blue?label=%E2%92%B6&labelColor=gray&color=magenta&cacheSeconds=86400)](https://www.antlr.org/) [![RFC4512](https://img.shields.io/badge/RFC-4512-blue)](https://datatracker.ietf.org/doc/html/rfc4512) [![Reference](https://pkg.go.dev/badge/github.com/JesseCoretta/go-rfc4512-antlr.svg)](https://pkg.go.dev/github.com/JesseCoretta/go-rfc4512-antlr) [![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-rfc4512-antlr/blob/main/LICENSE) [![Issues](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-rfc4512-antlr/issues) [![Experimental](https://img.shields.io/badge/experimental-blue?logoColor=blue&label=%F0%9F%A7%AA%20%F0%9F%94%AC&labelColor=blue&color=gray)](https://github.com/JesseCoretta/JesseCoretta/blob/main/EXPERIMENTAL.md)

Package antlr4512 implements Sections 2.5 and 4 of RFC 4512 by way of an unmodified ANTLR4-generated target.

## ANTLR

See the `_grammar/RFC4512.g4` file within this repository for ANTLR grammar. Using this file, the target was auto-generated using the following `antlr4` command:

```
$ antlr4 -Dlanguage=Go -package antlr4512 RFC4512.g4
```

For details on ANTLR, click the above ANTLR badge. This page was generated using **antlr4-go** version 4.13.1.

If alterations are desired, modify the above grammar file and regenerate code as needed. Remember it is NOT recommended for auto-generated code to be modified in any way. Code enhancements, convenience wrappers and other extensions devised by the adopter(s) should ONLY be made in separate files sharing the same package namespace.

## Features

  - Compliant
    - RFC 4512 syntax implementation honors relevant ABNF productions
  - Compatible
    - 389DS and OpenLDAP style schema formats supported
    - Auto-resolving OpenLDAP `objectidentifier` macro support
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
      - `ParseAttributeType`
      - `ParseObjectClass`
      - `ParseDITContentRule`
      - `ParseNameForm`
      - `ParseDITStructureRule`
  - Simple, easily understood type instances
    - Values are confined to primitive instances, easily read
    - No unnecessary pointers, no private fields
    - Just parse your instance and go!
    - Great for templating and external processing

## Cyclomatics Notice

Please note that due to the complexity and flexibility of the grammar design (and some questionable antlr4-go design decisions) the ANTLR-generated codebase for the RFC4512 grammar has some pretty high cyclomatic factors, with the highest factor at just above 100.

The upside of this is that a schema instance is usually only spawned once per runtime, and not in rapid succession through multiple threads. A schema is a source of truth, and conceptually it should be a _singular_ entity in a process, model or state machine.

Thus the performance of this package -- while still adequately fast -- was really the lowest priority in terms of design goals.

## Ordering of Fields

In addition to this package's extreme flexibility in terms of newlines and whitespace indenting, this package also allows flexibility in terms of field ordering within a given definition. This only applies to contiguous sequences of optional fields, such as `NAME`, `DESC` and others.

For example, consider the following `nameForm` instance derived from Section 2.7.1 of **draft-coretta-oiddir-schema**:

```
 nameForm ( 1.3.6.1.4.1.56521.101.2.7.1
	NAME `nRootArcForm`
	DESC 'root arc name form for a number form RDN'
	OC rootArc
	MUST n )
```

Observe the `NameFormDescription` ABNF production per Section 4.1.7.2 of RFC 4512:

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

Both `OC` and `MUST` are required clauses, while `NAME` and `DESC` are optional. If defined, these optional fields can be reversed order-wise (e.g.: `DESC` before `NAME`), but their placement must always precede the subsequent required `OC` and `MUST` clauses as defined in the above RFC 4512 excerpt.

This rule applies to any other definition type that contains a mixture of optional and required clauses, but is moot if said optional fields were not specified at all.  Required fields themselves CANNOT be ordered in any manner other than that specified within the above `NameFormDescription` ABNF production, et al.

## Macro Support

The OpenLDAP `objectidentifier` parameter is supported in parsed schema files.  This allows succinct aliases -- called macros -- to be used in place of cumbersome and error-prone numeric OIDs alone.

For instance, the RFC2307 alias macro `nisSchema` can be declared as follows in the relevant schema files:

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

When using macros, a reference to a macro must only be made after the needed macro has been properly defined. Failure to follow this rule will result in OID resolutions during the parsing process.

Note that macros can only be used for the following definition types:

  - `attributeType`
  - `objectClass`
  - `nameForm`
  - `dITContentRule`

Either a dot (".") or colon (":") may be used as the macro delimiter. In other words, either `nisSchema:0` or `nisSchema.0` are supported.

Macros also support nesting. These are normally referred to as nested or sub macros. For example, one could establish a `nisAttrs` sub macro like so:

```
objectidentifier nisAttrs nisSchema:1
```

Use of nesting requires dependency macros are defined prior to being referenced by other subordinate macros.
