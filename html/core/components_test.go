package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
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
