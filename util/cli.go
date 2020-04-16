package util

import (
	"strings"
)

func CLIDescription(s string) (r string) {
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			r += "\n\n"
		} else {
			r += strings.Replace(line, "\t", "", -1) + " "
		}
	}

	return WrapToMargin(strings.TrimSpace(r), 80)
}

func WrapToMargin(s string, width int) (r string) {
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		words := strings.Split(line, " ")
		mutNewLine := ""

		for _, word := range words {
			if len(mutNewLine)+len(word)+1 > width {
				r += strings.TrimSpace(mutNewLine) + "\n"
				mutNewLine = word
			} else {
				mutNewLine += " " + word
			}
		}

		r += strings.TrimSpace(mutNewLine) + "\n"
	}

	// Remove last new line.
	// ghost:ignore
	r = r[:len(r)-1]

	return
}

// CLIStringSlice is used to accept multiple values from a CLI argument like:
//
//   -foo value1 -foo value2
//
// Usage:
//
//   var foos CLIStringSlice
//   flag.Var(&foos, "foo", "Some description for this param.")
//
type CLIStringSlice []string

func (i *CLIStringSlice) String() string {
	return strings.Join(*i, ";")
}

func (i *CLIStringSlice) Set(value string) error {
	*i = append(*i, value)

	return nil
}
