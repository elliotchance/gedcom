package gedcom

import "time"

// Year is an approximation for the duration of a year.
//
// This should not be used in calculations that require more than month-level
// precision.
//
// A year is approximated at 365.25 to take into account leap years.
const Year = time.Duration(float64(time.Hour) * 24 * 365.25)

// Age represents an age of an individual at a point in time.
//
// An Age is often paired with another Age to allow a range of time. For example
// an event could be an absolute like a christening, or it could represent a
// range such as the start and end age during a residence that spans multiple
// years.
//
// You should avoid initializing the struct directly, and instead use the most
// appropriate constructor.
type Age struct {
	// Age is the duration of time from the point where the birth was
	// determined.
	//
	// The age may be an estimation, see IsEstimate. The age might also be
	// greater than the maximum living age of the individual, see IsAfterDeath.
	Age time.Duration

	// IsEstimate will be true when there was no birth event or the birth event
	// is a range (like "Between 1943 to 1947").
	IsEstimate bool

	// IsKnown will be true if the age can be determined (either as exact or an
	// estimation). The age cannot be determined if no estimated birth date is
	// found or an event does not contain a usable date.
	//
	// When IsKnown is false you should not use the value of Age.
	IsKnown bool

	// See the AgeConstraint constants for a full explanation.
	Constraint AgeConstraint
}

// NewUnknownAge returns a value that represents an age that is not known.
func NewUnknownAge() Age {
	return Age{}
}

// NewAge will initialise a new known exact or estimate age age with the
// provided attributes.
func NewAge(age time.Duration, isEstimate bool, constraint AgeConstraint) Age {
	return Age{
		Age:        age,
		IsKnown:    true,
		IsEstimate: isEstimate,
		Constraint: constraint,
	}
}

// NewAgeWithYears creates an age by using the number of years.
//
// This is not precise as a year is averaged out at 365.25 days (see Year
// constant) but is useful when only whole years matter.
func NewAgeWithYears(years float64, isEstimate bool, constraint AgeConstraint) Age {
	age := time.Duration(years * float64(Year))

	return NewAge(age, isEstimate, constraint)
}

// IsAfter is true if the right age is after (greater than) the left age. If
// both ages are equal this will return false.
func (age Age) IsAfter(age2 Age) bool {
	return age.Age > age2.Age
}

func constraintBetweenAges(estimatedBirthDate, estimatedDeathDate Date, age time.Duration) AgeConstraint {
	if age < 0 {
		return AgeConstraintBeforeBirth
	}

	if estimatedBirthDate.IsZero() || estimatedDeathDate.IsZero() {
		return AgeConstraintUnknown
	}

	if age > estimatedDeathDate.Time().Sub(estimatedBirthDate.Time()) {
		return AgeConstraintAfterDeath
	}

	return AgeConstraintLiving
}
