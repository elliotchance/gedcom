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
	sex   gedcom.Sex
}{
	{
		node:  individual("P1", "", "", ""),
		names: []*gedcom.NameNode{},
		sex:   gedcom.SexUnknown,
	},
	{
		node:  gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
		names: []*gedcom.NameNode{},
		sex:   gedcom.SexUnknown,
	},
	{
		node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
			gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
		},
		sex: gedcom.SexUnknown,
	},
	{
		node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
			gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewNodeWithChildren(nil, gedcom.TagVersion, "", "", []gedcom.Node{}),
			gedcom.NewNameNode(nil, "John /Doe/", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewNameNode(nil, "John /Doe/", "", []gedcom.Node{}),
		},
		sex: gedcom.SexUnknown,
	},
	{
		node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
			gedcom.NewNodeWithChildren(nil, gedcom.TagSex, "M", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{},
		sex:   gedcom.SexMale,
	},
	{
		node: gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
			gedcom.NewNameNode(nil, "Joan /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewNodeWithChildren(nil, gedcom.TagSex, "F", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Joan /Bloggs/", "", []gedcom.Node{}),
		},
		sex: gedcom.SexFemale,
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

	Sex((*gedcom.IndividualNode)(nil)).Returns(gedcom.SexUnknown)
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
			node:   individual("P1", "", "", ""),
			births: nil,
		},
		{
			node:   gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			births: nil,
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			}),
			births: []*gedcom.BirthNode{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			births: []*gedcom.BirthNode{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewBirthNode(nil, "bar", "", []gedcom.Node{}),
			}),
			births: []*gedcom.BirthNode{
				gedcom.NewBirthNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewBirthNode(nil, "bar", "", []gedcom.Node{}),
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
			node:     individual("P1", "", "", ""),
			baptisms: []*gedcom.BaptismNode{},
		},
		{
			node:     gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			baptisms: []*gedcom.BaptismNode{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []*gedcom.BaptismNode{
				gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []*gedcom.BaptismNode{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			baptisms: []*gedcom.BaptismNode{
				gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "bar", "", []gedcom.Node{}),
			}),
			baptisms: []*gedcom.BaptismNode{
				gedcom.NewBaptismNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewBaptismNode(nil, "bar", "", []gedcom.Node{}),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Baptisms(), test.baptisms)
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
			node:   individual("P1", "", "", ""),
			deaths: []*gedcom.DeathNode{},
		},
		{
			node:   gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			deaths: []*gedcom.DeathNode{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			deaths: []*gedcom.DeathNode{
				gedcom.NewDeathNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			}),
			deaths: []*gedcom.DeathNode{
				gedcom.NewDeathNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "bar", "", []gedcom.Node{}),
			}),
			deaths: []*gedcom.DeathNode{
				gedcom.NewDeathNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewDeathNode(nil, "bar", "", []gedcom.Node{}),
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
			node:    individual("P1", "", "", ""),
			burials: []*gedcom.BurialNode{},
		},
		{
			node:    gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			burials: []*gedcom.BurialNode{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "", "", []gedcom.Node{}),
			}),
			burials: []*gedcom.BurialNode{
				gedcom.NewBurialNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "", "", []gedcom.Node{}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			}),
			burials: []*gedcom.BurialNode{
				gedcom.NewBurialNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "bar", "", []gedcom.Node{}),
			}),
			burials: []*gedcom.BurialNode{
				gedcom.NewBurialNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewBurialNode(nil, "bar", "", []gedcom.Node{}),
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

	p1 := individual("P1", "", "", "")
	p2 := individual("P2", "", "", "")
	p3 := gedcom.NewIndividualNode(nil, "", "P3", nil)
	p4 := gedcom.NewIndividualNode(nil, "", "P4", nil)
	p5 := gedcom.NewIndividualNode(nil, "", "P5", nil)
	p6 := gedcom.NewIndividualNode(nil, "", "P6", nil)
	p7 := gedcom.NewIndividualNode(nil, "", "P7", nil)
	p8 := gedcom.NewIndividualNode(nil, "", "P8", nil)

	// P1 - P3
	//    |
	//  -----
	// P4   P5
	f1 := gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P1@", "", nil),
		gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P3@", "", nil),
		gedcom.NewNodeWithChildren(nil, gedcom.TagChild, "@P4@", "", nil),
		gedcom.NewNodeWithChildren(nil, gedcom.TagChild, "@P5@", "", nil),
	})

	// P1 - P2
	//    |
	//   P6
	f2 := gedcom.NewFamilyNode(nil, "F2", []gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P1@", "", nil),
		gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P2@", "", nil),
		gedcom.NewNodeWithChildren(nil, gedcom.TagChild, "@P6@", "", nil),
	})

	// P6 - ?
	//    |
	//   P7
	f3 := gedcom.NewFamilyNode(nil, "F3", []gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHusband, "@P6@", "", nil),
		gedcom.NewNodeWithChildren(nil, gedcom.TagChild, "@P7@", "", nil),
	})

	// ? - P3
	//   |
	//   P6
	f4 := gedcom.NewFamilyNode(nil, "F4", []gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagWife, "@P3@", "", nil),
		gedcom.NewNodeWithChildren(nil, gedcom.TagChild, "@P6@", "", nil),
	})

	return gedcom.NewDocumentWithNodes([]gedcom.Node{
		p1, p2, p3, p4, p5, p6, p7, p8,
		f1, f2, f3, f4,
	})
}

