package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"github.com/elliotchance/gedcom/util"
	"sort"
)

type PublishShowOptions struct {
	ShowIndividuals  bool
	ShowPlaces       bool
	ShowFamilies     bool
	ShowSurnames     bool
	ShowSources      bool
	ShowStatistics   bool
	LivingVisibility LivingVisibility
}

type Publisher struct {
	doc               *gedcom.Document
	options           *PublishShowOptions
	fileWriter        core.FileWriter
	GoogleAnalyticsID string
	indexLetters      []rune
	individuals       map[string]*gedcom.IndividualNode
}

// NewPublisher generates the pages to be rendered for a published website.
//
// The fileWriter will be responsible for rendering the pages to their
// destination. Which may be a file system, or somewhere else of your choosing.
// If you only wish to generate files you should use a DirectoryFileWriter.
func NewPublisher(doc *gedcom.Document, options *PublishShowOptions) *Publisher {
	return &Publisher{
		doc:          doc,
		options:      options,
		indexLetters: GetIndexLetters(doc, options.LivingVisibility),
		individuals:  GetIndividuals(doc),
	}
}

func (publisher *Publisher) Publish(fileWriter core.FileWriter, parallel int) (err error) {
	files := publisher.Files(parallel)
	util.WorkerPool(parallel, func(_ int) {
		for file := range files {
			fileErr := fileWriter.WriteFile(file)
			if fileErr != nil {
				err = fileErr
				break
			}
		}
	})

	return
}

func (publisher *Publisher) Files(channelSize int) chan *core.File {
	files := make(chan *core.File, channelSize)

	go func() {
		publisher.sendFiles(files)
		close(files)
	}()

	return files
}

func (publisher *Publisher) sendIndividualFiles(files chan *core.File) {
	if publisher.options.ShowIndividuals {
		for _, letter := range publisher.indexLetters {
			files <- core.NewFile(
				PageIndividuals(letter),
				NewIndividualListPage(publisher.doc, letter,
					publisher.GoogleAnalyticsID, publisher.options,
					publisher.indexLetters),
			)
		}

		for _, individual := range publisher.individuals {
			if individual.IsLiving() {
				switch publisher.options.LivingVisibility {
				case LivingVisibilityHide,
					LivingVisibilityPlaceholder:
					continue

				case LivingVisibilityShow:
					// Proceed.
				}
			}

			page := NewIndividualPage(publisher.doc, individual, publisher.GoogleAnalyticsID, publisher.options, publisher.indexLetters)
			pageName := PageIndividual(publisher.doc, individual, publisher.options.LivingVisibility)
			files <- core.NewFile(pageName, page)
		}
	}
}

func (publisher *Publisher) sendPlaceFiles(files chan *core.File) {
	if publisher.options.ShowPlaces {
		page := NewPlaceListPage(publisher.doc, publisher.GoogleAnalyticsID,
			publisher.options, publisher.indexLetters)
		files <- core.NewFile(PagePlaces(), page)

		// Sort the places so that the generated page names will be more
		// deterministic.
		places := GetPlaces(publisher.doc)
		placeKeys := []string{}

		for key := range places {
			placeKeys = append(placeKeys, key)
		}

		sort.Strings(placeKeys)

		for _, key := range placeKeys {
			place := places[key]
			page := NewPlacePage(publisher.doc, key, publisher.GoogleAnalyticsID, publisher.options, publisher.indexLetters)
			files <- core.NewFile(PagePlace(publisher.doc, place.PrettyName), page)
		}
	}
}

func (publisher *Publisher) sendFamilyFiles(files chan *core.File) {
	if publisher.options.ShowFamilies {
		files <- core.NewFile(
			PageFamilies(),
			NewFamilyListPage(publisher.doc, publisher.GoogleAnalyticsID, publisher.options, publisher.indexLetters),
		)
	}
}

func (publisher *Publisher) sendSurnameFiles(files chan *core.File) {
	if publisher.options.ShowSurnames {
		files <- core.NewFile(
			PageSurnames(),
			NewSurnameListPage(publisher.doc, publisher.GoogleAnalyticsID,
				publisher.options, publisher.indexLetters))
	}
}

func (publisher *Publisher) sendSourceFiles(files chan *core.File) {
	if publisher.options.ShowSources {
		files <- core.NewFile(PageSources(),
			NewSourceListPage(publisher.doc, publisher.GoogleAnalyticsID,
				publisher.options, publisher.indexLetters))

		for _, source := range publisher.doc.Sources() {
			page := NewSourcePage(publisher.doc, source,
				publisher.GoogleAnalyticsID, publisher.options,
				publisher.indexLetters)
			files <- core.NewFile(PageSource(source), page)
		}
	}
}

func (publisher *Publisher) sendStatisticsFiles(files chan *core.File) {
	if publisher.options.ShowStatistics {
		files <- core.NewFile(PageStatistics(),
			NewStatisticsPage(publisher.doc, publisher.GoogleAnalyticsID,
				publisher.options, publisher.indexLetters))
	}
}

func (publisher *Publisher) sendFiles(files chan *core.File) {
	publisher.sendIndividualFiles(files)
	publisher.sendPlaceFiles(files)
	publisher.sendFamilyFiles(files)
	publisher.sendSurnameFiles(files)
	publisher.sendSourceFiles(files)
	publisher.sendStatisticsFiles(files)
}
