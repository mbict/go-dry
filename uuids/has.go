package uuids

import "github.com/satori/go.uuid"

func Has(slice []uuid.UUID, value uuid.UUID) bool {
	for i := range slice {
		if uuid.Equal(slice[i], value) {
			return true
		}
	}
	return false
}
