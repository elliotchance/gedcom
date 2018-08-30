package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
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
			gedcom.NewSimpleNode(nil, gedcom.TagVersion, "", "", []gedcom.Node{}),
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
			gedcom.NewSimpleNode(nil, gedcom.TagSex, "M", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{},
		sex:   gedcom.SexMale,
	},
	{
		node: gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
			gedcom.NewNameNode(nil, "Joan /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewSimpleNode(nil, gedcom.TagSex, "F", "", []gedcom.Node{}),
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
}

func TestIndividualNode_Sex(t *testing.T) {
	for _, test := range individualTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Sex(), test.sex)
		})
	}
}

func TestIndividualNode_Births(t *testing.T) {
	var tests = []struct {
		node   *gedcom.IndividualNode
		births []*gedcom.BirthNode
	}{
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
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			births: []*gedcom.BirthNode{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
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
	var tests = []struct {
		node     *gedcom.IndividualNode
		baptisms []gedcom.Node
	}{
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
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "bar", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "bar", "", []gedcom.Node{}),
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
	var tests = []struct {
		node   *gedcom.IndividualNode
		deaths []gedcom.Node
	}{
		{
			node:   individual("P1", "", "", ""),
			deaths: []gedcom.Node{},
		},
		{
			node:   gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			deaths: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			deaths: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			}),
			deaths: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "bar", "", []gedcom.Node{}),
			}),
			deaths: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "bar", "", []gedcom.Node{}),
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
	var tests = []struct {
		node    *gedcom.IndividualNode
		burials []gedcom.Node
	}{
		{
			node:    individual("P1", "", "", ""),
			burials: []gedcom.Node{},
		},
		{
			node:    gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
			burials: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{}),
			}),
			burials: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
			}),
			burials: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "bar", "", []gedcom.Node{}),
			}),
			burials: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "bar", "", []gedcom.Node{}),
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
		gedcom.NewSimpleNode(nil, gedcom.TagHusband, "@P1@", "", nil),
		gedcom.NewSimpleNode(nil, gedcom.TagWife, "@P3@", "", nil),
		gedcom.NewSimpleNode(nil, gedcom.TagChild, "@P4@", "", nil),
		gedcom.NewSimpleNode(nil, gedcom.TagChild, "@P5@", "", nil),
	})

	// P1 - P2
	//    |
	//   P6
	f2 := gedcom.NewFamilyNode(nil, "F2", []gedcom.Node{
		gedcom.NewSimpleNode(nil, gedcom.TagHusband, "@P1@", "", nil),
		gedcom.NewSimpleNode(nil, gedcom.TagWife, "@P2@", "", nil),
		gedcom.NewSimpleNode(nil, gedcom.TagChild, "@P6@", "", nil),
	})

	// P6 - ?
	//    |
	//   P7
	f3 := gedcom.NewFamilyNode(nil, "F3", []gedcom.Node{
		gedcom.NewSimpleNode(nil, gedcom.TagHusband, "@P6@", "", nil),
		gedcom.NewSimpleNode(nil, gedcom.TagChild, "@P7@", "", nil),
	})

	// ? - P3
	//   |
	//   P6
	f4 := gedcom.NewFamilyNode(nil, "F4", []gedcom.Node{
		gedcom.NewSimpleNode(nil, gedcom.TagWife, "@P3@", "", nil),
		gedcom.NewSimpleNode(nil, gedcom.TagChild, "@P6@", "", nil),
	})

	return &gedcom.Document{
		Nodes: []gedcom.Node{
			p1, p2, p3, p4, p5, p6, p7, p8,
			f1, f2, f3, f4,
		},
	}
}

