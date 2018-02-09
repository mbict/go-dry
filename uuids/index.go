package uuids

import "github.com/satori/go.uuid"

func Index(slice []uuid.UUID, value uuid.UUID) int {
	for i := range slice {
		if uuid.Equal(slice[i], value) {
			return i
		}
	}
	return -1
}
