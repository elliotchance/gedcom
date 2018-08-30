package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func parse(s ...string) []gedcom.Node {
	doc, err := gedcom.NewDocumentFromString(strings.Join(s, "\n"))
	if err != nil {
		panic(err)
	}

	return doc.Nodes
}

func TestCompareNodes(t *testing.T) {
	tests := []struct {
		left, right gedcom.Node
		expected    *gedcom.NodeDiff
	}{
		// Nils
		{
			left:     nil,
			right:    nil,
			expected: &gedcom.NodeDiff{},
		},
		{
			left:  parse("0 @P1@ BIRT foo")[0],
			right: nil,
			expected: &gedcom.NodeDiff{
				Left: parse("0 @P1@ BIRT foo")[0],
			},
		},
		{
			left:  nil,
			right: parse("0 @P1@ BIRT foo")[0],
			expected: &gedcom.NodeDiff{
				Right: parse("0 @P1@ BIRT foo")[0],
			},
		},
		{
			left:  parse("0 INDI @P1@", "1 BIRT")[0],
			right: nil,
			expected: &gedcom.NodeDiff{
				Left:  parse("0 INDI @P1@")[0],
				Right: nil,
				Children: []*gedcom.NodeDiff{
					{
						Left: parse("0 BIRT")[0],
					},
				},
			},
		},
		{
			left:  nil,
			right: parse("0 INDI @P1@", "1 DEAT")[0],
			expected: &gedcom.NodeDiff{
				Left:  nil,
				Right: parse("0 INDI @P1@")[0],
				Children: []*gedcom.NodeDiff{
					{
						Right: parse("0 DEAT")[0],
					},
				},
			},
		},

		// Different root nodes.
		{
			left:  parse("0 INDI @P1@")[0],
			right: parse("0 INDI @P2@")[0],
			expected: &gedcom.NodeDiff{
				Left:  parse("0 INDI @P1@")[0],
				Right: parse("0 INDI @P2@")[0],
			},
		},
		{
			left:  parse("0 INDI @P1@")[0],
			right: parse("0 INDI @P2@", "1 DEAT")[0],
			expected: &gedcom.NodeDiff{
				Left:  parse("0 INDI @P1@")[0],
				Right: parse("0 INDI @P2@")[0],
				Children: []*gedcom.NodeDiff{
					{
						Right: parse("0 DEAT")[0],
					},
				},
			},
		},
		{
			left:  parse("0 INDI @P1@", "1 BIRT")[0],
			right: parse("0 INDI @P2@", "1 DEAT")[0],
			expected: &gedcom.NodeDiff{
				Left:  parse("0 INDI @P1@")[0],
				Right: parse("0 INDI @P2@")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left: parse("0 BIRT")[0],
					},
					{
						Right: parse("0 DEAT")[0],
					},
				},
			},
		},
		{
			left:  parse("0 INDI @P1@", "1 BIRT")[0],
			right: parse("0 INDI @P2@", "1 BIRT")[0],
			expected: &gedcom.NodeDiff{
				Left:  parse("0 INDI @P1@")[0],
				Right: parse("0 INDI @P2@")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left:  parse("0 BIRT")[0],
						Right: parse("0 BIRT")[0],
					},
				},
			},
		},

		// Example from the docs.
		{
			left: parse(
				"0 INDI @P3@",
				"1 NAME John /Smith/",
				"1 BIRT",
				"2 DATE 3 SEP 1943",
				"1 DEAT",
				"2 PLAC England",
				"1 BIRT",
				"2 DATE Abt. Oct 1943",
			)[0],
			right: parse(
				"0 INDI @P4@",
				"1 NAME J. /Smith/",
				"1 BIRT",
				"2 DATE Abt. Sep 1943",
				"1 DEAT",
				"2 DATE Aft. 2001",
				"1 BIRT",
				"2 DATE 3 SEP 1943",
				"2 PLAC Surry, England",
			)[0],
			expected: &gedcom.NodeDiff{
				Left:  parse("0 INDI @P3@")[0],
				Right: parse("0 INDI @P4@")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left:     parse("0 NAME John /Smith/")[0],
						Right:    nil,
						Children: nil,
					},
					{
						Left:  parse("0 BIRT")[0],
						Right: parse("0 BIRT")[0],
						Children: []*gedcom.NodeDiff{
							{
								Left:  parse("0 DATE 3 SEP 1943")[0],
								Right: parse("0 DATE 3 SEP 1943")[0],
							},
							{
								Left:  parse("0 DATE Abt. Oct 1943")[0],
								Right: nil,
							},
							{
								Left:  nil,
								Right: parse("0 DATE Abt. Sep 1943")[0],
							},
							{
								Left:  nil,
								Right: parse("0 PLAC Surry, England")[0],
							},
						},
					},
					{
						Left:  parse("0 DEAT")[0],
						Right: parse("0 DEAT")[0],
						Children: []*gedcom.NodeDiff{
							{
								Left:  parse("0 PLAC England")[0],
								Right: nil,
							},
							{
								Left:  nil,
								Right: parse("0 DATE Aft. 2001")[0],
							},
						},
					},
					{
						Left:     nil,
						Right:    parse("0 NAME J. /Smith/")[0],
						Children: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			actual := gedcom.CompareNodes(test.left, test.right)
			assert.Equal(t, test.expected.String(), actual.String())
		})
	}
}
