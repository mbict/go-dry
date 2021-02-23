package uuids

import "github.com/google/uuid"

func Has(slice []uuid.UUID, value uuid.UUID) bool {
	for i := range slice {
		if slice[i] == value {
			return true
		}
	}
	return false
}
