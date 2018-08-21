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
				gedcom.NewSimpleNode(nil, gedcom.TagVersion, "5.5", "", nil),
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
				gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
					gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{}),
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
			options := TransformOptions{}
			assert.Equal(t, Transform(test.doc, options), test.expected)
		})
	}
}