func TestIndividualNode_Parents(t *testing.T) {
	doc := getDocument()

	var tests = []struct {
		node    *gedcom.IndividualNode
		parents []*gedcom.FamilyNode
	}{
		{
			node:    nil,
			parents: nil,
		},
		{
			node:    doc.Individuals()[0],
			parents: []*gedcom.FamilyNode{},
		},
		{
			node:    doc.Individuals()[1],
			parents: []*gedcom.FamilyNode{},
		},
		{
			node:    doc.Individuals()[2],
			parents: []*gedcom.FamilyNode{},
		},
		{
			node:    doc.Individuals()[3],
			parents: []*gedcom.FamilyNode{doc.Families()[0]},
		},
		{
			node:    doc.Individuals()[4],
			parents: []*gedcom.FamilyNode{doc.Families()[0]},
		},
		{
			node:    doc.Individuals()[5],
			parents: []*gedcom.FamilyNode{doc.Families()[1], doc.Families()[3]},
		},
		{
			node:    doc.Individuals()[6],
			parents: []*gedcom.FamilyNode{doc.Families()[2]},
		},
		{
			node:    doc.Individuals()[7],
			parents: []*gedcom.FamilyNode{},
		},
	}

	for _, test := range tests {
		t.Run(gedcom.Pointer(test.node), func(t *testing.T) {
			for _, n := range doc.Nodes() {
				n.SetDocument(doc)
			}
			assert.Equal(t, test.node.Parents(), test.parents)
		})
	}
}

func TestIndividualNode_SpouseChildren(t *testing.T) {
	doc := getDocument()

	var tests = []struct {
		node     *gedcom.IndividualNode
		expected gedcom.SpouseChildren
	}{
		{
			node:     nil,
			expected: gedcom.SpouseChildren{},
		},
		{
			node: doc.Individuals()[0],
			expected: gedcom.SpouseChildren{
				doc.Individuals()[2]: {
					doc.Individuals()[3],
					doc.Individuals()[4],
				},
				doc.Individuals()[1]: {
					doc.Individuals()[5],
				},
			},
		},
		{
			node: doc.Individuals()[1],
			expected: gedcom.SpouseChildren{
				doc.Individuals()[0]: {
					doc.Individuals()[5],
				},
			},
		},
		{
			node: doc.Individuals()[2],
			expected: gedcom.SpouseChildren{
				doc.Individuals()[0]: {
					doc.Individuals()[3],
					doc.Individuals()[4],
				},
				nil: {
					doc.Individuals()[5],
				},
			},
		},
		{
			node:     doc.Individuals()[3],
			expected: gedcom.SpouseChildren{},
		},
		{
			node:     doc.Individuals()[4],
			expected: gedcom.SpouseChildren{},
		},
		{
			node: doc.Individuals()[5],
			expected: gedcom.SpouseChildren{
				nil: {
					doc.Individuals()[6],
				},
			},
		},
		{
			node:     doc.Individuals()[6],
			expected: gedcom.SpouseChildren{},
		},
		{
			node:     doc.Individuals()[7],
			expected: gedcom.SpouseChildren{},
		},
	}

	for _, test := range tests {
		t.Run(gedcom.Pointer(test.node), func(t *testing.T) {
			for _, n := range doc.Nodes() {
				n.SetDocument(doc)
			}
			assert.Equal(t, test.expected, test.node.SpouseChildren())
		})
	}
}

