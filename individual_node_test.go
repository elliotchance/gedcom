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
