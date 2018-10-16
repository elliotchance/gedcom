package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

var familyTests = []struct {
	doc     *gedcom.Document
	husband *gedcom.IndividualNode
	wife    *gedcom.IndividualNode
}{
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewFamilyNode(nil, "F1", nil),
		}),
		husband: nil,
		wife:    nil,
	},
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P1@", "", []gedcom.Node{}),
			}),
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
		}),
		husband: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
		wife:    nil,
	},
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewFamilyNode(nil, "F2", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P3@", "", []gedcom.Node{}),
			}),
			gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{}),
		}),
		husband: nil,
		wife:    gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{}),
	},
}

func TestFamilyNode_Husband(t *testing.T) {
	for _, test := range familyTests {
		t.Run("", func(t *testing.T) {
			node := test.doc.Nodes()[0].(*gedcom.FamilyNode)
			node.SetDocument(test.doc)
			assert.Equal(t, node.Husband(), test.husband)
		})
	}

	Husband := tf.Function(t, (*gedcom.FamilyNode).Husband)

	Husband((*gedcom.FamilyNode)(nil)).Returns((*gedcom.IndividualNode)(nil))
}

func TestFamilyNode_Wife(t *testing.T) {
	for _, test := range familyTests {
		t.Run("", func(t *testing.T) {
			node := test.doc.Nodes()[0].(*gedcom.FamilyNode)
			assert.Equal(t, node.Wife(), test.wife)
		})
	}

	Wife := tf.Function(t, (*gedcom.FamilyNode).Wife)

	Wife((*gedcom.FamilyNode)(nil)).Returns((*gedcom.IndividualNode)(nil))
}

func TestFamilyNode_Similarity(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		doc      *gedcom.Document
		expected float64
	}{
		// Empty cases.
		{
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				gedcom.NewFamilyNode(nil, "F1", nil),
				gedcom.NewFamilyNode(nil, "F2", nil),
			}),
			expected: 0.5,
		},
		{
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{}),
				gedcom.NewFamilyNode(nil, "F2", []gedcom.Node{}),
			}),
			expected: 0.5,
		},

		// Perfect cases.
		{
			// All details match exactly.
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P1@", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P2@", "", nil),
				}),
				gedcom.NewFamilyNode(nil, "F2", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P3@", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P4@", "", nil),
				}),
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("Elliot Rupert de Peyster /Chance/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Dina Victoria /Wyche/"),
					born("Abt. Feb 1837"),
					died("8 Apr 1923"),
				}),
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Elliot Rupert de Peyster /Chance/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P4", []gedcom.Node{
					name("Dina Victoria /Wyche/"),
					born("Abt. Feb 1837"),
					died("8 Apr 1923"),
				}),
			}),
			expected: 1.0,
		},

		// Almost perfect matches.
		{
			// Name is more/less complete.
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P1@", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P2@", "", nil),
				}),
				gedcom.NewFamilyNode(nil, "F2", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P3@", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P4@", "", nil),
				}),
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("Elliot Rupert de Peyster /Chance/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Dina Victoria /Wyche/"),
					born("Abt. Feb 1837"),
				}),
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Elliot R. d. P. /Chance/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P4", []gedcom.Node{
					name("Dina /Wyche/"),
					born("Bef. Mar 1837"),
				}),
			}),
			expected: 0.8904318416381887,
		},

		// These ones are way off.
		{
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P1@", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P2@", "", nil),
				}),
				gedcom.NewFamilyNode(nil, "F2", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P3@", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P4@", "", nil),
				}),
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("Elliot Rupert de Peyster /Chance/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Dina Victoria /Wyche/"),
					born("Abt. Feb 1837"),
				}),
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					born("1627"),
				}),
				gedcom.NewIndividualNode(nil, "", "P4", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
			}),
			expected: 0.37700025152486955,
		},
	}

	options := gedcom.NewSimilarityOptions()
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			family1 := test.doc.Families()[0]
			family1.SetDocument(test.doc)

			family2 := test.doc.Families()[1]
			family2.SetDocument(test.doc)

			got := family1.Similarity(family2, 0, options)

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestFamilyNode_Children(t *testing.T) {
	Children := tf.Function(t, (*gedcom.FamilyNode).Children)

	Children((*gedcom.FamilyNode)(nil)).Returns((gedcom.IndividualNodes)(nil))
}

func TestFamilyNode_HasChild(t *testing.T) {
	HasChild := tf.Function(t, (*gedcom.FamilyNode).HasChild)

	HasChild((*gedcom.FamilyNode)(nil), (*gedcom.IndividualNode)(nil)).Returns(false)
}
