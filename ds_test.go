package antlr4512

import (
	"testing"
)

var testStructureRules []string = []string{
	`( 0 NAME 'structureRule' DESC 'structural control' FORM formidableForm )`,
	`( 1 NAME 'subStructureRule' DESC 'substructural control' FORM formidableForm SUP 0 X-ORIGIN 'YOUR MOM' )`,
}

func TestDITStructureRule(t *testing.T) {
	for _, def := range testStructureRules {
		_, err := ParseDITStructureRule(def)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}
