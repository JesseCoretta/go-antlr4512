package antlr4512

import (
	"testing"
)

var testSyntaxes []string = []string{
	`( 1.3.6.1.4.1.1466.115.121.1.1 DESC 'ACI Item' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.2 DESC 'Access Point' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.3 DESC 'Attribute Type Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.4 DESC 'Audio' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.5 DESC 'Binary' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.6 DESC 'Bit String' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.7 DESC 'Boolean' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.11 DESC 'Country String' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.12 DESC 'DN' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.13 DESC 'Data Quality Syntax' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.14 DESC 'Delivery Method' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.15 DESC 'Directory String' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.16 DESC 'DIT Content Rule Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.17 DESC 'DIT Structure Rule Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.18 DESC 'DL Submit Permission' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.19 DESC 'DSA Quality Syntax' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.20 DESC 'DSE Type' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.21 DESC 'Enhanced Guide' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.22 DESC 'Facsimile Telephone Number' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.23 DESC 'Fax' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.24 DESC 'Generalized Time' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.25 DESC 'Guide' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.26 DESC 'IA5 String' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.27 DESC 'INTEGER' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.28 DESC 'JPEG' X-ORIGIN 'RFC4517' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.30 DESC 'Matching Rule Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.31 DESC 'Matching Rule Use Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.32 DESC 'Mail Preference' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.33 DESC 'MHS OR Address' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.34 DESC 'Name And Optional UID' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.35 DESC 'Name Form Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.36 DESC 'Numeric String' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.37 DESC 'Object Class Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.38 DESC 'OID' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.39 DESC 'Other Mailbox' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.40 DESC 'Octet String' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.41 DESC 'Postal Address' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.42 DESC 'Protocol Information' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.43 DESC 'Presentation Address' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.44 DESC 'Printable String' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.45 DESC 'Subtree Specification' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.46 DESC 'Supplier Information' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.47 DESC 'Supplier Or Consumer' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.48 DESC 'Supplier And Consumer' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.50 DESC 'Telephone Number' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.51 DESC 'Teletext Terminal Identifier' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.52 DESC 'Telex Number' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.53 DESC 'UTC Time' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.54 DESC 'LDAP Syntax Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.55 DESC 'Modify Rights' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.56 DESC 'LDAP Schema Definition' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.57 DESC 'LDAP Schema Description' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.58 DESC 'Substring Assertion' X-ORIGIN 'RFC4517' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.8 DESC 'Certificate' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.9 DESC 'Certificate List' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.10 DESC 'Certificate Pair' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.4.1.1466.115.121.1.49 DESC 'Supported Algorithm' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.1.15.1 DESC 'X.509 Certificate Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.1.15.2 DESC 'X.509 Certificate Assertion' X-ORIGIN 'RFC4523' )`,
	`( 1.3.6.1.1.15.3 DESC 'X.509 Certificate Pair Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.1.15.4 DESC 'X.509 Certificate Pair Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.1.15.5 DESC 'X.509 Certificate List Exact Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.1.15.6 DESC 'X.509 Certificate List Assertion' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.1.15.7 DESC 'X.509 Algorithm Identifier' X-ORIGIN 'RFC4523' X-NOT-HUMAN-READABLE 'TRUE' )`,
	`( 1.3.6.1.1.16.1 DESC 'UUID' X-ORIGIN 'RFC4530' )`,
}

func TestParseLDAPSyntax(t *testing.T) {

	//var lsc int
	for _, label := range []string{
		`ldapSyntax `,
		`ldapsyntax `,
		`ldapSyntax:`,
		`ldapsyntax:`,
		`ldapSyntax=`,
		`ldapsyntax=`,
		`ldapSyntaxes `,
		`ldapsyntaxes `,
		`ldapSyntaxes:`,
		`ldapsyntaxes:`,
		`ldapSyntaxes=`,
		`ldapsyntaxes=`,
	} {
		for _, ls := range testSyntaxes {
			if _, err := ParseInstance(label + ls); err != nil {
				t.Errorf("%s failed: %v", t.Name(), err)
				return
			}
			//lsc++
			//t.Logf("%s\n", i.P.LDAPSyntaxDescriptions().GetText())
		}
	}

	//t.Logf("Parsed %d ldapSyntax definitions\n", lsc)
}
