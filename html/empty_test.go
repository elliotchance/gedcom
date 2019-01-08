package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestEmpty_WriteTo(t *testing.T) {
	c := testComponent(t, "Empty")

	c(html.NewEmpty()).Returns(`&nbsp;`)
}
