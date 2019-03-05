package gedcom_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

var individualTests = []struct {
	node  *gedcom.IndividualNode
	names []*gedcom.NameNode
	sex   *gedcom.SexNode
}{
	{
		node:  individual(gedcom.NewDocument(), "P1", "", "", ""),
		names: []*gedcom.NameNode{},
		sex:   nil,
	},
	{
		node:  gedcom.NewDocument().AddIndividual("P1"),
		names: []*gedcom.NameNode{},
		sex:   nil,
	},
	{
		node: gedcom.NewDocument().AddIndividual("P1",
			gedcom.NewNameNode("Joe /Bloggs/"),
		),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode("Joe /Bloggs/"),
		},
		sex: nil,
	},
	{
		node: gedcom.NewDocument().AddIndividual("P1",
			gedcom.NewNameNode("Joe /Bloggs/"),
			gedcom.NewNode(gedcom.TagVersion, "", ""),
			gedcom.NewNameNode("John /Doe/"),
		),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode("Joe /Bloggs/"),
			gedcom.NewNameNode("John /Doe/"),
		},
		sex: nil,
	},
	{
		node: gedcom.NewDocument().AddIndividual("P1",
			gedcom.NewNode(gedcom.TagSex, "M", ""),
		),
		names: []*gedcom.NameNode{},
		sex:   gedcom.NewSexNode(gedcom.SexMale),
	},
	{
		node: gedcom.NewDocument().AddIndividual("P2",
			gedcom.NewNameNode("Joan /Bloggs/"),
			gedcom.NewNode(gedcom.TagSex, "F", ""),
		),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode("Joan /Bloggs/"),
		},
		sex: gedcom.NewSexNode(gedcom.SexFemale),
	},
}

func TestIndividualNode_Names(t *testing.T) {
	for _, test := range individualTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Names(), test.names)
		})
	}

	Names := tf.Function(t, (*gedcom.IndividualNode).Names)

	Names((*gedcom.IndividualNode)(nil)).Returns(([]*gedcom.NameNode)(nil))
}

func TestIndividualNode_Sex(t *testing.T) {
	for _, test := range individualTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Sex(), test.sex)
		})
	}

	Sex := tf.Function(t, (*gedcom.IndividualNode).Sex)

	Sex((*gedcom.IndividualNode)(nil)).Returns(nil)
}

func TestIndividualNode_Births(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node   *gedcom.IndividualNode
		births []*gedcom.BirthNode
	}{
		{
			node:   nil,
			births: nil,
		},
		{
			node:   individual(gedcom.NewDocument(), "P1", "", "", ""),
			births: nil,
		},
		{
			node:   gedcom.NewDocument().AddIndividual("P1"),
			births: nil,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode(""),
			),
			births: []*gedcom.BirthNode{
				gedcom.NewBirthNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode(""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			),
			births: []*gedcom.BirthNode{
				gedcom.NewBirthNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode("foo"),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewBirthNode("bar"),
			),
			births: []*gedcom.BirthNode{
				gedcom.NewBirthNode("foo"),
				gedcom.NewBirthNode("bar"),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Births(), test.births)
		})
	}
}

func TestIndividualNode_Baptisms(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.IndividualNode
		baptisms []*gedcom.BaptismNode
	}{
		{
			node:     nil,
			baptisms: []*gedcom.BaptismNode{},
		},
		{
			node:     individual(gedcom.NewDocument(), "P1", "", "", ""),
			baptisms: []*gedcom.BaptismNode{},
		},
		{
			node:     gedcom.NewDocument().AddIndividual("P1"),
			baptisms: []*gedcom.BaptismNode{},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBaptism, "", ""),
			),
			baptisms: []*gedcom.BaptismNode{
				gedcom.NewBaptismNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagLDSBaptism, "", ""),
			),
			baptisms: []*gedcom.BaptismNode{},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBaptism, "", ""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			),
			baptisms: []*gedcom.BaptismNode{
				gedcom.NewBaptismNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBaptism, "foo", ""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewNode(gedcom.TagBaptism, "bar", ""),
			),
			baptisms: []*gedcom.BaptismNode{
				gedcom.NewBaptismNode("foo"),
				gedcom.NewBaptismNode("bar"),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assertEqual(t, test.node.Baptisms(), test.baptisms)
		})
	}
}

func TestIndividualNode_Deaths(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node   *gedcom.IndividualNode
		deaths []*gedcom.DeathNode
	}{
		{
			node:   nil,
			deaths: []*gedcom.DeathNode{},
		},
		{
			node:   individual(gedcom.NewDocument(), "P1", "", "", ""),
			deaths: []*gedcom.DeathNode{},
		},
		{
			node:   gedcom.NewDocument().AddIndividual("P1"),
			deaths: []*gedcom.DeathNode{},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			),
			deaths: []*gedcom.DeathNode{
				gedcom.NewDeathNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewBirthNode(""),
			),
			deaths: []*gedcom.DeathNode{
				gedcom.NewDeathNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagDeath, "foo", ""),
				gedcom.NewNode(gedcom.TagBurial, "", ""),
				gedcom.NewNode(gedcom.TagDeath, "bar", ""),
			),
			deaths: []*gedcom.DeathNode{
				gedcom.NewDeathNode("foo"),
				gedcom.NewDeathNode("bar"),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Deaths(), test.deaths)
		})
	}
}

