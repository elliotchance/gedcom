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
				assert.Equal(t, gedcom.NodeGedcom(test.expected), gedcom.NodeGedcom(actual))
			} else {
				assert.EqualError(t, err, test.error)
				assert.Nil(t, actual)
			}
		})
	}
}
