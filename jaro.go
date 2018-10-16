package gedcom

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

// Theses constants are used with JaroWinkler. They have been calculated using
// the gedcomtune command.
const (
	DefaultJaroWinklerBoostThreshold = 0.0
	DefaultJaroWinklerPrefixSize     = 8
)

// JaroWinkler distance. JaroWinkler returns a number between 0 and 1 where 1
// means perfectly equal and 0 means completely different. It is commonly used
// on Record Linkage stuff, thus it tries to be accurate for real names and
// common typos. You should consider it on data such as person names and street
// names.
//
// JaroWinkler is sensitive to punctuation and capitalization of strings. If you
// are comparing names or places you may want to use StringSimilarity instead.
// StringSimilarity removes punctuation, extra whitespace and normalises
// capitalization before performing the comparison which almost always delivers
// more desirable results.
//
// JaroWinkler is a more accurate version of the Jaro algorithm. It works by
// boosting the score of exact matches at the beginning of the strings. By doing
// this, Winkler says that typos are less common to happen at the beginning. For
// this to happen, it introduces two more parameters: the boostThreshold and the
// prefixSize. These are commonly set to 0.7 and 4, respectively. However, you
// should use DefaultJaroWinklerBoostThreshold and DefaultJaroWinklerPrefixSize
// as they have been calcuated by the gedcomtune command.
//
// The code and comment above has been copied from:
//
//   https://github.com/xrash/smetrics/blob/master/jaro-winkler.go
//
// A big thanks to Felipe (@xrash) for the explanation and code. It you read
// this, I copied the code to avoid the need to have a dependency manager for
// this project.
func JaroWinkler(a, b string, boostThreshold float64, prefixSize int) float64 {
	j := jaro(a, b)

	if j <= boostThreshold {
		return j
	}

	aLen := len(a)
	bLen := len(b)
	prefixSize = minInt(prefixSize, aLen, bLen)

	var prefixMatch float64
	for i := 0; i < prefixSize; i++ {
		if a[i] == b[i] {
			prefixMatch++
		}
	}

	return j + 0.1*prefixMatch*(1.0-j)
}

func minInt(values ...int) int {
	sort.Ints(values)

	return values[0]
}

// jaro was copied from the same place as JaroWinkler.
func jaro(a, b string) float64 {
	aLen := len(a)
	la := float64(aLen)
	lb := float64(len(b))

	// match range = max(len(a), len(b)) / 2 - 1
	lAvg := math.Max(la, lb) / 2.0
	matchRange := math.Floor(lAvg) - 2
	matchRange = math.Max(0, matchRange)

	var matches, halfs float64
	transposed := make([]bool, len(b))

	for i := 0; i < aLen; i++ {
		start := int(math.Max(0, float64(i)-matchRange))
		end := int(math.Min(lb-1, float64(i)+matchRange))

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

	transposes := math.Floor(halfs / 2)
	aMatches := matches / la
	bMatches := matches / lb
	cMatches := (matches - transposes) / matches

	return avg(aMatches, bMatches, cMatches)
}

func sum(numbers ...float64) (total float64) {
	for _, number := range numbers {
		total += number
	}

	return
}

func avg(numbers ...float64) float64 {
	total := sum(numbers...)
	count := float64(len(numbers))

	return total / count
}

var alnumRegexp = regexp.MustCompile("[^a-z0-9 ]")

// StringSimilarity is a less sensitive version of a JaroWinkler string
// comparison. It is the ideal choice to compare strings that represent
// individual names, places, etc as it does not take into account
// capitalization, punctuation and extra spaces.
func StringSimilarity(a, b string, boostThreshold float64, prefixSize int) float64 {
	a = alnumRegexp.ReplaceAllString(strings.ToLower(a), "")
	b = alnumRegexp.ReplaceAllString(strings.ToLower(b), "")

	cleanA := CleanSpace(a)
	cleanB := CleanSpace(b)

	return JaroWinkler(cleanA, cleanB, boostThreshold, prefixSize)
}
