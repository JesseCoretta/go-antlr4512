package antlr4512

import (
	"testing"
)

var testContentRules []string = []string{
	`( 0.9.2342.19200300.100.4.13 NAME 'domainContent' DESC 'content control for domain objects' AUX ( dcObject $ domainRelatedObject ) MUST ( dc $ associatedDomain ) MAY description NOT street )`,
}

func TestDITContentRule(t *testing.T) {
	for _, def := range testContentRules {
		_, err := ParseDITContentRule(def)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}
