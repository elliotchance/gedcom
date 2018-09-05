// Package html is shared HTML rendering components that are shared by the other
// packages.
package html

import (
	"fmt"
)

// Page is the entire page wrapped that provides the HTML head and body.
type Page struct {
	title string
	body  fmt.Stringer
}

func NewPage(title string, body fmt.Stringer) *Page {
	return &Page{
		title: title,
		body:  body,
	}
}

func (c *Page) String() string {
	return fmt.Sprintf(`
    <html>
	<head>
		<title>%s</title>
		<link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css"
	rel="stylesheet"
	integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB"
	crossorigin="anonymous">

		<link rel="stylesheet"
	href="https://cdnjs.cloudflare.com/ajax/libs/octicons/4.4.0/font/octicons.css"/>
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js"
	integrity="sha384-smHYKdLADwkXOn1EmN1qk/HfnUcbVRZyYmZ4qpPea6sjB/pTJ0euyQp0Mk8ck+5T"
	crossorigin="anonymous"></script>
	</head>
	<body>
		<div class="container">
			%s
			%s
		</div>
	</body>
	</html>
	`, c.title, c.body, NewFooter())
}
