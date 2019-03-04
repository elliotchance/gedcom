package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestHorizontalRule_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "HorizontalRule")

	c(core.NewHorizontalRule()).Returns("<hr/>")
}
