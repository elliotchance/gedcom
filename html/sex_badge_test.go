package html_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestSexBadge_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "SexBadge")

	c(html.NewSexBadge(nil)).
		Returns("<span class=\"badge badge-info\">Unknown</span>")

	c(html.NewSexBadge(gedcom.NewSexNode(gedcom.SexUnknown))).
		Returns("<span class=\"badge badge-info\">Unknown</span>")

	c(html.NewSexBadge(gedcom.NewSexNode(gedcom.SexMale))).
		Returns("<span class=\"badge badge-primary\">Male</span>")

	c(html.NewSexBadge(gedcom.NewSexNode(gedcom.SexFemale))).
		Returns("<span class=\"badge badge-danger\">Female</span>")

	c(html.NewSexBadge(gedcom.NewSexNode(""))).
		Returns("<span class=\"badge badge-info\">Unknown</span>")

	c(html.NewSexBadge(gedcom.NewSexNode("Foo"))).
		Returns("<span class=\"badge badge-info\">Unknown</span>")
}
