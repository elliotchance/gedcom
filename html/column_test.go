package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestColumn_WriteTo(t *testing.T) {
	c := testComponent(t, "Column")

	c(html.NewColumn(12, html.NewText("foo"))).
		Returns("<div class=\"col-12\">foo</div>")

	c(html.NewColumn(6, html.NewText("bar"))).
		Returns("<div class=\"col-6\">bar</div>")
}
