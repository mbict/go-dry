package intslice

func Unique(elements []int) []int {
	found := map[int]bool{}
	result := []int{}

	for v := range elements {
		if found[elements[v]] != true {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
