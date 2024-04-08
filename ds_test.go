package antlr4512

import (
	"testing"
)
func TestDITStructureRule(t *testing.T) {
        for _, ds := range []string{
		`( 0 NAME 'structureRule' DESC 'structural control' FORM formidableForm )`,
        } {
                p, err := ParseInstance(ds)
                if err != nil {
                        t.Errorf("%s failed: %v", t.Name(), err)
                        return
                }
		_ = p.P.DITStructureRuleDescription().GetText()
		
        }
}
