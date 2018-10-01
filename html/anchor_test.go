package html

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestAnchor_String(t *testing.T) {
	String := tf.Function(t, (*Anchor).String)

	String(NewAnchor("foo")).Returns(`<a name="foo"/>`)
}
