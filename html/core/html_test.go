package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestHTML_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "HTML")

	c(core.NewHTML("foo")).Returns(`foo`)
	c(core.NewHTML(`"Fran & Freddie's Diner" <tasty@example.com>`)).
		Returns(`"Fran & Freddie's Diner" <tasty@example.com>`)
}
