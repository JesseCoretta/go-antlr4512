package antlr4512

import (
	"testing"
)

var testNameForms []string = []string{
	// Extracted from draft-coretta-oiddir-schema
	`( 1.3.6.1.4.1.56521.101.2.7.1 NAME 'nRootArcForm' DESC 'root arc name form for a number form RDN' OC rootArc MUST n )`,
	`( 1.3.6.1.4.1.56521.101.2.7.2 NAME 'nArcForm' DESC 'arc name form for a number form RDN' OC arc MUST n )`,
	`( 1.3.6.1.4.1.56521.101.2.7.3 NAME 'dotNotationArcForm' DESC 'arc name form for a numeric OID RDN' OC arc MUST dotNotation )`,

	// made-up
	`( 1.3.6.1.4.1.56521.999.5 NAME 'formidableForm' OC domain MUST dc MAY associatedDomain )`,
}

func TestNameForm(t *testing.T) {
	for _, def := range testNameForms {
		_, err := ParseNameForm(def)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}
