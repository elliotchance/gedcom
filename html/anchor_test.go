package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/tf"
	"testing"
)

func TestAnchor_String(t *testing.T) {
	String := tf.Function(t, (*html.Anchor).String)

	String(html.NewAnchor("foo")).Returns(`<a name="foo"/>`)
}
