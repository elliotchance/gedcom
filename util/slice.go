package util

// StringSliceContains finds if a string exist in a slice of strings.
func StringSliceContains(slice []string, find string) bool {
	for _, element := range slice {
		if element == find {
			return true
		}
	}

	return false
}
