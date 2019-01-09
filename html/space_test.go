package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestSpace_WriteTo(t *testing.T) {
	c := testComponent(t, "Space")

	c(html.NewSpace()).
		Returns("<div class=\"row\"><div class=\"col-12\">&nbsp;</div></div>")
}
