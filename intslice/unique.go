package intslice

func Unique(elements []int) []int {
	found := map[int]bool{}
	result := []int{}

	for v := range elements {
		if found[elements[v]] != true {
			found[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

func Has(slice []int, value int) bool {
	for i := range slice {
		if slice[i] == value {
			return true
		}
	}
	return false
}