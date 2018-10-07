package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

func TestIndividualNodes_Similarity(t *testing.T) {
	// ghost:ignore
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
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
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
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1843"),
					died("Apr 1907"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
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
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					died("Apr 1907"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("John /Smith/"),
					died("Apr 1907"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.875,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("John /Smith/"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.75,
		},

		// Similar matches but the same sized slice on both sides.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P4", []gedcom.Node{
					name("John /Smith/"),
					born("Abt. Jan 1843"),
					died("1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P5", []gedcom.Node{
					name("Jane /Doe/"),
					born("Bef. 1846"),
				}),
				gedcom.NewIndividualNode(nil, "", "P6", []gedcom.Node{
					name("Bob Thomas /Jones/"),
					buried("1927"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.872532146404072,
		},

		// The slices are different lengths. The same score should be returned
		// when different sizes slices are swapped.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P4", []gedcom.Node{
					name("Jane /Doe/"),
					born("Between 1845 and 1846"),
				}),
				gedcom.NewIndividualNode(nil, "", "P5", []gedcom.Node{
					name("John /Smith/"),
					born("Bef. 10 Jan 1843"),
					died("Abt. 1908"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.7754008744441251,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P4", []gedcom.Node{
					name("Jane /Doe/"),
					born("Between 1845 and 1846"),
				}),
				gedcom.NewIndividualNode(nil, "", "P5", []gedcom.Node{
					name("John /Smith/"),
					born("Bef. 10 Jan 1843"),
					died("Abt. 1908"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.7754008744441251,
		},

		// Whenever one slice is empty the result will always be 0.5.
		{
			a: gedcom.IndividualNodes{},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					buried("1927"),
				}),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
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
				gedcom.NewIndividualNode(nil, "", "P4", []gedcom.Node{
					name("Jane /Doe/"),
					born("Between 1845 and 1846"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P5", []gedcom.Node{
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
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
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
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					born("1627"),
				}),
			},
			minSimilarity: 0.95,
			expected:      0.5,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					name("John /Smith/"),
					born("4 Jan 1843"),
					died("17 Mar 1907"),
				}),
				gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
					name("Jane /Doe/"),
					born("Sep 1845"),
				}),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewIndividualNode(nil, "", "P3", []gedcom.Node{
					name("Bob /Jones/"),
					born("1627"),
				}),
			},
			minSimilarity: 0.0,
			expected:      0.45708333333333334,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			options := gedcom.NewSimilarityOptions()
			options.MinimumSimilarity = test.minSimilarity
			got := test.a.Similarity(test.b, options)

			assert.Equal(t, test.expected, got)
		})
	}
}

func TestIndividualNodes_Compare(t *testing.T) {
	elliot := individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
	john := individual("P2", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
	jane := individual("P3", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
	bob := individual("P4", "Bob /Jones/", "1749", "1810")

	// ghost:ignore
	tests := []struct {
		doc1, doc2 *gedcom.Document
		min        float64
		want       []gedcom.IndividualComparison
	}{
		{
			doc1: gedcom.NewDocument(),
			doc2: gedcom.NewDocument(),
			min:  0.0,
			want: []gedcom.IndividualComparison{},
		},
		{
			doc1: gedcom.NewDocumentWithNodes([]gedcom.Node{elliot}),
			doc2: gedcom.NewDocument(),
			min:  0.0,
			want: []gedcom.IndividualComparison{
				{elliot, nil, gedcom.SurroundingSimilarity{}},
			},
		},
		{
			doc1: gedcom.NewDocument(),
			doc2: gedcom.NewDocumentWithNodes([]gedcom.Node{elliot}),
			min:  0.0,
			want: []gedcom.IndividualComparison{
				{nil, elliot, gedcom.SurroundingSimilarity{}},
			},
		},
		{
			doc1: gedcom.NewDocumentWithNodes([]gedcom.Node{elliot}),
			doc2: gedcom.NewDocumentWithNodes([]gedcom.Node{elliot}),
			min:  0.0,
			want: []gedcom.IndividualComparison{
				{elliot, elliot, gedcom.SurroundingSimilarity{0.5, 1.0, 1.0, 1.0}},
			},
		},
		{
			doc1: gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, john, jane}),
			doc2: gedcom.NewDocumentWithNodes([]gedcom.Node{jane, elliot, john}),
			min:  0.0,
			want: []gedcom.IndividualComparison{
				{elliot, elliot, gedcom.SurroundingSimilarity{0.5, 1.0, 1.0, 1.0}},
				{john, john, gedcom.SurroundingSimilarity{0.5, 1.0, 1.0, 1.0}},
				{jane, jane, gedcom.SurroundingSimilarity{0.5, 1.0, 1.0, 1.0}},
			},
		},
		{
			doc1: gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, jane}),
			doc2: gedcom.NewDocumentWithNodes([]gedcom.Node{jane, john}),
			min:  0.0,
			want: []gedcom.IndividualComparison{
				// elliot and john match because the minimumSimilarity is so
				// low.
				{jane, jane, gedcom.SurroundingSimilarity{0.5, 1, 1.0, 1.0}},
				{elliot, john, gedcom.SurroundingSimilarity{0.5, 0.24743589743589745, 1.0, 1.0}},
			},
		},
		{
			doc1: gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, jane}),
			doc2: gedcom.NewDocumentWithNodes([]gedcom.Node{jane, john}),
			min:  0.75,
			want: []gedcom.IndividualComparison{
				{jane, jane, gedcom.SurroundingSimilarity{0.5, 1.0, 1.0, 1.0}},
				{elliot, nil, gedcom.SurroundingSimilarity{}},
				{nil, john, gedcom.SurroundingSimilarity{}},
			},
		},
		{
			doc1: gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, jane}),
			doc2: gedcom.NewDocumentWithNodes([]gedcom.Node{bob, john}),
			min:  0.9,
			want: []gedcom.IndividualComparison{
				{elliot, nil, gedcom.SurroundingSimilarity{}},
				{jane, nil, gedcom.SurroundingSimilarity{}},
				{nil, bob, gedcom.SurroundingSimilarity{}},
				{nil, john, gedcom.SurroundingSimilarity{}},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			for _, n := range test.doc1.Nodes() {
				n.SetDocument(test.doc1)
			}
			for _, n := range test.doc2.Nodes() {
				n.SetDocument(test.doc2)
			}

			options := gedcom.NewSimilarityOptions()
			options.MinimumWeightedSimilarity = test.min

			individuals1 := test.doc1.Individuals()
			individuals2 := test.doc2.Individuals()
			got := individuals1.Compare(individuals2,
				&gedcom.IndividualNodesCompareOptions{
					SimilarityOptions: options,
				})

			assert.Equal(t, test.want, got)
		})
	}
}
