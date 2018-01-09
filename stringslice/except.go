package stringslice

//Except returns a slice with strings from a who are not present in b
func Except(a []string, b []string) []string {
	var res []string
	for _, v := range a {
		if !Has(b, v) {
			res = append(res, v)
		}
	}
	return res
}
