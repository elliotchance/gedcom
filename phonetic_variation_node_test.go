package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/tag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPhoneticVariationNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewPhoneticVariationNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.PhoneticVariationNode)(nil))
	assert.Equal(t, tag.TagPhonetic, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestPhoneticVariationNode_Type(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.PhoneticVariationNode
		expected *gedcom.TypeNode
	}{
		{
			node:     nil,
			expected: nil,
		},
		{
			node:     gedcom.NewPhoneticVariationNode(""),
			expected: nil,
		},
		{
			node:     gedcom.NewPhoneticVariationNode(""),
			expected: nil,
		},
		{
			node: gedcom.NewPhoneticVariationNode("",
				gedcom.NewTypeNode(""),
			),
			expected: gedcom.NewTypeNode(""),
		},
		{
			node: gedcom.NewPhoneticVariationNode("",
				gedcom.NewNameNode(""),
			),
			expected: nil,
		},
		{
			node: gedcom.NewPhoneticVariationNode("",
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
