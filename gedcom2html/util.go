package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"strings"
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

func surnameStartsWith(individual *gedcom.IndividualNode, letter rune) bool {
	name := individual.Name().Format(gedcom.NameFormatIndex)
	if name == "" {
		name = "#"
	}

	lowerName := strings.ToLower(name)
	return rune(lowerName[0]) == letter
}

func individualForNode(node gedcom.Node) *gedcom.IndividualNode {
	for _, individual := range node.Document().Individuals() {
		if gedcom.HasNestedNode(individual, node) {
			return individual
		}
	}

	return nil
}
