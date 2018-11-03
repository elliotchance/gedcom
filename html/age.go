package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"time"
)

type Age struct {
	start, end gedcom.Age
}

func NewAge(start, end gedcom.Age) *Age {
	return &Age{
		start: start,
		end:   end,
	}
}

func (c *Age) string(age gedcom.Age) string {
	if !age.IsKnown {
		return ""
	}

	// Ages can be after death, such as a burial or probate. It does not make
	// sense to show these ages.
	if age.Constraint == gedcom.AgeConstraintAfterDeath {
		return ""
	}

	return age.String()
}

func (c *Age) String() (s string) {
	start := c.string(c.start)
	end := c.string(c.end)

	switch {
	case start == "" && end == "":
		s = ""

	case end == "":
		s = fmt.Sprintf(`after %s`, start)

	case start == end:
		s = start

	case start == "":
		s = fmt.Sprintf(`until %s`, end)

	case start != end:
		// If there is less than a year between the two dates (which is very
		// common because many dates only contain a year) we collapse it into a
		// single minimum value.
		//
		// The 1.05 is to account for slight rounding errors and leap years.
		//
		// We only need to check one direction because Age() and AgeAt()
		// guarantee that the end is greater than or equal to the start.
		yearAndABit := float64(gedcom.Year) * 1.05

		if c.end.Age-c.start.Age < time.Duration(yearAndABit) {
			s = start
		} else {
			s = fmt.Sprintf(`from %s to %s`, start, end)
		}
	}

	return
}
