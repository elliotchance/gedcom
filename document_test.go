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
	families    []*gedcom.FamilyNode
	p2          gedcom.Node
}{
	{
		doc:         gedcom.NewDocument(),
		individuals: gedcom.IndividualNodes{},
		p2:          nil,
		families:    []*gedcom.FamilyNode{},
	},
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		}),
		individuals: gedcom.IndividualNodes{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		},
		p2:       nil,
		families: []*gedcom.FamilyNode{},
	},
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
			gedcom.NewNodeWithChildren(nil, gedcom.TagVersion, "", "", []gedcom.Node{}),
			gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
				gedcom.NewNameNode(nil, "John /Doe/", "", []gedcom.Node{}),
			}),
		}),
		individuals: gedcom.IndividualNodes{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
			gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
				gedcom.NewNameNode(nil, "John /Doe/", "", []gedcom.Node{}),
			}),
		},
		p2: gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
			gedcom.NewNameNode(nil, "John /Doe/", "", []gedcom.Node{}),
		}),
		families: []*gedcom.FamilyNode{},
	},
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
			gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{}),
		}),
		individuals: gedcom.IndividualNodes{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		},
		p2: nil,
		families: []*gedcom.FamilyNode{
			gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{}),
		},
	},
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
			gedcom.NewFamilyNode(nil, "F3", []gedcom.Node{}),
		}),
		individuals: gedcom.IndividualNodes{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		},
		p2: nil,
		families: []*gedcom.FamilyNode{
			gedcom.NewFamilyNode(nil, "F3", []gedcom.Node{}),
		},
	},
}

func TestDocument_Individuals(t *testing.T) {
	for _, test := range documentTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.doc.Individuals(), test.individuals)
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
			assert.Equal(t, test.doc.Families(), test.families)
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

	assert.Len(t, doc.Nodes(), 0)
}

func TestDocument_AddNode(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		doc := gedcom.NewDocument()
		nameNode := gedcom.NewNameNode(doc, "foo", "", nil)

		doc.AddNode(nameNode)

		assert.Equal(t, doc.Nodes(), []gedcom.Node{nameNode})
	})

	t.Run("Two", func(t *testing.T) {
		doc := gedcom.NewDocument()
		nameNode1 := gedcom.NewNameNode(doc, "foo", "", nil)
		nameNode2 := gedcom.NewNameNode(doc, "foo", "", nil)

		doc.AddNode(nameNode1)
		doc.AddNode(nameNode2)

		assert.Equal(t, doc.Nodes(), []gedcom.Node{nameNode1, nameNode2})
	})

	t.Run("Duplicate", func(t *testing.T) {
		doc := gedcom.NewDocument()
		nameNode := gedcom.NewNameNode(doc, "foo", "", nil)

		doc.AddNode(nameNode)
		doc.AddNode(nameNode)

		assert.Equal(t, doc.Nodes(), []gedcom.Node{nameNode, nameNode})
	})

	t.Run("Nil", func(t *testing.T) {
		doc := gedcom.NewDocument()
		nameNode := gedcom.NewNameNode(doc, "foo", "", nil)

		doc.AddNode(nil)
		doc.AddNode(nameNode)

		assert.Equal(t, doc.Nodes(), []gedcom.Node{nameNode})
	})
}
