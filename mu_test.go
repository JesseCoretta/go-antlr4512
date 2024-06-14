package antlr4512

import (
	"testing"
)

var testMatchingRuleUses []string = []string{
	`( 2.5.13.1
    NAME 'distinguishedNameMatch'
    APPLIES ( aliasedObjectName
            $ associatedName
            $ collectiveAttributeSubentries
            $ creatorsName
            $ distinguishedName
            $ documentAuthor
            $ manager
            $ modifiersName
            $ secretary
            $ subschemaSubentry )
    X-ORIGIN 'RFC4517' )`,
}

func TestParseMatchingRuleUse(t *testing.T) {
	for _, def := range testMatchingRuleUses {
		_, err := ParseMatchingRuleUse(def)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}