func TestIndividualNode_LDSBaptisms(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.IndividualNode
		baptisms []gedcom.Node
	}{
		{
			node:     nil,
			baptisms: nil,
		},
		{
			node:     individual("P1", "", "", ""),
			baptisms: []gedcom.Node{},
		},
		{
			node:     gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			baptisms: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "bar", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "bar", "", []gedcom.Node{}),
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
		node     *gedcom.IndividualNode
		expected *gedcom.DateNode
	}{
		// Nil
		{
			node:     nil,
			expected: nil,
		},

		// No dates
		{
			node:     individual("P1", "", "", ""),
			expected: nil,
		},
		{
			node:     gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			expected: nil,
		},

		// A single date.
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "Abt. Dec 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Abt. Dec 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "Abt. Nov 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Abt. Nov 1980", "", nil),
		},

		// Multiple dates and other cases.
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "Abt. Jan 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Abt. Jan 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
					gedcom.NewDateNode(nil, "23 Mar 1979", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "23 Mar 1979", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "23 Mar 1979", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "23 Mar 1979", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := test.node.EstimatedBirthDate()

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
		node     *gedcom.IndividualNode
		expected *gedcom.DateNode
	}{
		// Nil
		{
			node:     nil,
			expected: nil,
		},

		// No dates
		{
			node:     individual("P1", "", "", ""),
			expected: nil,
		},
		{
			node:     gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			expected: nil,
		},

		// A single date.
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "Abt. Dec 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Abt. Dec 1980", "", nil),
		},

		// Multiple dates and other cases.
		{
			// Multiple death dates always returns the earliest.
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
					gedcom.NewDateNode(nil, "Mar 1980", "", nil),
					gedcom.NewDateNode(nil, "Jun 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Mar 1980", "", nil),
		},
		{
			// Multiple burial dates always returns the earliest.
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "3 Aug 1980", "", nil),
					gedcom.NewDateNode(nil, "Apr 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Apr 1980", "", nil),
		},
		{
			// Death is before burial.
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "3 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
		},
		{
			// Burial is before death.
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "3 Aug 1980", "", nil),
				}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "3 Aug 1980", "", nil),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := test.node.EstimatedDeathDate()

			if got == nil {
				assert.Nil(t, test.expected)
			} else {
				assert.Equal(t, got.SimpleNode, test.expected.SimpleNode)
			}
		})
	}
}

func born(value string) *gedcom.BirthNode {
	return gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
}

func died(value string) gedcom.Node {
	return gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
}

func name(value string) gedcom.Node {
	return gedcom.NewNameNode(nil, value, "", nil)
}