func TestIndividualNode_Burials(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node    *gedcom.IndividualNode
		burials []*gedcom.BurialNode
	}{
		{
			node:    nil,
			burials: []*gedcom.BurialNode{},
		},
		{
			node:    individual(gedcom.NewDocument(), "P1", "", "", ""),
			burials: []*gedcom.BurialNode{},
		},
		{
			node:    gedcom.NewDocument().AddIndividual("P1"),
			burials: []*gedcom.BurialNode{},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBurial, "", ""),
			),
			burials: []*gedcom.BurialNode{
				gedcom.NewBurialNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBurial, "", ""),
				gedcom.NewBirthNode(""),
			),
			burials: []*gedcom.BurialNode{
				gedcom.NewBurialNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBurial, "foo", ""),
				gedcom.NewNode(gedcom.TagBaptism, "", ""),
				gedcom.NewNode(gedcom.TagBurial, "bar", ""),
			),
			burials: []*gedcom.BurialNode{
				gedcom.NewBurialNode("foo"),
				gedcom.NewBurialNode("bar"),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Burials(), test.burials)
		})
	}
}

func getDocument() *gedcom.Document {
	// The following document has this tree:
	//
	//      ?  --- P3
	// P3 - P1 --- P2    P8
	//    |     |
	//  -----   |
	// P4   P5  P6 -- ?
	//              |
	//              P7
	//
	// - P3 and P2 are two wives of P1.
	// - P8 does not connect to anything.
	// - P3 is the alternate mother of P6.

	doc := gedcom.NewDocument()
	p1 := doc.AddIndividual("P1")
	p2 := doc.AddIndividual("P2")
	p3 := doc.AddIndividual("P3")
	p4 := doc.AddIndividual("P4")
	p5 := doc.AddIndividual("P5")
	p6 := doc.AddIndividual("P6")
	p7 := doc.AddIndividual("P7")
	doc.AddIndividual("P8")

	// P1 - P3
	//    |
	//  -----
	// P4   P5
	f1 := doc.AddFamilyWithHusbandAndWife("F1", p1, p3)
	f1.AddChild(p4)
	f1.AddChild(p5)

	// P1 - P2
	//    |
	//   P6
	f2 := doc.AddFamilyWithHusbandAndWife("F2", p1, p2)
	f2.AddChild(p6)

	// P6 - ?
	//    |
	//   P7
	f3 := doc.AddFamilyWithHusbandAndWife("F3", p6, nil)
	f3.AddChild(p7)

	// ? - P3
	//   |
	//   P6
	f4 := doc.AddFamilyWithHusbandAndWife("F4", nil, p3)
	f4.AddChild(p6)

	return doc
}

func TestIndividualNode_Parents(t *testing.T) {
	doc := getDocument()
	individuals := doc.Individuals()
	families := doc.Families()

	var tests = []struct {
		node    *gedcom.IndividualNode
		parents gedcom.FamilyNodes
	}{
		{
			node:    nil,
			parents: nil,
		},
		{
			node:    individuals.ByPointer("P1"),
			parents: gedcom.FamilyNodes{},
		},
		{
			node:    individuals.ByPointer("P2"),
			parents: gedcom.FamilyNodes{},
		},
		{
			node:    individuals.ByPointer("P3"),
			parents: gedcom.FamilyNodes{},
		},
		{
			node:    individuals.ByPointer("P4"),
			parents: gedcom.FamilyNodes{families.ByPointer("F1")},
		},
		{
			node:    individuals.ByPointer("P5"),
			parents: gedcom.FamilyNodes{families.ByPointer("F1")},
		},
		{
			node: individuals.ByPointer("P6"),
			parents: gedcom.FamilyNodes{
				families.ByPointer("F2"),
				families.ByPointer("F4"),
			},
		},
		{
			node:    individuals.ByPointer("P7"),
			parents: gedcom.FamilyNodes{families.ByPointer("F3")},
		},
		{
			node:    individuals.ByPointer("P8"),
			parents: gedcom.FamilyNodes{},
		},
	}

	for _, test := range tests {
		t.Run(gedcom.Pointer(test.node), func(t *testing.T) {
			assertEqual(t, test.parents, test.node.Parents())
		})
	}
}

func TestIndividualNode_SpouseChildren(t *testing.T) {
	doc := getDocument()
	individuals := doc.Individuals()
	families := doc.Families()

	var tests = map[string]struct {
		node     *gedcom.IndividualNode
		expected gedcom.SpouseChildren
	}{
		"Nil": {
			node:     nil,
			expected: gedcom.SpouseChildren{},
		},
		"P1": {
			node: individuals.ByPointer("P1"),
			expected: gedcom.SpouseChildren{
				individuals.ByPointer("P3"): {
					families.ByPointer("F1").Children().ByPointer("P4"),
					families.ByPointer("F1").Children().ByPointer("P5"),
				},
				individuals.ByPointer("P2"): {
					families.ByPointer("F2").Children().ByPointer("P6"),
				},
			},
		},
		"P2": {
			node: individuals.ByPointer("P2"),
			expected: gedcom.SpouseChildren{
				individuals.ByPointer("P1"): {
					families.ByPointer("F2").Children().ByPointer("P6"),
				},
			},
		},
		"P3": {
			node: individuals.ByPointer("P3"),
			expected: gedcom.SpouseChildren{
				individuals.ByPointer("P1"): {
					families.ByPointer("F1").Children().ByPointer("P4"),
					families.ByPointer("F1").Children().ByPointer("P5"),
				},
				nil: {
					families.ByPointer("F2").Children().ByPointer("P6"),
				},
			},
		},
		"P4": {
			node:     individuals.ByPointer("P4"),
			expected: gedcom.SpouseChildren{},
		},
		"P5": {
			node:     individuals.ByPointer("P5"),
			expected: gedcom.SpouseChildren{},
		},
		"P6": {
			node: individuals.ByPointer("P6"),
			expected: gedcom.SpouseChildren{
				nil: {
					families.ByPointer("F3").Children().ByPointer("P7"),
				},
			},
		},
		"P7": {
			node:     individuals.ByPointer("P7"),
			expected: gedcom.SpouseChildren{},
		},
		"P8": {
			node:     individuals.ByPointer("P8"),
			expected: gedcom.SpouseChildren{},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			assertEqual(t, test.expected, test.node.SpouseChildren())
		})
	}
}

