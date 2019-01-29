package html_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestIndividualButton_WriteTo(t *testing.T) {
	doc := gedcom.NewDocument()
	elliot := individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")

	c := testComponent(t, "IndividualButton")

	c(html.NewIndividualButton(doc, elliot, html.LivingVisibilityPlaceholder)).
		Returns("<button class=\"btn btn-outline-info btn-block\" onclick=\"location.href='elliot-chance.html'\" type=\"button\"><strong>Elliot Chance</strong><br/><em>b.</em> 4 Jan 1843&nbsp;&nbsp;&nbsp;<em>d.</em> 17 Mar 1907&nbsp;</button>")
}
