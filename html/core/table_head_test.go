package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestTableHead_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "TableHead")

	c(core.NewTableHead()).
		Returns("<thead><tr></tr></thead>")

	c(core.NewTableHead("a")).
		Returns("<thead><tr><th scope=\"col\">a</th></tr></thead>")

	c(core.NewTableHead("a", "b")).
		Returns("<thead><tr><th scope=\"col\">a</th><th scope=\"col\">b</th></tr></thead>")
}
