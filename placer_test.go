package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaces(t *testing.T) {
	// ghost:ignore
	tests := []struct {
		nodes gedcom.Nodes
		want  []*gedcom.PlaceNode
	}{
		{nil, nil},
		{
			gedcom.Nodes{
				gedcom.NewNode(gedcom.TagVersion, "foo", ""),
			},
			nil,
		},
		{
			gedcom.Nodes{
				gedcom.NewNameNode("foo bar",
					gedcom.NewPlaceNode("Australia"),
				),
			},
			[]*gedcom.PlaceNode{
				gedcom.NewPlaceNode("Australia"),
			},
		},
		{
			gedcom.Nodes{
				gedcom.NewNameNode("foo bar",
					gedcom.NewPlaceNode("Australia"),
					gedcom.NewPlaceNode("United States"),
				),
			},
			[]*gedcom.PlaceNode{
				gedcom.NewPlaceNode("Australia"),
				gedcom.NewPlaceNode("United States"),
			},
		},
		{
			gedcom.Nodes{
				gedcom.NewNameNode("foo bar",
					gedcom.NewPlaceNode("Australia"),
					gedcom.NewPlaceNode("United States"),
				),
				gedcom.NewNameNode("foo bar",
					gedcom.NewPlaceNode("England"),
				),
			},
			[]*gedcom.PlaceNode{
				gedcom.NewPlaceNode("Australia"),
				gedcom.NewPlaceNode("United States"),
				gedcom.NewPlaceNode("England"),
			},
		},
		{
			gedcom.Nodes{
				gedcom.NewNameNode("foo bar"),
				gedcom.NewNameNode("foo bar",
					gedcom.NewPlaceNode("Australia"),
				),
			},
			[]*gedcom.PlaceNode{
				gedcom.NewPlaceNode("Australia"),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Places(test.nodes...))
		})
	}
}
