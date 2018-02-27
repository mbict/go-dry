package ints


func Index(slice []int, value int) int {
	for i := range slice {
		if slice[i] == value {
			return i
		}
	}
	return -1
}
