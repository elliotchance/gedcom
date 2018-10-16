package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPhoneticVariationNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewPhoneticVariationNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.PhoneticVariationNode)(nil))
	assert.Equal(t, gedcom.TagPhonetic, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
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
			node:     gedcom.NewPhoneticVariationNode(nil, "", "", nil),
			expected: nil,
		},
		{
			node:     gedcom.NewPhoneticVariationNode(nil, "", "", []gedcom.Node{}),
			expected: nil,
		},
		{
			node: gedcom.NewPhoneticVariationNode(nil, "", "", []gedcom.Node{
				gedcom.NewTypeNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewTypeNode(nil, "", "", []gedcom.Node{}),
		},
		{
			node: gedcom.NewPhoneticVariationNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewPhoneticVariationNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewTypeNode(nil, "1", "", []gedcom.Node{}),
				gedcom.NewTypeNode(nil, "2", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewTypeNode(nil, "1", "", []gedcom.Node{}),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Type(), test.expected)
		})
	}
}
