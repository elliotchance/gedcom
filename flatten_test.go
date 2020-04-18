package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	dateNode1 := gedcom.NewDateNode("1851")
	dateNode2 := gedcom.NewDateNode("1856")
	birthNode := gedcom.NewBirthNode("",
		dateNode1,
	)
	deathNode := gedcom.NewBirthNode("",
		dateNode1,
		dateNode2,
	)
	individualNode := gedcom.NewDocument().AddIndividual("P221",
		birthNode,
		deathNode,
	)

	t.Run("Nil", func(t *testing.T) {
		doc := gedcom.NewDocument()
		actual := gedcom.Flatten(doc, nil)

		assert.Nil(t, actual)
	})

	t.Run("NoChildren", func(t *testing.T) {
		doc := gedcom.NewDocument()
		actual := gedcom.Flatten(doc, dateNode1)

		if assert.Len(t, actual, 1) {
			assert.True(t, actual[0] == dateNode1)
		}
	})

	t.Run("Birth", func(t *testing.T) {
		doc := gedcom.NewDocument()
		actual := gedcom.Flatten(doc, birthNode)

		if assert.Len(t, actual, 2) {
			assert.True(t, actual[0] == birthNode)
			assert.True(t, actual[1] == dateNode1)
		}
	})

	t.Run("Death", func(t *testing.T) {
		doc := gedcom.NewDocument()
		actual := gedcom.Flatten(doc, deathNode)

		if assert.Len(t, actual, 3) {
			assert.True(t, actual[0] == deathNode)
			assert.True(t, actual[1] == dateNode1)
			assert.True(t, actual[2] == dateNode2)
		}
	})

	t.Run("Individual", func(t *testing.T) {
		doc := gedcom.NewDocument()
		actual := gedcom.Flatten(doc, individualNode)

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
		doc := gedcom.NewDocument()
		actual := gedcom.Flatten(doc, gedcom.DeepCopy(individualNode, doc))

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
