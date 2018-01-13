package stringslice

import "strings"

func Index(slice []string, value string) int {
	for i := range slice {
		if slice[i] == value {
			return i
		}
	}
	return -1
}

// HasCI is case insensitve version of Has
func IndexCI(slice []string, value string) int {
	for i := range slice {
		if strings.EqualFold(slice[i], value) {
			return i
		}
	}
	return -1
}