func TestIndividualNode_LDSBaptisms(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.IndividualNode
		baptisms gedcom.Nodes
	}{
		{
			node:     nil,
			baptisms: nil,
		},
		{
			node:     individual(gedcom.NewDocument(), "P1", "", "", ""),
			baptisms: gedcom.Nodes{},
		},
		{
			node:     gedcom.NewDocument().AddIndividual("P1"),
			baptisms: gedcom.Nodes{},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagLDSBaptism, "", ""),
			),
			baptisms: gedcom.Nodes{
				gedcom.NewNode(gedcom.TagLDSBaptism, "", ""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBaptism, "", ""),
			),
			baptisms: gedcom.Nodes{},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagLDSBaptism, "", ""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			),
			baptisms: gedcom.Nodes{
				gedcom.NewNode(gedcom.TagLDSBaptism, "", ""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagLDSBaptism, "foo", ""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewNode(gedcom.TagLDSBaptism, "bar", ""),
			),
			baptisms: gedcom.Nodes{
				gedcom.NewNode(gedcom.TagLDSBaptism, "foo", ""),
				gedcom.NewNode(gedcom.TagLDSBaptism, "bar", ""),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.LDSBaptisms(), test.baptisms)
		})
	}
}

func TestIndividualNode_EstimatedBirthDate(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node         *gedcom.IndividualNode
		expected     *gedcom.DateNode
		isBirthEvent bool
	}{
		// Nil
		{
			node:         nil,
			expected:     nil,
			isBirthEvent: false,
		},

		// No dates
		{
			node:         individual(gedcom.NewDocument(), "P1", "", "", ""),
			expected:     nil,
			isBirthEvent: false,
		},
		{
			node:         gedcom.NewDocument().AddIndividual("P1"),
			expected:     nil,
			isBirthEvent: false,
		},

		// A single date.
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("1 Aug 1980"),
				),
			),
			expected:     gedcom.NewDateNode("1 Aug 1980"),
			isBirthEvent: true,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBaptism, "", "",
					gedcom.NewDateNode("Abt. Dec 1980"),
				),
			),
			expected:     gedcom.NewDateNode("Abt. Dec 1980"),
			isBirthEvent: false,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagLDSBaptism, "", "",
					gedcom.NewDateNode("Abt. Nov 1980"),
				),
			),
			expected:     gedcom.NewDateNode("Abt. Nov 1980"),
			isBirthEvent: false,
		},

		// Multiple dates and other cases.
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("1 Aug 1980"),
				),
				gedcom.NewNode(gedcom.TagBaptism, "", "",
					gedcom.NewDateNode("Abt. Jan 1980"),
				),
			),
			expected:     gedcom.NewDateNode("1 Aug 1980"),
			isBirthEvent: true,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("1 Aug 1980"),
					gedcom.NewDateNode("23 Mar 1979"),
				),
			),
			expected:     gedcom.NewDateNode("23 Mar 1979"),
			isBirthEvent: true,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("1 Aug 1980"),
				),
				gedcom.NewNode(gedcom.TagLDSBaptism, "", "",
					gedcom.NewDateNode("23 Mar 1979"),
				),
			),
			expected:     gedcom.NewDateNode("1 Aug 1980"),
			isBirthEvent: true,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode(""),
				gedcom.NewNode(gedcom.TagLDSBaptism, "", ""),
			),
			expected:     nil,
			isBirthEvent: false,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode(""),
				gedcom.NewNode(gedcom.TagLDSBaptism, "", "",
					gedcom.NewDateNode("1 Aug 1980"),
				),
			),
			expected:     gedcom.NewDateNode("1 Aug 1980"),
			isBirthEvent: false,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got, isBirthEvent := test.node.EstimatedBirthDate()

			assert.Equal(t, test.isBirthEvent, isBirthEvent)

			if got == nil {
				assert.Nil(t, test.expected)
			} else {
				assert.Equal(t, got.SimpleNode, test.expected.SimpleNode)
			}
		})
	}
}

