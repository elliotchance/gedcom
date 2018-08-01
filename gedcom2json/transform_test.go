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
		doc:      &gedcom.Document{},
		expected: []interface{}{},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.TagVersion, "5.5", "", nil),
			},
		},
		expected: []interface{}{
			map[string]interface{}{
				"tag": "VERS",
				"val": "5.5",
			},
		},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
				}),
			},
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
			options := gedcom.TransformOptions{}
			assert.Equal(t, gedcom.Transform(test.doc, options), test.expected)
		})
	}
}
