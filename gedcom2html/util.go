package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

func colorForIndividual(individual *gedcom.IndividualNode) string {
	if individual == nil {
		return "black"
	}

	switch individual.Sex() {
	case gedcom.SexMale:
		return maleColor
	case gedcom.SexFemale:
		return femaleColor
	}

	return "black"
}

func colorClassForSex(sex gedcom.Sex) string {
	switch sex {
	case gedcom.SexMale:
		return "primary"

	case gedcom.SexFemale:
		return "danger"

	case gedcom.SexUnknown:
		return "info"

	default:
		return "info"
	}
}

func colorClassForIndividual(individual *gedcom.IndividualNode) string {
	if individual == nil {
		return "info"
	}

	return colorClassForSex(individual.Sex())
}

func getBirth(individual *gedcom.IndividualNode) (birthDate string, birthPlace string) {
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

func getDeath(individual *gedcom.IndividualNode) (deathDate string, deathPlace string) {
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
