package uuids

import "github.com/google/uuid"

func Index(slice []uuid.UUID, value uuid.UUID) int {
	for i := range slice {
		if slice[i] == value {
			return i
		}
	}
	return -1
}
