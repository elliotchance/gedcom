package gedcom

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// A duration that only considers whole-day resolution.
type Duration struct {
	Duration time.Duration

	// IsEstimate and IsKnown work the same way as described in Age.
	IsEstimate, IsKnown bool
}

func NewExactDuration(duration time.Duration) Duration {
	return NewDuration(duration, true, false)
}

func NewDuration(duration time.Duration, isKnown, isEstimate bool) Duration {
	return Duration{
		// Durations must always be positive.
		Duration:   positiveDuration(duration),
		IsEstimate: isEstimate,
		IsKnown:    isKnown,
	}
}

func pluralize(value int, word string) string {
	switch value {
	case 0:
		return ""

	case 1:
		return "one " + word

	default:
		return fmt.Sprintf("%d %ss", value, word)
	}
}

func (d Duration) String() string {
	oneDay := time.Duration(24 * time.Hour)
	oneMonth := time.Duration(30.4166 * float64(oneDay))
	oneYear := time.Duration(365 * float64(oneDay))

	if d.Duration < oneDay {
		return "one day"
	}

	var parts []string

	if years := int(d.Duration / oneYear); years != 0 {
		d.Duration -= time.Duration(years) * oneYear
		parts = append(parts, pluralize(years, "year"))
	}

	if months := int(d.Duration / oneMonth); months != 0 {
		d.Duration -= time.Duration(months) * oneMonth
		parts = append(parts, pluralize(months, "month"))
	}

	if days := int(math.Ceil(float64(d.Duration) / float64(oneDay))); days != 0 {
		parts = append(parts, pluralize(days, "day"))
	}

	return strings.Join(parts, " and ")
}
