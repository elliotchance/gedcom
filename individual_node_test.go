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
		node:  gedcom.NewIndividualNode("", "P1", nil),
		names: []*gedcom.NameNode{},
		sex:   gedcom.SexUnknown,
	},
	{
		node:  gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
		names: []*gedcom.NameNode{},
		sex:   gedcom.SexUnknown,
	},
	{
		node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
			gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
		},
		sex: gedcom.SexUnknown,
	},
	{
		node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
			gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewSimpleNode(gedcom.TagVersion, "", "", []gedcom.Node{}),
			gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
		},
		sex: gedcom.SexUnknown,
	},
	{
		node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.TagSex, "M", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{},
		sex:   gedcom.SexMale,
	},
	{
		node: gedcom.NewIndividualNode("", "P2", []gedcom.Node{
			gedcom.NewNameNode("Joan /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewSimpleNode(gedcom.TagSex, "F", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode("Joan /Bloggs/", "", []gedcom.Node{}),
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
		births []gedcom.Node
	}{
		{
			node:   gedcom.NewIndividualNode("", "P1", nil),
			births: []gedcom.Node{},
		},
		{
			node:   gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			births: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{}),
			}),
			births: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			births: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBirth, "bar", "", []gedcom.Node{}),
			}),
			births: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBirth, "bar", "", []gedcom.Node{}),
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
			node:     gedcom.NewIndividualNode("", "P1", nil),
			baptisms: []gedcom.Node{},
		},
		{
			node:     gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			baptisms: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBaptism, "bar", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBaptism, "bar", "", []gedcom.Node{}),
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
			node:   gedcom.NewIndividualNode("", "P1", nil),
			deaths: []gedcom.Node{},
		},
		{
			node:   gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			deaths: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			deaths: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{}),
			}),
			deaths: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagDeath, "bar", "", []gedcom.Node{}),
			}),
			deaths: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagDeath, "bar", "", []gedcom.Node{}),
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
			node:    gedcom.NewIndividualNode("", "P1", nil),
			burials: []gedcom.Node{},
		},
		{
			node:    gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			burials: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{}),
			}),
			burials: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{}),
			}),
			burials: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBurial, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBurial, "bar", "", []gedcom.Node{}),
			}),
			burials: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBurial, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagBurial, "bar", "", []gedcom.Node{}),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Burials(), test.burials)
		})
	}
}

