package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestHorizontalRuleRow_WriteTo(t *testing.T) {
	c := testComponent(t, "HorizontalRuleRow")

	c(html.NewHorizontalRuleRow()).Returns("<div class=\"row\"><div class=\"col-12\"><hr/></div></div>")
}
