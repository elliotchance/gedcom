package html

import (
	"github.com/elliotchance/gedcom"
	"strings"
)

type place struct {
	PrettyName string
	country    string
	nodes      gedcom.Nodes
}

func prettyPlaceName(s string) string {
	s = strings.Replace(s, ",,", ",", -1)
	s = strings.Replace(s, ",,", ",", -1)
	s = strings.Replace(s, ",", ", ", -1)
	s = strings.Trim(s, ", ")

	return strings.TrimSpace(s)
}
