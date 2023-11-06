package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestHTML_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "HTML")

	c(core.NewHTML("foo")).Returns(`foo`)
	c(core.NewHTML(`"Fran & Freddie's Diner" <tasty@example.com>`)).
		Returns(`"Fran & Freddie's Diner" <tasty@example.com>`)
}
