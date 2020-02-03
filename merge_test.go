package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

func TestMergeNodes(t *testing.T) {
	doc1 := gedcom.NewDocument()
	doc2 := gedcom.NewDocument()

	for testName, test := range map[string]struct {
		left, right, expected gedcom.Node
		error                 string
	}{
		"1": {
			left: gedcom.NewNode(tag.TagBirth, "", "P1",
				gedcom.NewDateNode("3 Sep 1943"),
			),
			right: gedcom.NewNode(tag.TagBirth, "", "P2",
				gedcom.NewDateNode("14 Apr 1947"),
			),
			expected: gedcom.NewNode(tag.TagBirth, "", "P1",
				gedcom.NewDateNode("3 Sep 1943"),
				gedcom.NewDateNode("14 Apr 1947"),
			),
		},
		"2": {
			left: gedcom.NewNode(tag.TagBirth, "", "P1",
				gedcom.NewDateNode("3 Sep 1943"),
			),
			right: gedcom.NewNode(tag.TagBirth, "", "P1",
				gedcom.NewDateNode("3 Sep 1943"),
			),
			expected: gedcom.NewNode(tag.TagBirth, "", "P1",
				gedcom.NewDateNode("3 Sep 1943"),
			),
		},
		"3": {
			left: gedcom.NewNode(tag.TagBirth, "", "P1",
				gedcom.NewDateNode("3 Sep 1943"),
			),
			right: gedcom.NewDateNode("3 Sep 1943"),
			error: "cannot merge BIRT and DATE nodes",
		},
		"4": {
			left: doc1.AddIndividual("P1",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("3 Sep 1943"),
				),
				gedcom.NewResidenceNode("",
					gedcom.NewPlaceNode("Sydney, Australia"),
					gedcom.NewDateNode("3 Sep 1943"),
				),
			),
			right: doc2.AddIndividual("P2",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("14 Apr 1947"),
				),
				gedcom.NewResidenceNode("",
					gedcom.NewPlaceNode("Queensland, Australia"),
				),
			),
			expected: doc1.AddIndividual("P1",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("14 Apr 1947"),
					gedcom.NewDateNode("3 Sep 1943"),
				),
				gedcom.NewResidenceNode("",
					gedcom.NewPlaceNode("Sydney, Australia"),
					gedcom.NewDateNode("3 Sep 1943"),
				),
				gedcom.NewResidenceNode("",
					gedcom.NewPlaceNode("Queensland, Australia"),
				),
			),
		},
		"NilLeft": {
			right: gedcom.NewNode(tag.TagBirth, "", "P2",
				gedcom.NewDateNode("14 Apr 1947"),
			),
			error: "left is nil",
		},
		"NilRight": {
			left: gedcom.NewNode(tag.TagBirth, "", "P2",
				gedcom.NewDateNode("14 Apr 1947"),
			),
			error: "right is nil",
		},
		"NilBoth": {
			error: "left is nil",
		},
		"SameIndividuals": {
			left:     elliot,
			right:    elliot,
			expected: elliot,
		},
		"EqualNodes": {
			left: gedcom.NewNode(tag.TagBirth, "", "P1",
				gedcom.NewUniqueIDNode("EE13561DDB204985BFFDEEBF82A5226C"),
			),
			right: gedcom.NewNode(tag.TagBirth, "", "P2",
				gedcom.NewUniqueIDNode("EE13561DDB204985BFFDEEBF82A5226C5B2E"),
			),
			expected: gedcom.NewNode(tag.TagBirth, "", "P1",
				gedcom.NewUniqueIDNode("EE13561DDB204985BFFDEEBF82A5226C"),
			),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			actual, err := gedcom.MergeNodes(test.left, test.right)

			if test.error == "" {
				assert.NoError(t, err)
				assert.Equal(t, gedcom.GEDCOMString(test.expected, 0),
					gedcom.GEDCOMString(actual, 0))
			} else {
				assert.EqualError(t, err, test.error)
				assert.Nil(t, actual)
			}
		})
	}
}

