package gedcom_test

import (
	"strings"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
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

		// Do not merge Residence (based on Equals).
		{
			left: parse(
				"0 INDI @P3@",
				"1 NAME John /Smith/",
				"1 RESI",
				"2 PLAC England",
				"1 RESI",
				"2 DATE Abt. Oct 1943",
			)[0],
			right: parse(
				"0 INDI @P3@",
				"1 NAME John /Smith/",
				"1 RESI",
				"2 DATE About Oct 1943",
				"1 RESI",
				"2 PLAC Yorkshire, England",
			)[0],
			expected: &gedcom.NodeDiff{
				Left:  parse("0 INDI @P3@")[0],
				Right: parse("0 INDI @P3@")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left:     parse("0 NAME John /Smith/")[0],
						Right:    parse("0 NAME John /Smith/")[0],
						Children: nil,
					},
					{
						Left:  parse("0 RESI")[0],
						Right: nil,
						Children: []*gedcom.NodeDiff{
							{
								Left:  parse("0 PLAC England")[0],
								Right: nil,
							},
						},
					},
					{
						Left:  parse("0 RESI")[0],
						Right: parse("0 RESI")[0],
						Children: []*gedcom.NodeDiff{
							{
								Left:  parse("0 DATE Abt. Oct 1943")[0],
								Right: parse("0 DATE About Oct 1943")[0],
							},
						},
					},
					{
						Left:  nil,
						Right: parse("0 RESI")[0],
						Children: []*gedcom.NodeDiff{
							{
								Left:  nil,
								Right: parse("0 PLAC Yorkshire, England")[0],
							},
						},
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
				"0 INDI @P3@",
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
				Right: parse("0 INDI @P3@")[0],
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

func TestNodeDiff_IsDeepEqual(t *testing.T) {
	d := &gedcom.NodeDiff{
		Left:  parse("0 INDI @P3@")[0],
		Right: parse("0 INDI @P3@")[0],
		Children: []*gedcom.NodeDiff{
			{
				Left:  parse("0 NAME John /Smith/")[0],
				Right: parse("0 NAME John /Smith/")[0],
			},
			{
				Left:  parse("0 BIRT")[0],
				Right: parse("0 BIRT")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left: parse("0 DATE Abt. Oct 1943")[0],
					},
					{
						Left:  parse("0 DATE 3 SEP 1943")[0],
						Right: parse("0 DATE 3 SEP 1943")[0],
					},
					{
						Right: parse("0 DATE Abt. Sep 1943")[0],
					},
				},
			},
			{
				Left:  parse("0 DEAT")[0],
				Right: parse("0 DEAT")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left:  parse("0 PLAC England")[0],
						Right: parse("0 PLAC England")[0],
					},
				},
			},
			{
				Right: parse("0 NAME J. /Smith/")[0],
			},
		},
	}

	assert.Equal(t, strings.TrimSpace(`
LR 0 INDI @P3@
LR 1 NAME John /Smith/
LR 1 BIRT
L  2 DATE Abt. Oct 1943
LR 2 DATE 3 SEP 1943
 R 2 DATE Abt. Sep 1943
LR 1 DEAT
LR 2 PLAC England
 R 1 NAME J. /Smith/`),
		d.String())

	tests := []struct {
		line     string
		nd       *gedcom.NodeDiff
		expected bool
	}{
		{"LR 0 INDI @P3@", d, false},
		{"LR 0 NAME John /Smith/", d.Children[0], true},

		{"LR 0 BIRT", d.Children[1], false},
		{"L  0 DATE Abt. Oct 1943", d.Children[1].Children[0], false},
		{"LR 0 DATE 3 SEP 1943", d.Children[1].Children[1], true},
		{" R 0 DATE Abt. Sep 1943", d.Children[1].Children[2], false},

		{"LR 0 DEAT", d.Children[2], true},
		{"LR 0 PLAC England", d.Children[2].Children[0], true},

		{" R 0 NAME J. /Smith/", d.Children[3], false},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			firstLine := strings.Split(test.nd.String(), "\n")[0]
			assert.Equal(t, test.line, firstLine)
			assert.Equal(t, test.expected, test.nd.IsDeepEqual())
		})
	}
}

