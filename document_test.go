package gedcom_test

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var documentTests = []struct {
	doc         *gedcom.Document
	individuals []*gedcom.IndividualNode
	families    []*gedcom.FamilyNode
	p2          gedcom.Node
}{
	{
		doc:         &gedcom.Document{},
		individuals: []*gedcom.IndividualNode{},
		p2:          nil,
		families:    []*gedcom.FamilyNode{},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
				}),
			},
		},
		individuals: []*gedcom.IndividualNode{
			gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		},
		p2:       nil,
		families: []*gedcom.FamilyNode{},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
				}),
				gedcom.NewSimpleNode(gedcom.TagVersion, "", "", []gedcom.Node{}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
				}),
			},
		},
		individuals: []*gedcom.IndividualNode{
			gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			}),
			gedcom.NewIndividualNode("", "P2", []gedcom.Node{
				gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
			}),
		},
		p2: gedcom.NewIndividualNode("", "P2", []gedcom.Node{
			gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
		}),
		families: []*gedcom.FamilyNode{},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
				}),
				gedcom.NewFamilyNode("F1", []gedcom.Node{}),
			},
		},
		individuals: []*gedcom.IndividualNode{
			gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		},
		p2: nil,
		families: []*gedcom.FamilyNode{
			gedcom.NewFamilyNode("F1", []gedcom.Node{}),
		},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
				}),
				gedcom.NewFamilyNode("F3", []gedcom.Node{}),
			},
		},
		individuals: []*gedcom.IndividualNode{
			gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		},
		p2: nil,
		families: []*gedcom.FamilyNode{
			gedcom.NewFamilyNode("F3", []gedcom.Node{}),
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
			assert.Equal(t, test.doc.NodeByPointer("P2"), test.p2,
				fmt.Sprintf("%+#v", test))
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
