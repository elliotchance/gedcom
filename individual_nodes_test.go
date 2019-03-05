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
	elliot = individual(gedcom.NewDocument(), "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
	john   = individual(gedcom.NewDocument(), "P2", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
	jane   = individual(gedcom.NewDocument(), "P2", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
	bob    = individual(gedcom.NewDocument(), "P4", "Bob /Jones/", "1749", "1810")
	harry  = individual(gedcom.NewDocument(), "P5", "Harry /Gold/", "1889", "1936")
)

var individualNodesTests = map[string]struct {
	Doc1, Doc2 *gedcom.Document

	MinimumWeightedSimilarity float64
	PreferPointerAbove        float64

	WantCompare gedcom.IndividualComparisons
	WantMerge   gedcom.IndividualNodes
}{
	"BothDocumentsEmpty": {
		Doc1: gedcom.NewDocument(),
		Doc2: gedcom.NewDocument(),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare:               gedcom.IndividualComparisons{},
	},
	"Doc2Empty": {
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot}),
		Doc2: gedcom.NewDocument(),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(elliot, nil, nil),
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
		},
	},
	"Doc1Empty": {
		Doc1: gedcom.NewDocument(),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot}),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(nil, elliot, nil),
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
		},
	},
	"SameIndividualInBothDocuments": {
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot}),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot}),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(elliot, elliot, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)),
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
		},
	},
	"SameIndividualsInDifferentOrder": {
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot, john, jane}),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{jane, elliot, john}),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(elliot, elliot, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)),
			gedcom.NewIndividualComparison(john, john, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)),
			gedcom.NewIndividualComparison(jane, jane, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)),
		},
		WantMerge: gedcom.IndividualNodes{
			elliot,
			john,
			jane,
		},
	},
	"ZeroMinimumSimilarity": {
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot, jane}),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{jane, john}),
		MinimumWeightedSimilarity: 0.0,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			// elliot and john match because the minimumSimilarity is so
			// low.
			gedcom.NewIndividualComparison(jane, jane, gedcom.NewSurroundingSimilarity(0.5, 1, 1.0, 1.0)),
			gedcom.NewIndividualComparison(elliot, john, gedcom.NewSurroundingSimilarity(0.5, 0.24743589743589745, 1.0, 1.0)),
		},
		WantMerge: gedcom.IndividualNodes{
			jane,
			gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("4 Jan 1803"), // john
					gedcom.NewDateNode("4 Jan 1843"), // elliot
				),
				gedcom.NewDeathNode("",
					gedcom.NewDateNode("17 Mar 1877"), // john
					gedcom.NewDateNode("17 Mar 1907"), // elliot
				),
				gedcom.NewNameNode("John /Smith/"),
			),
		},
	},
	"OneMatch": {
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot, jane}),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{jane, john}),
		MinimumWeightedSimilarity: 0.75,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(jane, jane, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)),
			gedcom.NewIndividualComparison(elliot, nil, nil),
			gedcom.NewIndividualComparison(nil, john, nil),
		},
		WantMerge: gedcom.IndividualNodes{
			jane,
			elliot,
			john,
		},
	},
	"NoMatches": {
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot, jane}),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{bob, john}),
		MinimumWeightedSimilarity: 0.9,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(elliot, nil, nil),
			gedcom.NewIndividualComparison(jane, nil, nil),
			gedcom.NewIndividualComparison(nil, bob, nil),
			gedcom.NewIndividualComparison(nil, john, nil),
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
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{elliot, jane}),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{bob, john}),
		MinimumWeightedSimilarity: 0.9,
		PreferPointerAbove:        0.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(jane, john, gedcom.NewSurroundingSimilarity(0.5, 0.8209932199959546, 1.0, 1.0)),
			gedcom.NewIndividualComparison(elliot, nil, nil),
			gedcom.NewIndividualComparison(nil, bob, nil),
		},
		WantMerge: gedcom.IndividualNodes{
			gedcom.NewDocument().AddIndividual("P2",
				gedcom.NewNameNode("Jane /Doe/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("4 Jan 1803"), // john
					gedcom.NewDateNode("3 Mar 1803"), // jane
				),
				gedcom.NewDeathNode("",
					gedcom.NewDateNode("17 Mar 1877"),  // john
					gedcom.NewDateNode("14 June 1877"), // jane
				),
				gedcom.NewNameNode("John /Smith/"),
			),
			elliot,
			bob,
		},
	},
	"AlwaysUseUID1": {
		// Harry and John will always match because of the shared unique
		// identifier.
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			elliot,
			setUID(individual(gedcom.NewDocument(), "P5", "Harry /Gold/", "1889", "1936"), "EE13561DDB204985BFFDEEBF82A5226C"),
		}),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			bob,
			setUID(individual(gedcom.NewDocument(), "P2", "John /Smith/", "4 Jan 1803", "17 Mar 1877"), "EE13561DDB204985BFFDEEBF82A5226C5B2E"),
		}),
		MinimumWeightedSimilarity: 0.9,
		PreferPointerAbove:        0.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(harry, john, gedcom.NewSurroundingSimilarity(0.5, 0.15, 1.0, 1.0)),
			gedcom.NewIndividualComparison(elliot, nil, nil),
			gedcom.NewIndividualComparison(nil, bob, nil),
		},
		WantMerge: gedcom.IndividualNodes{
			gedcom.NewDocument().AddIndividual("P5", // P5 = harry
				gedcom.NewNameNode("Harry /Gold/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("4 Jan 1803"), // john
					gedcom.NewDateNode("1889"),       // harry
				),
				gedcom.NewDeathNode("",
					gedcom.NewDateNode("17 Mar 1877"), // john
					gedcom.NewDateNode("1936"),        // harry
				),
				gedcom.NewUniqueIDNode("EE13561DDB204985BFFDEEBF82A5226C"),
				gedcom.NewNameNode("John /Smith/"),
			),
			elliot,
			bob,
		},
	},
	"AlwaysUseUID2": {
		// This is the same as above, but we use the opposite PreferPointerAbove
		// value to prove that it doesn't affect unique identifier matches.
		Doc1: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			elliot,
			setUID(individual(gedcom.NewDocument(), "P5", "Harry /Gold/", "1889", "1936"), "EE13561DDB204985BFFDEEBF82A5226C"),
		}),
		Doc2: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			bob,
			setUID(individual(gedcom.NewDocument(), "P2", "John /Smith/", "4 Jan 1803", "17 Mar 1877"), "EE13561DDB204985BFFDEEBF82A5226C5B2E"),
		}),
		MinimumWeightedSimilarity: 0.9,
		PreferPointerAbove:        1.0,
		WantCompare: gedcom.IndividualComparisons{
			gedcom.NewIndividualComparison(harry, john, gedcom.NewSurroundingSimilarity(0.5, 0.15, 1.0, 1.0)),
			gedcom.NewIndividualComparison(elliot, nil, nil),
			gedcom.NewIndividualComparison(nil, bob, nil),
		},
		WantMerge: gedcom.IndividualNodes{
			gedcom.NewDocument().AddIndividual("P5", // P5 = harry
				gedcom.NewNameNode("Harry /Gold/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("4 Jan 1803"), // john
					gedcom.NewDateNode("1889"),       // harry
				),
				gedcom.NewDeathNode("",
					gedcom.NewDateNode("17 Mar 1877"), // john
					gedcom.NewDateNode("1936"),        // harry
				),
				gedcom.NewUniqueIDNode("EE13561DDB204985BFFDEEBF82A5226C"),
				gedcom.NewNameNode("John /Smith/"),
			),
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
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      1.0,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("Apr 1907")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("Jane /DOE/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("Apr 1907")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      1.0,
		},

		// Exact matches, but missing information on both sides. These
		// specifically should NOT return 1.0 as it would throw out the real
		// similarities. See the docs for explanation.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewDeathNode("", gedcom.NewDateNode("Apr 1907")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewDeathNode("", gedcom.NewDateNode("Apr 1907")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.875,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("John /Smith/"),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.75,
		},

		// Similar matches but the same sized slice on both sides.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				),
				gedcom.NewDocument().AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBurialNode("", gedcom.NewDateNode("1927")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P4",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Abt. Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("1907")),
				),
				gedcom.NewDocument().AddIndividual("P5",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Bef. 1846")),
				),
				gedcom.NewDocument().AddIndividual("P6",
					gedcom.NewNameNode("Bob Thomas /Jones/"),
					gedcom.NewBurialNode("", gedcom.NewDateNode("1927")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.872532146404072,
		},

		// The slices are different lengths. The same score should be returned
		// when different sizes slices are swapped.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				),
				gedcom.NewDocument().AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBurialNode("", gedcom.NewDateNode("1927")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P4",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Between 1845 and 1846")),
				),
				gedcom.NewDocument().AddIndividual("P5",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Bef. 10 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("Abt. 1908")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.7754008744441251,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P4",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Between 1845 and 1846")),
				),
				gedcom.NewDocument().AddIndividual("P5",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Bef. 10 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("Abt. 1908")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				),
				gedcom.NewDocument().AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBurialNode("", gedcom.NewDateNode("1927")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.7754008744441251,
		},

		// Whenever one slice is empty the result will always be 0.5.
		{
			a: gedcom.IndividualNodes{},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				),
				gedcom.NewDocument().AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBurialNode("", gedcom.NewDateNode("1927")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				),
				gedcom.NewDocument().AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBurialNode("", gedcom.NewDateNode("1927")),
				),
			},
			b:             gedcom.IndividualNodes{},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},

		// These ones are just way off and should not be considered matches.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P4",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Between 1845 and 1846")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P5",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Bef. 10 Jan 1943")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("Abt. 2008")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("1627")),
				),
			},
			minSimilarity: gedcom.DefaultMinimumSimilarity,
			expected:      0.5,
		},

		// Different values for minimumSimilarity.
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("1627")),
				),
			},
			minSimilarity: 0.95,
			expected:      0.5,
		},
		{
			a: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P1",
					gedcom.NewNameNode("John /Smith/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("4 Jan 1843")),
					gedcom.NewDeathNode("", gedcom.NewDateNode("17 Mar 1907")),
				),
				gedcom.NewDocument().AddIndividual("P2",
					gedcom.NewNameNode("Jane /Doe/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("Sep 1845")),
				),
			},
			b: gedcom.IndividualNodes{
				gedcom.NewDocument().AddIndividual("P3",
					gedcom.NewNameNode("Bob /Jones/"),
					gedcom.NewBirthNode("", gedcom.NewDateNode("1627")),
				),
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
	simplifyErrors := cmp.Transformer("Errors", func(in error) string {
		if in == nil {
			return ""
		}

		return in.Error()
	})

	// IgnoreUnexported tell the diff engine to ignore unexported fields for the
	// following types.
	diff := cmp.Diff(expected, actual, cmpopts.IgnoreUnexported(
		gedcom.SimpleNode{},
		gedcom.IndividualNode{},
		gedcom.IndividualComparison{},
		gedcom.FamilyNode{},
		gedcom.DateNode{},
		gedcom.ChildNode{},
	), simplifyErrors)
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

	i1 := individual(gedcom.NewDocument(), "P1", "Elliot /Chance/", "", "")
	i2 := individual(gedcom.NewDocument(), "P2", "Joe /Bloggs/", "", "")

	Nodes(nil).Returns(nil)
	Nodes(gedcom.IndividualNodes{}).Returns(nil)
	Nodes(gedcom.IndividualNodes{i1, i2}).Returns(gedcom.Nodes{i1, i2})
}

func TestIndividualNodes_Merge(t *testing.T) {
	for testName, test := range individualNodesTests {
		t.Run(testName, func(t *testing.T) {
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
	i.AddNode(gedcom.NewUniqueIDNode(uid))

	return i
}
