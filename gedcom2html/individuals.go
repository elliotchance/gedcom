package main

import (
	"strings"
	"github.com/elliotchance/gedcom"
	"fmt"
	"regexp"
)

var individualMap map[string]*gedcom.IndividualNode

func getIndividuals(document *gedcom.Document) map[string]*gedcom.IndividualNode {
	if individualMap == nil {
		individualMap = map[string]*gedcom.IndividualNode{}

		for _, individual := range document.Individuals() {
			name := individual.Name().String()

			if d := individual.FirstNodeWithTagPath(gedcom.TagBirth, gedcom.TagDate); d != nil {
				name += "-" + d.Value()
			}

			if d := individual.FirstNodeWithTagPath(gedcom.TagDeath, gedcom.TagDate); d != nil {
				name += "-" + d.Value()
			}

			key := getUniqueKey(regexp.MustCompile("[^a-z_0-9-]+").
				ReplaceAllString(strings.ToLower(name), "-"))

			individualMap[key] = individual
		}
	}

	return individualMap
}

func getUniqueKey(s string) string {
	i := -1
	for {
		i += 1

		testString := s
		if i > 0 {
			testString = fmt.Sprintf("%s-%d", s, i)
		}

		if _, ok := individualMap[testString]; ok {
			continue
		}

		if _, ok := placesMap[testString]; ok {
			continue
		}

		return testString
	}

	// This should not be possible
	panic(s)
}
