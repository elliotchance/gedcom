package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/stretchr/testify/assert"
)

func TestNodeSet_Add(t *testing.T) {
	ns := gedcom.NodeSet{}
	nameNode := gedcom.NewNameNode("")

	assert.False(t, ns.Has(nameNode))
	ns.Add(nameNode)
	assert.True(t, ns.Has(nameNode))
}

func TestNodeSet_Has(t *testing.T) {
	ns := gedcom.NodeSet{}
	nameNode1 := gedcom.NewNameNode("")
	nameNode2 := gedcom.NewNameNode("")

	assert.False(t, ns.Has(nameNode1))
	assert.False(t, ns.Has(nameNode2))

	ns.Add(nameNode2)

	assert.False(t, ns.Has(nameNode1))
	assert.True(t, ns.Has(nameNode2))

	ns.Add(nameNode1)

	assert.True(t, ns.Has(nameNode1))
	assert.True(t, ns.Has(nameNode2))
}
