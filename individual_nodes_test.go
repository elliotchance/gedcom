package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

var (
	// John and Jane share the same pointer on purpose. They will be used for
	// pointer comparisons.
	elliot = individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
	john   = individual("P2", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
	jane   = individual("P2", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
	bob    = individual("P4", "Bob /Jones/", "1749", "1810")
	harry  = individual("P5", "Harry /Gold/", "1889", "1936")
)

var individualNodesTests = map[string]struct {
	Doc1, Doc2 *gedcom.Document

	MinimumWeightedSimilarity float64
	PreferPointerAbove        float64

	WantCompare gedcom.IndividualComparisons
	WantMerge   gedcom.IndividualNodes
}{
	"BothDocumentsEmpty": {
		Doc1:                      gedcom.NewDocument(),
		Doc2:                      gedcom.NewDocument(),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare:               gedcom.IndividualComparisons{},
	},
	"Doc2Empty": {
		Doc1:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot}),
		Doc2:                      gedcom.NewDocument(),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			{elliot, nil, nil},
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
		},
	},
	"Doc1Empty": {
		Doc1:                      gedcom.NewDocument(),
		Doc2:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot}),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			{nil, elliot, nil},
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
		},
	},
	"SameIndividualInBothDocuments": {
		Doc1:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot}),
		Doc2:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot}),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			{elliot, elliot, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)},
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
		},
	},
	"SameIndividualsInDifferentOrder": {
		Doc1:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, john, jane}),
		Doc2:                      gedcom.NewDocumentWithNodes([]gedcom.Node{jane, elliot, john}),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			{elliot, elliot, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)},
			{john, john, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)},
			{jane, jane, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)},
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
			john,
			jane,
		},
	},
	"ZeroMinimumSimilarity": {
		Doc1:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, jane}),
		Doc2:                      gedcom.NewDocumentWithNodes([]gedcom.Node{jane, john}),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			// elliot and john match because the minimumSimilarity is so
			// low.
			{jane, jane, gedcom.NewSurroundingSimilarity(0.5, 1, 1.0, 1.0)},
			{elliot, john, gedcom.NewSurroundingSimilarity(0.5, 0.24743589743589745, 1.0, 1.0)},
		},
		WantMerge: gedcom.IndividualNodes{
			jane,
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "4 Jan 1803", "", nil), // john
					gedcom.NewDateNode(nil, "4 Jan 1843", "", nil), // elliot
				}),
				gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "17 Mar 1877", "", nil), // john
					gedcom.NewDateNode(nil, "17 Mar 1907", "", nil), // elliot
				}),
				gedcom.NewNameNode(nil, "John /Smith/", "", nil),
			}),
		},
	},
	"OneMatch": {
		Doc1:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, jane}),
		Doc2:                      gedcom.NewDocumentWithNodes([]gedcom.Node{jane, john}),
		MinimumWeightedSimilarity: 0.75,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			{jane, jane, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)},
			{elliot, nil, nil},
			{nil, john, nil},
		},
		WantMerge: gedcom.IndividualNodes{
			jane,
			elliot,
			john,
		},
	},
	"NoMatches": {
		Doc1:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, jane}),
		Doc2:                      gedcom.NewDocumentWithNodes([]gedcom.Node{bob, john}),
		MinimumWeightedSimilarity: 0.9,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			{elliot, nil, nil},
			{jane, nil, nil},
			{nil, bob, nil},
			{nil, john, nil},
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
			jane,
			bob,
			john,
		},
	},
	"AlwaysUsePointer": {
		// John and Jane are both P2. Even though they are completely different
		// we force pointers to match with a prefer value of 0.0.
		Doc1:                      gedcom.NewDocumentWithNodes([]gedcom.Node{elliot, jane}),
		Doc2:                      gedcom.NewDocumentWithNodes([]gedcom.Node{bob, john}),
		MinimumWeightedSimilarity: 0.9,
		PreferPointerAbove:        0.0,
		WantCompare: gedcom.IndividualComparisons{
			{jane, john, gedcom.NewSurroundingSimilarity(0.5, 0.8209932199959546, 1.0, 1.0)},
			{elliot, nil, nil},
			{nil, bob, nil},
		},
		WantMerge: gedcom.IndividualNodes{
			gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
				gedcom.NewNameNode(nil, "Jane /Doe/", "", nil),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "4 Jan 1803", "", nil), // john
					gedcom.NewDateNode(nil, "3 Mar 1803", "", nil), // jane
				}),
				gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "17 Mar 1877", "", nil),  // john
					gedcom.NewDateNode(nil, "14 June 1877", "", nil), // jane
				}),
				gedcom.NewNameNode(nil, "John /Smith/", "", nil),
			}),
			elliot,
			bob,
		},
	},
	"AlwaysUsePointerUID": {
		// John and Jane are both P2. Even though they are completely different
		// we force pointers to match with a prefer value of 0.0.
		Doc1: gedcom.NewDocumentWithNodes([]gedcom.Node{
			elliot,
			setUID(individual("P5", "Harry /Gold/", "1889", "1936"), "EE13561DDB204985BFFDEEBF82A5226C"),
		}),
		Doc2: gedcom.NewDocumentWithNodes([]gedcom.Node{
			bob,
			setUID(individual("P2", "John /Smith/", "4 Jan 1803", "17 Mar 1877"), "EE13561DDB204985BFFDEEBF82A5226C5B2E"),
		}),
		MinimumWeightedSimilarity: 0.9,
		PreferPointerAbove:        0.0,
		WantCompare: gedcom.IndividualComparisons{
			{harry, john, gedcom.NewSurroundingSimilarity(0.5, 0.15, 1.0, 1.0)},
			{elliot, nil, nil},
			{nil, bob, nil},
		},
		WantMerge: gedcom.IndividualNodes{
			gedcom.NewIndividualNode(nil, "", "P5", []gedcom.Node{ // P5 = harry
				gedcom.NewNameNode(nil, "Harry /Gold/", "", nil),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "4 Jan 1803", "", nil), // john
					gedcom.NewDateNode(nil, "1889", "", nil),       // harry
				}),
				gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "17 Mar 1877", "", nil), // john
					gedcom.NewDateNode(nil, "1936", "", nil),        // harry
				}),
				gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C", "", nil),
				gedcom.NewNameNode(nil, "John /Smith/", "", nil),
			}),
			elliot,
			bob,
		},
	},
}

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
	for testName, test := range individualNodesTests {
		t.Run(testName, func(t *testing.T) {
			for _, n := range test.Doc1.Nodes() {
				n.SetDocument(test.Doc1)
			}

			for _, n := range test.Doc2.Nodes() {
				n.SetDocument(test.Doc2)
			}

			similarityOptions := gedcom.NewSimilarityOptions()
			similarityOptions.MinimumWeightedSimilarity = test.MinimumWeightedSimilarity
			similarityOptions.PreferPointerAbove = test.PreferPointerAbove

			compareOptions := gedcom.NewIndividualNodesCompareOptions()
			compareOptions.SimilarityOptions = similarityOptions

			individuals1 := test.Doc1.Individuals()
			individuals2 := test.Doc2.Individuals()
			got := individuals1.Compare(individuals2, compareOptions)

			// The comparison results (got) will include the options from above.
			// However, the fixture for this test does not provide the
			// compareOptions as it would make the fixture verbose and
			// confusing. Instead we set the Options on each of the comparison
			// results so that the deep equal passes.
			for _, x := range test.WantCompare {
				if x.Similarity != nil {
					x.Similarity.Options = similarityOptions
				}
			}

			assertEqual(t, test.WantCompare, got)
		})
	}
}