func TestNodeDiff_Sort(t *testing.T) {
	d := &gedcom.NodeDiff{
		Left:  parse("0 INDI @P3@")[0],
		Right: parse("0 INDI @P3@")[0],
		Children: []*gedcom.NodeDiff{
			{
				Left:  parse("0 BURI")[0],
				Right: parse("0 BURI")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left: parse("0 DATE Abt. 1943")[0],
					},
				},
			},
			{
				Left: parse("0 NAME John /Smith/")[0],
			},
			{
				Left:  parse("0 DEAT")[0],
				Right: parse("0 DEAT")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left:  parse("0 PLAC England")[0],
						Right: parse("0 PLAC England")[0],
					},
				},
			},
			{
				Left: parse("0 BIRT")[0],
				Children: []*gedcom.NodeDiff{
					{
						Right: parse("0 PLAC England")[0],
					},
					{
						Left: parse("0 DATE 1943")[0],
					},
					{
						Right: parse("0 DATE 3 SEP 1942")[0],
					},
				},
			},
			{
				Left: parse("0 SEX M")[0],
			},
			{
				Right: parse("0 NAME John R /Smith/")[0],
			},
			{
				Left:  parse("0 RESI")[0],
				Right: parse("0 RESI")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left: parse("0 DATE 3 Mar 1937")[0],
					},
				},
			},
			{
				Left:  parse("0 RESI")[0],
				Right: parse("0 RESI")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left: parse("0 DATE Aft. Sep 1920")[0],
					},
				},
			},
		},
	}

	assert.Equal(t, strings.TrimSpace(`
LR 0 INDI @P3@
LR 1 BURI
L  2 DATE Abt. 1943
L  1 NAME John /Smith/
LR 1 DEAT
LR 2 PLAC England
L  1 BIRT
 R 2 PLAC England
L  2 DATE 1943
 R 2 DATE 3 SEP 1942
L  1 SEX M
 R 1 NAME John R /Smith/
LR 1 RESI
L  2 DATE 3 Mar 1937
LR 1 RESI
L  2 DATE Aft. Sep 1920`), d.String())

	d.Sort()

	assert.Equal(t, strings.TrimSpace(`
LR 0 INDI @P3@
L  1 NAME John /Smith/
 R 1 NAME John R /Smith/
L  1 SEX M
L  1 BIRT
 R 2 DATE 3 SEP 1942
L  2 DATE 1943
 R 2 PLAC England
LR 1 RESI
L  2 DATE Aft. Sep 1920
LR 1 RESI
L  2 DATE 3 Mar 1937
LR 1 DEAT
LR 2 PLAC England
LR 1 BURI
L  2 DATE Abt. 1943`), d.String())
}

func TestNodeDiff_LeftNode(t *testing.T) {
	d := &gedcom.NodeDiff{
		Left:  parse("0 INDI @P3@")[0],
		Right: parse("0 INDI @P4@")[0],
		Children: []*gedcom.NodeDiff{
			{
				Left: parse("0 BURI")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left:  parse("0 DATE 1943")[0],
						Right: parse("0 DATE Abt. 1943")[0],
					},
				},
			},
			{
				Right: parse("0 NAME John /Smith/")[0],
			},
		},
	}

	assert.Equal(t, `0 INDI @P3@
1 BURI
2 DATE 1943
1 NAME John /Smith/
`, gedcom.NodeGedcom(d.LeftNode()))
}

func TestNodeDiff_RightNode(t *testing.T) {
	d := &gedcom.NodeDiff{
		Left:  parse("0 INDI @P3@")[0],
		Right: parse("0 INDI @P4@")[0],
		Children: []*gedcom.NodeDiff{
			{
				Left: parse("0 BURI")[0],
				Children: []*gedcom.NodeDiff{
					{
						Left:  parse("0 DATE 1943")[0],
						Right: parse("0 DATE Abt. 1943")[0],
					},
				},
			},
			{
				Left:  parse("0 NAME J /Smith/")[0],
				Right: parse("0 NAME John /Smith/")[0],
			},
		},
	}

	assert.Equal(t, `0 INDI @P4@
1 BURI
2 DATE Abt. 1943
1 NAME John /Smith/
`, gedcom.NodeGedcom(d.RightNode()))
}
