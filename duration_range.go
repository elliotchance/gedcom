package gedcom

type DurationRange struct {
	Min, Max Duration
}

func NewDurationRange(min, max Duration) DurationRange {
	return DurationRange{
		Min: min,
		Max: max,
	}
}

func (dr DurationRange) Age() (Age, Age) {
	min := NewAge(dr.Min.Duration, dr.Min.IsEstimate, AgeConstraintUnknown)
	max := NewAge(dr.Max.Duration, dr.Max.IsEstimate, AgeConstraintUnknown)

	return min, max
}
