package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/tag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNoteNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewNoteNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.NoteNode)(nil))
	assert.Equal(t, tag.TagNote, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}