func TestMergeNodeSlices(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right, expected gedcom.Nodes
		mergeFn               gedcom.MergeFunction
	}{
		"1": {
			left:  gedcom.Nodes{},
			right: gedcom.Nodes{},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: gedcom.Nodes{},
		},
		"2": {
			left: gedcom.Nodes{
				gedcom.NewDateNode("3 Sep 1943"),
			},
			right: gedcom.Nodes{
				gedcom.NewDateNode("14 Apr 1947"),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: gedcom.Nodes{
				gedcom.NewDateNode("3 Sep 1943"),
				gedcom.NewDateNode("14 Apr 1947"),
			},
		},
		"3": {
			left: gedcom.Nodes{
				gedcom.NewDateNode("3 Sep 1943"),
				gedcom.NewPlaceNode("Sydney, Australia"),
			},
			right: gedcom.Nodes{
				gedcom.NewPlaceNode("Queensland, Australia"),
				gedcom.NewDateNode("14 Apr 1947"),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: gedcom.Nodes{
				gedcom.NewDateNode("3 Sep 1943"),
				gedcom.NewPlaceNode("Sydney, Australia"),
				gedcom.NewPlaceNode("Queensland, Australia"),
				gedcom.NewDateNode("14 Apr 1947"),
			},
		},
		"4": {
			left: gedcom.Nodes{
				gedcom.NewDateNode("3 Sep 1943"),
				gedcom.NewPlaceNode("Sydney, Australia"),
			},
			right: gedcom.Nodes{
				gedcom.NewPlaceNode("Queensland, Australia"),
				gedcom.NewDateNode("14 Apr 1947"),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				if left.Tag().Is(right.Tag()) {
					return left.ShallowCopy()
				}

				return nil
			},
			expected: gedcom.Nodes{
				gedcom.NewDateNode("3 Sep 1943"),
				gedcom.NewPlaceNode("Sydney, Australia"),
			},
		},
		"5": {
			// This test is to make sure the nodes from the right are not merged
			// more than into either a left node or a node already placed from
			// the right.
			left: gedcom.Nodes{
				gedcom.NewDateNode("3 Sep 1943"),
			},
			right: gedcom.Nodes{
				gedcom.NewDateNode("14 Apr 1947"),
				gedcom.NewDateNode("15 Apr 1947"),
				gedcom.NewDateNode("16 Apr 1947"),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				// Always consider it a merge.
				return left.ShallowCopy()
			},
			expected: gedcom.Nodes{
				// 14 Apr 1947 was merged into 3 Sep 1943. The rest would be
				// appended.
				gedcom.NewDateNode("3 Sep 1943"),
				gedcom.NewDateNode("15 Apr 1947"),
				gedcom.NewDateNode("16 Apr 1947"),
			},
		},
		"LeftNil": {
			right: gedcom.Nodes{
				gedcom.NewDateNode("14 Apr 1947"),
				gedcom.NewDateNode("15 Apr 1947"),
				gedcom.NewDateNode("16 Apr 1947"),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: gedcom.Nodes{
				gedcom.NewDateNode("14 Apr 1947"),
				gedcom.NewDateNode("15 Apr 1947"),
				gedcom.NewDateNode("16 Apr 1947"),
			},
		},
		"RightNil": {
			left: gedcom.Nodes{
				gedcom.NewDateNode("14 Apr 1947"),
				gedcom.NewDateNode("15 Apr 1947"),
				gedcom.NewDateNode("16 Apr 1947"),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: gedcom.Nodes{
				gedcom.NewDateNode("14 Apr 1947"),
				gedcom.NewDateNode("15 Apr 1947"),
				gedcom.NewDateNode("16 Apr 1947"),
			},
		},
		"BothNil": {
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			actual := gedcom.MergeNodeSlices(test.left, test.right, test.mergeFn)

			actualDoc := gedcom.NewDocumentWithNodes(actual)
			expectedDoc := gedcom.NewDocumentWithNodes(test.expected)

			assert.Equal(t, expectedDoc.String(), actualDoc.String())
		})
	}
}

var mergeDocumentsTests = map[string]struct {
	left, right               *gedcom.Document
	mergeFn                   gedcom.MergeFunction
	expectedDoc               *gedcom.Document
	minimumSimilarity         float64
	expectedDocAndIndividuals *gedcom.Document
}{
	"LeftNil": {
		right: gedcom.NewDocument(),
		mergeFn: func(left, right gedcom.Node) gedcom.Node {
			return nil
		},
		expectedDoc:               gedcom.NewDocument(),
		expectedDocAndIndividuals: gedcom.NewDocument(),
	},
	"RightNil": {
		left: gedcom.NewDocument(),
		mergeFn: func(left, right gedcom.Node) gedcom.Node {
			return nil
		},
		expectedDoc:               gedcom.NewDocument(),
		expectedDocAndIndividuals: gedcom.NewDocument(),
	},
	"BothNil": {
		mergeFn: func(left, right gedcom.Node) gedcom.Node {
			return nil
		},
		expectedDoc:               gedcom.NewDocument(),
		expectedDocAndIndividuals: gedcom.NewDocument(),
	},
	"Merge1": {
		left: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
		}),
		right: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewPlaceNode("Queensland, Australia"),
			gedcom.NewDateNode("14 Apr 1947"),
		}),
		mergeFn: func(left, right gedcom.Node) gedcom.Node {
			return nil
		},
		expectedDoc: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
			gedcom.NewPlaceNode("Queensland, Australia"),
			gedcom.NewDateNode("14 Apr 1947"),
		}),
		expectedDocAndIndividuals: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
			gedcom.NewPlaceNode("Queensland, Australia"),
			gedcom.NewDateNode("14 Apr 1947"),
		}),
	},
	"Merge2": {
		left: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
		}),
		right: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewPlaceNode("Queensland, Australia"),
			gedcom.NewDateNode("14 Apr 1947"),
		}),
		mergeFn: func(left, right gedcom.Node) gedcom.Node {
			if left.Tag().Is(right.Tag()) {
				return left.ShallowCopy()
			}

			return nil
		},
		expectedDoc: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
		}),
		expectedDocAndIndividuals: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
		}),
	},
	"Merge3": {
		left: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			elliot,
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
			jane,
		}),
		right: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewPlaceNode("Sydney, Australia"),
			jane,
			john,
			gedcom.NewDateNode("3 Sep 1943"),
		}),
		mergeFn: gedcom.EqualityMergeFunction,
		expectedDoc: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			elliot,
			gedcom.NewDateNode("3 Sep 1943"),
			jane,
			gedcom.NewPlaceNode("Sydney, Australia"),
			john,
		}),
		minimumSimilarity: 0.1,
		expectedDocAndIndividuals: gedcom.NewDocumentWithNodes(gedcom.Nodes{
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
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
		}),
	},
	"Merge4": {
		left: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			elliot,
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
			jane,
		}),
		right: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewPlaceNode("Sydney, Australia"),
			jane,
			john,
			gedcom.NewDateNode("3 Sep 1943"),
		}),
		mergeFn:           gedcom.EqualityMergeFunction,
		minimumSimilarity: 0.9,
		expectedDoc: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			elliot,
			gedcom.NewDateNode("3 Sep 1943"),
			jane,
			gedcom.NewPlaceNode("Sydney, Australia"),
			john,
		}),
		expectedDocAndIndividuals: gedcom.NewDocumentWithNodes(gedcom.Nodes{
			jane,
			elliot,
			john,
			gedcom.NewDateNode("3 Sep 1943"),
			gedcom.NewPlaceNode("Sydney, Australia"),
		}),
	},
}

