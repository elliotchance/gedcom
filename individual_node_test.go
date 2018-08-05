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
				SpouseChildren: map[*gedcom.IndividualNode][]*gedcom.IndividualNode{
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
				SpouseChildren: map[*gedcom.IndividualNode][]*gedcom.IndividualNode{
					p1: {p6},
				},
			},
		},
		{
			node: p3,
			descent: &gedcom.Descent{
				Parents:    []*gedcom.FamilyNode{},
				Individual: p3,
				SpouseChildren: map[*gedcom.IndividualNode][]*gedcom.IndividualNode{
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
				SpouseChildren: map[*gedcom.IndividualNode][]*gedcom.IndividualNode{},
			},
		},
		{
			node: p5,
			descent: &gedcom.Descent{
				Parents:        []*gedcom.FamilyNode{f1},
				Individual:     p5,
				SpouseChildren: map[*gedcom.IndividualNode][]*gedcom.IndividualNode{},
			},
		},
		{
			node: p6,
			descent: &gedcom.Descent{
				Parents:    []*gedcom.FamilyNode{f2, f4},
				Individual: p6,
				SpouseChildren: map[*gedcom.IndividualNode][]*gedcom.IndividualNode{
					nil: {p7},
				},
			},
		},
		{
			node: p7,
			descent: &gedcom.Descent{
				Parents:        []*gedcom.FamilyNode{f3},
				Individual:     p7,
				SpouseChildren: map[*gedcom.IndividualNode][]*gedcom.IndividualNode{},
			},
		},
		{
			node: p8,
			descent: &gedcom.Descent{
				Parents:        []*gedcom.FamilyNode{},
				Individual:     p8,
				SpouseChildren: map[*gedcom.IndividualNode][]*gedcom.IndividualNode{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.node.Pointer(), func(t *testing.T) {
			assert.Equal(t, test.node.Descent(doc), test.descent)
		})
	}
}
