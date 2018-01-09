package ints

//Except returns a slice with int from a who are not present in b
func Except(a []int, b []int) []int {
	var res []int
	for _, v := range a {
		if !Has(b, v) {
			res = append(res, v)
		}
	}
	return res
}
