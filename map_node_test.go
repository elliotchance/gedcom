package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMapNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewMapNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.MapNode)(nil))
	assert.Equal(t, gedcom.TagMap, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestPlaceNode_Latitude(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.MapNode
		expected *gedcom.LatitudeNode
	}{
		{
			node:     nil,
			expected: nil,
		},
		{
			node:     gedcom.NewMapNode(nil, "", "", nil),
			expected: nil,
		},
		{
			node:     gedcom.NewMapNode(nil, "", "", []gedcom.Node{}),
			expected: nil,
		},
		{
			node: gedcom.NewMapNode(nil, "", "", []gedcom.Node{
				gedcom.NewLatitudeNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewLatitudeNode(nil, "", "", []gedcom.Node{}),
		},
		{
			node: gedcom.NewMapNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewMapNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewLatitudeNode(nil, "1", "", []gedcom.Node{}),
				gedcom.NewLatitudeNode(nil, "2", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewLatitudeNode(nil, "1", "", []gedcom.Node{}),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Latitude(), test.expected)
		})
	}
}

func TestPlaceNode_Longitude(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.MapNode
		expected *gedcom.LongitudeNode
	}{
		{
			node:     nil,
			expected: nil,
		},
		{
			node:     gedcom.NewMapNode(nil, "", "", nil),
			expected: nil,
		},
		{
			node:     gedcom.NewMapNode(nil, "", "", []gedcom.Node{}),
			expected: nil,
		},
		{
			node: gedcom.NewMapNode(nil, "", "", []gedcom.Node{
				gedcom.NewLongitudeNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewLongitudeNode(nil, "", "", []gedcom.Node{}),
		},
		{
			node: gedcom.NewMapNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewMapNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewLongitudeNode(nil, "1", "", []gedcom.Node{}),
				gedcom.NewLongitudeNode(nil, "2", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewLongitudeNode(nil, "1", "", []gedcom.Node{}),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Longitude(), test.expected)
		})
	}
}
