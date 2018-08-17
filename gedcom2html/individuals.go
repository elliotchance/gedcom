package main

import (
	"github.com/elliotchance/gedcom"
	"regexp"
	"strings"
)

var individualMap map[string]*gedcom.IndividualNode

var alnumOrDashRegexp = regexp.MustCompile("[^a-z_0-9-]+")

func getIndividuals(document *gedcom.Document) map[string]*gedcom.IndividualNode {
	if individualMap == nil {
		individualMap = map[string]*gedcom.IndividualNode{}

		for _, individual := range document.Individuals() {
			name := individual.Name().String()

			key := getUniqueKey(alnumOrDashRegexp.
				ReplaceAllString(strings.ToLower(name), "-"))

			individualMap[key] = individual
		}
	}

	return individualMap
}
