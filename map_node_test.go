package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/tag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMapNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewMapNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.MapNode)(nil))
	assert.Equal(t, tag.TagMap, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
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
			node:     gedcom.NewMapNode(""),
			expected: nil,
		},
		{
			node:     gedcom.NewMapNode(""),
			expected: nil,
		},
		{
			node: gedcom.NewMapNode("",
				gedcom.NewLatitudeNode(""),
			),
			expected: gedcom.NewLatitudeNode(""),
		},
		{
			node: gedcom.NewMapNode("",
				gedcom.NewNameNode(""),
			),
			expected: nil,
		},
		{
			node: gedcom.NewMapNode("",
				gedcom.NewNameNode(""),
				gedcom.NewLatitudeNode("1"),
				gedcom.NewLatitudeNode("2"),
			),
			expected: gedcom.NewLatitudeNode("1"),
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
			node:     gedcom.NewMapNode(""),
			expected: nil,
		},
		{
			node:     gedcom.NewMapNode(""),
			expected: nil,
		},
		{
			node: gedcom.NewMapNode("",
				gedcom.NewLongitudeNode(""),
			),
			expected: gedcom.NewLongitudeNode(""),
		},
		{
			node: gedcom.NewMapNode("",
				gedcom.NewNameNode(""),
			),
			expected: nil,
		},
		{
			node: gedcom.NewMapNode("",
				gedcom.NewNameNode(""),
				gedcom.NewLongitudeNode("1"),
				gedcom.NewLongitudeNode("2"),
			),
			expected: gedcom.NewLongitudeNode("1"),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Longitude(), test.expected)
		})
	}
}