func baptised(value string) gedcom.Node {
	return gedcom.NewNodeWithChildren(nil, gedcom.TagBaptism, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
}

func buried(value string) gedcom.Node {
	return gedcom.NewNodeWithChildren(nil, gedcom.TagBurial, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
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
			b:        individual("P1", "", "", ""),
			expected: 0.5,
		},
		{
			a:        individual("P1", "", "", ""),
			b:        nil,
			expected: 0.5,
		},
		{
			a:        individual("P1", "", "", ""),
			b:        individual("P1", "", "", ""),
			expected: 0.25,
		},
		{
			a:        gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			b:        gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			expected: 0.25,
		},

		// Perfect cases.
		{
			// All details match exactly.
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 1.0,
		},
		{
			// Extra names, but one name is still a perfect match.
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				name("Elliot Rupert /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot R d P /Chance/"),
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 1.0,
		},
		{
			// Name are not senstive to case or whitespace.
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("elliot /CHANCE/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 1.0,
		},

		// Almost perfect matches.
		{
			// Name is more/less complete.
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 0.9831720430107527,
		},
		{
			// Last name is similar.
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chaunce/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 0.997883064516129,
		},
		{
			// Birth date is less specific.
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 0.9999701394487746,
		},
		{
			// Death date is less specific.
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("Mar 1907"),
			}),
			expected: 0.999999792635061,
		},

		// Estimated birth/death.
		{
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				baptised("Abt. 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("Mar 1907"),
			}),
			expected: 0.9933556126223895,
		},
		{
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				baptised("Abt. 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				buried("Aft. 20 Mar 1907"),
			}),
			expected: 0.9933539537028769,
		},

		// Missing dates.
		{
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				died("Abt. 1907"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert /Chance/"),
				died("1909"),
			}),
			expected: 0.7470609318996415,
		},
		{
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				baptised("after Sep 1823"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert /Chance/"),
				born("Between 1822 and 1823"),
			}),
			expected: 0.8443154512111212,
		},
		{
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Elliot Rupert /Chance/"),
			}),
			expected: 0.7331720430107527,
		},

		// These ones are way off.
		{
			a: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Jane /Doe/"),
				born("Sep 1845"),
			}),
			b: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				name("Bob /Jones/"),
				born("1627"),
			}),
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
		doc      *gedcom.Document
		expected *gedcom.SurroundingSimilarity
	}{
		"EmptyIndividuals": {
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "", "", ""),
				individual("P2", "", "", ""),
			}),

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
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
			}),
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
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chance/", "Abt. 1843", "Abt. 1910"),
			}),
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
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Joe /Bloggs/", "1945", "2000"),
			}),

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
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				family("F1", "P3", "P4", "P1"),
				family("F2", "P5", "P6", "P2"),
			}),
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
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				family("F1", "P3", "P4", "P1"),
				family("F2", "P5", "P6", "P2"),
			}),
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
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				family("F1", "P3", "", "P1"),
				family("F2", "P5", "P6", "P2"),
			}),
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
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				family("F1", "", "", "P1"),
				family("F2", "P5", "P6", "P2"),
			}),
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
			doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				individual("P7", "Jane /Bloggs/", "8 Mar 1803", "14 June 1877"),
				individual("P8", "Jane /Bloggs/", "8 Mar 1803", "14 June 1877"),
				family("F1", "P3", "P4", "P1"),
				family("F2", "P5", "P6", "P2"),
				family("F3", "P1", "P7"),
				family("F4", "P2", "P8"),
			}),
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
			for _, n := range test.doc.Nodes() {
				n.SetDocument(test.doc)
			}

			a := test.doc.Individuals()[0]
			b := test.doc.Individuals()[1]
			s := a.SurroundingSimilarity(b, options, false)

			assert.Equal(t, test.expected, s)
		})
	}
}

func individual(pointer, fullName, birth, death string) *gedcom.IndividualNode {
	nodes := []gedcom.Node{}

	if fullName != "" {
		nodes = append(nodes, name(fullName))
	}

	if birth != "" {
		nodes = append(nodes, born(birth))
	}

	if death != "" {
		nodes = append(nodes, died(death))
	}

	return gedcom.NewIndividualNode(nil, "", pointer, nodes)
}

