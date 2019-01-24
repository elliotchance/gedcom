package gedcom

import (
	"fmt"
	"regexp"
	"strings"
)

type UUID string

var uuidRegexp = regexp.MustCompile(`^[0-9a-f]{32}$`)

// NewUUIDFromString returns a valid UUID in one of the forms:
//
//   EE13561DDB204985BFFDEEBF82A5226C
//   e0d4d387-618a-4713-ab3b-5fa3500b7a75
//
// An error is returned if there is more than 32 hexadecimal characters.
func NewUUIDFromString(s string) (UUID, error) {
	s2 := strings.Replace(s, "-", "", -1)
	s2 = strings.ToLower(s2)
	if !uuidRegexp.MatchString(s2) {
		return "", fmt.Errorf("invalid UUID: " + s)
	}

	uuid := fmt.Sprintf("%s-%s-%s-%s-%s",
		s2[:8], s2[8:12], s2[12:16], s2[16:20], s2[20:])

	return UUID(uuid), nil
}

// String returns a lowercase, hyphenated UUID in the form of:
//
//   e0d4d387-618a-4713-ab3b-5fa3500b7a75
//
func (uuid UUID) String() string {
	return string(uuid)
}

// Equals is true only if both UUIDs represent the same value.
func (uuid UUID) Equals(uuid2 UUID) bool {
	return string(uuid) == string(uuid2)
}
