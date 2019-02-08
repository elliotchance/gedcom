package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var transformTests = []struct {
	doc      func(*gedcom.Document)
	expected []interface{}
}{
	{
		doc:      func(document *gedcom.Document) {},
		expected: []interface{}{},
	},
	{
		doc: func(document *gedcom.Document) {
			document.AddNode(gedcom.NewNode(gedcom.TagVersion, "5.5", ""))
		},
		expected: []interface{}{
			map[string]interface{}{
				"tag": "VERS",
				"val": "5.5",
			},
		},
	},
	{
		doc: func(document *gedcom.Document) {
			document.AddIndividual("P1",
				gedcom.NewNameNode("Joe /Bloggs/"),
			)
		},
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
			doc := gedcom.NewDocument()
			test.doc(doc)
			assert.Equal(t, Transform(doc, options), test.expected)
		})
	}
}
