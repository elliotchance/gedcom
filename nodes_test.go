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
	{gedcom.NewNameNode(nil, "", "", nil), gedcom.TagHeader, []gedcom.Node{}},
	{
		gedcom.NewNameNode(nil, "", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
		}),
		gedcom.TagHeader,
		[]gedcom.Node{},
	},
	{
		gedcom.NewNameNode(nil, "", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
		}),
		gedcom.TagSurname,
		[]gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
		},
	},
	{
		gedcom.NewNameNode(nil, "", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
		}),
		gedcom.TagSurname,
		[]gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
		},
	},
	{
		gedcom.NewNameNode(nil, "", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
		}),
		gedcom.TagSurname,
		[]gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
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
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
			}),
			[]gedcom.Tag{},
			[]gedcom.Node{},
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname},
			[]gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil),
				}),
			},
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			[]gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil),
			},
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "1", nil),
				}),
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "2", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			[]gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagText, "", "1", nil),
				gedcom.NewSimpleNode(nil, gedcom.TagText, "", "2", nil),
			},
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "1", nil),
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "2", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagText},
			[]gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagText, "", "1", nil),
				gedcom.NewSimpleNode(nil, gedcom.TagText, "", "2", nil),
			},
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagGivenName, gedcom.TagText},
			[]gedcom.Node{},
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil),
				}),
			}),
			[]gedcom.Tag{gedcom.TagSurname, gedcom.TagSurname},
			[]gedcom.Node{},
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagText, "", "", nil),
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
	surname := gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil)
	givenName := gedcom.NewSimpleNode(nil, gedcom.TagGivenName, "", "", nil)

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
			gedcom.NewNameNode(nil, "", "", nil),
			surname,
			false,
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
			surname,
			false,
		},

		// Other cases.
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				surname,
			}),
			surname,
			true,
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				surname,
			}),
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "", "", nil),
			false,
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				givenName,
			}),
			surname,
			false,
		},
		{
			gedcom.NewNameNode(nil, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagGivenName, "", "", []gedcom.Node{
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
