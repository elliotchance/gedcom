package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestBigTitle_WriteTo(t *testing.T) {
	c := testComponent(t, "BigTitle")

	c(html.NewBigTitle(1, html.NewText("foo"))).
		Returns("<div class=\"row\"><div class=\"col-12\"><h1 class=\"text-center\">foo</h1></div></div>")

	c(html.NewBigTitle(2, html.NewText("bar <"))).
		Returns("<div class=\"row\"><div class=\"col-12\"><h2 class=\"text-center\">bar &lt;</h2></div></div>")
}
