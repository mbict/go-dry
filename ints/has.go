package ints

func Has(slice []int, value int) bool {
	for i := range slice {
		if slice[i] == value {
			return true
		}
	}
	return false
}
