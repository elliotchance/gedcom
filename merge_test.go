package gedcom_test

import (
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
			left: gedcom.NewBirthNode(doc1, "", "P1", []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
			}),
			right: gedcom.NewBirthNode(doc2, "", "P2", []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			}),
			expected: gedcom.NewBirthNode(doc1, "", "P1", []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			}),
		},
		"2": {
			left: gedcom.NewBirthNode(doc1, "", "P1", []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
			}),
			right: gedcom.NewBirthNode(doc1, "", "P1", []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
			}),
			expected: gedcom.NewBirthNode(doc1, "", "P1", []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
			}),
		},
		"3": {
			left: gedcom.NewBirthNode(doc1, "", "P1", []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
			}),
			right: gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
			error: "cannot merge BIRT and DATE nodes",
		},
		"4": {
			left: gedcom.NewIndividualNode(doc1, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(doc1, "", "", []gedcom.Node{
					gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				}),
				gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(doc1, "Sydney, Australia", "", nil),
					gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				}),
			}),
			right: gedcom.NewIndividualNode(doc2, "", "P2", []gedcom.Node{
				gedcom.NewBirthNode(doc2, "", "", []gedcom.Node{
					gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
				}),
				gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(doc2, "Queensland, Australia", "", nil),
				}),
			}),
			expected: gedcom.NewIndividualNode(doc1, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(doc1, "", "", []gedcom.Node{
					gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
					gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				}),
				gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(doc1, "Sydney, Australia", "", nil),
					gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				}),
				gedcom.NewResidenceNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(doc2, "Queensland, Australia", "", nil),
				}),
			}),
		},
		"NilLeft": {
			right: gedcom.NewBirthNode(doc2, "", "P2", []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			}),
			error: "left is nil",
		},
		"NilRight": {
			left: gedcom.NewBirthNode(doc2, "", "P2", []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			}),
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
	doc1 := gedcom.NewDocument()
	doc2 := gedcom.NewDocument()

	for testName, test := range map[string]struct {
		left, right, expected []gedcom.Node
		mergeFn               gedcom.MergeFunction
	}{
		"1": {
			left:  []gedcom.Node{},
			right: []gedcom.Node{},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: []gedcom.Node{},
		},
		"2": {
			left: []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
			},
			right: []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			},
		},
		"3": {
			left: []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				gedcom.NewPlaceNode(doc1, "Sydney, Australia", "", nil),
			},
			right: []gedcom.Node{
				gedcom.NewPlaceNode(doc1, "Queensland, Australia", "", nil),
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				gedcom.NewPlaceNode(doc1, "Sydney, Australia", "", nil),
				gedcom.NewPlaceNode(doc1, "Queensland, Australia", "", nil),
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			},
		},
		"4": {
			left: []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
				gedcom.NewPlaceNode(doc1, "Sydney, Australia", "", nil),
			},
			right: []gedcom.Node{
				gedcom.NewPlaceNode(doc1, "Queensland, Australia", "", nil),
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				if left.Tag().Is(right.Tag()) {
					return left.ShallowCopy()
				}

				return nil
			},
			expected: []gedcom.Node{
				gedcom.NewDateNode(doc2, "3 Sep 1943", "", nil),
				gedcom.NewPlaceNode(doc1, "Sydney, Australia", "", nil),
			},
		},
		"5": {
			// This test is to make sure the nodes from the right are not merged
			// more than into either a left node or a node already placed from
			// the right.
			left: []gedcom.Node{
				gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
			},
			right: []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "15 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "16 Apr 1947", "", nil),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				// Always consider it a merge.
				return left.ShallowCopy()
			},
			expected: []gedcom.Node{
				// 14 Apr 1947 was merged into 3 Sep 1943. The rest would be
				// appended.
				gedcom.NewDateNode(doc2, "3 Sep 1943", "", nil),
				gedcom.NewDateNode(doc2, "15 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "16 Apr 1947", "", nil),
			},
		},
		"LeftNil": {
			right: []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "15 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "16 Apr 1947", "", nil),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "15 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "16 Apr 1947", "", nil),
			},
		},
		"RightNil": {
			left: []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "15 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "16 Apr 1947", "", nil),
			},
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: []gedcom.Node{
				gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "15 Apr 1947", "", nil),
				gedcom.NewDateNode(doc2, "16 Apr 1947", "", nil),
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
		left: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
		}),
		right: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewPlaceNode(nil, "Queensland, Australia", "", nil),
			gedcom.NewDateNode(nil, "14 Apr 1947", "", nil),
		}),
		mergeFn: func(left, right gedcom.Node) gedcom.Node {
			return nil
		},
		expectedDoc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			gedcom.NewPlaceNode(nil, "Queensland, Australia", "", nil),
			gedcom.NewDateNode(nil, "14 Apr 1947", "", nil),
		}),
		expectedDocAndIndividuals: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			gedcom.NewPlaceNode(nil, "Queensland, Australia", "", nil),
			gedcom.NewDateNode(nil, "14 Apr 1947", "", nil),
		}),
	},
	"Merge2": {
		left: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
		}),
		right: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewPlaceNode(nil, "Queensland, Australia", "", nil),
			gedcom.NewDateNode(nil, "14 Apr 1947", "", nil),
		}),
		mergeFn: func(left, right gedcom.Node) gedcom.Node {
			if left.Tag().Is(right.Tag()) {
				return left.ShallowCopy()
			}

			return nil
		},
		expectedDoc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
		}),
		expectedDocAndIndividuals: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
		}),
	},
	"Merge3": {
		left: gedcom.NewDocumentWithNodes([]gedcom.Node{
			elliot,
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			jane,
		}),
		right: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			jane,
			john,
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
		}),
		mergeFn: gedcom.EqualityMergeFunction,
		expectedDoc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			elliot,
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			jane,
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			john,
		}),
		minimumSimilarity: 0.1,
		expectedDocAndIndividuals: gedcom.NewDocumentWithNodes([]gedcom.Node{
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
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
		}),
	},
	"Merge4": {
		left: gedcom.NewDocumentWithNodes([]gedcom.Node{
			elliot,
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			jane,
		}),
		right: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			jane,
			john,
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
		}),
		mergeFn:           gedcom.EqualityMergeFunction,
		minimumSimilarity: 0.9,
		expectedDoc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			elliot,
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			jane,
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			john,
		}),
		expectedDocAndIndividuals: gedcom.NewDocumentWithNodes([]gedcom.Node{
			jane,
			elliot,
			john,
			gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
			gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
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
			left:  gedcom.NewPlaceNode(nil, "Queensland, Australia", "", nil),
			right: gedcom.NewIndividualNode(nil, "", "", nil),
		},
		"RightIsNotIndividual": {
			left:  gedcom.NewIndividualNode(nil, "", "", nil),
			right: gedcom.NewPlaceNode(nil, "Queensland, Australia", "", nil),
		},
		"BothAreNotIndividuals": {
			left:  gedcom.NewPlaceNode(nil, "Queensland, Australia", "", nil),
			right: gedcom.NewPlaceNode(nil, "Queensland, Australia", "", nil),
		},
		"MergeAlmostExact": {
			left: gedcom.NewIndividualNode(doc, "", "", []gedcom.Node{
				gedcom.NewNameNode(doc, "Elliot /Chance/", "", nil),
				gedcom.NewBirthNode(doc, "", "", []gedcom.Node{
					gedcom.NewDateNode(doc, "3 Sep 1943", "", nil),
				}),
			}),
			right: gedcom.NewIndividualNode(doc, "", "", []gedcom.Node{
				gedcom.NewNameNode(doc, "Elliot /Chance/", "", nil),
				gedcom.NewBirthNode(doc, "", "", []gedcom.Node{
					gedcom.NewDateNode(doc, "7 Sep 1943", "", nil),
				}),
			}),
			similarity: 0.8666640123954465,
			expected: gedcom.NewIndividualNode(doc, "", "", []gedcom.Node{
				gedcom.NewNameNode(doc, "Elliot /Chance/", "", nil),
				gedcom.NewBirthNode(doc, "", "", []gedcom.Node{
					gedcom.NewDateNode(doc, "3 Sep 1943", "", nil),
					gedcom.NewDateNode(doc, "7 Sep 1943", "", nil),
				}),
			}),
		},
		"MergeNotSimilar": {
			left: gedcom.NewIndividualNode(doc, "", "", []gedcom.Node{
				gedcom.NewNameNode(doc, "John /Smith/", "", nil),
			}),
			right: gedcom.NewIndividualNode(doc, "", "", []gedcom.Node{
				gedcom.NewNameNode(doc, "Jane /Doe/", "", nil),
			}),

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
	doc := gedcom.NewDocument()

	for testName, test := range mergeDocumentsTests {
		t.Run(testName, func(t *testing.T) {
			beforeLeft := test.left.String()
			beforeRight := test.right.String()

			if test.left != nil {
				for _, n := range test.left.Nodes() {
					n.SetDocument(doc)
				}
			}

			if test.right != nil {
				for _, n := range test.right.Nodes() {
					n.SetDocument(doc)
				}
			}

			options := gedcom.NewIndividualNodesCompareOptions()
			options.SimilarityOptions.MinimumWeightedSimilarity = test.minimumSimilarity

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
