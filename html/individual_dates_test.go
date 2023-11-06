package html_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html"
)

func TestIndividualDates_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "IndividualDates")

	doc := gedcom.NewDocument()
	elliot := individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")

	c(html.NewIndividualDates(elliot, html.LivingVisibilityPlaceholder)).
		Returns("<em>b.</em> 4 Jan 1843&nbsp;&nbsp;&nbsp;<em>d.</em> 17 Mar 1907")
}
