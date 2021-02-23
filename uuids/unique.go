package uuids

import "github.com/google/uuid"

func Unique(elements []uuid.UUID) []uuid.UUID {
	found := map[uuid.UUID]bool{}
	result := []uuid.UUID{}

	for v := range elements {
		if found[elements[v]] != true {
			found[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