func family(pointer, husband, wife string, children ...string) *gedcom.FamilyNode {
	nodes := []gedcom.Node{}

	if husband != "" {
		nodes = append(nodes, gedcom.NewNodeWithChildren(nil,
			gedcom.TagHusband, "@"+husband+"@", "", nil))
	}

	if wife != "" {
		nodes = append(nodes, gedcom.NewNodeWithChildren(nil,
			gedcom.TagWife, "@"+wife+"@", "", nil))
	}

	for _, child := range children {
		nodes = append(nodes, gedcom.NewNodeWithChildren(nil,
			gedcom.TagChild, "@"+child+"@", "", nil))
	}

	return gedcom.NewFamilyNode(nil, pointer, nodes)
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

	Families((*gedcom.IndividualNode)(nil)).Returns(([]*gedcom.FamilyNode)(nil))
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

	IsLiving(gedcom.NewIndividualNode(nil, "", "", nil)).Returns(true)

	IsLiving(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewDeathNode(nil, "", "", nil),
	})).Returns(false)

	IsLiving(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "3 Sep 1845", "", nil),
		}),
	})).Returns(false)

	IsLiving(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "3 Sep 1945", "", nil),
		}),
	})).Returns(true)

	doc := gedcom.NewDocument()
	IsLiving(gedcom.NewIndividualNode(doc, "", "", []gedcom.Node{
		gedcom.NewBirthNode(doc, "", "", []gedcom.Node{
			gedcom.NewDateNode(doc, "3 Sep 1945", "", nil),
		}),
	})).Returns(true)

	doc.MaxLivingAge = 25
	IsLiving(gedcom.NewIndividualNode(doc, "", "", []gedcom.Node{
		gedcom.NewBirthNode(doc, "", "", []gedcom.Node{
			gedcom.NewDateNode(doc, "3 Sep 1945", "", nil),
		}),
	})).Returns(false)

	doc.MaxLivingAge = 0
	IsLiving(gedcom.NewIndividualNode(doc, "", "", []gedcom.Node{
		gedcom.NewBirthNode(doc, "", "", []gedcom.Node{
			gedcom.NewDateNode(doc, "3 Sep 1945", "", nil),
		}),
	})).Returns(true)
}

func TestIndividualNode_Children(t *testing.T) {
	Children := tf.Function(t, (*gedcom.IndividualNode).Children)

	Children((*gedcom.IndividualNode)(nil)).Returns(gedcom.IndividualNodes{})
}

func TestIndividualNode_AllEvents(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node   *gedcom.IndividualNode
		events []gedcom.Node
	}{
		{
			node:   individual("P1", "", "", ""),
			events: nil,
		},
		{
			node:   gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			events: nil,
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			}),
			events: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagNote, "", "", []gedcom.Node{}),
			}),
			events: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			events: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewBirthNode(nil, "bar", "", []gedcom.Node{}),
			}),
			events: []gedcom.Node{
				gedcom.NewBirthNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewBirthNode(nil, "bar", "", []gedcom.Node{}),
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
	date3Sep1953 := gedcom.NewDateNode(nil, "3 Sep 1953", "", nil)
	place1 := gedcom.NewPlaceNode(nil, "Australia", "", nil)

	individual := gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			date3Sep1953,
		}),
		gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
			gedcom.NewPlaceNode(nil, "United Kingdom", "", nil),
		}),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			place1,
		}),
	})

	date, place := individual.Birth()

	assert.Equal(t, date3Sep1953, date)
	assert.Equal(t, place1, place)
}

func TestIndividualNode_Death(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode(nil, "3 Sep 1953", "", nil)
	place1 := gedcom.NewPlaceNode(nil, "Australia", "", nil)

	individual := gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
			date3Sep1953,
		}),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewPlaceNode(nil, "United Kingdom", "", nil),
		}),
		gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
			place1,
		}),
	})

	date, place := individual.Death()

	assert.Equal(t, date3Sep1953, date)
	assert.Equal(t, place1, place)
}

func TestIndividualNode_Baptism(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode(nil, "3 Sep 1953", "", nil)
	place1 := gedcom.NewPlaceNode(nil, "Australia", "", nil)

	individual := gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{
			date3Sep1953,
		}),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewPlaceNode(nil, "United Kingdom", "", nil),
		}),
		gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{
			place1,
		}),
	})

	date, place := individual.Baptism()

	assert.Equal(t, date3Sep1953, date)
	assert.Equal(t, place1, place)
}

