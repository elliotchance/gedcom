package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleNode_ChildNodes(t *testing.T) {
	node := gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil)

	assert.Len(t, node.Nodes(), 0)
}

func TestIsNil(t *testing.T) {
	IsNil := tf.Function(t, gedcom.IsNil)

	IsNil((*gedcom.SimpleNode)(nil)).Returns(true)
	IsNil(gedcom.NewBirthNode(nil, "", "", nil)).Returns(false)
	IsNil((*gedcom.NameNode)(nil)).Returns(true)
	IsNil(gedcom.NewNameNode(nil, "", "", nil)).Returns(false)

	// Untyped nil is a special case that cannot be tested above.
	assert.True(t, gedcom.IsNil(nil))
}

func TestSimpleNode_Equals(t *testing.T) {
	Equals := tf.Function(t, (*gedcom.SimpleNode).Equals)

	s0 := (*gedcom.SimpleNode)(nil)

	// These are the same.
	s1 := gedcom.NewSimpleNode(nil, gedcom.TagName, "", "", nil)
	s2 := gedcom.NewNameNode(nil, "", "", nil)

	// These are different in some way from each other.
	s3 := gedcom.NewSimpleNode(nil, gedcom.TagName, "", "a", nil)
	s4 := gedcom.NewNameNode(nil, "a", "", nil)
	s5 := gedcom.NewNameNode(nil, "", "b", nil)
	s6 := gedcom.NewSimpleNode(nil, gedcom.TagVersion, "", "a", nil)

	// Nils
	Equals(s0, s0).Returns(false)
	Equals(s0, s1).Returns(false)
	Equals(s1, s0).Returns(false)

	Equals(s1, s1).Returns(true)
	Equals(s1, s2).Returns(true)

	Equals(s1, s3).Returns(false)
	Equals(s1, s4).Returns(false)
	Equals(s1, s5).Returns(false)

	Equals(s3, s3).Returns(true)
	Equals(s3, s4).Returns(false)
	Equals(s3, s5).Returns(false)
	Equals(s3, s6).Returns(false)
	Equals(s6, s3).Returns(false)
	Equals(s6, s4).Returns(false)
	Equals(s6, s5).Returns(false)
	Equals(s6, s6).Returns(true)
}
