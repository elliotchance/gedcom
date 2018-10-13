package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var transformTests = []struct {
	doc      *gedcom.Document
	expected []interface{}
}{
	{
		doc:      gedcom.NewDocument(),
		expected: []interface{}{},
	},
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagVersion, "5.5", ""),
		}),
		expected: []interface{}{
			map[string]interface{}{
				"tag": "VERS",
				"val": "5.5",
			},
		},
	},
	{
		doc: gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		}),
		expected: []interface{}{
			map[string]interface{}{
				"tag": "INDI",
				"ptr": "P1",
				"nodes": []interface{}{
					map[string]interface{}{
						"tag": "NAME",
						"val": "Joe /Bloggs/",
					},
				},
			},
		},
	},
}

func TestTransform(t *testing.T) {
	for _, test := range transformTests {
		t.Run("", func(t *testing.T) {
			options := TransformOptions{}
			assert.Equal(t, Transform(test.doc, options), test.expected)
		})
	}
}
