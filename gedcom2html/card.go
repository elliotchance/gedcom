package main

import (
	"fmt"
	"strconv"
	"github.com/elliotchance/gedcom/html"
)

const (
	noBadgeCount = -1
)

// card is a simple box with a header and body section.
type card struct {
	title string
	body  fmt.Stringer
	count int
}

func newCard(title string, count int, body fmt.Stringer) *card {
	return &card{
		title: title,
		body:  body,
		count: count,
	}
}

func (c *card) String() string {
	var count fmt.Stringer = newEmpty()
	if c.count != noBadgeCount {
		count = newBadgePill("secondary", "float-right", strconv.Itoa(c.count))
	}

	return html.NewDiv("card", html.NewComponents(
		html.NewHeading(5, "card-header", c.title+count.String()),
		c.body,
	)).String()
}
