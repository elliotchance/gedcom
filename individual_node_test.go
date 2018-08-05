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
