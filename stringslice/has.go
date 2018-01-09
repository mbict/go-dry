package stringslice

func Has(slice []string, value string) bool {
	for i := range slice {
		if slice[i] == value {
			return true
		}
	}
	return false
}
