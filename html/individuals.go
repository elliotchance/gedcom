package html

import (
	"regexp"
	"strings"

	"github.com/elliotchance/gedcom/v39"
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
