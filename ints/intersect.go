package ints

//Intersect returns a slice with int's who are present in both slices
func Intersect(a []int, b []int) []int {
	var res []int
	for _, v := range b {
		if Has(a, v) {
			res = append(res, v)
		}
	}
	return res
}
