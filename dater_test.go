package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDates(t *testing.T) {
	tests := []struct {
		nodes gedcom.Node
		want  gedcom.DateNodes
	}{
		{nil, nil},
		{
			gedcom.NewNodeWithChildren(nil, gedcom.TagVersion, "foo", "", nil),
			nil,
		},
		{
			gedcom.NewNameNode(nil, "foo bar", "", []gedcom.Node{
				gedcom.NewDateNode(nil, "2 Sep 1981", "", nil),
			}),
			gedcom.DateNodes{
				gedcom.NewDateNode(nil, "2 Sep 1981", "", nil),
			},
		},
		{
			gedcom.NewNameNode(nil, "foo bar", "", []gedcom.Node{
				gedcom.NewDateNode(nil, "2 Sep 1981", "", nil),
				gedcom.NewDateNode(nil, "3 Sep 1981", "", nil),
			}),
			gedcom.DateNodes{
				gedcom.NewDateNode(nil, "2 Sep 1981", "", nil),
				gedcom.NewDateNode(nil, "3 Sep 1981", "", nil),
			},
		},
		{
			gedcom.NewNameNode(nil, "foo bar", "", nil),
			nil,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Dates(test.nodes))
		})
	}
}
