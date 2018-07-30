package gedcom

import "strings"

// DateConstraint describes if a date is constrained by a particular range. See
// Date for full explanation.
type DateConstraint int

const (
	// There is no constraint. The date is at the value specified.
	DateConstraintExact = DateConstraint(iota)

	// The date is approximate. There is no defined error margin (how much the
	// date may be off by) but in mose cases it is save to assume it is
	// proportional to how precise the date is. That is, usually a date that
	// provides a day, month and year will have a smaller error margin than a
	// day that only provides a year.
	DateConstraintAbout

	// The real date is before the specified date value. This often (but not
	// always) follows the same proportional rules as DateConstraintApprox.
	DateConstraintBefore

	// The real date is after the specified date value. This often (but not
	// always) follows the same proportional rules as DateConstraintApprox.
	DateConstraintAfter
)

// DateConstraintFromString will return the constraint for the provided keyword
// based on the keywords described in the specification of DateNode.
//
// If the word is not understood DateConstraintExact will be returned.
//
// This function is not case sensitive so "before" and "Before" are treated the
// same.
func DateConstraintFromString(word string) DateConstraint {
	// All words are lower case.
	lowerWord := strings.ToLower(word)

	switch {
	case wordInWords(lowerWord, DateWordsAbout):
		return DateConstraintAbout

	case wordInWords(lowerWord, DateWordsAfter):
		return DateConstraintAfter

	case wordInWords(lowerWord, DateWordsBefore):
		return DateConstraintBefore
	}

	// In all other cases default to Exact.
	return DateConstraintExact
}

func wordInWords(word, words string) bool {
	// Be careful to convert the words to lowercase as the capitalization is
	// different for the first value for DateConstraint.String()
	for _, w := range strings.Split(strings.ToLower(words), "|") {
		if w == word {
			return true
		}
	}

	return false
}

// String returns the constraint abbreviation for non-exact dates. Exact dates
// will return an empty string.
//
// The chosen (default) abbreviation is the first value in the appropriate
// DateWords constant.
//
// If the constraint is invalid then an empty string will be returned.
func (constraint DateConstraint) String() string {
	words := ""

	switch constraint {
	case DateConstraintAbout:
		words = DateWordsAbout

	case DateConstraintAfter:
		words = DateWordsAfter

	case DateConstraintBefore:
		words = DateWordsBefore
	}

	return strings.Split(words, "|")[0]
}
