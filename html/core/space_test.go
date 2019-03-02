package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestSpace_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Space")

	c(core.NewSpace()).
		Returns("<div class=\"row\"><div class=\"col-12\">&nbsp;</div></div>")
}
