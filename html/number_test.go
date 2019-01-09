package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestNumber_WriteTo(t *testing.T) {
	c := testComponent(t, "Number")

	c(html.NewNumber(0)).Returns("0")

	c(html.NewNumber(123)).Returns("123")

	c(html.NewNumber(4987342)).Returns("4,987,342")

	c(html.NewNumber(-28534)).Returns("-28,534")
}
