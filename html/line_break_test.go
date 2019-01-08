package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestLineBreak_WriteTo(t *testing.T) {
	c := testComponent(t, "LineBreak")

	c(html.NewLineBreak()).Returns("<br/>")
}