func TestIndividualNode_EstimatedDeathDate(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node         *gedcom.IndividualNode
		expected     *gedcom.DateNode
		isDeathEvent bool
	}{
		// Nil
		{
			node:         nil,
			expected:     nil,
			isDeathEvent: false,
		},

		// No dates
		{
			node:         individual(gedcom.NewDocument(), "P1", "", "", ""),
			expected:     nil,
			isDeathEvent: false,
		},
		{
			node:         gedcom.NewDocument().AddIndividual("P1"),
			expected:     nil,
			isDeathEvent: false,
		},

		// A single date.
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagDeath, "", "",
					gedcom.NewDateNode("1 Aug 1980"),
				),
			),
			expected:     gedcom.NewDateNode("1 Aug 1980"),
			isDeathEvent: true,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBurial, "", "",
					gedcom.NewDateNode("Abt. Dec 1980"),
				),
			),
			expected:     gedcom.NewDateNode("Abt. Dec 1980"),
			isDeathEvent: false,
		},

		// Multiple dates and other cases.
		{
			// Multiple death dates always returns the earliest.
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagDeath, "", "",
					gedcom.NewDateNode("1 Aug 1980"),
					gedcom.NewDateNode("Mar 1980"),
					gedcom.NewDateNode("Jun 1980"),
				),
			),
			expected:     gedcom.NewDateNode("Mar 1980"),
			isDeathEvent: true,
		},
		{
			// Multiple burial dates always returns the earliest.
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagBurial, "", "",
					gedcom.NewDateNode("3 Aug 1980"),
					gedcom.NewDateNode("Apr 1980"),
				),
			),
			expected:     gedcom.NewDateNode("Apr 1980"),
			isDeathEvent: false,
		},
		{
			// Death is before burial.
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagDeath, "", "",
					gedcom.NewDateNode("1 Aug 1980"),
				),
				gedcom.NewNode(gedcom.TagBurial, "", "",
					gedcom.NewDateNode("3 Aug 1980"),
				),
			),
			expected:     gedcom.NewDateNode("1 Aug 1980"),
			isDeathEvent: true,
		},
		{
			// Burial is before death.
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNode(gedcom.TagDeath, "", "",
					gedcom.NewDateNode("3 Aug 1980"),
				),
				gedcom.NewNode(gedcom.TagBurial, "", "",
					gedcom.NewDateNode("1 Aug 1980"),
				),
			),
			expected:     gedcom.NewDateNode("3 Aug 1980"),
			isDeathEvent: true,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got, isDeathEvent := test.node.EstimatedDeathDate()

			assert.Equal(t, test.isDeathEvent, isDeathEvent)

			if got == nil {
				assert.Nil(t, test.expected)
			} else {
				assert.Equal(t, got.SimpleNode, test.expected.SimpleNode)
			}
		})
	}
}

