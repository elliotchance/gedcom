package html_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/elliotchance/gedcom/v39/html/core"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testComponent(t *testing.T, name string) func(args ...interface{}) *tf.F {
	return tf.NamedFunction(t, name+"_WriteTo", func(c core.Component) string {
		buf := bytes.NewBuffer(nil)
		n, err := c.WriteHTMLTo(buf)
		assert.NoError(t, err)

		data := buf.Bytes()
		assert.Equal(t, int64(len(data)), n)

		return string(data)
	})
}

func assertTextByXPath(t *testing.T, body, query string, expected []string) {
	doc, err := htmlquery.Parse(strings.NewReader(body))
	require.NoError(t, err)

	nodes, err := htmlquery.QueryAll(doc, query)
	require.NoError(t, err)

	var nodeTexts []string
	for _, node := range nodes {
		nodeTexts = append(nodeTexts, node.Data)
	}

	assert.Equal(t, nodeTexts, expected)
}
