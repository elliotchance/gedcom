package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	dateNode1 := gedcom.NewDateNode(nil, "1851", "", nil)
	dateNode2 := gedcom.NewDateNode(nil, "1856", "", nil)
	birthNode := gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
		dateNode1,
	})
	deathNode := gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
		dateNode1,
		dateNode2,
	})
	individualNode := gedcom.NewIndividualNode(nil, "", "P221", []gedcom.Node{
		birthNode,
		deathNode,
	})

	t.Run("Nil", func(t *testing.T) {
		actual := gedcom.DeepCopy(nil)

		assert.Nil(t, actual)
	})

	t.Run("NoChildren", func(t *testing.T) {
		actual := gedcom.DeepCopy(dateNode1)

		assert.True(t, actual != dateNode1)
	})

	t.Run("Birth", func(t *testing.T) {
		actual := gedcom.DeepCopy(birthNode)

		assert.True(t, actual != birthNode)
		assert.True(t, actual.Nodes()[0] != dateNode1)
	})

	t.Run("Death", func(t *testing.T) {
		actual := gedcom.DeepCopy(deathNode)

		assert.True(t, actual != deathNode)
		assert.True(t, actual.Nodes()[0] != dateNode1)
		assert.True(t, actual.Nodes()[1] != dateNode2)
	})

	t.Run("Individual", func(t *testing.T) {
		actual := gedcom.DeepCopy(individualNode)

		assert.True(t, actual != individualNode)
		assert.True(t, actual.Nodes()[0] != birthNode)
		assert.True(t, actual.Nodes()[0].Nodes()[0] != dateNode1)
		assert.True(t, actual.Nodes()[1] != deathNode)
		assert.True(t, actual.Nodes()[1].Nodes()[0] != dateNode1)
		assert.True(t, actual.Nodes()[1].Nodes()[1] != dateNode2)
	})
}
