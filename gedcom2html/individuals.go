package main

import (
	"strings"
	"github.com/elliotchance/gedcom"
	"regexp"
)

var individualMap map[string]*gedcom.IndividualNode

func getIndividuals(document *gedcom.Document) map[string]*gedcom.IndividualNode {
	if individualMap == nil {
		individualMap = map[string]*gedcom.IndividualNode{}

		for _, individual := range document.Individuals() {
			name := individual.Name().String()

			key := getUniqueKey(regexp.MustCompile("[^a-z_0-9-]+").
				ReplaceAllString(strings.ToLower(name), "-"))

			individualMap[key] = individual
		}
	}

	return individualMap
}
