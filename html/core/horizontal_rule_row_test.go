package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestHorizontalRuleRow_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "HorizontalRuleRow")

	c(core.NewHorizontalRuleRow()).Returns("<div class=\"row\"><div class=\"col-12\"><hr/></div></div>")
}
