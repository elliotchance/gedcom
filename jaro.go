package gedcom

import (
	"math"
)

// JaroWinkler distance. JaroWinkler returns a number between 0 and 1 where 1
// means perfectly equal and 0 means completely different. It is commonly used
// on Record Linkage stuff, thus it tries to be accurate for real names and
// common typos. You should consider it on data such as person names and street
// names.
//
// JaroWinkler is a more accurate version of the Jaro algorithm. It works by
// boosting the score of exact matches at the beginning of the strings. By doing
// this, Winkler says that typos are less common to happen at the beginning. For
// this to happen, it introduces two more parameters: the boostThreshold and the
// prefixSize. These are commonly set to 0.7 and 4, respectively.
//
// The code and comment above has been copied from:
//
//   https://github.com/xrash/smetrics/blob/master/jaro-winkler.go
//
// A big thanks to Felipe (@xrash) for the explanation and code. It you read
// this, I copied the code to avoid the need to have a dependency manager for
// this project.
func JaroWinkler(a, b string) float64 {
	boostThreshold := 0.7
	prefixSize := 4

	j := jaro(a, b)

	if j <= boostThreshold {
		return j
	}

	prefixSize = int(math.Min(float64(len(a)), math.Min(float64(prefixSize), float64(len(b)))))

	var prefixMatch float64
	for i := 0; i < prefixSize; i++ {
		if a[i] == b[i] {
			prefixMatch++
		}
	}

	return j + 0.1*prefixMatch*(1.0-j)
}

// jaro was copied from the same place as JaroWinkler.
func jaro(a, b string) float64 {
	la := float64(len(a))
	lb := float64(len(b))

	// match range = max(len(a), len(b)) / 2 - 1
	matchRange := int(math.Floor(math.Max(la, lb)/2.0)) - 1
	matchRange = int(math.Max(0, float64(matchRange-1)))
	var matches, halfs float64
	transposed := make([]bool, len(b))

	for i := 0; i < len(a); i++ {
		start := int(math.Max(0, float64(i-matchRange)))
		end := int(math.Min(lb-1, float64(i+matchRange)))

		for j := start; j <= end; j++ {
			if transposed[j] {
				continue
			}

			if a[i] == b[j] {
				if i != j {
					halfs++
				}
				matches++
				transposed[j] = true
				break
			}
		}
	}

	if matches == 0 {
		return 0
	}

	transposes := math.Floor(float64(halfs / 2))

	return ((matches / la) + (matches / lb) + (matches-transposes)/matches) / 3.0
}
