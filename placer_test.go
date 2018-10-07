package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaces(t *testing.T) {
	// ghost:ignore
	tests := []struct {
		nodes []gedcom.Node
		want  []*gedcom.PlaceNode
	}{
		{nil, nil},
		{
			[]gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagVersion, "foo", "", nil),
			},
			nil,
		},
		{
			[]gedcom.Node{
				gedcom.NewNameNode(nil, "foo bar", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Australia", "", nil),
				}),
			},
			[]*gedcom.PlaceNode{
				gedcom.NewPlaceNode(nil, "Australia", "", nil),
			},
		},
		{
			[]gedcom.Node{
				gedcom.NewNameNode(nil, "foo bar", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Australia", "", nil),
					gedcom.NewPlaceNode(nil, "United States", "", nil),
				}),
			},
			[]*gedcom.PlaceNode{
				gedcom.NewPlaceNode(nil, "Australia", "", nil),
				gedcom.NewPlaceNode(nil, "United States", "", nil),
			},
		},
		{
			[]gedcom.Node{
				gedcom.NewNameNode(nil, "foo bar", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Australia", "", nil),
					gedcom.NewPlaceNode(nil, "United States", "", nil),
				}),
				gedcom.NewNameNode(nil, "foo bar", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "England", "", nil),
				}),
			},
			[]*gedcom.PlaceNode{
				gedcom.NewPlaceNode(nil, "Australia", "", nil),
				gedcom.NewPlaceNode(nil, "United States", "", nil),
				gedcom.NewPlaceNode(nil, "England", "", nil),
			},
		},
		{
			[]gedcom.Node{
				gedcom.NewNameNode(nil, "foo bar", "", nil),
				gedcom.NewNameNode(nil, "foo bar", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Australia", "", nil),
				}),
			},
			[]*gedcom.PlaceNode{
				gedcom.NewPlaceNode(nil, "Australia", "", nil),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Places(test.nodes...))
		})
	}
}
