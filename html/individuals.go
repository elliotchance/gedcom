package html

import (
	"github.com/elliotchance/gedcom"
	"regexp"
	"strings"
)

var alnumOrDashRegexp = regexp.MustCompile("[^a-z_0-9-]+")

func GetIndividuals(document *gedcom.Document, placesMap map[string]*place) map[string]*gedcom.IndividualNode {
	individualMap := map[string]*gedcom.IndividualNode{}

	for _, individual := range document.Individuals() {
		name := individual.Name().String()

		key := getUniqueKey(individualMap, alnumOrDashRegexp.
			ReplaceAllString(strings.ToLower(name), "-"), placesMap)

		individualMap[key] = individual
	}

	return individualMap
}
