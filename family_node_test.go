package gedcom_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elliotchance/gedcom"
)

var familyTests = []struct {
	doc     *gedcom.Document
	husband *gedcom.IndividualNode
	wife    *gedcom.IndividualNode
}{
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewFamilyNode("F1", nil),
			},
		},
		husband: nil,
		wife:    nil,
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewFamilyNode("F1", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagHusband, "@P1@", "", []gedcom.Node{}),
				}),
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			},
		},
		husband: gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
		wife:    nil,
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewFamilyNode("F2", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagWife, "@P3@", "", []gedcom.Node{}),
				}),
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{}),
			},
		},
		husband: nil,
		wife:    gedcom.NewIndividualNode("", "P3", []gedcom.Node{}),
	},
}

func TestFamilyNode_Husband(t *testing.T) {
	for _, test := range familyTests {
		t.Run("", func(t *testing.T) {
			node := test.doc.Nodes[0].(*gedcom.FamilyNode)
			assert.Equal(t, node.Husband(test.doc), test.husband)
		})
	}
}

func TestFamilyNode_Wife(t *testing.T) {
	for _, test := range familyTests {
		t.Run("", func(t *testing.T) {
			node := test.doc.Nodes[0].(*gedcom.FamilyNode)
			assert.Equal(t, node.Wife(test.doc), test.wife)
		})
	}
}
