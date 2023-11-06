package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestHorizontalRule_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "HorizontalRule")

	c(core.NewHorizontalRule()).Returns("<hr/>")
}
