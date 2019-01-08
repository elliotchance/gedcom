package html_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestEventDate_WriteTo(t *testing.T) {
	c := testComponent(t, "EventDate")

	c(html.NewEventDate("foo", gedcom.DateNodes{})).Returns(``)

	c(html.NewEventDate("foo", gedcom.DateNodes{
		gedcom.NewDateNode(nil, "3 Sep 1945", "", nil),
	})).Returns(`<em>foo</em> 3 Sep 1945`)

	c(html.NewEventDate("bar", gedcom.DateNodes{
		gedcom.NewDateNode(nil, "17 Sep 1945", "", nil),
		gedcom.NewDateNode(nil, "3 Sep 1945", "", nil),
	})).Returns(`<em>bar</em> 17 Sep 1945`)
}
