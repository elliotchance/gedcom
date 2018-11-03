package html

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestText_String(t *testing.T) {
	String := tf.Function(t, (*Text).String)

	String(NewText("foo")).Returns(`foo`)
	String(NewText(`"Fran & Freddie's Diner" <tasty@example.com>`)).
		Returns(`&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;`)
}
