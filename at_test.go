package antlr4512

import (
	"testing"
)

func TestAttrOpts(t *testing.T) {
	for _, at := range []string{
		`l;lang-de;lang-jp`,
		`cn;lang-en`,
		`sn`,
		`2.5.4.3;lang-cn`,
		`userCertificate;binary`,
	} {
		_, err := ParseInstance(at)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}

var testAttributes []string = []string{
	`( 2.5.21.6 NAME 'objectClasses' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.37 USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.18.10 NAME 'subschemaSubentry' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.21.5 NAME 'attributeTypes' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.3 USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.21.4 NAME 'matchingRules' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.30 USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.21.8 NAME 'matchingRuleUse' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.31 USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 1.3.6.1.4.1.1466.101.120.16 NAME 'ldapSyntaxes' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.54 USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.21.2 NAME 'dITContentRules' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.16 USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.21.1 NAME 'dITStructureRules' EQUALITY integerFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.17 USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.21.7 NAME 'nameForms' EQUALITY objectIdentifierFirstComponentMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.35 USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 1.3.6.1.4.1.1466.101.120.6 NAME 'altServer' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 USAGE dSAOperation X-ORIGIN 'RFC4512' )`,
	`( 1.3.6.1.4.1.1466.101.120.5 NAME 'namingContexts' SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 USAGE dSAOperation X-ORIGIN 'RFC4512' )`,
	`( 1.3.6.1.4.1.1466.101.120.13 NAME 'supportedControl' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`,
	`( 1.3.6.1.4.1.1466.101.120.7 NAME 'supportedExtension' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`,
	`( 1.3.6.1.4.1.4203.1.3.5 NAME 'supportedFeatures' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 USAGE dSAOperation X-ORIGIN 'RFC4512' )`,
	`( 1.3.6.1.4.1.1466.101.120.15 NAME 'supportedLDAPVersion' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 USAGE dSAOperation X-ORIGIN 'RFC4512' )`,
	`( 1.3.6.1.4.1.1466.101.120.14 NAME 'supportedSASLMechanisms' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 USAGE dSAOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.4.1 NAME 'aliasedObjectName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE X-ORIGIN 'RFC4512' )`,
	`( 2.5.4.0 NAME 'objectClass' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4512' )`,
	`( 2.5.18.3 NAME 'creatorsName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.18.1 NAME 'createTimestamp' EQUALITY generalizedTimeMatch ORDERING generalizedTimeOrderingMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.18.4 NAME 'modifiersName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.18.2 NAME 'modifyTimestamp' EQUALITY generalizedTimeMatch ORDERING generalizedTimeOrderingMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.21.9 NAME 'structuralObjectClass' EQUALITY objectIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.21.10 NAME 'governingStructureRule' EQUALITY integerMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4512' )`,
	`( 2.5.4.15 NAME 'businessCategory' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.41 NAME 'name' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.6 NAME 'c' SUP name SYNTAX 1.3.6.1.4.1.1466.115.121.1.11 SINGLE-VALUE X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.3 NAME ( 'cn' 'commonName' ) DESC 'RFC4519: common name(s) for which the entity is known by' SUP name X-ORIGIN 'RFC4519' )`,
	`( 0.9.2342.19200300.100.1.25 NAME ( 'dc' 'domainComponent' ) EQUALITY caseIgnoreIA5Match SUBSTR caseIgnoreIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 SINGLE-VALUE X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.13 NAME 'description' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.27 NAME 'destinationIndicator' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.49 NAME 'distinguishedName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.46 NAME 'dnQualifier' EQUALITY caseIgnoreMatch ORDERING caseIgnoreOrderingMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.47 NAME 'enhancedSearchGuide' SYNTAX 1.3.6.1.4.1.1466.115.121.1.21 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.23 NAME 'facsimileTelephoneNumber' SYNTAX 1.3.6.1.4.1.1466.115.121.1.22 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.44 NAME 'generationQualifier' SUP name X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.42 NAME 'givenName' SUP name X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.51 NAME 'houseIdentifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.43 NAME 'initials' SUP name X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.25 NAME 'internationalISDNNumber' EQUALITY numericStringMatch SUBSTR numericStringSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.7 NAME ( 'l' 'localityName' ) SUP name X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.31 NAME 'member' SUP distinguishedName X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.10 NAME ( 'o' 'organizationName' ) SUP name X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.11 NAME ( 'ou' 'organizationalUnitName' ) SUP name X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.32 NAME 'owner' SUP distinguishedName X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.19 NAME 'physicalDeliveryOfficeName' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.16 NAME 'postalAddress' EQUALITY caseIgnoreListMatch SUBSTR caseIgnoreListSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.17 NAME 'postalCode' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.18 NAME 'postOfficeBox' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.28 NAME 'preferredDeliveryMethod' SYNTAX 1.3.6.1.4.1.1466.115.121.1.14 SINGLE-VALUE X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.26 NAME 'registeredAddress' SUP postalAddress SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.33 NAME 'roleOccupant' SUP distinguishedName X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.14 NAME 'searchGuide' SYNTAX 1.3.6.1.4.1.1466.115.121.1.25 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.34 NAME 'seeAlso' SUP distinguishedName X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.5 NAME 'serialNumber' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.44 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.4 NAME ( 'sn' 'surname' ) SUP name X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.8 NAME ( 'st' 'stateOrProvinceName' ) SUP name X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.9 NAME ( 'street' 'streetAddress' ) EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.20 NAME 'telephoneNumber' EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.22 NAME 'teletexTerminalIdentifier' SYNTAX 1.3.6.1.4.1.1466.115.121.1.51 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.21 NAME 'telexNumber' SYNTAX 1.3.6.1.4.1.1466.115.121.1.52 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.12 NAME 'title' SUP name X-ORIGIN 'RFC4519' )`,
	`( 0.9.2342.19200300.100.1.1 NAME ( 'uid' 'userid' ) DESC 'RFC4519: user identifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.50 NAME 'uniqueMember' EQUALITY uniqueMemberMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.34 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.35 NAME 'userPassword' EQUALITY octetStringMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.24 NAME 'x121Address' EQUALITY numericStringMatch SUBSTR numericStringSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.45 NAME 'x500UniqueIdentifier' EQUALITY bitStringMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.6 X-ORIGIN 'RFC4519' )`,
	`( 2.5.4.36 NAME 'userCertificate' DESC 'X.509 user certificate' EQUALITY certificateExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.8 X-ORIGIN 'RFC4523' )`,
	`( 2.5.4.37 NAME 'cACertificate' DESC 'X.509 CA certificate' EQUALITY certificateExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.8 X-ORIGIN 'RFC4523' )`,
	`( 2.5.4.40 NAME 'crossCertificatePair' DESC 'X.509 cross certificate pair' EQUALITY certificatePairExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.10 X-ORIGIN 'RFC4523' )`,
	`( 2.5.4.39 NAME 'certificateRevocationList' DESC 'X.509 certificate revocation list' EQUALITY certificateListExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.9 X-ORIGIN 'RFC4523' )`,
	`( 2.5.4.38 NAME 'authorityRevocationList' DESC 'X.509 authority revocation list' EQUALITY certificateListExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.9 X-ORIGIN 'RFC4523' )`,
	`( 2.5.4.53 NAME 'deltaRevocationList' DESC 'X.509 delta revocation list' EQUALITY certificateListExactMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.9 X-ORIGIN 'RFC4523' )`,
	`( 2.5.4.52 NAME 'supportedAlgorithms' DESC 'X.509 supported algorithms' EQUALITY algorithmIdentifierMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.49 X-ORIGIN 'RFC4523' )`,
	`( 0.9.2342.19200300.100.1.37 NAME 'associatedDomain' EQUALITY caseIgnoreIA5Match SUBSTR caseIgnoreIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.38 NAME 'associatedName' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.48 NAME 'buildingName' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.43 NAME ( 'co' 'friendlyCountryName' ) EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.14 NAME 'documentAuthor' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.11 NAME 'documentIdentifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.15 NAME 'documentLocation' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.56 NAME 'documentPublisher' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.12 NAME 'documentTitle' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.13 NAME 'documentVersion' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.5 NAME ( 'drink' 'favouriteDrink' ) EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.20 NAME ( 'homePhone' 'homeTelephoneNumber' ) EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.39 NAME 'homePostalAddress' EQUALITY caseIgnoreListMatch SUBSTR caseIgnoreListSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.9 NAME 'host' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.4 NAME 'info' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{2048} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.3 NAME ( 'mail' 'rfc822Mailbox' ) EQUALITY caseIgnoreIA5Match SUBSTR caseIgnoreIA5SubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.26{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.10 NAME 'manager' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.41 NAME ( 'mobile' 'mobileTelephoneNumber' ) EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.45 NAME 'organizationalStatus' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.42 NAME ( 'pager' 'pagerTelephoneNumber' ) EQUALITY telephoneNumberMatch SUBSTR telephoneNumberSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.40 NAME 'personalTitle' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.6 NAME 'roomNumber' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.21 NAME 'secretary' EQUALITY distinguishedNameMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.44 NAME 'uniqueIdentifier' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 0.9.2342.19200300.100.1.8 NAME 'userClass' EQUALITY caseIgnoreMatch SUBSTR caseIgnoreSubstringsMatch SYNTAX 1.3.6.1.4.1.1466.115.121.1.15{256} X-ORIGIN 'RFC4524' )`,
	`( 1.3.6.1.1.16.4 NAME 'entryUUID' DESC 'UUID of the entry' EQUALITY uuidMatch ORDERING uuidOrderingMatch SYNTAX 1.3.6.1.1.16.1 SINGLE-VALUE NO-USER-MODIFICATION USAGE directoryOperation X-ORIGIN 'RFC4530' )`,
}

func TestAttributeType(t *testing.T) {
	for _, def := range testAttributes {
		_, err := ParseAttributeType(def)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}
