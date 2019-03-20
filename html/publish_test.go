package html_test

import (
	"testing"

	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

var publisherTests = map[string]struct {
	options *html.PublishShowOptions
	files   []string
}{
	"Empty": {
		options: &html.PublishShowOptions{
			LivingVisibility: html.LivingVisibilityShow,
		},
	},
	"All": {
		options: &html.PublishShowOptions{
			ShowIndividuals:  true,
			ShowPlaces:       true,
			ShowFamilies:     true,
			ShowSurnames:     true,
			ShowSources:      true,
			ShowStatistics:   true,
			LivingVisibility: html.LivingVisibilityShow,
		},
		files: []string{
			"individuals-e.html",
			"elliot-chance.html",
			"places.html",
			"families.html",
			"surnames.html",
			"statistics.html",
		},
	},
}

func TestPublisher_Files(t *testing.T) {
	for testName, test := range publisherTests {
		t.Run(testName, func(t *testing.T) {
			doc := gedcom.NewDocument()
			p1 := doc.AddIndividual("P1")
			p1.AddName("Elliot /Chance/")
			doc.AddFamilyWithHusbandAndWife("F1", p1, nil)

			publisher := html.NewPublisher(doc, test.options)

			var files []string
			for file := range publisher.Files(1) {
				files = append(files, file.Name)
			}

			assert.Equal(t, test.files, files)
		})
	}
}
