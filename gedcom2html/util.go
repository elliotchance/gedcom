package main

import (
	"github.com/elliotchance/gedcom"
	"strings"
)

func colorForIndividual(individual *gedcom.IndividualNode) string {
	switch individual.Sex() {
	case gedcom.SexMale:
		return maleColor
	case gedcom.SexFemale:
		return femaleColor
	}

	return "white"
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

	birthNode := individual.FirstNodeWithTag(gedcom.TagBirth)
	if birthNode != nil {
		birthDateNode := birthNode.FirstNodeWithTag(gedcom.TagDate)
		if birthDateNode != nil {
			birthDate = birthDateNode.Value()
		}

		birthPlaceNode := birthNode.FirstNodeWithTag(gedcom.TagPlace)
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

	deathNode := individual.FirstNodeWithTag(gedcom.TagDeath)
	if deathNode != nil {
		deathDateNode := deathNode.FirstNodeWithTag(gedcom.TagDate)
		if deathDateNode != nil {
			deathDate = deathDateNode.Value()
		}

		deathPlaceNode := deathNode.FirstNodeWithTag(gedcom.TagPlace)
		if deathPlaceNode != nil {
			deathPlace = deathPlaceNode.Value()
		}
	}

	return
}

func prettyPlaceName(s string) string {
	s = strings.Replace(s, ",,", ",", -1)
	s = strings.Replace(s, ",,", ",", -1)
	s = strings.Replace(s, ",", ", ", -1)

	return s
}
