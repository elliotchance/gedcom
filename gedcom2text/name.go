package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

var usedFileNames map[string]bool

func outputFileName(individual *gedcom.IndividualNode) string {
	names := individual.Names()
	if len(names) == 0 {
		return ""
	}

	// Include the birth/death information to make the name more unique.
	birth := ""
	if node := individual.FirstNodeWithTag(gedcom.TagBirth); node != nil {
		if node2 := node.FirstNodeWithTag(gedcom.TagDate); node2 != nil {
			birth = node2.Value()
		}
	}

	death := ""
	if node := individual.FirstNodeWithTag(gedcom.TagDeath); node != nil {
		if node2 := node.FirstNodeWithTag(gedcom.TagDate); node2 != nil {
			death = node2.Value()
		}
	}

	// Sanitise the file name.
	fileName := fmt.Sprintf("%s (%s - %s)", names[0].String(), birth, death)

	// If the file name has already been used we mush make it unique.
	if usedFileNames == nil {
		usedFileNames = map[string]bool{}
	}

	i := 0
	for {
		if _, ok := usedFileNames[fileName]; !ok {
			break
		}

		if i == 0 {
			fileName += ", "
		}
		fileName += "I"
		i += 1
	}

	usedFileNames[fileName] = true

	fqn := fmt.Sprintf("%s/%s.txt", optionSplitDir, fileName)
	fmt.Println(fqn)

	return fqn
}
