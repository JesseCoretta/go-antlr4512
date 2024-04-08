package antlr4512

import (
	"testing"
)
func TestDITContentRule(t *testing.T) {
        for _, dc := range []string{
		`( 0.9.2342.19200300.100.4.13 NAME 'domainContent' DESC 'content control for domain objects' AUX ( dcObject $ domainRelatedObject ) MUST ( dc $ associatedDomain ) MAY description NOT street )`,
        } {
                p, err := ParseInstance(dc)
                if err != nil {
                        t.Errorf("%s failed: %v", t.Name(), err)
                        return
                }
		_ = p.P.DITContentRuleDescription().GetText()
		
        }
}
