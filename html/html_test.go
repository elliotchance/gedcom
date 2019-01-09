package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestHTML_WriteTo(t *testing.T) {
	c := testComponent(t, "HTML")

	c(html.NewHTML("foo")).Returns(`foo`)
	c(html.NewHTML(`"Fran & Freddie's Diner" <tasty@example.com>`)).
		Returns(`"Fran & Freddie's Diner" <tasty@example.com>`)
}
