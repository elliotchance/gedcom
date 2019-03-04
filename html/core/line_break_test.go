package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestLineBreak_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "LineBreak")

	c(core.NewLineBreak()).Returns("<br/>")
}