func TestIndividualNode_Descent(t *testing.T) {
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

	p1 := gedcom.NewIndividualNode("", "P1", nil)
	p2 := gedcom.NewIndividualNode("", "P2", nil)
	p3 := gedcom.NewIndividualNode("", "P3", nil)
	p4 := gedcom.NewIndividualNode("", "P4", nil)
	p5 := gedcom.NewIndividualNode("", "P5", nil)
	p6 := gedcom.NewIndividualNode("", "P6", nil)
	p7 := gedcom.NewIndividualNode("", "P7", nil)
	p8 := gedcom.NewIndividualNode("", "P8", nil)

	// P1 - P3
	//    |
	//  -----
	// P4   P5
	f1 := gedcom.NewFamilyNode("F1", []gedcom.Node{
		gedcom.NewSimpleNode(gedcom.TagHusband, "@P1@", "", nil),
		gedcom.NewSimpleNode(gedcom.TagWife, "@P3@", "", nil),
		gedcom.NewSimpleNode(gedcom.TagChild, "@P4@", "", nil),
		gedcom.NewSimpleNode(gedcom.TagChild, "@P5@", "", nil),
	})

	// P1 - P2
	//    |
	//   P6
	f2 := gedcom.NewFamilyNode("F2", []gedcom.Node{
		gedcom.NewSimpleNode(gedcom.TagHusband, "@P1@", "", nil),
		gedcom.NewSimpleNode(gedcom.TagWife, "@P2@", "", nil),
		gedcom.NewSimpleNode(gedcom.TagChild, "@P6@", "", nil),
	})

	// P6 - ?
	//    |
	//   P7
	f3 := gedcom.NewFamilyNode("F3", []gedcom.Node{
		gedcom.NewSimpleNode(gedcom.TagHusband, "@P6@", "", nil),
		gedcom.NewSimpleNode(gedcom.TagChild, "@P7@", "", nil),
	})

	// ? - P3
	//   |
	//   P6
	f4 := gedcom.NewFamilyNode("F4", []gedcom.Node{
		gedcom.NewSimpleNode(gedcom.TagWife, "@P3@", "", nil),
		gedcom.NewSimpleNode(gedcom.TagChild, "@P6@", "", nil),
	})

	doc := &gedcom.Document{
		Nodes: []gedcom.Node{
			p1, p2, p3, p4, p5, p6, p7, p8,
			f1, f2, f3, f4,
		},
	}

	var tests = []struct {
		node    *gedcom.IndividualNode
		descent *gedcom.Descent
	}{
		{
			node: p1,
			descent: &gedcom.Descent{
				Parents:    []*gedcom.FamilyNode{},
				Individual: p1,
				SpouseChildren: map[*gedcom.IndividualNode]gedcom.IndividualNodes{
					p3: {p4, p5},
					p2: {p6},
				},
			},
		},
		{
			node: p2,
			descent: &gedcom.Descent{
				Parents:    []*gedcom.FamilyNode{},
				Individual: p2,
				SpouseChildren: map[*gedcom.IndividualNode]gedcom.IndividualNodes{
					p1: {p6},
				},
			},
		},
		{
			node: p3,
			descent: &gedcom.Descent{
				Parents:    []*gedcom.FamilyNode{},
				Individual: p3,
				SpouseChildren: map[*gedcom.IndividualNode]gedcom.IndividualNodes{
					p1:  {p4, p5},
					nil: {p6},
				},
			},
		},
		{
			node: p4,
			descent: &gedcom.Descent{
				Parents:        []*gedcom.FamilyNode{f1},
				Individual:     p4,
				SpouseChildren: map[*gedcom.IndividualNode]gedcom.IndividualNodes{},
			},
		},
		{
			node: p5,
			descent: &gedcom.Descent{
				Parents:        []*gedcom.FamilyNode{f1},
				Individual:     p5,
				SpouseChildren: map[*gedcom.IndividualNode]gedcom.IndividualNodes{},
			},
		},
		{
			node: p6,
			descent: &gedcom.Descent{
				Parents:    []*gedcom.FamilyNode{f2, f4},
				Individual: p6,
				SpouseChildren: map[*gedcom.IndividualNode]gedcom.IndividualNodes{
					nil: {p7},
				},
			},
		},
		{
			node: p7,
			descent: &gedcom.Descent{
				Parents:        []*gedcom.FamilyNode{f3},
				Individual:     p7,
				SpouseChildren: map[*gedcom.IndividualNode]gedcom.IndividualNodes{},
			},
		},
		{
			node: p8,
			descent: &gedcom.Descent{
				Parents:        []*gedcom.FamilyNode{},
				Individual:     p8,
				SpouseChildren: map[*gedcom.IndividualNode]gedcom.IndividualNodes{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.node.Pointer(), func(t *testing.T) {
			assert.Equal(t, test.node.Descent(doc), test.descent)
		})
	}
}

func TestIndividualNode_LDSBaptisms(t *testing.T) {
	var tests = []struct {
		node     *gedcom.IndividualNode
		baptisms []gedcom.Node
	}{
		{
			node:     gedcom.NewIndividualNode("", "P1", nil),
			baptisms: []gedcom.Node{},
		},
		{
			node:     gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			baptisms: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "bar", "", []gedcom.Node{}),
			}),
			baptisms: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "foo", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "bar", "", []gedcom.Node{}),
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
			node:     gedcom.NewIndividualNode("", "P1", nil),
			expected: nil,
		},
		{
			node:     gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			expected: nil,
		},

		// A single date.
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("1 Aug 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode("Abt. Dec 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("Abt. Dec 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode("Abt. Nov 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("Abt. Nov 1980", "", nil),
		},

		// Multiple dates and other cases.
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
				}),
				gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode("Abt. Jan 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("Abt. Jan 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
					gedcom.NewDateNode("23 Mar 1979", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("23 Mar 1979", "", nil),
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
				}),
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode("23 Mar 1979", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("23 Mar 1979", "", nil),
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(gedcom.TagLDSBaptism, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("1 Aug 1980", "", nil),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.EstimatedBirthDate(), test.expected)
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
			node:     gedcom.NewIndividualNode("", "P1", nil),
			expected: nil,
		},
		{
			node:     gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			expected: nil,
		},

		// A single date.
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("1 Aug 1980", "", nil),
		},
		{
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode("Abt. Dec 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("Abt. Dec 1980", "", nil),
		},

		// Multiple dates and other cases.
		{
			// Multiple death dates always returns the earliest.
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
					gedcom.NewDateNode("Mar 1980", "", nil),
					gedcom.NewDateNode("Jun 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("Mar 1980", "", nil),
		},
		{
			// Multiple burial dates always returns the earliest.
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode("3 Aug 1980", "", nil),
					gedcom.NewDateNode("Apr 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("Apr 1980", "", nil),
		},
		{
			// Death is before burial.
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
				}),
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode("3 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("1 Aug 1980", "", nil),
		},
		{
			// Burial is before death.
			node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode("3 Aug 1980", "", nil),
				}),
				gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{
					gedcom.NewDateNode("1 Aug 1980", "", nil),
				}),
			}),
			expected: gedcom.NewDateNode("3 Aug 1980", "", nil),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.EstimatedDeathDate(), test.expected)
		})
	}
}

