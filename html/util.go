package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

func GetBirth(individual *gedcom.IndividualNode) (birthDate string, birthPlace string) {
	if individual == nil {
		return
	}

	birthNode := gedcom.First(gedcom.NodesWithTag(individual, gedcom.TagBirth))
	if birthNode != nil {
		birthDateNode := gedcom.First(gedcom.NodesWithTag(birthNode, gedcom.TagDate))
		if birthDateNode != nil {
			birthDate = birthDateNode.Value()
		}

		birthPlaceNode := gedcom.First(gedcom.NodesWithTag(birthNode, gedcom.TagPlace))
		if birthPlaceNode != nil {
			birthPlace = birthPlaceNode.Value()
		}
	}

	return
}

func GetDeath(individual *gedcom.IndividualNode) (deathDate string, deathPlace string) {
	if individual == nil {
		return
	}

	deathNode := gedcom.First(gedcom.NodesWithTag(individual, gedcom.TagDeath))
	if deathNode != nil {
		deathDateNode := gedcom.First(gedcom.NodesWithTag(deathNode, gedcom.TagDate))
		if deathDateNode != nil {
			deathDate = deathDateNode.Value()
		}

		deathPlaceNode := gedcom.First(gedcom.NodesWithTag(deathNode, gedcom.TagPlace))
		if deathPlaceNode != nil {
			deathPlace = deathPlaceNode.Value()
		}
	}

	return
}

func Sprintf(format string, args ...interface{}) string {
	newArgs := make([]interface{}, len(args))
	for i, arg := range args {
		if a, ok := arg.(fmt.Stringer); ok {
			newArgs[i] = a.String()
		} else {
			newArgs[i] = arg
		}
	}

	return fmt.Sprintf(format, newArgs...)
}
