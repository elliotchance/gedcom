package util

import "strings"

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
		newLine := ""

		for _, word := range words {
			if len(newLine)+len(word)+1 > width {
				r += strings.TrimSpace(newLine) + "\n"
				newLine = word
			} else {
				newLine += " " + word
			}
		}

		r += strings.TrimSpace(newLine) + "\n"
	}

	// Remove last new line
	r = r[:len(r)-1]

	return
}
