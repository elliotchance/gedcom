package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLatitudeNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewLatitudeNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.LatitudeNode)(nil))
	assert.Equal(t, gedcom.TagLatitude, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
}
