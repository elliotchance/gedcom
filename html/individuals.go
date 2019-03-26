package html

import (
	"github.com/elliotchance/gedcom"
	"regexp"
	"strings"
)

var alnumOrDashRegexp = regexp.MustCompile("[^a-z_0-9-]+")

func GetIndividuals(document *gedcom.Document) map[string]*gedcom.IndividualNode {
	document.PublishIndividualMapMutex.Lock()

	if document.PublishIndividualMap != nil {
		return document.PublishIndividualMap
	}

	for _, individual := range document.Individuals() {
		name := individual.Name().String()

		key := getUniqueKey(document, alnumOrDashRegexp.
			ReplaceAllString(strings.ToLower(name), "-"))

		document.PublishIndividualSyncedMap.Store(key, individual)
	}

	document.PublishIndividualMap = map[string]*gedcom.IndividualNode{}

	document.PublishIndividualSyncedMap.Range(func(key, value interface{}) bool {
		document.PublishIndividualMap[key.(string)] = value.(*gedcom.IndividualNode)

		return true
	})

	document.PublishIndividualMapMutex.Unlock()

	return document.PublishIndividualMap
}
