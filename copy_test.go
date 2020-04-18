package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	dateNode1 := gedcom.NewDateNode("1851")
	dateNode2 := gedcom.NewDateNode("1856")
	birthNode := gedcom.NewBirthNode("",
		dateNode1,
	)
	deathNode := gedcom.NewBirthNode("",
		dateNode1,
		dateNode2,
	)
	doc := gedcom.NewDocument()
	individualNode := doc.AddIndividual("P221",
		birthNode,
		deathNode,
	)

	t.Run("Nil", func(t *testing.T) {
		actual := gedcom.DeepCopy(nil, doc)

		assert.Nil(t, actual)
	})

	t.Run("NoChildren", func(t *testing.T) {
		actual := gedcom.DeepCopy(dateNode1, doc)

		assert.True(t, actual != dateNode1)
	})

	t.Run("Birth", func(t *testing.T) {
		actual := gedcom.DeepCopy(birthNode, doc)

		assert.True(t, actual != birthNode)
		assert.True(t, actual.Nodes()[0] != dateNode1)
	})

	t.Run("Death", func(t *testing.T) {
		actual := gedcom.DeepCopy(deathNode, doc)

		assert.True(t, actual != deathNode)
		assert.True(t, actual.Nodes()[0] != dateNode1)
		assert.True(t, actual.Nodes()[1] != dateNode2)
	})

	t.Run("Individual", func(t *testing.T) {
		actual := gedcom.DeepCopy(individualNode, doc)

		assert.True(t, actual != individualNode)
		assert.True(t, actual.Nodes()[0] != birthNode)
		assert.True(t, actual.Nodes()[0].Nodes()[0] != dateNode1)
		assert.True(t, actual.Nodes()[1] != deathNode)
		assert.True(t, actual.Nodes()[1].Nodes()[0] != dateNode1)
		assert.True(t, actual.Nodes()[1].Nodes()[1] != dateNode2)
	})
}
