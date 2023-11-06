package html_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html"
)

func TestEventDates_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "EventDates")

	c(html.NewEventDates([]*html.EventDate{})).Returns(``)

	c(html.NewEventDates([]*html.EventDate{
		html.NewEventDate("foo", gedcom.DateNodes{
			gedcom.NewDateNode("3 Sep 1945"),
		}),
	})).Returns("<em>foo</em> 3 Sep 1945")

	c(html.NewEventDates([]*html.EventDate{
		html.NewEventDate("foo", gedcom.DateNodes{
			gedcom.NewDateNode("3 Sep 1945"),
		}),
		html.NewEventDate("bar", gedcom.DateNodes{
			gedcom.NewDateNode("17 Sep 1945"),
			gedcom.NewDateNode("3 Sep 1945"),
		}),
	})).Returns("<em>foo</em> 3 Sep 1945&nbsp;&nbsp;&nbsp;<em>bar</em> 17 Sep 1945")
}
