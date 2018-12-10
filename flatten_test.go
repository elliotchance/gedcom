package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
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
		actual := gedcom.Flatten(nil)

		assert.Nil(t, actual)
	})

	t.Run("NoChildren", func(t *testing.T) {
		actual := gedcom.Flatten(dateNode1)

		if assert.Len(t, actual, 1) {
			assert.True(t, actual[0] == dateNode1)
		}
	})

	t.Run("Birth", func(t *testing.T) {
		actual := gedcom.Flatten(birthNode)

		if assert.Len(t, actual, 2) {
			assert.True(t, actual[0] == birthNode)
			assert.True(t, actual[1] == dateNode1)
		}
	})

	t.Run("Death", func(t *testing.T) {
		actual := gedcom.Flatten(deathNode)

		if assert.Len(t, actual, 3) {
			assert.True(t, actual[0] == deathNode)
			assert.True(t, actual[1] == dateNode1)
			assert.True(t, actual[2] == dateNode2)
		}
	})

	t.Run("Individual", func(t *testing.T) {
		actual := gedcom.Flatten(individualNode)

		if assert.Len(t, actual, 6) {
			assert.True(t, actual[0] == individualNode)
			assert.True(t, actual[1] == birthNode)
			assert.True(t, actual[2] == dateNode1)
			assert.True(t, actual[3] == deathNode)
			assert.True(t, actual[4] == dateNode1)
			assert.True(t, actual[5] == dateNode2)
		}
	})

	t.Run("DeepCopy", func(t *testing.T) {
		actual := gedcom.Flatten(gedcom.DeepCopy(individualNode))

		if assert.Len(t, actual, 6) {
			assert.True(t, actual[0] != individualNode)
			assert.True(t, actual[1] != birthNode)
			assert.True(t, actual[2] != dateNode1)
			assert.True(t, actual[3] != deathNode)
			assert.True(t, actual[4] != dateNode1)
			assert.True(t, actual[5] != dateNode2)
		}
	})
}
