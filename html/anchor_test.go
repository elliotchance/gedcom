package html_test

import (
	"github.com/elliotchance/tf"
	"testing"
	"github.com/elliotchance/gedcom/html"
)

func TestAnchor_String(t *testing.T) {
	String := tf.Function(t, (*html.Anchor).String)

	String(html.NewAnchor("foo")).Returns(`<a name="foo"/>`)
}
