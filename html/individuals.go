package html

import (
	"github.com/elliotchance/gedcom"
	"regexp"
	"strings"
	"sync"
)

var individualSyncedMap sync.Map // map[string]*gedcom.IndividualNode
var individualMap map[string]*gedcom.IndividualNode
var individualMapSync sync.Mutex

var alnumOrDashRegexp = regexp.MustCompile("[^a-z_0-9-]+")

func GetIndividuals(document *gedcom.Document) map[string]*gedcom.IndividualNode {
	if individualMap != nil {
		return individualMap
	}

	individualMapSync.Lock()

	for _, individual := range document.Individuals() {
		name := individual.Name().String()

		key := getUniqueKey(alnumOrDashRegexp.
			ReplaceAllString(strings.ToLower(name), "-"))

		individualSyncedMap.Store(key, individual)
	}

	individualMap = map[string]*gedcom.IndividualNode{}

	individualSyncedMap.Range(func(key, value interface{}) bool {
		individualMap[key.(string)] = value.(*gedcom.IndividualNode)

		return true
	})

	individualMapSync.Unlock()

	return individualMap
}
