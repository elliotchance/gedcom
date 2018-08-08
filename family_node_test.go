package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestFamilyNode_Similarity(t *testing.T) {
	var tests = []struct {
		doc      *gedcom.Document
		expected float64
	}{
		// Empty cases.
		{
			doc: &gedcom.Document{
				Nodes: []gedcom.Node{
					gedcom.NewFamilyNode("F1", nil),
					gedcom.NewFamilyNode("F2", nil),
				},
			},
			expected: 0.5,
		},
		{
			doc: &gedcom.Document{
				Nodes: []gedcom.Node{
					gedcom.NewFamilyNode("F1", []gedcom.Node{}),
					gedcom.NewFamilyNode("F2", []gedcom.Node{}),
				},
			},
			expected: 0.5,
		},

		// Perfect cases.
		{
			// All details match exactly.
			doc: &gedcom.Document{
				Nodes: []gedcom.Node{
					gedcom.NewFamilyNode("F1", []gedcom.Node{
						gedcom.NewSimpleNode(gedcom.TagHusband, "@P1@", "", nil),
						gedcom.NewSimpleNode(gedcom.TagWife, "@P2@", "", nil),
					}),
					gedcom.NewFamilyNode("F2", []gedcom.Node{
						gedcom.NewSimpleNode(gedcom.TagHusband, "@P3@", "", nil),
						gedcom.NewSimpleNode(gedcom.TagWife, "@P4@", "", nil),
					}),
					gedcom.NewIndividualNode("", "P1", []gedcom.Node{
						name("Elliot Rupert de Peyster /Chance/"),
						born("4 Jan 1843"),
						died("17 Mar 1907"),
					}),
					gedcom.NewIndividualNode("", "P2", []gedcom.Node{
						name("Dina Victoria /Wyche/"),
						born("Abt. Feb 1837"),
						died("8 Apr 1923"),
					}),
					gedcom.NewIndividualNode("", "P3", []gedcom.Node{
						name("Elliot Rupert de Peyster /Chance/"),
						born("4 Jan 1843"),
						died("17 Mar 1907"),
					}),
					gedcom.NewIndividualNode("", "P4", []gedcom.Node{
						name("Dina Victoria /Wyche/"),
						born("Abt. Feb 1837"),
						died("8 Apr 1923"),
					}),
				},
			},
			expected: 1.0,
		},

		// Almost perfect matches.
		{
			// Name is more/less complete.
			doc: &gedcom.Document{
				Nodes: []gedcom.Node{
					gedcom.NewFamilyNode("F1", []gedcom.Node{
						gedcom.NewSimpleNode(gedcom.TagHusband, "@P1@", "", nil),
						gedcom.NewSimpleNode(gedcom.TagWife, "@P2@", "", nil),
					}),
					gedcom.NewFamilyNode("F2", []gedcom.Node{
						gedcom.NewSimpleNode(gedcom.TagHusband, "@P3@", "", nil),
						gedcom.NewSimpleNode(gedcom.TagWife, "@P4@", "", nil),
					}),
					gedcom.NewIndividualNode("", "P1", []gedcom.Node{
						name("Elliot Rupert de Peyster /Chance/"),
						born("4 Jan 1843"),
						died("17 Mar 1907"),
					}),
					gedcom.NewIndividualNode("", "P2", []gedcom.Node{
						name("Dina Victoria /Wyche/"),
						born("Abt. Feb 1837"),
					}),
					gedcom.NewIndividualNode("", "P3", []gedcom.Node{
						name("Elliot R. d. P. /Chance/"),
						born("4 Jan 1843"),
						died("17 Mar 1907"),
					}),
					gedcom.NewIndividualNode("", "P4", []gedcom.Node{
						name("Dina /Wyche/"),
						born("Bef. Mar 1837"),
					}),
				},
			},
			expected: 0.8347467653467052,
		},

		// These ones are way off.
		{
			doc: &gedcom.Document{
				Nodes: []gedcom.Node{
					gedcom.NewFamilyNode("F1", []gedcom.Node{
						gedcom.NewSimpleNode(gedcom.TagHusband, "@P1@", "", nil),
						gedcom.NewSimpleNode(gedcom.TagWife, "@P2@", "", nil),
					}),
					gedcom.NewFamilyNode("F2", []gedcom.Node{
						gedcom.NewSimpleNode(gedcom.TagHusband, "@P3@", "", nil),
						gedcom.NewSimpleNode(gedcom.TagWife, "@P4@", "", nil),
					}),
					gedcom.NewIndividualNode("", "P1", []gedcom.Node{
						name("Elliot Rupert de Peyster /Chance/"),
						born("4 Jan 1843"),
						died("17 Mar 1907"),
					}),
					gedcom.NewIndividualNode("", "P2", []gedcom.Node{
						name("Dina Victoria /Wyche/"),
						born("Abt. Feb 1837"),
					}),
					gedcom.NewIndividualNode("", "P3", []gedcom.Node{
						name("Bob /Jones/"),
						born("1627"),
					}),
					gedcom.NewIndividualNode("", "P4", []gedcom.Node{
						name("Jane /Doe/"),
						born("Sep 1845"),
					}),
				},
			},
			expected: 0.36400720887980753,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			family1 := test.doc.Families()[0]
			family2 := test.doc.Families()[1]
			got := family1.Similarity(test.doc, test.doc, family2, 0)

			assert.Equal(t, test.expected, got)
		})
	}
}
