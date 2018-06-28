package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"testing"
	"github.com/stretchr/testify/assert"
)

var transformTests = []struct {
	doc      *gedcom.Document
	expected map[string]interface{}
}{
	{
		doc: &gedcom.Document{},
		expected: map[string]interface{}{
			"individuals": map[string]interface{}{},
			"other":       []interface{}{},
		},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(gedcom.Version, "5.5", "", nil),
			},
		},
		expected: map[string]interface{}{
			"individuals": map[string]interface{}{},
			"other": []interface{}{
				map[string]interface{}{
					"tag": "VERS",
					"val": "5.5",
				},
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
		expected: map[string]interface{}{
			"individuals": map[string]interface{}{
				"P1": map[string]interface{}{
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
			"other": []interface{}{},
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
