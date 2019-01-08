package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestIndividualDates_WriteTo(t *testing.T) {
	c := testComponent(t, "IndividualDates")

	c(html.NewIndividualDates(elliot, true)).
		Returns("<em>b.</em> 4 Jan 1843&nbsp;&nbsp;&nbsp;<em>d.</em> 17 Mar 1907")
}
