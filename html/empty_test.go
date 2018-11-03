package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/tf"
	"testing"
)

func TestEmpty_String(t *testing.T) {
	String := tf.Function(t, (*html.Empty).String)

	String(html.NewEmpty()).Returns(`&nbsp;`)
}
