package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTypeNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewTypeNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.TypeNode)(nil))
	assert.Equal(t, gedcom.TagType, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}
