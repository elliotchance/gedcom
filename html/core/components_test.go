package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestComponents_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Components")

	c(core.NewComponents()).
		Returns("")

	c(core.NewComponents(core.NewText("foo"))).
		Returns("foo")

	c(core.NewComponents(core.NewText("foo"), core.NewText("bar"))).
		Returns("foobar")
}
