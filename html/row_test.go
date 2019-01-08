package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestRow_WriteTo(t *testing.T) {
	c := testComponent(t, "Row")

	c(html.NewRow(html.NewColumn(12, html.NewText("hi")))).
		Returns("<div class=\"row\"><div class=\"col-12\">hi</div></div>")

	c(html.NewRow(
		html.NewColumn(6, html.NewText("hi")),
		html.NewColumn(6, html.NewText("there")),
	)).Returns("<div class=\"row\"><div class=\"col-6\">hi</div><div class=\"col-6\">there</div></div>")
}