func assertEqual(t *testing.T, expected, actual interface{}) bool {
	diff := cmp.Diff(expected, actual, cmpopts.IgnoreUnexported(gedcom.SimpleNode{}, gedcom.IndividualNode{}))
	if diff != "" {
		assert.Fail(t, diff)
	}

	return diff == ""
}

func TestNewIndividualNodesCompareOptions(t *testing.T) {
	actual := gedcom.NewIndividualNodesCompareOptions()

	assert.Equal(t, actual.SimilarityOptions, gedcom.NewSimilarityOptions())
}

func TestIndividualNodes_Nodes(t *testing.T) {
	Nodes := tf.Function(t, gedcom.IndividualNodes.Nodes)

	i1 := individual("P1", "Elliot /Chance/", "", "")
	i2 := individual("P2", "Joe /Bloggs/", "", "")

	Nodes(nil).Returns(nil)
	Nodes(gedcom.IndividualNodes{}).Returns(nil)
	Nodes(gedcom.IndividualNodes{i1, i2}).Returns([]gedcom.Node{i1, i2})
}

func TestIndividualNodes_Merge(t *testing.T) {
	for testName, test := range individualNodesTests {
		t.Run(testName, func(t *testing.T) {
			for _, n := range test.Doc1.Nodes() {
				n.SetDocument(test.Doc1)
			}

			for _, n := range test.Doc2.Nodes() {
				n.SetDocument(test.Doc2)
			}

			similarityOptions := gedcom.NewSimilarityOptions()
			similarityOptions.MinimumWeightedSimilarity = test.MinimumWeightedSimilarity
			similarityOptions.PreferPointerAbove = test.PreferPointerAbove

			compareOptions := gedcom.NewIndividualNodesCompareOptions()
			compareOptions.SimilarityOptions = similarityOptions

			individuals1 := test.Doc1.Individuals()
			individuals2 := test.Doc2.Individuals()
			got, err := individuals1.Merge(individuals2, compareOptions)

			assert.NoError(t, err)
			assertIndividualNodes(t, test.WantMerge, got)
		})
	}
}

func assertIndividualNodes(t *testing.T, expected, actual gedcom.IndividualNodes) {
	assert.Equal(t, expected.GEDCOMString(0), actual.GEDCOMString(0))
}

func setUID(i *gedcom.IndividualNode, uid string) *gedcom.IndividualNode {
	i.AddNode(gedcom.NewUniqueIDNode(i.Document(), uid, "", nil))

	return i
}
