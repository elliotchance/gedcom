package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var nodesWithTagTests = []struct {
	node gedcom.Node
	tag  gedcom.Tag
	want []gedcom.Node
}{
	{nil, gedcom.TagHeader, []gedcom.Node{}},
	{gedcom.NewNameNode("", "", nil), gedcom.TagHeader, []gedcom.Node{}},
	{
		gedcom.NewNameNode("", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
		}),
		gedcom.TagHeader,
		[]gedcom.Node{},
	},
	{
		gedcom.NewNameNode("", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
		}),
		gedcom.TagSurname,
		[]gedcom.Node{
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
		},
	},
	{
		gedcom.NewNameNode("", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.TagHeader, "", "", nil),
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
		}),
		gedcom.TagSurname,
		[]gedcom.Node{
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
		},
	},
	{
		gedcom.NewNameNode("", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
		}),
		gedcom.TagSurname,
		[]gedcom.Node{
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
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
	tests := []struct {
		node    gedcom.Node
		tagPath []gedcom.Tag
		want    []gedcom.Node
	}{
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			}),
			[]gedcom.Tag{},
			[]gedcom.Node{},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname},
			[]gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "", nil),
				}),
			},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			[]gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagText, "", "", nil),
			},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "1", nil),
				}),
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "2", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			[]gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagText, "", "1", nil),
				gedcom.NewSimpleNode(gedcom.TagText, "", "2", nil),
			},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "1", nil),
					gedcom.NewSimpleNode(gedcom.TagText, "", "2", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			[]gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagText, "", "1", nil),
				gedcom.NewSimpleNode(gedcom.TagText, "", "2", nil),
			},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagGivenName, gedcom.TagText},
			[]gedcom.Node{},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagSurname},
			[]gedcom.Node{},
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagGivenName},
			[]gedcom.Node{},
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
	surname := gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil)
	givenName := gedcom.NewSimpleNode(gedcom.TagGivenName, "", "", nil)

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
			gedcom.NewNameNode("", "", nil),
			surname,
			false,
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{}),
			surname,
			false,
		},

		// Other cases.
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				surname,
			}),
			surname,
			true,
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				surname,
			}),
			gedcom.NewSimpleNode(gedcom.TagSurname, "", "", nil),
			false,
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				givenName,
			}),
			surname,
			false,
		},
		{
			gedcom.NewNameNode("", "", []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagGivenName, "", "", []gedcom.Node{
					givenName,
				}),
			}),
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
