package rfc4512antlr

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
