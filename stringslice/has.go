package stringslice

import "strings"

func Has(slice []string, value string) bool {
	for i := range slice {
		if slice[i] == value {
			return true
		}
	}
	return false
}

// HasCI is case insensitve version of has
func HasCI(slice []string, value string) bool {
	for i := range slice {
		if strings.EqualFold(slice[i], value) {
			return true
		}
	}
	return false
}
