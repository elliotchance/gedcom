package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

var nodesWithTagTests = []struct {
	node gedcom.Node
	tag  tag.Tag
	want gedcom.Nodes
}{
	{nil, tag.TagHeader, nil},
	{gedcom.NewNameNode(""), tag.TagHeader, gedcom.Nodes{}},
	{
		gedcom.NewNameNode("",
			gedcom.NewNode(tag.TagSurname, "", ""),
		),
		tag.TagHeader,
		gedcom.Nodes{},
	},
	{
		gedcom.NewNameNode("",
			gedcom.NewNode(tag.TagSurname, "", ""),
		),
		tag.TagSurname,
		gedcom.Nodes{
			gedcom.NewNode(tag.TagSurname, "", ""),
		},
	},
	{
		gedcom.NewNameNode("",
			gedcom.NewNode(tag.TagHeader, "", ""),
			gedcom.NewNode(tag.TagSurname, "", ""),
		),
		tag.TagSurname,
		gedcom.Nodes{
			gedcom.NewNode(tag.TagSurname, "", ""),
		},
	},
	{
		gedcom.NewNameNode("",
			gedcom.NewNode(tag.TagSurname, "", ""),
			gedcom.NewNode(tag.TagSurname, "", ""),
		),
		tag.TagSurname,
		gedcom.Nodes{
			gedcom.NewNode(tag.TagSurname, "", ""),
			gedcom.NewNode(tag.TagSurname, "", ""),
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
		tagPath []tag.Tag
		want    gedcom.Nodes
	}{
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(tag.TagSurname, "", ""),
			),
			[]tag.Tag{},
			gedcom.Nodes{},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", ""),
				),
			),
			[]tag.Tag{tag.TagSurname},
			gedcom.Nodes{
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", ""),
				),
			},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", ""),
				),
			),
			[]tag.Tag{tag.TagSurname, tag.TagText},
			gedcom.Nodes{
				gedcom.NewNode(tag.TagText, "", ""),
			},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", "1"),
				),
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", "2"),
				),
			),
			[]tag.Tag{tag.TagSurname, tag.TagText},
			gedcom.Nodes{
				gedcom.NewNode(tag.TagText, "", "1"),
				gedcom.NewNode(tag.TagText, "", "2"),
			},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", "1"),
					gedcom.NewNode(tag.TagText, "", "2"),
				),
			),
			[]tag.Tag{tag.TagSurname, tag.TagText},
			gedcom.Nodes{
				gedcom.NewNode(tag.TagText, "", "1"),
				gedcom.NewNode(tag.TagText, "", "2"),
			},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", ""),
				),
			),
			[]tag.Tag{tag.TagGivenName, tag.TagText},
			gedcom.Nodes{},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", ""),
				),
			),
			[]tag.Tag{tag.TagSurname, tag.TagSurname},
			gedcom.Nodes{},
		},
		{
			gedcom.NewNameNode("",
				gedcom.NewNode(tag.TagSurname, "", "",
					gedcom.NewNode(tag.TagText, "", ""),
				),
			),
			[]tag.Tag{tag.TagSurname, tag.TagGivenName},
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
	surname := gedcom.NewNode(tag.TagSurname, "", "")
	givenName := gedcom.NewNode(tag.TagGivenName, "", "")

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
			gedcom.NewNode(tag.TagSurname, "", ""),
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
				gedcom.NewNode(tag.TagGivenName, "", "",
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