func TestIndividualNode_Similarity(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		a, b     *gedcom.IndividualNode
		expected float64
	}{
		// Empty cases.
		{
			a:        nil,
			b:        nil,
			expected: 0.5,
		},
		{
			a:        nil,
			b:        individual(gedcom.NewDocument(), "P1", "", "", ""),
			expected: 0.5,
		},
		{
			a:        individual(gedcom.NewDocument(), "P1", "", "", ""),
			b:        nil,
			expected: 0.5,
		},
		{
			a:        individual(gedcom.NewDocument(), "P1", "", "", ""),
			b:        individual(gedcom.NewDocument(), "P1", "", "", ""),
			expected: 0.25,
		},
		{
			a:        gedcom.NewDocument().AddIndividual("P1"),
			b:        gedcom.NewDocument().AddIndividual("P1"),
			expected: 0.25,
		},

		// Perfect cases.
		{
			// All details match exactly.
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			expected: 1.0,
		},
		{
			// Extra names, but one name is still a perfect match.
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewNameNode("Elliot Rupert /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot R d P /Chance/"),
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			expected: 1.0,
		},
		{
			// Name are not senstive to case or whitespace.
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("elliot /CHANCE/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			expected: 1.0,
		},

		// Almost perfect matches.
		{
			// Name is more/less complete.
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			expected: 0.9831720430107527,
		},
		{
			// Last name is similar.
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chaunce/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			expected: 0.997883064516129,
		},
		{
			// Birth date is less specific.
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			expected: 0.9999701394487746,
		},
		{
			// Death date is less specific.
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("Mar 1907")),
			),
			expected: 0.999999792635061,
		},

		// Estimated birth/death.
		{
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBaptismNode("", gedcom.NewDateNode("Abt. 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("Mar 1907")),
			),
			expected: 0.9933556126223895,
		},
		{
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBaptismNode("", gedcom.NewDateNode("Abt. 1843")),
				gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
				gedcom.NewBurialNode("", gedcom.NewDateNode("Aft. 20 Mar 1907")),
			),
			expected: 0.9933539537028769,
		},

		// Missing dates.
		{
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewDeathNode("", gedcom.NewDateNode("Abt. 1907")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert /Chance/"),
				gedcom.NewDeathNode("", gedcom.NewDateNode("1909")),
			),
			expected: 0.7470609318996415,
		},
		{
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
				gedcom.NewBaptismNode("", gedcom.NewDateNode("after Sep 1823")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert /Chance/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("Between 1822 and 1823")),
			),
			expected: 0.8443154512111212,
		},
		{
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot Rupert /Chance/"),
			),
			expected: 0.7331720430107527,
		},

		// These ones are way off.
		{
			a: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Jane /Doe/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
			),
			b: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Bob /Jones/"),
				gedcom.NewBirthNode("", gedcom.NewDateNode("1627")),
			),
			expected: 0.38125,
		},
	}

	options := gedcom.NewSimilarityOptions()
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := test.a.Similarity(test.b, options)

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestIndividualNode_SurroundingSimilarity(t *testing.T) {
	// ghost:ignore
	var tests = map[string]struct {
		doc      func(*gedcom.Document)
		expected *gedcom.SurroundingSimilarity
	}{
		"EmptyIndividuals": {
			doc: func(doc *gedcom.Document) {
				doc.AddIndividual("P1")
				doc.AddIndividual("P2")
			},

			// These are the real values, but they are not calculated because
			// the weighted similarity would be less than the minimum threshold.
			//
			//   expected: &gedcom.SurroundingSimilarity{
			//     ParentsSimilarity:    0.5,
			//     IndividualSimilarity: 0.25,
			//     SpousesSimilarity:    1.0,
			//     ChildrenSimilarity:   1.0,
			//     Options:              gedcom.NewSimilarityOptions(),
			//   },
			//
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.0,
				IndividualSimilarity: 0.0,
				SpousesSimilarity:    0.0,
				ChildrenSimilarity:   0.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},

		// Only matching individuals, but they are exact matches.
		"Matching1": {
			doc: func(doc *gedcom.Document) {
				individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P2", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
			},
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.5,
				IndividualSimilarity: 1.0,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},

		// Only matching individuals, but they are similar matches.
		"Matching2": {
			doc: func(doc *gedcom.Document) {
				individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P2", "Elliot /Chance/", "Abt. 1843", "Abt. 1910")
			},
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.5,
				IndividualSimilarity: 0.7433558199873285,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},

		// Only matching individuals and they are way off.
		"Matching3": {
			doc: func(doc *gedcom.Document) {
				individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P2", "Joe /Bloggs/", "1945", "2000")
			},

			// These are the real values, but they are not calculated because
			// the weighted similarity would be less than the minimum threshold.
			//
			//   expected: &gedcom.SurroundingSimilarity{
			//     ParentsSimilarity:    0.5,
			//     IndividualSimilarity: 0.20128205128205132,
			//     SpousesSimilarity:    1.0,
			//     ChildrenSimilarity:   1.0,
			//     Options:              gedcom.NewSimilarityOptions(),
			//   },
			//
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.0,
				IndividualSimilarity: 0.0,
				SpousesSimilarity:    0.0,
				ChildrenSimilarity:   0.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},

		// Parents and individuals match exactly.
		"Parents1": {
			doc: func(doc *gedcom.Document) {
				individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P2", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P4", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
				individual(doc, "P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
				family(doc, "F1", "P3", "P4", "P1")
				family(doc, "F2", "P5", "P6", "P2")
			},
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    1.0,
				IndividualSimilarity: 1.0,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},

		// Parents and individuals are very similar.
		"Parents2": {
			doc: func(doc *gedcom.Document) {
				individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877")
				individual(doc, "P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
				family(doc, "F1", "P3", "P4", "P1")
				family(doc, "F2", "P5", "P6", "P2")
			},
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.9981481481481481,
				IndividualSimilarity: 0.9950549450549451,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},

		// One parent is missing, otherwise exactly the same.
		"Parents3": {
			doc: func(doc *gedcom.Document) {
				individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877")
				individual(doc, "P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
				family(doc, "F1", "P3", "", "P1")
				family(doc, "F2", "P5", "P6", "P2")
			},
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.75,
				IndividualSimilarity: 0.9950549450549451,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},

		// Both parents are missing on one side, otherwise exactly the same.
		"Parents4": {
			doc: func(doc *gedcom.Document) {
				individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877")
				individual(doc, "P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
				family(doc, "F1", "", "", "P1")
				family(doc, "F2", "P5", "P6", "P2")
			},
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.5,
				IndividualSimilarity: 0.9950549450549451,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},

		// Parents, individual and spouses match exactly.
		"Parents5": {
			doc: func(doc *gedcom.Document) {
				individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P2", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
				individual(doc, "P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P4", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
				individual(doc, "P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
				individual(doc, "P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
				individual(doc, "P7", "Jane /Bloggs/", "8 Mar 1803", "14 June 1877")
				individual(doc, "P8", "Jane /Bloggs/", "8 Mar 1803", "14 June 1877")
				family(doc, "F1", "P3", "P4", "P1")
				family(doc, "F2", "P5", "P6", "P2")
				family(doc, "F3", "P1", "P7")
				family(doc, "F4", "P2", "P8")
			},
			expected: &gedcom.SurroundingSimilarity{
				ParentsSimilarity:    1.0,
				IndividualSimilarity: 1.0,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
				Options:              gedcom.NewSimilarityOptions(),
			},
		},
	}

	options := gedcom.NewSimilarityOptions()
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			doc := gedcom.NewDocument()
			test.doc(doc)
			a := doc.Individuals()[0]
			b := doc.Individuals()[1]
			s := a.SurroundingSimilarity(b, options, false)

			assert.Equal(t, test.expected, s)
		})
	}
}

func individual(doc *gedcom.Document, pointer, fullName, birth, death string) *gedcom.IndividualNode {
	nodes := gedcom.Nodes{}

	if fullName != "" {
		nodes = append(nodes, gedcom.NewNameNode(fullName))
	}

	if birth != "" {
		nodes = append(nodes, gedcom.NewBirthNode("", gedcom.NewDateNode(birth)))
	}

	if death != "" {
		nodes = append(nodes, gedcom.NewDeathNode("", gedcom.NewDateNode(death)))
	}

	return doc.AddIndividual(pointer, nodes...)
}

func family(doc *gedcom.Document, pointer, husband, wife string, children ...string) *gedcom.FamilyNode {
	f := doc.AddFamily(pointer)

	if husband != "" {
		f.SetHusband(doc.Individuals().ByPointer(husband))
	}

	if wife != "" {
		f.SetWife(doc.Individuals().ByPointer(wife))
	}

	for _, child := range children {
		f.AddChild(doc.Individuals().ByPointer(child))
	}

	return f
}

func TestIndividualNode_Name(t *testing.T) {
	Name := tf.Function(t, (*gedcom.IndividualNode).Name)

	Name((*gedcom.IndividualNode)(nil)).Returns((*gedcom.NameNode)(nil))
}

func TestIndividualNode_Spouses(t *testing.T) {
	Spouses := tf.Function(t, (*gedcom.IndividualNode).Spouses)

	Spouses((*gedcom.IndividualNode)(nil)).Returns((gedcom.IndividualNodes)(nil))
}

func TestIndividualNode_Families(t *testing.T) {
	Families := tf.Function(t, (*gedcom.IndividualNode).Families)

	Families((*gedcom.IndividualNode)(nil)).Returns((gedcom.FamilyNodes)(nil))
}

func TestIndividualNode_FamilyWithSpouse(t *testing.T) {
	FamilyWithSpouse := tf.Function(t, (*gedcom.IndividualNode).FamilyWithSpouse)

	FamilyWithSpouse((*gedcom.IndividualNode)(nil), (*gedcom.IndividualNode)(nil)).Returns((*gedcom.FamilyNode)(nil))
}

func TestIndividualNode_FamilyWithUnknownSpouse(t *testing.T) {
	FamilyWithUnknownSpouse := tf.Function(t, (*gedcom.IndividualNode).FamilyWithUnknownSpouse)

	FamilyWithUnknownSpouse((*gedcom.IndividualNode)(nil)).Returns((*gedcom.FamilyNode)(nil))
}

func TestIndividualNode_IsLiving(t *testing.T) {
	IsLiving := tf.Function(t, (*gedcom.IndividualNode).IsLiving)

	IsLiving(nil).Returns(false)

	IsLiving(gedcom.NewDocument().AddIndividual("")).Returns(true)

	IsLiving(gedcom.NewDocument().AddIndividual("",
		gedcom.NewDeathNode(""),
	)).Returns(false)

	IsLiving(gedcom.NewDocument().AddIndividual("",
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("3 Sep 1845"),
		),
	)).Returns(false)

	IsLiving(gedcom.NewDocument().AddIndividual("",
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("3 Sep 1945"),
		),
	)).Returns(true)

	doc := gedcom.NewDocument()
	IsLiving(doc.AddIndividual("",
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("3 Sep 1945"),
		),
	)).Returns(true)

	doc.MaxLivingAge = 25
	IsLiving(doc.AddIndividual("",
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("3 Sep 1945"),
		),
	)).Returns(false)

	doc.MaxLivingAge = 0
	IsLiving(doc.AddIndividual("",
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("3 Sep 1945"),
		),
	)).Returns(true)
}

