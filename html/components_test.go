package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestComponents_WriteTo(t *testing.T) {
	c := testComponent(t, "Components")

	c(html.NewComponents()).
		Returns("")

	c(html.NewComponents(html.NewText("foo"))).
		Returns("foo")

	c(html.NewComponents(html.NewText("foo"), html.NewText("bar"))).
		Returns("foobar")
}
