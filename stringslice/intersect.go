package stringslice

//Intersect returns a slice with strings who are present in both slices
func Intersect(a []string, b []string) []string {
	var res []string
	for _, v := range b {
		if Has(a, v) {
			res = append(res, v)
		}
	}
	return res
}
