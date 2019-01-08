package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestText_WriteTo(t *testing.T) {
	c := testComponent(t, "Text")

	c(html.NewText("foo")).Returns(`foo`)

	c(html.NewText(`"Fran & Freddie's Diner" <tasty@example.com>`)).
		Returns(`&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;`)

	c(html.NewText("  foo")).Returns(`  foo`)

	c(html.NewText("&nbsp;&nbsp;foo")).Returns(`&nbsp;&nbsp;foo`)
}
