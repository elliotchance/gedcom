package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

func TestNewFormatNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewFormatNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.FormatNode)(nil))
	assert.Equal(t, tag.TagFormat, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}
