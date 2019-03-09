package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChildNode_Individual(t *testing.T) {
	Individual := tf.NamedFunction(t, "ChildNode_Individual",
		(*gedcom.ChildNode).Individual)

	Individual(nil).Returns(nil)

	doc := gedcom.NewDocument()
	p1 := doc.AddIndividual("P1")
	p2 := doc.AddIndividual("P2")
	f1 := doc.AddFamilyWithHusbandAndWife("F1", p1, nil)
	f1.AddChild(p2)

	assert.Equal(t, f1.Children()[0].Individual(), p2)
}
