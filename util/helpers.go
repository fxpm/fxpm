package util

// SliceContainsString checks if a slice contains the provided
// string value.
func SliceContainsString(slice []string, value string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}

	return false
}
