package gedcom_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/elliotchance/gedcom"
)

func TestNodesWithTag(t *testing.T) {
	tests := []struct {
		node gedcom.Node
		tag  gedcom.Tag
		want []gedcom.Node
	}{
		{nil, gedcom.TagHeader, []gedcom.Node{}},
		{gedcom.NewNameNode("", "", nil), gedcom.TagHeader, []gedcom.Node{}},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			}),
			gedcom.TagHeader,
			[]gedcom.Node{},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			}),
			gedcom.TagSurname,
			[]gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagHeader, "", "", nil),
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			}),
			gedcom.TagSurname,
			[]gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			}),
			gedcom.TagSurname,
			[]gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.NodesWithTag(test.node, test.tag))
		})
	}
}
