package uuids

import "github.com/google/uuid"

//Intersect returns a slice with uuids's who are present in both slices
func Intersect(a []uuid.UUID, b []uuid.UUID) []uuid.UUID {
	var res []uuid.UUID
	for _, v := range b {
		if Has(a, v) {
			res = append(res, v)
		}
	}
	return res
}
