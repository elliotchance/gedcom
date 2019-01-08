package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestLines_WriteTo(t *testing.T) {
	c := testComponent(t, "Lines")

	c(html.NewLines()).Returns("")

	c(html.NewLines(
		html.NewText("foo"),
	)).Returns("foo")

	c(html.NewLines(
		html.NewText("foo"),
		html.NewText("bar"),
	)).Returns("foo<br/>bar")

	c(html.NewLines(
		html.NewText("foo"),
		html.NewText("bar"),
		html.NewText("baz"),
	)).Returns("foo<br/>bar<br/>baz")
}