func TestIndividualNode_Burial(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode(nil, "3 Sep 1953", "", nil)
	place1 := gedcom.NewPlaceNode(nil, "Australia", "", nil)

	individual := gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewBurialNode(nil, "", "", []gedcom.Node{
			date3Sep1953,
		}),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewPlaceNode(nil, "United Kingdom", "", nil),
		}),
		gedcom.NewBurialNode(nil, "", "", []gedcom.Node{
			place1,
		}),
	})

	date, place := individual.Burial()

	assert.Equal(t, date3Sep1953, date)
	assert.Equal(t, place1, place)
}

func TestIndividualNode_AgeAt(t *testing.T) {
	tests := []struct {
		individual string
		event      string
		start      string
		end        string
	}{
		{
			// No dates at all.
			individual: " - ",
			event:      "",
			start:      "= ?",
			end:        "= ?",
		},
		{
			// Event does not have any dates.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "",
			start:      "= ?",
			end:        "= ?",
		},
		{
			// Missing birth date for individual.
			individual: " - 2 Mar 2001",
			event:      "3 Sep 1945",
			start:      "= ?",
			end:        "= ?",
		},
		{
			// Approximate birth date makes the age an estimate.
			individual: "Abt. 1934 - 2 Mar 2001",
			event:      "3 Sep 1945",
			start:      "1945 - 1934 = ~10.7",
			end:        "1945 - 1934 = ~11.7",
		},
		{
			// Non-exact birth date makes the age an estimate.
			individual: "1934 - 2 Mar 2001",
			event:      "3 Sep 1945",
			start:      "1945 - 1934 = ~10.7",
			end:        "1945 - 1934 = ~11.7",
		},
		{
			// Event has an exact date. This is the best case scenario.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "12 Jan 1973",
			start:      "1973 - 1945 = 27.4",
			end:        "1973 - 1945 = 27.4",
		},
		{
			// Event has multiple exact dates. We must assume the min and max
			// dates create a possible range.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "12 Jan 1973, 14 Nov 1970, 7 Dec 1975",
			start:      "1970 - 1945 = 25.2",
			end:        "1975 - 1945 = 30.3",
		},
		{
			// Like the previous example we have several dates but not all of
			// them are exact.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "Abt. Mar 1973, After 14 Nov 1970, Abt. 1975",
			start:      "1970 - 1945 = ~25.2",
			end:        "1975 - 1945 = ~30.3",
		},
		{
			// There are two date ranges that partially overlap each other to
			// create a single larger range.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "Between 1965 and 1969, Between 1963 and 1967",
			start:      "1963 - 1945 = ~17.3",
			end:        "1969 - 1945 = ~24.3",
		},
		{
			// One of the dates is invalid.
			individual: "3 Sep 1945 - 2 Mar 2001",
			event:      "foo bar, Between 1963 and 1967",
			start:      "1963 - 1945 = ~17.3",
			end:        "1967 - 1945 = ~22.3",
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			individualParts := strings.Split(test.individual, "-")
			individual := gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, individualParts[0], "", nil),
				}),
				gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, individualParts[1], "", nil),
				}),
			})

			eventDates := []gedcom.Node{}
			if test.event != "" {
				for _, dateString := range strings.Split(test.event, ",") {
					dateNode := gedcom.NewDateNode(nil, dateString, "", nil)
					eventDates = append(eventDates, dateNode)
				}
			}

			event := gedcom.NewResidenceNode(nil, "", "", eventDates)

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

	String(gedcom.NewIndividualNode(nil, "", "", nil)).
		Returns("(no name)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "", "", nil),
	})).Returns("(no name)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
	})).Returns("Elliot Chance")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewBirthNode(nil, "", "", nil),
	})).Returns("Elliot Chance")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "3 Apr 1983", "", nil),
		}),
	})).Returns("Elliot Chance (b. 3 Apr 1983)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewDeathNode(nil, "", "", nil),
	})).Returns("Elliot Chance")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "19 Nov 2007", "", nil),
		}),
	})).Returns("Elliot Chance (d. 19 Nov 2007)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "John /Smith/", "", nil),
		gedcom.NewDeathNode(nil, "", "", nil),
		gedcom.NewBirthNode(nil, "", "", nil),
	})).Returns("John Smith")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "John /Smith/", "", nil),
		gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "19 Nov 2007", "", nil),
		}),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "7 Aug 1971", "", nil),
		}),
	})).Returns("John Smith (b. 7 Aug 1971, d. 19 Nov 2007)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Jane /Doe/", "", nil),
		gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "14 Jun 2007", "", nil),
		}),
	})).Returns("Jane Doe (bap. 14 Jun 2007)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Jane /Doe/", "", nil),
		gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "14 Jun 2007", "", nil),
		}),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "7 Jun 2007", "", nil),
		}),
	})).Returns("Jane Doe (b. 7 Jun 2007)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Jane /Doe/", "", nil),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "7 Jun 2007", "", nil),
		}),
		gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "14 Jun 2007", "", nil),
		}),
	})).Returns("Jane Doe (b. 7 Jun 2007)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Jane /Doe/", "", nil),
		gedcom.NewBurialNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "14 Jun 2007", "", nil),
		}),
	})).Returns("Jane Doe (bur. 14 Jun 2007)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Jane /Doe/", "", nil),
		gedcom.NewBurialNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "14 Jun 2007", "", nil),
		}),
		gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "7 Jun 2007", "", nil),
		}),
	})).Returns("Jane Doe (d. 7 Jun 2007)")

	String(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "Jane /Doe/", "", nil),
		gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "7 Jun 2007", "", nil),
		}),
		gedcom.NewBurialNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "14 Jun 2007", "", nil),
		}),
	})).Returns("Jane Doe (d. 7 Jun 2007)")
}

