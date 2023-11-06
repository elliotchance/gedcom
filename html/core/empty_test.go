package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestEmpty_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Empty")

	c(core.NewEmpty()).Returns(`&nbsp;`)
}
