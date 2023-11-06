package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestHorizontalRuleRow_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "HorizontalRuleRow")

	c(core.NewHorizontalRuleRow()).Returns("<div class=\"row\"><div class=\"col-12\"><hr/></div></div>")
}
