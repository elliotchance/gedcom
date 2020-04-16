package core

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

func (c *Page) WriteHTMLTo(w io.Writer) (int64, error) {
	googleAnalytics := NewGoogleAnalytics(c.googleAnalyticsID)
	footer := NewFooterRow()
	title := NewTag("title", nil, NewText(c.title))

	mutN := appendString(w, `<html><head><meta charset="UTF-8">`)
	mutN += appendComponent(w, googleAnalytics)
	mutN += appendComponent(w, title)
	mutN += appendString(w, `<link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" rel="stylesheet">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/octicons/4.4.0/font/octicons.css"/>
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js"></script>
	</head>
	<body>
		<div class="container">`)

	mutN += appendComponent(w, c.body)
	mutN += appendComponent(w, footer)
	mutN += appendString(w, `</div></body></html>`)

	return mutN, nil
}
