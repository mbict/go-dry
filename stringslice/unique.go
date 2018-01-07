package stringslice

func Unique(elements []string) []string {
	found := map[string]bool{}
	result := []string{}

	for v := range elements {
		if found[elements[v]] != true {
			found[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}


func Has(slice []string, value string) bool {
	for i := range slice {
		if slice[i] == value {
			return true
		}
	}
	return false
}