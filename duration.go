package gedcom

import (
	"time"
	"fmt"
	"math"
	"strings"
)

// A duration that only considers whole-day resolution.
type Duration time.Duration

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
	oneDay := Duration(24 * time.Hour)
	oneMonth := Duration(30.4166 * float64(oneDay))
	oneYear := Duration(365 * float64(oneDay))

	if d < oneDay {
		return "one day"
	}

	var parts []string

	if years := int(d / oneYear); years != 0 {
		d -= Duration(years) * oneYear
		parts = append(parts, pluralize(years, "year"))
	}

	if months := int(d / oneMonth); months != 0 {
		d -= Duration(months) * oneMonth
		parts = append(parts, pluralize(months, "month"))
	}

	if days := int(math.Ceil(float64(d)/float64(oneDay))); days != 0 {
		parts = append(parts, pluralize(days, "day"))
	}

	return strings.Join(parts, " and ")
}
