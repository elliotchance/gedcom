package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestHeading_WriteTo(t *testing.T) {
	c := testComponent(t, "Heading")

	c(html.NewHeading(1, "", html.NewText("hi"))).Returns("<h1>hi</h1>")

	c(html.NewHeading(3, "hi there", html.NewText("a&b"))).
		Returns("<h3 class=\"hi there\">a&amp;b</h3>")
}
