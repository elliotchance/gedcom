package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRomanizedVariationNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewRomanizedVariationNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.RomanizedVariationNode)(nil))
	assert.Equal(t, gedcom.TagRomanized, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestRomanizedVariationNode_Type(t *testing.T) {
	var tests = []struct {
		node     *gedcom.RomanizedVariationNode
		expected *gedcom.TypeNode
	}{
		{
			node:     nil,
			expected: nil,
		},
		{
			node:     gedcom.NewRomanizedVariationNode(nil, "", "", nil),
			expected: nil,
		},
		{
			node:     gedcom.NewRomanizedVariationNode(nil, "", "", []gedcom.Node{}),
			expected: nil,
		},
		{
			node: gedcom.NewRomanizedVariationNode(nil, "", "", []gedcom.Node{
				gedcom.NewTypeNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewTypeNode(nil, "", "", []gedcom.Node{}),
		},
		{
			node: gedcom.NewRomanizedVariationNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewRomanizedVariationNode(nil, "", "", []gedcom.Node{
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
