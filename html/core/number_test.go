package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestNumber_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Number")

	c(core.NewNumber(0)).Returns("0")

	c(core.NewNumber(123)).Returns("123")

	c(core.NewNumber(4987342)).Returns("4,987,342")

	c(core.NewNumber(-28534)).Returns("-28,534")
}
