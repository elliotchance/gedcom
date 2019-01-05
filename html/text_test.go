package html_test

import (
	html "github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/tf"
	"testing"
)

func TestText_String(t *testing.T) {
	String := tf.Function(t, (*html.Text).String)

	String(html.NewText("foo")).Returns(`foo`)

	String(html.NewText(`"Fran & Freddie's Diner" <tasty@example.com>`)).
		Returns(`&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;`)

	String(html.NewText("  foo")).Returns(`  foo`)

	String(html.NewText("&nbsp;&nbsp;foo")).Returns(`&nbsp;&nbsp;foo`)
}
