package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestHeading_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Heading")

	c(core.NewHeading(1, "", core.NewText("hi"))).Returns("<h1>hi</h1>")

	c(core.NewHeading(3, "hi there", core.NewText("a&b"))).
		Returns("<h3 class=\"hi there\">a&amp;b</h3>")
}
