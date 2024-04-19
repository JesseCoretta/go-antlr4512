package antlr4512

import (
	"fmt"
	"testing"
)

func ExampleParseObjectClass() {
        var raw string = `( 2.5.6.0
		NAME 'top'
		ABSTRACT MUST objectClass X-ORIGIN 'RFC4512' )`

        oc, err := ParseObjectClass(raw)
        if err != nil {
                fmt.Errorf("%v", err)
                return
        }

        fmt.Printf("%s\n", oc.OID)
        // Output: 2.5.6.0
}

func TestObjectClass(t *testing.T) {
	for _, oc := range []string{
		`( 2.5.6.0 NAME 'top' ABSTRACT MUST objectClass X-ORIGIN 'RFC4512' )`,
		`( 2.5.6.1 NAME 'alias' SUP top STRUCTURAL MUST aliasedObjectName X-ORIGIN 'RFC4512' )`,
		`( 2.5.20.1 NAME 'subschema' AUXILIARY MAY ( dITStructureRules $ nameForms $ dITContentRules $ objectClasses $ attributeTypes $ matchingRules $ matchingRuleUse ) X-ORIGIN 'RFC4512' )`,
		`( 1.3.6.1.4.1.1466.101.120.111 NAME 'extensibleObject' SUP top AUXILIARY X-ORIGIN 'RFC4512' )`,
		`( 2.5.6.11 NAME 'applicationProcess' SUP top STRUCTURAL MUST cn MAY ( seeAlso $ ou $ l $ description ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.2 NAME 'country' SUP top STRUCTURAL MUST c MAY ( searchGuide $ description ) X-ORIGIN 'RFC4519' )`,
		`( 1.3.6.1.4.1.1466.344 NAME 'dcObject' SUP top AUXILIARY MUST dc X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.14 NAME 'device' SUP top STRUCTURAL MUST cn MAY ( serialNumber $ seeAlso $ owner $ ou $ o $ l $ description ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.9 NAME 'groupOfNames' SUP top STRUCTURAL MUST ( member $ cn ) MAY ( businessCategory $ seeAlso $ owner $ ou $ o $ description ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.17 NAME 'groupOfUniqueNames' SUP top STRUCTURAL MUST ( uniqueMember $ cn ) MAY ( businessCategory $ seeAlso $ owner $ ou $ o $ description ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.3 NAME 'locality' SUP top STRUCTURAL MAY ( street $ seeAlso $ searchGuide $ st $ l $ description ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.4 NAME 'organization' SUP top STRUCTURAL MUST o MAY ( userPassword $ searchGuide $ seeAlso $ businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l $ description ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.6 NAME 'person' SUP top STRUCTURAL MUST ( sn $ cn ) MAY ( userPassword $ telephoneNumber $ seeAlso $ description ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.7 NAME 'organizationalPerson' SUP person STRUCTURAL MAY ( title $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ ou $ st $ l ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.8 NAME 'organizationalRole' SUP top STRUCTURAL MUST cn MAY ( x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ seeAlso $ roleOccupant $ preferredDeliveryMethod $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ ou $ st $ l $ description ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.5 NAME 'organizationalUnit' SUP top STRUCTURAL MUST ou MAY ( businessCategory $ description $ destinationIndicator $ facsimileTelephoneNumber $ internationalISDNNumber $ l $ physicalDeliveryOfficeName $ postalAddress $ postalCode $ postOfficeBox $ preferredDeliveryMethod $ registeredAddress $ searchGuide $ seeAlso $ st $ street $ telephoneNumber $ teletexTerminalIdentifier $ telexNumber $ userPassword $ x121Address ) X-ORIGIN 'RFC4519' )`,
		`( 2.5.6.10 NAME 'residentialPerson' SUP person STRUCTURAL MUST l MAY ( businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationalISDNNumber $ facsimileTelephoneNumber $ preferredDeliveryMethod $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l ) X-ORIGIN 'RFC4519' )`,
		`( 1.3.6.1.1.3.1 NAME 'uidObject' SUP top AUXILIARY MUST uid X-ORIGIN 'RFC4519' )`,
		` ( 2.5.6.21 NAME 'pkiUser' DESC 'X.509 PKI User' SUP top AUXILIARY MAY userCertificate X-ORIGIN 'RFC4523' )`,
		` ( 2.5.6.22 NAME 'pkiCA' DESC 'X.509 PKI Certificate Authority' SUP top AUXILIARY MAY ( cACertificate $ certificateRevocationList $ authorityRevocationList $ crossCertificatePair ) X-ORIGIN 'RFC4523' )`,
		`( 2.5.6.19 NAME 'cRLDistributionPoint' DESC 'X.509 CRL distribution point' SUP top STRUCTURAL MUST cn MAY ( certificateRevocationList $ authorityRevocationList $ deltaRevocationList ) X-ORIGIN 'RFC4523' )`,
		`( 2.5.6.23 NAME 'deltaCRL' DESC 'X.509 delta CRL' SUP top AUXILIARY MAY deltaRevocationList X-ORIGIN 'RFC4523' )`,
		`( 2.5.6.15 NAME 'strongAuthenticationUser' DESC 'X.521 strong authentication user' SUP top AUXILIARY MUST userCertificate X-ORIGIN 'RFC4523' )`,
		`( 2.5.6.18 NAME 'userSecurityInformation' DESC 'X.521 user security information' SUP top AUXILIARY MAY ( supportedAlgorithms ) X-ORIGIN 'RFC4523' )`,
		`( 2.5.6.16 NAME 'certificationAuthority' DESC 'X.509 certificate authority' SUP top AUXILIARY MUST ( authorityRevocationList $ certificateRevocationList $ cACertificate ) MAY crossCertificatePair X-ORIGIN 'RFC4523' )`,
		`( 2.5.6.16.2 NAME 'certificationAuthority-V2' DESC 'X.509 certificate authority, version 2' SUP certificationAuthority AUXILIARY MAY deltaRevocationList X-ORIGIN 'RFC4523' )`,
		`( 0.9.2342.19200300.100.4.5 NAME 'account' SUP top STRUCTURAL MUST uid MAY ( description $ seeAlso $ l $ o $ ou $ host ) X-ORIGIN 'RFC4524' )`,
		`( 0.9.2342.19200300.100.4.6 NAME 'document' SUP top STRUCTURAL MUST documentIdentifier MAY ( cn $ description $ seeAlso $ l $ o $ ou $ documentTitle $ documentVersion $ documentAuthor $ documentLocation $ documentPublisher ) X-ORIGIN 'RFC4524' )`,
		`( 0.9.2342.19200300.100.4.9 NAME 'documentSeries' SUP top STRUCTURAL MUST cn MAY ( description $ l $ o $ ou $ seeAlso $ telephoneNumber ) X-ORIGIN 'RFC4524' )`,
		`( 0.9.2342.19200300.100.4.13 NAME 'domain' SUP top STRUCTURAL MUST dc MAY ( userPassword $ searchGuide $ seeAlso $ businessCategory $ x121Address $ registeredAddress $ destinationIndicator $ preferredDeliveryMethod $ telexNumber $ teletexTerminalIdentifier $ telephoneNumber $ internationaliSDNNumber $ facsimileTelephoneNumber $ street $ postOfficeBox $ postalCode $ postalAddress $ physicalDeliveryOfficeName $ st $ l $ description $ o $ associatedName ) X-ORIGIN 'RFC4524' )`,
		`( 0.9.2342.19200300.100.4.17 NAME 'domainRelatedObject' SUP top AUXILIARY MUST associatedDomain X-ORIGIN 'RFC4524' )`,
		`( 0.9.2342.19200300.100.4.18 NAME 'friendlyCountry' SUP country STRUCTURAL MUST co X-ORIGIN 'RFC4524' )`,
		`( 0.9.2342.19200300.100.4.14 NAME 'rFC822localPart' SUP domain STRUCTURAL MAY ( cn $ description $ destinationIndicator $ facsimileTelephoneNumber $ internationaliSDNNumber $ physicalDeliveryOfficeName $ postalAddress $ postalCode $ postOfficeBox $ preferredDeliveryMethod $ registeredAddress $ seeAlso $ sn $ street $ telephoneNumber $ teletexTerminalIdentifier $ telexNumber $ x121Address ) X-ORIGIN 'RFC4524' )`,
	} {
		_, err := ParseObjectClass(oc)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}
