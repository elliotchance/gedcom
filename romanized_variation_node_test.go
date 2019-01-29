package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRomanizedVariationNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewRomanizedVariationNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.RomanizedVariationNode)(nil))
	assert.Equal(t, gedcom.TagRomanized, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestRomanizedVariationNode_Type(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.RomanizedVariationNode
		expected *gedcom.TypeNode
	}{
		{
			node:     nil,
			expected: nil,
		},
		{
			node:     gedcom.NewRomanizedVariationNode(""),
			expected: nil,
		},
		{
			node:     gedcom.NewRomanizedVariationNode(""),
			expected: nil,
		},
		{
			node: gedcom.NewRomanizedVariationNode("",
				gedcom.NewTypeNode(""),
			),
			expected: gedcom.NewTypeNode(""),
		},
		{
			node: gedcom.NewRomanizedVariationNode("",
				gedcom.NewNameNode(""),
			),
			expected: nil,
		},
		{
			node: gedcom.NewRomanizedVariationNode("",
				gedcom.NewNameNode(""),
				gedcom.NewTypeNode("1"),
				gedcom.NewTypeNode("2"),
			),
			expected: gedcom.NewTypeNode("1"),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Type(), test.expected)
		})
	}
}
