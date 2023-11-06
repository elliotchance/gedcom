package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestLineBreak_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "LineBreak")

	c(core.NewLineBreak()).Returns("<br/>")
}
