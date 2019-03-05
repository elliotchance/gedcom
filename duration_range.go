package gedcom

import "time"

type DurationRange struct {
	Min, Max   Duration
	IsEstimate bool
}

func NewDurationRange(min, max Duration, isEstimate bool) DurationRange {
	return DurationRange{
		Min:        min,
		Max:        max,
		IsEstimate: isEstimate,
	}
}

func (dr DurationRange) Age() (Age, Age) {
	min := NewAge(time.Duration(dr.Min), dr.IsEstimate, AgeConstraintUnknown)
	max := NewAge(time.Duration(dr.Max), dr.IsEstimate, AgeConstraintUnknown)

	return min, max
}
