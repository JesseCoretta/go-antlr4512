package antlr4512

import (
	"testing"
)

var testMatchingRules []string = []string{
	`( 2.5.13.22 NAME 'presentationAddressMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.43 X-ORIGIN 'RFC2256' )`,
	`( 2.5.13.24 NAME 'protocolInformationMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.42 X-ORIGIN 'RFC2252' )`,
	`( 1.3.6.1.4.1.1466.109.114.3 NAME 'caseIgnoreIA5SubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.109.114.2 NAME 'caseIgnoreIA5Match' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.109.114.1 NAME 'caseExactIA5Match' SYNTAX 1.3.6.1.4.1.1466.115.121.1.26 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.0 NAME 'objectIdentifierMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.1 NAME 'distinguishedNameMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.12 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.2 NAME 'caseIgnoreMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.3 NAME 'caseIgnoreOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.4 NAME 'caseIgnoreSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.5 NAME 'caseExactMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.6 NAME 'caseExactOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.7 NAME 'caseExactSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.8 NAME 'numericStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.9 NAME 'numericStringOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.36 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.10 NAME 'numericStringSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.11 NAME 'caseIgnoreListMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.41 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.12 NAME 'caseIgnoreListSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.13 NAME 'booleanMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.7 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.14 NAME 'integerMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.15 NAME 'integerOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.16 NAME 'bitStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.6 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.17 NAME 'octetStringMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.18 NAME 'octetStringOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.40 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.20 NAME 'telephoneNumberMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.50 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.21 NAME 'telephoneNumberSubstringsMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.58 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.23 NAME 'uniqueMemberMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.34 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.27 NAME 'generalizedTimeMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.28 NAME 'generalizedTimeOrderingMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.24 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.29 NAME 'integerFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.27 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.30 NAME 'objectIdentifierFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.38 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.31 NAME 'directoryStringFirstComponentMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.32 NAME 'wordMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.33 NAME 'keywordMatch' SYNTAX 1.3.6.1.4.1.1466.115.121.1.15 X-ORIGIN 'RFC4517' )`,
	`( 2.5.13.34 NAME 'certificateExactMatch' SYNTAX 1.3.6.1.1.15.1 X-ORIGIN 'RFC4523' )`,
	`( 2.5.13.35 NAME 'certificateMatch' SYNTAX 1.3.6.1.1.15.2 X-ORIGIN 'RFC4523' )`,
	`( 2.5.13.36 NAME 'certificatePairExactMatch' SYNTAX 1.3.6.1.1.15.3 X-ORIGIN 'RFC4523' )`,
	`( 2.5.13.37 NAME 'certificatePairMatch' SYNTAX 1.3.6.1.1.15.4 X-ORIGIN 'RFC4523' )`,
	`( 2.5.13.38 NAME 'certificateListExactMatch' SYNTAX 1.3.6.1.1.15.5 X-ORIGIN 'RFC4523' )`,
	`( 2.5.13.39 NAME 'certificateListMatch' SYNTAX 1.3.6.1.1.15.6 X-ORIGIN 'RFC4523' )`,
	`( 2.5.13.40 NAME 'algorithmIdentifierMatch' SYNTAX 1.3.6.1.1.15.7 X-ORIGIN 'RFC4523' )`,
	`( 1.3.6.1.1.16.2 NAME 'uuidMatch' SYNTAX 1.3.6.1.1.16.1 X-ORIGIN 'RFC4530' )`,
	`( 1.3.6.1.1.16.3 NAME 'uuidOrderingMatch' SYNTAX 1.3.6.1.1.16.1 X-ORIGIN 'RFC4530' )`,
}

func TestParseMatchingRule(t *testing.T) {
	for _, def := range testMatchingRules {
		_, err := ParseMatchingRule(def)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}
