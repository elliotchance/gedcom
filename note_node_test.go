package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/stretchr/testify/assert"
)

func TestNewNoteNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewNoteNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.NoteNode)(nil))
	assert.Equal(t, gedcom.TagNote, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}
