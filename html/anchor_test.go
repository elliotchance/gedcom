package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestAnchor_WriteTo(t *testing.T) {
	c := testComponent(t, "Anchor")

	c(html.NewAnchor("foo")).Returns(`<a name="foo"/>`)
}
