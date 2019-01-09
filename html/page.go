// Package html is shared HTML rendering components that are shared by the other
// packages.
package html

import (
	"io"
)

// Page is the entire page wrapped that provides the HTML head and body.
type Page struct {
	title             string
	body              Component
	googleAnalyticsID string
}

func NewPage(title string, body Component, googleAnalyticsID string) *Page {
	return &Page{
		title:             title,
		body:              body,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *Page) WriteTo(w io.Writer) (int64, error) {
	googleAnalytics := NewGoogleAnalytics(c.googleAnalyticsID)
	footer := NewFooterRow()

	n := appendString(w, `<html><head><meta charset="UTF-8">`)
	n += appendComponent(w, googleAnalytics)
	n += appendComponent(w, NewTag("title", nil, NewText(c.title)))
	n += appendString(w, `<link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" rel="stylesheet">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/octicons/4.4.0/font/octicons.css"/>
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js"></script>
	</head>
	<body>
		<div class="container">`)

	n += appendComponent(w, c.body)
	n += appendComponent(w, footer)
	n += appendString(w, `</div></body></html>`)

	return n, nil
}
