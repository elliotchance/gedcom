package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestTableHead_WriteTo(t *testing.T) {
	c := testComponent(t, "TableHead")

	c(html.NewTableHead()).
		Returns("<thead><tr></tr></thead>")

	c(html.NewTableHead("a")).
		Returns("<thead><tr><th scope=\"col\">a</th></tr></thead>")

	c(html.NewTableHead("a", "b")).
		Returns("<thead><tr><th scope=\"col\">a</th><th scope=\"col\">b</th></tr></thead>")
}
