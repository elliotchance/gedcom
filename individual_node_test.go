package gedcom_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elliotchance/gedcom"
)

var individualTests = []struct {
	node  *gedcom.IndividualNode
	names []*gedcom.NameNode
}{
	{
		node:  gedcom.NewIndividualNode("", "P1", nil),
		names: []*gedcom.NameNode{},
	},
	{
		node:  gedcom.NewIndividualNode("", "P1", []gedcom.Node{}),
		names: []*gedcom.NameNode{},
	},
	{
		node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
			gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
		},
	},
	{
		node: gedcom.NewIndividualNode("", "P1", []gedcom.Node{
			gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewSimpleNode(gedcom.Version, "", "", []gedcom.Node{}),
			gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
		}),
		names: []*gedcom.NameNode{
			gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
		},
	},
}

func TestIndividualNode_Names(t *testing.T) {
	for _, test := range individualTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Names(), test.names)
		})
	}
}
