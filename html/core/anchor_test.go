package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestAnchor_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Anchor")

	c(core.NewAnchor("foo")).Returns(`<a name="foo"/>`)
}