func born(value string) gedcom.Node {
	return gedcom.NewSimpleNode(gedcom.TagBirth, "", "", []gedcom.Node{
		gedcom.NewDateNode(value, "", []gedcom.Node{}),
	})
}

func died(value string) gedcom.Node {
	return gedcom.NewSimpleNode(gedcom.TagDeath, "", "", []gedcom.Node{
		gedcom.NewDateNode(value, "", []gedcom.Node{}),
	})
}

func name(value string) gedcom.Node {
	return gedcom.NewNameNode(value, "", nil)
}

func baptised(value string) gedcom.Node {
	return gedcom.NewSimpleNode(gedcom.TagBaptism, "", "", []gedcom.Node{
		gedcom.NewDateNode(value, "", []gedcom.Node{}),
	})
}

func buried(value string) gedcom.Node {
	return gedcom.NewSimpleNode(gedcom.TagBurial, "", "", []gedcom.Node{
		gedcom.NewDateNode(value, "", []gedcom.Node{}),
	})
}

func TestIndividualNode_Similarity(t *testing.T) {
	var tests = []struct {
		a, b     *gedcom.IndividualNode
		expected float64
	}{
		// Empty cases.
		{
			a:        gedcom.NewIndividualNode("", "P1", nil),
			b:        gedcom.NewIndividualNode("", "P1", nil),
			expected: 0.3333333333333333,
		},
		{
			a:        gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			b:        gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
			expected: 0.3333333333333333,
		},

		// Perfect cases.
		{
			// All details match exactly.
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 1.0,
		},
		{
			// Extra names, but one name is still a perfect match.
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				name("Elliot Rupert /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot R d P /Chance/"),
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 1.0,
		},
		{
			// Name are not senstive to case or whitespace.
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("elliot /CHANCE/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 1.0,
		},

		// Almost perfect matches.
		{
			// Name is more/less complete.
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 0.9663440860215053,
		},
		{
			// Last name is similar.
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chaunce/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 0.995766129032258,
		},
		{
			// Birth date is less specific.
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("Jan 1843"),
				died("17 Mar 1907"),
			}),
			expected: 0.999996416733853,
		},
		{
			// Death date is less specific.
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("Mar 1907"),
			}),
			expected: 0.9999999751162073,
		},

		// Estimated birth/death.
		{
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				baptised("Abt. 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				died("Mar 1907"),
			}),
			expected: 0.9992026735146867,
		},
		{
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				baptised("Abt. 1843"),
				died("17 Mar 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				born("4 Jan 1843"),
				buried("Aft. 20 Mar 1907"),
			}),
			expected: 0.9992024744443452,
		},

		// Missing dates.
		{
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				died("Abt. 1907"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert /Chance/"),
				died("1909"),
			}),
			expected: 0.7863440860215053,
		},
		{
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
				baptised("after Sep 1823"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert /Chance/"),
				born("Between 1822 and 1823"),
			}),
			expected: 0.7980146283388829,
		},
		{
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert de Peyster /Chance/"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Elliot Rupert /Chance/"),
			}),
			expected: 0.633010752688172,
		},

		// These ones are way off.
		{
			a: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Jane /Doe/"),
				born("Sep 1845"),
			}),
			b: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				name("Bob /Jones/"),
				born("1627"),
			}),
			expected: 0.3194444444444444,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.a.Similarity(test.b), test.expected)
		})
	}
}
