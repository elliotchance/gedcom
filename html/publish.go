package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"github.com/elliotchance/gedcom/util"
	"sort"
	"strings"
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
	placesMap         map[string]*place
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

		// placesMap can be nil because we handle found the places yet.
		individuals: GetIndividuals(doc, nil),
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
					publisher.indexLetters, publisher.placesMap),
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

			page := NewIndividualPage(publisher.doc, individual,
				publisher.GoogleAnalyticsID, publisher.options,
				publisher.indexLetters, publisher.placesMap)
			pageName := PageIndividual(publisher.doc, individual,
				publisher.options.LivingVisibility, publisher.placesMap)
			files <- core.NewFile(pageName, page)
		}
	}
}

func (publisher *Publisher) sendPlaceFiles(files chan *core.File) {
	if publisher.options.ShowPlaces {
		// Sort the places so that the generated page names will be more
		// deterministic.
		places := publisher.Places()

		page := NewPlaceListPage(publisher.doc, publisher.GoogleAnalyticsID,
			publisher.options, publisher.indexLetters, places)
		files <- core.NewFile(PagePlaces(), page)

		var placeKeys []string

		for key := range places {
			placeKeys = append(placeKeys, key)
		}

		sort.Strings(placeKeys)

		for _, key := range placeKeys {
			place := places[key]
			page := NewPlacePage(publisher.doc, key,
				publisher.GoogleAnalyticsID, publisher.options,
				publisher.indexLetters, publisher.placesMap)
			files <- core.NewFile(
				PagePlace(place.PrettyName, publisher.placesMap), page)
		}
	}
}

func (publisher *Publisher) sendFamilyFiles(files chan *core.File) {
	if publisher.options.ShowFamilies {
		files <- core.NewFile(
			PageFamilies(),
			NewFamilyListPage(publisher.doc, publisher.GoogleAnalyticsID,
				publisher.options, publisher.indexLetters, publisher.placesMap),
		)
	}
}

func (publisher *Publisher) sendSurnameFiles(files chan *core.File) {
	if publisher.options.ShowSurnames {
		files <- core.NewFile(
			PageSurnames(),
			NewSurnameListPage(publisher.doc, publisher.GoogleAnalyticsID,
				publisher.options, publisher.indexLetters, publisher.placesMap))
	}
}

func (publisher *Publisher) sendSourceFiles(files chan *core.File) {
	if publisher.options.ShowSources {
		files <- core.NewFile(PageSources(),
			NewSourceListPage(publisher.doc, publisher.GoogleAnalyticsID,
				publisher.options, publisher.indexLetters, publisher.placesMap))

		for _, source := range publisher.doc.Sources() {
			page := NewSourcePage(publisher.doc, source,
				publisher.GoogleAnalyticsID, publisher.options,
				publisher.indexLetters, publisher.placesMap)
			files <- core.NewFile(PageSource(source), page)
		}
	}
}

func (publisher *Publisher) sendStatisticsFiles(files chan *core.File) {
	if publisher.options.ShowStatistics {
		files <- core.NewFile(PageStatistics(),
			NewStatisticsPage(publisher.doc, publisher.GoogleAnalyticsID,
				publisher.options, publisher.indexLetters, publisher.placesMap))
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

func (publisher *Publisher) Places() map[string]*place {
	if publisher.placesMap == nil {
		publisher.placesMap = map[string]*place{}

		// Get all of the unique place names.
		for placeTag, node := range publisher.doc.Places() {
			prettyName := prettyPlaceName(placeTag.Value())

			if prettyName == "" {
				prettyName = "(none)"
			}

			key := alnumOrDashRegexp.
				ReplaceAllString(strings.ToLower(prettyName), "-")

			if _, ok := publisher.placesMap[key]; !ok {
				country := placeTag.Country()
				if country == "" {
					country = "(unknown)"
				}

				publisher.placesMap[key] = &place{
					PrettyName: prettyName,
					country:    country,
					nodes:      gedcom.Nodes{},
				}
			}

			publisher.placesMap[key].nodes = append(publisher.placesMap[key].nodes, node)
		}

		for key := range publisher.placesMap {
			// Make sure the events are sorted otherwise the pages will be
			// different.
			sort.Slice(publisher.placesMap[key].nodes, func(i, j int) bool {
				left := publisher.placesMap[key].nodes[i]
				right := publisher.placesMap[key].nodes[j]

				// Years.
				leftYears := gedcom.Years(left)
				rightYears := gedcom.Years(right)

				if leftYears != rightYears {
					return leftYears < rightYears
				}

				// Tag.
				leftTag := left.Tag().String()
				rightTag := right.Tag().String()

				if leftTag != rightTag {
					return leftTag < rightTag
				}

				// Individual name.
				leftIndividual := individualForNode(publisher.doc, left)
				rightIndividual := individualForNode(publisher.doc, right)

				if leftIndividual != nil && rightIndividual != nil {
					leftName := gedcom.String(leftIndividual.Name())
					rightName := gedcom.String(rightIndividual.Name())

					return leftName < rightName
				}

				// Value.
				valueLeft := gedcom.Value(left)
				valueRight := gedcom.Value(right)

				return valueLeft < valueRight
			})
		}
	}

	return publisher.placesMap
}
