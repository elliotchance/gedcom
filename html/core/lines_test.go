package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestLines_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Lines")

	c(core.NewLines()).Returns("")

	c(core.NewLines(
		core.NewText("foo"),
	)).Returns("foo")

	c(core.NewLines(
		core.NewText("foo"),
		core.NewText("bar"),
	)).Returns("foo<br/>bar")

	c(core.NewLines(
		core.NewText("foo"),
		core.NewText("bar"),
		core.NewText("baz"),
	)).Returns("foo<br/>bar<br/>baz")
}
