package html_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html"
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
			"individuals-c.html",
			"elliot-chance.html",
			"places.html",
			"families.html",
			"surnames.html",
			"sources.html",
			"statistics.html",
		},
	},
	"HideLettersForInvisible": {
		options: &html.PublishShowOptions{
			ShowIndividuals:  true,
			ShowPlaces:       true,
			ShowFamilies:     true,
			ShowSurnames:     true,
			ShowSources:      true,
			ShowStatistics:   true,
			LivingVisibility: html.LivingVisibilityHide,
		},
		files: []string{
			"places.html",
			"families.html",
			"surnames.html",
			"sources.html",
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
			p1.AddBirthDate("2019")
			doc.AddFamilyWithHusbandAndWife("F1", p1, nil)

			publisher := html.NewPublisher(doc, test.options)

			assert.True(t, p1.IsLiving())

			var files []string
			for file := range publisher.Files(1) {
				files = append(files, file.Name)
			}

			assert.Equal(t, test.files, files)
		})
	}
}
