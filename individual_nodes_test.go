package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndividualNodes_Similarity(t *testing.T) {
	var tests = []struct {
		a, b          gedcom.IndividualNodes
		minSimilarity float64
		expected      float64
	}{
		// Exact matches.
		{
			a:             gedcom.IndividualNodes{},
			b:             gedcom.IndividualNodes{},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      1.0,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      1.0,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1843"),
					died("Apr 1907"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("Jane /DOE/"),
					born("Sep 1843"),
					died("Apr 1907"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      1.0,
		},

		// Exact matches, but missing information on both sides. These
		// specifically should NOT return 1.0 as it would throw out the real
		// similarities. See the docs for explanation.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					died("Apr 1907"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("John /Smith/"),
					died("Apr 1907"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.8333333333333334,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("John /Smith/"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},

		// Similar matches but the same sized slice on both sides.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P4", []gedcom.Node{
					name("John /Smith/"),
					born("Abt. Jan 1843"),
					died("1907"),
				}),
				gedcom.NewIndividualNode("", "P5", []gedcom.Node{
					name("Jane /Doe/"),
					born("Bef. 1846"),
				}),
				gedcom.NewIndividualNode("", "P6", []gedcom.Node{
					name("Bob Thomas /Jones/"),
					buried("1927"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.8464260797907109,
		},

		// The slices are different lengths. The same score should be returned
		// when different sizes slices are swapped.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P4", []gedcom.Node{
					name("Jane /Doe/"),
					born("Between 1845 and 1846"),
				}),
				gedcom.NewIndividualNode("", "P5", []gedcom.Node{
					name("John /Smith/"),
					born("Bef. 10 Jan 1843"),
					died("Abt. 1908"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.7758258827110728,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P4", []gedcom.Node{
					name("Jane /Doe/"),
					born("Between 1845 and 1846"),
				}),
				gedcom.NewIndividualNode("", "P5", []gedcom.Node{
					name("John /Smith/"),
					born("Bef. 10 Jan 1843"),
					died("Abt. 1908"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.7758258827110728,
		},

		// Whenever one slice is empty the result will always be 0.5.
		{
			a: gedcom.IndividualNodes{},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			b:             gedcom.IndividualNodes{},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},

		// These ones are just way off and should not be considered matches.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P4", []gedcom.Node{
					name("Jane /Doe/"),
					born("Between 1845 and 1846"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P5", []gedcom.Node{
					name("John /Smith/"),
					born("Bef. 10 Jan 1943"),
					died("Abt. 2008"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					born("1627"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},

		// Different values for minimumSimilarity.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					born("1627"),
				}),
			},
			minSimilarity: 0.95,
			expected:      0.5,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode("", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					born("1627"),
				}),
			},
			minSimilarity: 0.0,
			expected:      0.4219135802469136,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.a.Similarity(test.b, test.minSimilarity), test.expected)
		})
	}
}
