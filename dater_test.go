package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDates(t *testing.T) {
	tests := []struct {
		nodes []gedcom.Node
		want  gedcom.DateNodes
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
					gedcom.NewDateNode("2 Sep 1981"),
				),
			},
			gedcom.DateNodes{
				gedcom.NewDateNode("2 Sep 1981"),
			},
		},
		{
			gedcom.Nodes{
				gedcom.NewNameNode("foo bar",
					gedcom.NewDateNode("2 Sep 1981"),
					gedcom.NewDateNode("3 Sep 1981"),
				),
			},
			gedcom.DateNodes{
				gedcom.NewDateNode("2 Sep 1981"),
				gedcom.NewDateNode("3 Sep 1981"),
			},
		},
		{
			gedcom.Nodes{
				gedcom.NewNameNode("foo bar"),
			},
			nil,
		},
		{
			gedcom.Nodes{
				gedcom.NewNameNode("foo bar",
					gedcom.NewDateNode("bar baz"),
					gedcom.NewDateNode("3 Sep 1981"),
				),
			},
			[]*gedcom.DateNode{
				gedcom.NewDateNode("bar baz"),
				gedcom.NewDateNode("3 Sep 1981"),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Dates(test.nodes...))
		})
	}
}