func TestIndividualNode_Children(t *testing.T) {
	Children := tf.Function(t, (*gedcom.IndividualNode).Children)

	Children((*gedcom.IndividualNode)(nil)).Returns(gedcom.ChildNodes{})
}

func TestIndividualNode_AllEvents(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node   *gedcom.IndividualNode
		events gedcom.Nodes
	}{
		{
			node:   individual(gedcom.NewDocument(), "P1", "", "", ""),
			events: nil,
		},
		{
			node:   gedcom.NewDocument().AddIndividual("P1"),
			events: nil,
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode(""),
			),
			events: gedcom.Nodes{
				gedcom.NewBirthNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode(""),
				gedcom.NewNode(gedcom.TagNote, "", ""),
			),
			events: gedcom.Nodes{
				gedcom.NewBirthNode(""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode(""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			),
			events: gedcom.Nodes{
				gedcom.NewBirthNode(""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			},
		},
		{
			node: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode("foo"),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewBirthNode("bar"),
			),
			events: gedcom.Nodes{
				gedcom.NewBirthNode("foo"),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewBirthNode("bar"),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.AllEvents(), test.events)
		})
	}
}

func TestIndividualNode_Birth(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode("3 Sep 1953")
	place1 := gedcom.NewPlaceNode("Australia")

	individual := gedcom.NewDocument().AddIndividual("",
		gedcom.NewBirthNode("",
			date3Sep1953,
		),
		gedcom.NewDeathNode("",
			gedcom.NewPlaceNode("United Kingdom"),
		),
		gedcom.NewBirthNode("",
			place1,
		),
	)

	date, place := individual.Birth()

	assert.Equal(t, date3Sep1953, date)
	assert.Equal(t, place1, place)
}

func TestIndividualNode_Death(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode("3 Sep 1953")
	place1 := gedcom.NewPlaceNode("Australia")

	individual := gedcom.NewDocument().AddIndividual("",
		gedcom.NewDeathNode("",
			date3Sep1953,
		),
		gedcom.NewBirthNode("",
			gedcom.NewPlaceNode("United Kingdom"),
		),
		gedcom.NewDeathNode("",
			place1,
		),
	)

	date, place := individual.Death()

	assert.Equal(t, date3Sep1953, date)
	assert.Equal(t, place1, place)
}

func TestIndividualNode_Baptism(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode("3 Sep 1953")
	place1 := gedcom.NewPlaceNode("Australia")

	individual := gedcom.NewDocument().AddIndividual("",
		gedcom.NewBaptismNode("", date3Sep1953),
		gedcom.NewBirthNode("",
			gedcom.NewPlaceNode("United Kingdom"),
		),
		gedcom.NewBaptismNode("", place1),
	)

	date, place := individual.Baptism()

	assert.Equal(t, date3Sep1953, date)
	assert.Equal(t, place1, place)
}

func TestIndividualNode_Burial(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode("3 Sep 1953")
	place1 := gedcom.NewPlaceNode("Australia")

	individual := gedcom.NewDocument().AddIndividual("",
		gedcom.NewBurialNode("",
			date3Sep1953,
		),
		gedcom.NewBirthNode("",
			gedcom.NewPlaceNode("United Kingdom"),
		),
		gedcom.NewBurialNode("",
			place1,
		),
	)

	date, place := individual.Burial()

	assert.Equal(t, date3Sep1953, date)
	assert.Equal(t, place1, place)
}

