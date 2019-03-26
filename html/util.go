package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
	"strings"
)

const symbolLetter = '#'

func write(w io.Writer, data []byte) (int64, error) {
	n, err := w.Write(data)

	return int64(n), err
}

func writeString(w io.Writer, data string) (int64, error) {
	return write(w, []byte(data))
}

func appendString(w io.Writer, data string) int64 {
	n, err := writeString(w, data)
	if err != nil {
		panic(err)
	}

	return n
}

func appendComponent(w io.Writer, component core.Component) int64 {
	n, err := component.WriteHTMLTo(w)
	if err != nil {
		panic(err)
	}

	return n
}

func writeSprintf(w io.Writer, format string, args ...interface{}) (int64, error) {
	return writeString(w, fmt.Sprintf(format, args...))
}

func appendSprintf(w io.Writer, format string, args ...interface{}) int64 {
	n, err := writeSprintf(w, format, args...)
	if err != nil {
		panic(err)
	}

	return n
}

func writeNothing() (int64, error) {
	return 0, nil
}

func PageIndividuals(firstLetter rune) string {
	if firstLetter == symbolLetter {
		return "individuals-symbol.html"
	}

	return fmt.Sprintf("individuals-%c.html", firstLetter)
}

func PageIndividual(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility) string {
	if individual.IsLiving() {
		switch visibility {
		case LivingVisibilityHide, LivingVisibilityPlaceholder:
			return "#"

		case LivingVisibilityShow:
			// Proceed.
		}
	}

	individuals := GetIndividuals(document)

	for key, value := range individuals {
		if value.Is(individual) {
			return fmt.Sprintf("%s.html", key)
		}
	}

	return "#"
}

func PagePlaces() string {
	return "places.html"
}

func PagePlace(document *gedcom.Document, place string) string {
	places := GetPlaces(document)

	for key, value := range places {
		if value.PrettyName == place {
			return fmt.Sprintf("%s.html", key)
		}
	}

	return "#"
}

func PageFamilies() string {
	return "families.html"
}

func PageSources() string {
	return "sources.html"
}

func PageSource(source *gedcom.SourceNode) string {
	return fmt.Sprintf("%s.html", source.Pointer())
}

func PageStatistics() string {
	return "statistics.html"
}

func PageSurnames() string {
	return "surnames.html"
}

func colorForIndividual(individual *gedcom.IndividualNode) string {
	if individual == nil {
		return "black"
	}

	sex := individual.Sex()
	switch {
	case sex.IsMale():
		return IndividualMaleColor
	case sex.IsFemale():
		return IndividualFemaleColor
	}

	return "black"
}

func colorClassForSex(sex *gedcom.SexNode) string {
	switch {
	case sex.IsMale():
		return "primary"

	case sex.IsFemale():
		return "danger"
	}

	return "info"
}

func colorClassForIndividual(individual *gedcom.IndividualNode) string {
	if individual == nil {
		return "info"
	}

	return colorClassForSex(individual.Sex())
}

func getUniqueKey(individualMap map[string]*gedcom.IndividualNode, s string) string {
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
	firstLetter := rune(lowerName[0])

	return firstLetter == letter
}

func individualForNode(doc *gedcom.Document, node gedcom.Node) *gedcom.IndividualNode {
	for _, individual := range doc.Individuals() {
		if gedcom.HasNestedNode(individual, node) {
			return individual
		}
	}

	return nil
}
