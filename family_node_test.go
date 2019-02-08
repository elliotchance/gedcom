package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

var familyTests = []struct {
	doc     func(*gedcom.Document)
	husband *gedcom.IndividualNode
	wife    *gedcom.IndividualNode
}{
	{
		doc: func(doc *gedcom.Document) {
			doc.AddFamily("F1")
		},
		husband: nil,
		wife:    nil,
	},
	{
		doc: func(doc *gedcom.Document) {
			elliot := individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
			f1 := doc.AddFamily("F1")
			f1.SetHusband(elliot)
		},
		husband: elliot,
		wife:    nil,
	},
	{
		doc: func(doc *gedcom.Document) {
			jane := individual(doc, "P2", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
			f2 := doc.AddFamily("F2")
			f2.SetWife(jane)
		},
		husband: nil,
		wife:    jane,
	},
}

func TestFamilyNode_Husband(t *testing.T) {
	for _, test := range familyTests {
		t.Run("", func(t *testing.T) {
			doc := gedcom.NewDocument()
			test.doc(doc)
			node := doc.Families()[0]

			actualHusband := node.Husband()
			if test.husband == nil {
				assert.Nil(t, actualHusband)
			} else {
				if assert.NotNil(t, actualHusband) {
					assertEqual(t, actualHusband.Individual(), test.husband)
				}
			}
		})
	}

	Husband := tf.Function(t, (*gedcom.FamilyNode).Husband)

	Husband((*gedcom.FamilyNode)(nil)).Returns((*gedcom.HusbandNode)(nil))
}

func TestFamilyNode_Wife(t *testing.T) {
	for _, test := range familyTests {
		t.Run("", func(t *testing.T) {
			doc := gedcom.NewDocument()
			test.doc(doc)
			node := doc.Families()[0]

			actualWife := node.Wife()
			if test.wife == nil {
				assert.Nil(t, actualWife)
			} else {
				if assert.NotNil(t, actualWife) {
					assertEqual(t, actualWife.Individual(), test.wife)
				}
			}
		})
	}

	Wife := tf.Function(t, (*gedcom.FamilyNode).Wife)

	Wife((*gedcom.FamilyNode)(nil)).Returns((*gedcom.WifeNode)(nil))
}

func TestFamilyNode_Similarity(t *testing.T) {
	// ghost:ignore
	var tests = map[string]struct {
		doc      func(*gedcom.Document)
		expected float64
	}{
		// Empty cases.
		"Empty1": {
			doc: func(doc *gedcom.Document) {
				doc.AddFamily("F1")
				doc.AddFamily("F2")
			},
			expected: 0.5,
		},
		"Empty2": {
			doc: func(doc *gedcom.Document) {
				doc.AddFamily("F1")
				doc.AddFamily("F2")
			},
			expected: 0.5,
		},

		// Perfect cases.
		"Perfect": {
			// All details match exactly.
			doc: func(doc *gedcom.Document) {
				p1 := doc.AddIndividual("P1",
					gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				)
				p2 := doc.AddIndividual("P2",
					gedcom.NewNameNode("Dina Victoria /Wyche/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Abt. Feb 1837")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("8 Apr 1923")),
				)
				p3 := doc.AddIndividual("P3",
					gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				)
				p4 := doc.AddIndividual("P4",
					gedcom.NewNameNode("Dina Victoria /Wyche/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Abt. Feb 1837")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("8 Apr 1923")),
				)

				doc.AddFamilyWithHusbandAndWife("F1", p1, p2)
				doc.AddFamilyWithHusbandAndWife("F2", p3, p4)
			},
			expected: 1.0,
		},

		// Almost perfect matches.
		"AlmostPerfect": {
			// Name is more/less complete.
			doc: func(doc *gedcom.Document) {
				p1 := doc.AddIndividual("P1",
					gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				)
				p2 := doc.AddIndividual("P2",
					gedcom.NewNameNode("Dina Victoria /Wyche/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Abt. Feb 1837")),
				)
				p3 := doc.AddIndividual("P3",
					gedcom.NewNameNode("Elliot R. d. P. /Chance/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				)
				p4 := doc.AddIndividual("P4",
					gedcom.NewNameNode("Dina /Wyche/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Bef. Mar 1837")),
				)

				doc.AddFamilyWithHusbandAndWife("F1", p1, p2)
				doc.AddFamilyWithHusbandAndWife("F2", p3, p4)
			},
			expected: 0.8904318416381887,
		},

		// These ones are way off.
		"WayOff": {
			doc: func(doc *gedcom.Document) {
				p1 := doc.AddIndividual("P1",
					gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				)
				p2 := doc.AddIndividual("P2",
					gedcom.NewNameNode("Dina Victoria /Wyche/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Abt. Feb 1837")),
				)
				p3 := doc.AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("1627")),
				)
				p4 := doc.AddIndividual("P4",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				)

				doc.AddFamilyWithHusbandAndWife("F1", p1, p2)
				doc.AddFamilyWithHusbandAndWife("F2", p3, p4)
			},
			expected: 0.37700025152486955,
		},
	}

	options := gedcom.NewSimilarityOptions()
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			doc := gedcom.NewDocument()
			test.doc(doc)
			family1 := doc.Families().ByPointer("F1")
			family2 := doc.Families().ByPointer("F2")

			got := family1.Similarity(family2, 0, options)

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestFamilyNode_Children(t *testing.T) {
	Children := tf.Function(t, (*gedcom.FamilyNode).Children)

	Children((*gedcom.FamilyNode)(nil)).Returns((gedcom.ChildNodes)(nil))
}

func TestFamilyNode_HasChild(t *testing.T) {
	HasChild := tf.Function(t, (*gedcom.FamilyNode).HasChild)

	HasChild((*gedcom.FamilyNode)(nil), (*gedcom.IndividualNode)(nil)).Returns(false)
}
