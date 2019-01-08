package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestHorizontalRule_WriteTo(t *testing.T) {
	c := testComponent(t, "HorizontalRule")

	c(html.NewHorizontalRule()).Returns("<hr/>")
}