func TestIndividualNode_AgeAt(t *testing.T) {
	tests := map[string]struct {
		individual string
		event      string
		start      string
		end        string
	}{
		"Missing1": {
			// No dates at all.
			individual: " - ",
			event:      "",
			start:      "= ?",
			end:        "= ?",
		},
		"Missing2": {
			// Event does not have any dates.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "",
			start:      "= ?",
			end:        "= ?",
		},
		"Missing3": {
			// Missing birth date for individual.
			individual: " - 2 Mar 2001",
			event:      "3 Sep 1945",
			start:      "= ?",
			end:        "= ?",
		},
		"Approx1": {
			// Approximate birth date makes the age an estimate.
			individual: "Abt. 1934 - 2 Mar 2001",
			event:      "3 Sep 1945",
			start:      "1945 - 1934 = ~10.7",
			end:        "1945 - 1934 = ~11.7",
		},
		"Approx2": {
			// Non-exact birth date makes the age an estimate.
			individual: "1934 - 2 Mar 2001",
			event:      "3 Sep 1945",
			start:      "1945 - 1934 = ~10.7",
			end:        "1945 - 1934 = ~11.7",
		},
		"Exact1": {
			// Event has an exact date. This is the best case scenario.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "12 Jan 1973",
			start:      "1973 - 1945 = 27.4",
			end:        "1973 - 1945 = 27.4",
		},
		"Exact2": {
			// Event has multiple exact dates. We must assume the min and max
			// dates create a possible range.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "12 Jan 1973, 14 Nov 1970, 7 Dec 1975",
			start:      "1970 - 1945 = 25.2",
			end:        "1975 - 1945 = 30.3",
		},
		"Approx3": {
			// Like the previous example we have several dates but not all of
			// them are exact.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "Abt. Mar 1973, After 14 Nov 1970, Abt. 1975",
			start:      "1970 - 1945 = ~25.2",
			end:        "1975 - 1945 = ~30.3",
		},
		"Approx4": {
			// There are two date ranges that partially overlap each other to
			// create a single larger range.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "Between 1965 and 1969, Between 1963 and 1967",
			start:      "1963 - 1945 = ~17.3",
			end:        "1969 - 1945 = ~24.3",
		},
		"Invalid1": {
			// One of the dates is invalid.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "foo bar, Between 1963 and 1967",
			start:      "1963 - 1945 = ~17.3",
			end:        "1967 - 1945 = ~22.3",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			individualParts := strings.Split(test.individual, "-")
			individual := gedcom.NewDocument().AddIndividual("",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode(individualParts[0]),
				),
				gedcom.NewDeathNode("",
					gedcom.NewDateNode(individualParts[1]),
				),
			)

			eventDates := gedcom.Nodes{}
			if test.event != "" {
				for _, dateString := range strings.Split(test.event, ",") {
					dateNode := gedcom.NewDateNode(dateString)
					eventDates = append(eventDates, dateNode)
				}
			}

			event := gedcom.NewResidenceNode("", eventDates...)

			startYears, startIsEstimate, startIsKnown := parseAgeYears(test.start)
			endYears, endIsEstimate, endIsKnown := parseAgeYears(test.end)

			start, end := individual.AgeAt(event)

			if startIsKnown {
				assertAge(t, start, startYears, startIsEstimate, gedcom.AgeConstraintLiving)
			} else {
				assert.False(t, start.IsKnown, "start is known")
			}

			if endIsKnown {
				assertAge(t, end, endYears, endIsEstimate, gedcom.AgeConstraintLiving)
			} else {
				assert.False(t, end.IsKnown, "end is known")
			}
		})
	}
}

func parseAgeYears(s string) (years float64, isEstimate, isKnown bool) {
	isKnown = true
	value := strings.Split(s, "=")[1]
	value = strings.TrimSpace(value)

	if value[0] == '~' {
		isEstimate = true
		value = value[1:]
	}

	if value[0] == '?' {
		isKnown = false
		return
	}

	fmt.Sscanf(value, "%f", &years)

	return
}

func TestIndividualNode_String(t *testing.T) {
	String := tf.NamedFunction(t, "IndividualNode_String", (*gedcom.IndividualNode).String)

	String(gedcom.NewDocument().AddIndividual("")).
		Returns("(no name)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode(""),
	)).Returns("(no name)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Elliot /Chance/"),
	)).Returns("Elliot Chance")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewBirthNode(""),
	)).Returns("Elliot Chance")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("3 Apr 1983"),
		),
	)).Returns("Elliot Chance (b. 3 Apr 1983)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewDeathNode(""),
	)).Returns("Elliot Chance")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewDeathNode("",
			gedcom.NewDateNode("19 Nov 2007"),
		),
	)).Returns("Elliot Chance (d. 19 Nov 2007)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("John /Smith/"),
		gedcom.NewDeathNode(""),
		gedcom.NewBirthNode(""),
	)).Returns("John Smith")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("John /Smith/"),
		gedcom.NewDeathNode("",
			gedcom.NewDateNode("19 Nov 2007"),
		),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("7 Aug 1971"),
		),
	)).Returns("John Smith (b. 7 Aug 1971, d. 19 Nov 2007)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Jane /Doe/"),
		gedcom.NewBaptismNode("",
			gedcom.NewDateNode("14 Jun 2007"),
		),
	)).Returns("Jane Doe (bap. 14 Jun 2007)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Jane /Doe/"),
		gedcom.NewBaptismNode("",
			gedcom.NewDateNode("14 Jun 2007"),
		),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("7 Jun 2007"),
		),
	)).Returns("Jane Doe (b. 7 Jun 2007)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Jane /Doe/"),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("7 Jun 2007"),
		),
		gedcom.NewBaptismNode("",
			gedcom.NewDateNode("14 Jun 2007"),
		),
	)).Returns("Jane Doe (b. 7 Jun 2007)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Jane /Doe/"),
		gedcom.NewBurialNode("",
			gedcom.NewDateNode("14 Jun 2007"),
		),
	)).Returns("Jane Doe (bur. 14 Jun 2007)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Jane /Doe/"),
		gedcom.NewBurialNode("",
			gedcom.NewDateNode("14 Jun 2007"),
		),
		gedcom.NewDeathNode("",
			gedcom.NewDateNode("7 Jun 2007"),
		),
	)).Returns("Jane Doe (d. 7 Jun 2007)")

	String(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode("Jane /Doe/"),
		gedcom.NewDeathNode("",
			gedcom.NewDateNode("7 Jun 2007"),
		),
		gedcom.NewBurialNode("",
			gedcom.NewDateNode("14 Jun 2007"),
		),
	)).Returns("Jane Doe (d. 7 Jun 2007)")
}

