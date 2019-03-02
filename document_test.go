package gedcom_test

import (
	"errors"
	"testing"

	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

var documentTests = []struct {
	doc         *gedcom.Document
	individuals gedcom.IndividualNodes
	families    gedcom.FamilyNodes
	p2          gedcom.Node
}{
	{
		doc:         gedcom.NewDocument(),
		individuals: gedcom.IndividualNodes{},
		p2:          nil,
		families:    gedcom.FamilyNodes{},
	},
	{
		doc: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			),
		}),
		individuals: gedcom.IndividualNodes{
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			),
		},
		p2:       nil,
		families: gedcom.FamilyNodes{},
	},
	{
		doc: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			),
			gedcom.NewNode(gedcom.TagVersion, "", ""),
			gedcom.NewDocument().AddIndividual("P2",
				gedcom.NewNameNode("John /Doe/"),
			),
		}),
		individuals: gedcom.IndividualNodes{
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			),
			gedcom.NewDocument().AddIndividual("P2",
				gedcom.NewNameNode("John /Doe/"),
			),
		},
		p2: gedcom.NewDocument().AddIndividual("P2",
			gedcom.NewNameNode("John /Doe/"),
		),
		families: gedcom.FamilyNodes{},
	},
	{
		doc: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			),
			gedcom.NewDocument().AddFamily("F1"),
		}),
		individuals: gedcom.IndividualNodes{
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			),
		},
		p2: nil,
		families: gedcom.FamilyNodes{
			gedcom.NewDocument().AddFamily("F1"),
		},
	},
	{
		doc: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			),
			gedcom.NewDocument().AddFamily("F3"),
		}),
		individuals: gedcom.IndividualNodes{
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			),
		},
		p2: nil,
		families: gedcom.FamilyNodes{
			gedcom.NewDocument().AddFamily("F3"),
		},
	},
}

func TestDocument_Individuals(t *testing.T) {
	for _, test := range documentTests {
		t.Run("", func(t *testing.T) {
			assertEqual(t, test.doc.Individuals(), test.individuals)
		})
	}
}

func TestDocument_NodeByPointer(t *testing.T) {
	for _, test := range documentTests {
		t.Run("", func(t *testing.T) {
			node := test.doc.NodeByPointer("P2")
			assertNodeEqual(t, test.p2, node, fmt.Sprintf("%+#v", test))
		})
	}
}

func TestDocument_Families(t *testing.T) {
	for _, test := range documentTests {
		t.Run("", func(t *testing.T) {
			assertEqual(t, test.doc.Families(), test.families)
		})
	}
}

func TestNewDocumentFromString(t *testing.T) {
	for _, test := range []struct {
		ged      string
		expected *gedcom.Document
		err      error
	}{
		{
			"",
			gedcom.NewDocument(),
			nil,
		},
		{
			"AAA",
			nil,
			errors.New("line 1: could not parse: AAA"),
		},
		{
			"0 INDI\nAAB",
			nil,
			errors.New("line 2: could not parse: AAB"),
		},
		{
			"0 INDI\n\nAAA",
			nil,
			errors.New("line 3: could not parse: AAA"),
		},
	} {
		t.Run(test.ged, func(t *testing.T) {
			result, err := gedcom.NewDocumentFromString(test.ged)

			if test.expected != nil {
				test.expected.MaxLivingAge = gedcom.DefaultMaxLivingAge
			}

			assert.Equal(t, err, test.err)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestNewDocument(t *testing.T) {
	doc := gedcom.NewDocument()

	t.Run("StartsWithZeroNodes", func(t *testing.T) {
		assert.Len(t, doc.Nodes(), 0)
	})

	t.Run("ImplementsGEDCOMStringer", func(t *testing.T) {
		assert.Implements(t, (*gedcom.GEDCOMStringer)(nil), doc)
	})
}

func TestDocument_AddNode(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		doc := gedcom.NewDocument()
		nameNode := gedcom.NewNameNode("foo")

		doc.AddNode(nameNode)

		assert.Equal(t, doc.Nodes(), gedcom.Nodes{nameNode})
	})

	t.Run("Two", func(t *testing.T) {
		doc := gedcom.NewDocument()
		nameNode1 := gedcom.NewNameNode("foo")
		nameNode2 := gedcom.NewNameNode("foo")

		doc.AddNode(nameNode1)
		doc.AddNode(nameNode2)

		assert.Equal(t, doc.Nodes(), gedcom.Nodes{nameNode1, nameNode2})
	})

	t.Run("Duplicate", func(t *testing.T) {
		doc := gedcom.NewDocument()
		nameNode := gedcom.NewNameNode("foo")

		doc.AddNode(nameNode)
		doc.AddNode(nameNode)

		assert.Equal(t, doc.Nodes(), gedcom.Nodes{nameNode, nameNode})
	})

	t.Run("Nil", func(t *testing.T) {
		doc := gedcom.NewDocument()
		nameNode := gedcom.NewNameNode("foo")

		doc.AddNode(nil)
		doc.AddNode(nameNode)

		assert.Equal(t, doc.Nodes(), gedcom.Nodes{nameNode})
	})
}

// These are just some random examples. More extensive testing is on the
// individual nodes.
var documentWarningTests = map[string]struct {
	doc      func(doc *gedcom.Document)
	expected []string
}{
	"EmptyDocument": {
		func(doc *gedcom.Document) {},
		nil,
	},
	"ChildBornBeforeMother": {
		func(doc *gedcom.Document) {
			p1 := doc.AddIndividual("P1").
				AddName("Jenny /Chance/").
				AddBirthDate("16 May 1989")

			p2 := doc.AddIndividual("P2").
				AddName("Elliot /Chance/").
				AddBirthDate("3 Apr 1961")

			doc.AddFamilyWithHusbandAndWife("F1", nil, p1).
				AddChild(p2)
		},
		[]string{
			"The child Elliot Chance (b. 3 Apr 1961) was born before their mother Jenny Chance (b. 16 May 1989).",
		},
	},
	"UnparsableDate": {
		func(doc *gedcom.Document) {
			doc.AddIndividual("P1").
				AddName("John /Chance/").
				AddBirthDate("foo bar")

			doc.AddIndividual("P2").
				AddName("Elliot /Chance/").
				AddBirthDate("around world war 2").
				SetSex(gedcom.SexMale)
		},
		[]string{
			`Unparsable date "foo bar"`,
			`Unparsable date "around world war 2"`,
		},
	},
	"DeathBeforeBirth": {
		func(doc *gedcom.Document) {
			doc.AddIndividual("P1").
				AddName("Jenny /Chance/").
				AddBirthDate("16 May 1989").
				AddDeathDate("16 May 1979")
		},
		[]string{
			"The death (16 May 1979) was before the birth (16 May 1989) of Jenny Chance (b. 16 May 1989, d. 16 May 1979).",
		},
	},
}

func TestDocument_Warnings(t *testing.T) {
	for testName, test := range documentWarningTests {
		t.Run(testName, func(t *testing.T) {
			doc := gedcom.NewDocument()
			test.doc(doc)

			assertEqual(t, doc.Warnings().Strings(), test.expected)
		})
	}
}
