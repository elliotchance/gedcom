package core_test

import (
	"bytes"
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
	"github.com/stretchr/testify/assert"
)

func TestFooterRow_WriteHTMLTo(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	core.NewFooterRow().WriteHTMLTo(buf)

	component := string(buf.Bytes())

	assert.Contains(t, component, "<div class=\"row\">")
	assert.Contains(t, component,
		"Generated with <a href=\"https://github.com/elliotchance/gedcom\">github.com/elliotchance/gedcom</a>")
}
