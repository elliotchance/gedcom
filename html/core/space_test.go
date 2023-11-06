package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestSpace_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Space")

	c(core.NewSpace()).
		Returns("<div class=\"row\"><div class=\"col-12\">&nbsp;</div></div>")
}
