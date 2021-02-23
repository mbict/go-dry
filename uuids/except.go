package uuids

import "github.com/google/uuid"

//Except returns a slice with uuid's from a who are not present in b
func Except(a []uuid.UUID, b []uuid.UUID) []uuid.UUID {
	var res []uuid.UUID
	for _, v := range a {
		if !Has(b, v) {
			res = append(res, v)
		}
	}
	return res
}
