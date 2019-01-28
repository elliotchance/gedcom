package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestSimpleNode_ChildNodes(t *testing.T) {
	node := gedcom.NewNodeWithChildren(nil, gedcom.TagText, "", "", nil)

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
	Equals := tf.NamedFunction(t, "SimpleNode_Equals", (*gedcom.SimpleNode).Equals)

	left := []*gedcom.SimpleNode{
		(*gedcom.SimpleNode)(nil),
		gedcom.NewNode(nil, gedcom.TagVersion, "", "").(*gedcom.SimpleNode),
		gedcom.NewNode(nil, gedcom.TagVersion, "a", "").(*gedcom.SimpleNode),
		gedcom.NewNode(nil, gedcom.TagVersion, "", "b").(*gedcom.SimpleNode),
		gedcom.NewNode(nil, gedcom.TagVersion, "a", "b").(*gedcom.SimpleNode),
	}

	right := []gedcom.Node{
		(*gedcom.SimpleNode)(nil),
		gedcom.NewNode(nil, gedcom.TagVersion, "", "").(*gedcom.SimpleNode),
		gedcom.NewNode(nil, gedcom.TagVersion, "a", "").(*gedcom.SimpleNode),
		gedcom.NewNode(nil, gedcom.TagVersion, "", "b").(*gedcom.SimpleNode),
		gedcom.NewNode(nil, gedcom.TagVersion, "a", "b").(*gedcom.SimpleNode),
		gedcom.NewNameNode(nil, "", "", nil),
	}

	const N = false
	const Y = true

	expected := [][]bool{
		{N, N, N, N, N, N},
		{N, Y, N, N, N, N},
		{N, N, Y, N, N, N},
		{N, N, N, Y, N, N},
		{N, N, N, N, Y, N},
	}

	for i, l := range left {
		for j, r := range right {
			Equals(l, r).Returns(expected[i][j])
		}
	}
}

func TestSimpleNode_Tag(t *testing.T) {
	Tag := tf.Function(t, (*gedcom.SimpleNode).Tag)

	Tag((*gedcom.SimpleNode)(nil)).Returns(gedcom.Tag{})
}

func TestSimpleNode_Value(t *testing.T) {
	Value := tf.Function(t, (*gedcom.SimpleNode).Value)

	Value((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_Pointer(t *testing.T) {
	Pointer := tf.Function(t, (*gedcom.SimpleNode).Pointer)

	Pointer((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_Document(t *testing.T) {
	Document := tf.Function(t, (*gedcom.SimpleNode).Document)

	Document((*gedcom.SimpleNode)(nil)).Returns((*gedcom.Document)(nil))
}

func TestSimpleNode_SetDocument(t *testing.T) {
	(*gedcom.SimpleNode)(nil).SetDocument(nil)
}

func TestSimpleNode_Nodes(t *testing.T) {
	Nodes := tf.Function(t, (*gedcom.SimpleNode).Nodes)

	Nodes((*gedcom.SimpleNode)(nil)).Returns(([]gedcom.Node)(nil))
}

func TestSimpleNode_String(t *testing.T) {
	String := tf.Function(t, (*gedcom.SimpleNode).String)

	String((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_GEDCOMString(t *testing.T) {
	root := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
		}),
	})

	assert.Equal(t, root.GEDCOMString(0), `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`)
}

func TestSimpleNode_GEDCOMLine(t *testing.T) {
	GEDCOMLine := tf.NamedFunction(t, "SimpleNode_GEDCOMLine",
		(*gedcom.SimpleNode).GEDCOMLine)

	GEDCOMLine(gedcom.NewBirthNode(nil, "foo", "72", nil).SimpleNode, 0).
		Returns("0 @72@ BIRT foo")

	GEDCOMLine(gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "bar", "baz", nil).(*gedcom.DeathNode).SimpleNode, 3).Returns("3 @baz@ DEAT bar")

	GEDCOMLine(gedcom.NewDateNode(nil, "3 SEP 1945", "", nil).SimpleNode, 2).
		Returns("2 DATE 3 SEP 1945")

	GEDCOMLine(gedcom.NewBirthNode(nil, "foo", "72", nil).SimpleNode, -1).
		Returns("@72@ BIRT foo")
}

func TestSimpleNode_SetNodes(t *testing.T) {
	birth := gedcom.NewBirthNode(nil, "foo", "72", nil)
	assert.Nil(t, birth.Nodes())

	birth.SetNodes([]gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1945", "", nil),
	})
	assert.Equal(t, []gedcom.Node{
		gedcom.NewDateNode(nil, "3 SEP 1945", "", nil),
	}, birth.Nodes())

	birth.SetNodes(nil)
	assert.Nil(t, birth.Nodes())
}
