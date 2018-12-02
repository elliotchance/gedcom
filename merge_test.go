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
					gedcom.NewDateNode(doc1, "3 Sep 1943", "", nil),
					gedcom.NewDateNode(doc2, "14 Apr 1947", "", nil),
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

func TestMergeDocuments(t *testing.T) {
	for testName, test := range map[string]struct {
		left, right, expected *gedcom.Document
		mergeFn               gedcom.MergeFunction
	}{
		"LeftNil": {
			right: gedcom.NewDocument(),
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: gedcom.NewDocument(),
		},
		"RightNil": {
			left: gedcom.NewDocument(),
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: gedcom.NewDocument(),
		},
		"BothNil": {
			mergeFn: func(left, right gedcom.Node) gedcom.Node {
				return nil
			},
			expected: gedcom.NewDocument(),
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
			expected: gedcom.NewDocumentWithNodes([]gedcom.Node{
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
			expected: gedcom.NewDocumentWithNodes([]gedcom.Node{
				gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
				gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
			}),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			actual := gedcom.MergeDocuments(test.left, test.right, test.mergeFn)

			assert.Equal(t, test.expected.String(), actual.String())
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
			similarity: 0.6253333333333334,
		},
	} {
		t.Run(testName, func(t *testing.T) {
			options := gedcom.NewSimilarityOptions()

			leftIndividual, leftOK := test.left.(*gedcom.IndividualNode)
			rightIndividual, rightOK := test.right.(*gedcom.IndividualNode)

			// Only check the similarity if it makes sense.
			if leftOK && rightOK {
				similarity := leftIndividual.SurroundingSimilarity(rightIndividual, options)

				assert.Equal(t, test.similarity, similarity.WeightedSimilarity(options))
			}

			mergerFn := gedcom.IndividualBySurroundingSimilarityMergeFunction(0.75, options)
			actual := mergerFn(test.left, test.right)

			assertNodeEqual(t, test.expected, actual)
		})
	}
}
