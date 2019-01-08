package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestBadgePill_WriteTo(t *testing.T) {
	c := testComponent(t, "BadgePill")

	c(html.NewBadgePill("green", "myclass", html.NewText("123"))).
		Returns(`<span class="badge badge-pill badge-green myclass">123</span>`)
}
