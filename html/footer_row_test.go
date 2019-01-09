package html_test

import (
	"bytes"
	"github.com/elliotchance/gedcom/html"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFooterRow_WriteTo(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	html.NewFooterRow().WriteTo(buf)

	component := string(buf.Bytes())

	assert.Contains(t, component, "<div class=\"row\">")
	assert.Contains(t, component,
		"Generated with <a href=\"https://github.com/elliotchance/gedcom\">github.com/elliotchance/gedcom</a>")
}
