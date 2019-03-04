package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestText_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Text")

	c(core.NewText("foo")).Returns(`foo`)

	c(core.NewText(`"Fran & Freddie's Diner" <tasty@example.com>`)).
		Returns(`&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;`)

	c(core.NewText("  foo")).Returns(`  foo`)

	c(core.NewText("&nbsp;&nbsp;foo")).Returns(`&nbsp;&nbsp;foo`)
}
