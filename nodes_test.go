package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

var nodesWithTagTests = []struct {
	node gedcom.Node
	tag  gedcom.Tag
	want gedcom.Nodes
}{
	{nil, gedcom.TagHeader, nil},
	{gedcom.NewNameNode(""), gedcom.TagHeader, gedcom.Nodes{}},
	{
		gedcom.NewNameNode("",
			gedcom.NewNode(gedcom.TagSurname, "", ""),
		),
		gedcom.TagHeader,
		gedcom.Nodes{},
	},
	{
		gedcom.NewNameNode("",
			gedcom.NewNode(gedcom.TagSurname, "", ""),
		),
		gedcom.TagSurname,
		gedcom.Nodes{
			gedcom.NewNode(gedcom.TagSurname, "", ""),
		},
	},
	{
		gedcom.NewNameNode("",
			gedcom.NewNode(gedcom.TagHeader, "", ""),
			gedcom.NewNode(gedcom.TagSurname, "", ""),
		),
		gedcom.TagSurname,
		gedcom.Nodes{
			gedcom.NewNode(gedcom.TagSurname, "", ""),
		},
	},
	{
		gedcom.NewNameNode("",
			gedcom.NewNode(gedcom.TagSurname, "", ""),
			gedcom.NewNode(gedcom.TagSurname, "", ""),
		),
		gedcom.TagSurname,
		gedcom.Nodes{
			gedcom.NewNode(gedcom.TagSurname, "", ""),
			gedcom.NewNode(gedcom.TagSurname, "", ""),
		},
	},
}

func TestNodesWithTag(t *testing.T) {
	for _, test := range nodesWithTagTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.NodesWithTag(test.node, test.tag))
		})
	}
}

func TestNodesWithTagPath(t *testing.T) {
	// ghost:ignore
	tests := []struct {
		node    gedcom.Node
		tagPath []gedcom.Tag
		want    gedcom.Nodes
	}{
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagSurname, "", ""),
			),
			[]gedcom.Tag{},
			gedcom.Nodes{},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", ""),
				),
			),
			[]gedcom.Tag{gedcom.TagSurname},
			gedcom.Nodes{
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", ""),
				),
			},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", ""),
				),
			),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			gedcom.Nodes{
				gedcom.NewNode(gedcom.TagText, "", ""),
			},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", "1"),
				),
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", "2"),
				),
			),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			gedcom.Nodes{
				gedcom.NewNode(gedcom.TagText, "", "1"),
				gedcom.NewNode(gedcom.TagText, "", "2"),
			},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", "1"),
					gedcom.NewNode(gedcom.TagText, "", "2"),
				),
			),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			gedcom.Nodes{
				gedcom.NewNode(gedcom.TagText, "", "1"),
				gedcom.NewNode(gedcom.TagText, "", "2"),
			},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", ""),
				),
			),
			[]gedcom.Tag{gedcom.TagGivenName, gedcom.TagText},
			gedcom.Nodes{},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", ""),
				),
			),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagSurname},
			gedcom.Nodes{},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagSurname, "", "",
					gedcom.NewNode(gedcom.TagText, "", ""),
				),
			),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagGivenName},
			gedcom.Nodes{},
		},
	}

	// It must satisfy all the tests for NodesWithTag.
	for _, test := range nodesWithTagTests {
		t.Run("", func(t *testing.T) {
			result := gedcom.NodesWithTagPath(test.node, test.tag)
			assert.Equal(t, test.want, result)
		})
	}

	// Now more complex paths.
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := gedcom.NodesWithTagPath(test.node, test.tagPath...)
			assert.Equal(t, test.want, result)
		})
	}
}

func TestHasNestedNode(t *testing.T) {
	surname := gedcom.NewNode(gedcom.TagSurname, "", "")
	givenName := gedcom.NewNode(gedcom.TagGivenName, "", "")

	// ghost:ignore
	tests := []struct {
		node       gedcom.Node
		lookingFor gedcom.Node
		want       bool
	}{
		// Nil parameters.
		{
			nil,
			nil,
			false,
		},
		{
			nil,
			surname,
			false,
		},
		{
			surname,
			nil,
			false,
		},

		// No children.
		{
			gedcom.NewNameNode(""),
			surname,
			false,
		},
		{
			gedcom.NewNameNode(""),
			surname,
			false,
		},

		// Other cases.
		{
			gedcom.NewNameNode("",
				surname,
			),
			surname,
			true,
		},
		{
			gedcom.NewNameNode("",
				surname,
			),
			gedcom.NewNode(gedcom.TagSurname, "", ""),
			false,
		},
		{
			gedcom.NewNameNode("",
				givenName,
			),
			surname,
			false,
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(gedcom.TagGivenName, "", "",
					givenName,
				),
			),
			givenName,
			true,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := gedcom.HasNestedNode(test.node, test.lookingFor)
			assert.Equal(t, test.want, result)
		})
	}
}

func TestNodes_CastTo(t *testing.T) {
	CastTo := tf.Function(t, gedcom.Nodes.CastTo)

	CastTo(nil, (*gedcom.NameNode)(nil)).
		Returns([]*gedcom.NameNode{})

	CastTo(gedcom.Nodes{}, (*gedcom.NameNode)(nil)).
		Returns([]*gedcom.NameNode{})

	name := gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/")
	CastTo(gedcom.Nodes{name}, (*gedcom.NameNode)(nil)).
		Returns([]*gedcom.NameNode{name})
}
