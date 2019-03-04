package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestEmpty_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Empty")

	c(core.NewEmpty()).Returns(`&nbsp;`)
}
