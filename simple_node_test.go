package gedcom_test

import (
	"testing"
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

func TestSimpleNode_ChildNodes(t *testing.T) {
	t.Run("NilReturnsArray", func(t *testing.T) {
		node := gedcom.NewSimpleNode(gedcom.Text, "", "", nil)

		assert.NotNil(t, node.Nodes())
		assert.Len(t, node.Nodes(), 0)
	})
}