func TestIndividualNode_FamilySearchIDs(t *testing.T) {
	FamilySearchIDs := tf.NamedFunction(t, "IndividualNode_FamilySearchIDs",
		(*gedcom.IndividualNode).FamilySearchIDs)

	FamilySearchIDs(gedcom.NewIndividualNode(nil, "", "", nil)).
		Returns(nil)

	FamilySearchIDs(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "", "", nil),
	})).Returns(nil)

	FamilySearchIDs(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "", "", nil),
		gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID1, "LZDP-V7V"),
	})).Returns([]*gedcom.FamilySearchIDNode{
		gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID1, "LZDP-V7V"),
	})

	FamilySearchIDs(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID2, "AZDP-V7V"),
		gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID1, "BZDP-V7V"),
		gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID2, "CZDP-V7V"),
	})).Returns([]*gedcom.FamilySearchIDNode{
		gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID1, "BZDP-V7V"),
		gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID2, "AZDP-V7V"),
		gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID2, "CZDP-V7V"),
	})
}

func TestIndividualNode_UniqueIDs(t *testing.T) {
	UniqueIDs := tf.NamedFunction(t, "IndividualNode_UniqueIDs",
		(*gedcom.IndividualNode).UniqueIDs)

	UniqueIDs(gedcom.NewIndividualNode(nil, "", "", nil)).
		Returns(nil)

	UniqueIDs(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "", "", nil),
	})).Returns(nil)

	UniqueIDs(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewNameNode(nil, "", "", nil),
		gedcom.NewUniqueIDNode(nil, "LZDP-V7V", "", nil),
	})).Returns([]*gedcom.UniqueIDNode{
		gedcom.NewUniqueIDNode(nil, "LZDP-V7V", "", nil),
	})

	UniqueIDs(gedcom.NewIndividualNode(nil, "", "", []gedcom.Node{
		gedcom.NewUniqueIDNode(nil, "AZDP-V7V", "", nil),
		gedcom.NewNameNode(nil, "", "", nil),
		gedcom.NewUniqueIDNode(nil, "BZDP-V7V", "", nil),
	})).Returns([]*gedcom.UniqueIDNode{
		gedcom.NewUniqueIDNode(nil, "AZDP-V7V", "", nil),
		gedcom.NewUniqueIDNode(nil, "BZDP-V7V", "", nil),
	})
}