func TestIndividualNode_Parents(t *testing.T) {
	doc := getDocument()

	var tests = []struct {
		node    *gedcom.IndividualNode
		parents []*gedcom.FamilyNode
	}{
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
		t.Run(test.node.Pointer(), func(t *testing.T) {
			for _, n := range doc.Nodes {
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
		t.Run(test.node.Pointer(), func(t *testing.T) {
			for _, n := range doc.Nodes {
				n.SetDocument(doc)
			}
			assert.Equal(t, test.expected, test.node.SpouseChildren())
		})
	}
}

func TestIndividualNode_LDSBaptisms(t *testing.T) {
	var tests = []struct {
		node     *gedcom.IndividualNode
		baptisms []gedcom.Node
	}{
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
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "bar", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "bar", "", []gedcom.Node{}),
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
	var tests = []struct {
		node     *gedcom.IndividualNode
		expected *gedcom.DateNode
	}{
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
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "Abt. Dec 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Abt. Dec 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{
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
				gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{
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
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "23 Mar 1979", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "23 Mar 1979", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagLDSBaptism, "", "", []gedcom.Node{
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
	var tests = []struct {
		node     *gedcom.IndividualNode
		expected *gedcom.DateNode
	}{
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
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "Abt. Dec 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Abt. Dec 1980", "", nil),
		},

		// Multiple dates and other cases.
		{
			// Multiple death dates always returns the earliest.
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{
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
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "3 Aug 1980", "", nil),
					gedcom.NewDateNode(nil, "Apr 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "Apr 1980", "", nil),
		},
		{
			// Death is before burial.
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
				}),
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "3 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode(nil, "1 Aug 1980", "", nil),
		},
		{
			// Burial is before death.
			node: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "3 Aug 1980", "", nil),
				}),
				gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{
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
	return gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
}

func name(value string) gedcom.Node {
	return gedcom.NewNameNode(nil, value, "", nil)
}

func baptised(value string) gedcom.Node {
	return gedcom.NewSimpleNode(nil, gedcom.TagBaptism, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
}

func buried(value string) gedcom.Node {
	return gedcom.NewSimpleNode(nil, gedcom.TagBurial, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
}

func TestIndividualNode_Similarity(t *testing.T) {
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
	var tests = []struct {
		doc      *gedcom.Document
		expected gedcom.SurroundingSimilarity
	}{
		// Empty individuals.
		{
			doc: document(
				individual("P1", "", "", ""),
				individual("P2", "", "", ""),
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.5,
				IndividualSimilarity: 0.25,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},

		// Only matching individuals, but they are exact matches.
		{
			doc: document(
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.5,
				IndividualSimilarity: 1.0,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},

		// Only matching individuals, but they are similar matches.
		{
			doc: document(
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chance/", "Abt. 1843", "Abt. 1910"),
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.5,
				IndividualSimilarity: 0.7433558199873285,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},

		// Only matching individuals and they are way off.
		{
			doc: document(
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Joe /Bloggs/", "1945", "2000"),
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.5,
				IndividualSimilarity: 0.20128205128205132,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},

		// Parents and individuals match exactly.
		{
			doc: document(
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				family("F1", "P3", "P4", "P1"),
				family("F2", "P5", "P6", "P2"),
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    1.0,
				IndividualSimilarity: 1.0,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},

		// Parents and individuals are very similar.
		{
			doc: document(
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				family("F1", "P3", "P4", "P1"),
				family("F2", "P5", "P6", "P2"),
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.9981481481481481,
				IndividualSimilarity: 0.9950549450549451,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},

		// One parent is missing, otherwise exactly the same.
		{
			doc: document(
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				family("F1", "P3", "", "P1"),
				family("F2", "P5", "P6", "P2"),
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.75,
				IndividualSimilarity: 0.9950549450549451,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},

		// Both parents are missing on one side, otherwise exactly the same.
		{
			doc: document(
				individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907"),
				individual("P2", "Elliot /Chaunce/", "4 Jan 1843", "17 Mar 1907"),
				individual("P3", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P4", "Jane /Doey/", "3 Mar 1803", "14 June 1877"),
				individual("P5", "John /Smith/", "4 Jan 1803", "17 Mar 1877"),
				individual("P6", "Jane /Doe/", "3 Mar 1803", "14 June 1877"),
				family("F1", "", "", "P1"),
				family("F2", "P5", "P6", "P2"),
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    0.5,
				IndividualSimilarity: 0.9950549450549451,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},

		// Parents, individual and spouses match exactly.
		{
			doc: document(
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
			),
			expected: gedcom.SurroundingSimilarity{
				ParentsSimilarity:    1.0,
				IndividualSimilarity: 1.0,
				SpousesSimilarity:    1.0,
				ChildrenSimilarity:   1.0,
			},
		},
	}

	options := gedcom.NewSimilarityOptions()
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			for _, n := range test.doc.Nodes {
				n.SetDocument(test.doc)
			}

			a := test.doc.Individuals()[0]
			b := test.doc.Individuals()[1]
			s := a.SurroundingSimilarity(b, options)

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
		nodes = append(nodes, gedcom.NewSimpleNode(nil,
			gedcom.TagHusband, "@"+husband+"@", "", nil))
	}

	if wife != "" {
		nodes = append(nodes, gedcom.NewSimpleNode(nil,
			gedcom.TagWife, "@"+wife+"@", "", nil))
	}

	for _, child := range children {
		nodes = append(nodes, gedcom.NewSimpleNode(nil,
			gedcom.TagChild, "@"+child+"@", "", nil))
	}

	return gedcom.NewFamilyNode(nil, pointer, nodes)
}

func document(nodes ...gedcom.Node) *gedcom.Document {
	return &gedcom.Document{
		Nodes: nodes,
	}
}
