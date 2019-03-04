package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestBigTitle_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "BigTitle")

	c(core.NewBigTitle(1, core.NewText("foo"))).
		Returns("<div class=\"row\"><div class=\"col-12\"><h1 class=\"text-center\">foo</h1></div></div>")

	c(core.NewBigTitle(2, core.NewText("bar <"))).
		Returns("<div class=\"row\"><div class=\"col-12\"><h2 class=\"text-center\">bar &lt;</h2></div></div>")
}