func TestIndividualNode_FamilySearchIDs(t *testing.T) {
	FamilySearchIDs := tf.NamedFunction(t, "IndividualNode_FamilySearchIDs",
		(*gedcom.IndividualNode).FamilySearchIDs)

	FamilySearchIDs(gedcom.NewDocument().AddIndividual("")).
		Returns(nil)

	FamilySearchIDs(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode(""),
	)).Returns(nil)

	FamilySearchIDs(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode(""),
		gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID1, "LZDP-V7V"),
	)).Returns([]*gedcom.FamilySearchIDNode{
		gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID1, "LZDP-V7V"),
	})

	FamilySearchIDs(gedcom.NewDocument().AddIndividual("",
		gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID2, "AZDP-V7V"),
		gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID1, "BZDP-V7V"),
		gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID2, "CZDP-V7V"),
	)).Returns([]*gedcom.FamilySearchIDNode{
		gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID1, "BZDP-V7V"),
		gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID2, "AZDP-V7V"),
		gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID2, "CZDP-V7V"),
	})
}

func TestIndividualNode_UniqueIDs(t *testing.T) {
	UniqueIDs := tf.NamedFunction(t, "IndividualNode_UniqueIDs",
		(*gedcom.IndividualNode).UniqueIDs)

	UniqueIDs(gedcom.NewDocument().AddIndividual("")).
		Returns(nil)

	UniqueIDs(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode(""),
	)).Returns(nil)

	UniqueIDs(gedcom.NewDocument().AddIndividual("",
		gedcom.NewNameNode(""),
		gedcom.NewUniqueIDNode("LZDP-V7V"),
	)).Returns([]*gedcom.UniqueIDNode{
		gedcom.NewUniqueIDNode("LZDP-V7V"),
	})

	UniqueIDs(gedcom.NewDocument().AddIndividual("",
		gedcom.NewUniqueIDNode("AZDP-V7V"),
		gedcom.NewNameNode(""),
		gedcom.NewUniqueIDNode("BZDP-V7V"),
	)).Returns([]*gedcom.UniqueIDNode{
		gedcom.NewUniqueIDNode("AZDP-V7V"),
		gedcom.NewUniqueIDNode("BZDP-V7V"),
	})
}

var individualWarningTests = map[string]struct {
	doc      func(doc *gedcom.Document)
	expected []string
}{
	"None": {
		func(doc *gedcom.Document) {
			individual(doc, "P1", "Elliot /Chance/", "", "16 May 1989")
		},
		nil,
	},
	"DeathBeforeBurial": {
		func(doc *gedcom.Document) {
			elliot := individual(doc, "P1", "Elliot /Chance/", "", "1 May 1989")
			elliot.AddBurialDate("16 May 1989")
		},
		nil,
	},
	"BurialBeforeDeath": {
		func(doc *gedcom.Document) {
			elliot := individual(doc, "P1", "Elliot /Chance/", "", "16 May 1989")
			elliot.AddBurialDate("1 May 1989")
		},
		[]string{"The burial (1 May 1989) was before the death (16 May 1989) of Elliot Chance (d. 16 May 1989)."},
	},
	"BirthBeforeBaptism": {
		func(doc *gedcom.Document) {
			elliot := individual(doc, "P1", "Elliot /Chance/", "1 May 1989", "")
			elliot.AddBaptismDate("16 May 1989")
		},
		nil,
	},
	"BaptismBeforeBirth": {
		func(doc *gedcom.Document) {
			elliot := individual(doc, "P1", "Elliot /Chance/", "16 May 1989", "")
			elliot.AddBaptismDate("1 May 1989")
		},
		[]string{"The baptism (1 May 1989) was before the birth (16 May 1989) of Elliot Chance (b. 16 May 1989)."},
	},
	"TooOldUnknownSex": {
		func(doc *gedcom.Document) {
			individual(doc, "P1", "Elliot /Chance/", "16 May 1889", "2000")
		},
		[]string{"Elliot Chance (b. 16 May 1889, d. 2000) was 111 years old at the time of their death."},
	},
	"TooOldMale": {
		func(doc *gedcom.Document) {
			elliot := individual(doc, "P1", "Elliot /Chance/", "16 May 1889", "2000")
			elliot.SetSex(gedcom.SexMale)
		},
		[]string{"Elliot Chance (b. 16 May 1889, d. 2000) was 111 years old at the time of his death."},
	},
	"NotTooOldWithoutDeath": {
		func(doc *gedcom.Document) {
			individual(doc, "P1", "Elliot /Chance/", "16 May 1789", "")
		},
		nil,
	},
	"NotTooOldFemale": {
		func(doc *gedcom.Document) {
			elliot := individual(doc, "P1", "Jane /Chance/", "12 Sep 1913", "7 Feb 2001")
			elliot.SetSex(gedcom.SexFemale)
		},
		nil,
	},
}

func TestIndividualNode_Warnings(t *testing.T) {
	for testName, test := range individualWarningTests {
		t.Run(testName, func(t *testing.T) {
			doc := gedcom.NewDocument()
			test.doc(doc)

			p1 := doc.Individuals().ByPointer("P1")
			assertEqual(t, p1.Warnings().Strings(), test.expected)
		})
	}
}
