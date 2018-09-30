package html

import (
	"testing"
	"github.com/elliotchance/tf"
)

func TestAnchor_String(t *testing.T) {
	String := tf.Function(t, (*Anchor).String)

	String(NewAnchor("foo")).Returns(`<a name="foo"/>`)
}
