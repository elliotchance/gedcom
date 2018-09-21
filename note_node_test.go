package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNoteNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewNoteNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.NoteNode)(nil))
	assert.Equal(t, gedcom.TagNote, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}
