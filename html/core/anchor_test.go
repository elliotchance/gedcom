package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestAnchor_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Anchor")

	c(core.NewAnchor("foo")).Returns(`<a name="foo"/>`)
}
