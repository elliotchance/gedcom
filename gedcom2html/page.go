package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

// page is the entire page wrapped that provides the HTML head and body.
type page struct {
	title string
	body  fmt.Stringer
}

func newPage(title string, body fmt.Stringer) *page {
	return &page{
		title: title,
		body:  body,
	}
}

func (c *page) String() string {
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
	`, c.title, c.body, newFooter())
}

func pageIndividuals() string {
	return "individuals.html"
}

func pageIndividual(document *gedcom.Document, individual *gedcom.IndividualNode) string {
	individuals := getIndividuals(document)

	for key, value := range individuals {
		if value.Is(individual) {
			return fmt.Sprintf("%s.html", key)
		}
	}

	return "#"
}

func pagePlaces() string {
	return "places.html"
}

func pagePlace(document *gedcom.Document, place string) string {
	places := getPlaces(document)

	for key, value := range places {
		if value.prettyName == place {
			return fmt.Sprintf("%s.html", key)
		}
	}

	return "#"
}

func pageFamilies() string {
	return "families.html"
}
