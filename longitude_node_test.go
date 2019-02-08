package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLongitudeNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewLongitudeNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.LongitudeNode)(nil))
	assert.Equal(t, gedcom.TagLongitude, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}