func TestMergeDocuments(t *testing.T) {
	for testName, test := range mergeDocumentsTests {
		t.Run(testName, func(t *testing.T) {
			beforeLeft := test.left.String()
			beforeRight := test.right.String()

			actual := gedcom.MergeDocuments(test.left, test.right, test.mergeFn)

			// Make sure the original documents were not modified.
			assert.Equal(t, beforeLeft, test.left.String())
			assert.Equal(t, beforeRight, test.right.String())

			assert.Equal(t, test.expectedDoc.String(), actual.String())
		})
	}
}

func TestIndividualBySurroundingSimilarityMergeFunction(t *testing.T) {
	doc := gedcom.NewDocument()

	for testName, test := range map[string]struct {
		left, right, expected gedcom.Node

		// This is just for reference to make sure the thresholds are correct.
		// If the algorithm for similarity changes these values can be safely
		// updated for the cases.
		similarity float64
	}{
		"LeftIsNotIndividual": {
			left:  gedcom.NewPlaceNode("Queensland, Australia"),
			right: gedcom.NewDocument().AddIndividual(""),
		},
		"RightIsNotIndividual": {
			left:  gedcom.NewDocument().AddIndividual(""),
			right: gedcom.NewPlaceNode("Queensland, Australia"),
		},
		"BothAreNotIndividuals": {
			left:  gedcom.NewPlaceNode("Queensland, Australia"),
			right: gedcom.NewPlaceNode("Queensland, Australia"),
		},
		"MergeAlmostExact": {
			left: doc.AddIndividual("",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("3 Sep 1943"),
				),
			),
			right: doc.AddIndividual("",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("7 Sep 1943"),
				),
			),
			similarity: 0.8666640123954465,
			expected: doc.AddIndividual("",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("3 Sep 1943"),
					gedcom.NewDateNode("7 Sep 1943"),
				),
			),
		},
		"MergeNotSimilar": {
			left: doc.AddIndividual("",
				gedcom.NewNameNode("John /Smith/"),
			),
			right: doc.AddIndividual("",
				gedcom.NewNameNode("Jane /Doe/"),
			),

			// It would be 0.6253333333333334. However, because its below the
			// minimum threshold it doesn't bother calculating the real value.
			similarity: 0.0,
		},
	} {
		t.Run(testName, func(t *testing.T) {
			options := gedcom.NewSimilarityOptions()

			leftIndividual, leftOK := test.left.(*gedcom.IndividualNode)
			rightIndividual, rightOK := test.right.(*gedcom.IndividualNode)

			// Only check the similarity if it makes sense.
			if leftOK && rightOK {
				similarity := leftIndividual.SurroundingSimilarity(rightIndividual, options, false)

				assert.Equal(t, test.similarity, similarity.WeightedSimilarity())
			}

			mergerFn := gedcom.IndividualBySurroundingSimilarityMergeFunction(0.75, options)
			actual := mergerFn(test.left, test.right)

			assertNodeEqual(t, test.expected, actual)
		})
	}
}

func TestMergeDocumentsAndIndividuals(t *testing.T) {
	for testName, test := range mergeDocumentsTests {
		t.Run(testName, func(t *testing.T) {
			beforeLeft := test.left.String()
			beforeRight := test.right.String()

			options := gedcom.NewIndividualNodesCompareOptions()
			options.SimilarityOptions.MinimumWeightedSimilarity = test.minimumSimilarity
			options.SimilarityOptions.PreferPointerAbove = 1.0

			actual, err := gedcom.MergeDocumentsAndIndividuals(
				test.left, test.right, test.mergeFn, options)

			// Make sure the original documents were not modified.
			assert.Equal(t, beforeLeft, test.left.String())
			assert.Equal(t, beforeRight, test.right.String())

			assert.NoError(t, err)
			assert.Equal(t, test.expectedDocAndIndividuals.String(), actual.String())
		})
	}
}
