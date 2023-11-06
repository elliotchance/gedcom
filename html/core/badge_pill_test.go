package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestBadgePill_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "BadgePill")

	c(core.NewBadgePill("green", "myclass", core.NewText("123"))).
		Returns(`<span class="badge badge-pill badge-green myclass">123</span>`)
}
