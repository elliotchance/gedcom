package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

func TestNewFormatNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewFormatNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.FormatNode)(nil))
	assert.Equal(t, gedcom.TagFormat, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}
