package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleNode_ChildNodes(t *testing.T) {
	t.Run("NilReturnsArray", func(t *testing.T) {
		node := gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil)

		assert.NotNil(t, node.Nodes())
		assert.Len(t, node.Nodes(), 0)
	})
}
