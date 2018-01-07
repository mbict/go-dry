package stringslice

func Unique(elements []string) []string {
	found := map[string]bool{}
	result := []string{}

	for v := range elements {
		if found[elements[v]] != true {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
