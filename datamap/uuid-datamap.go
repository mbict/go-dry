package datamap

import (
	"github.com/google/uuid"
)

type UUIDDatamap struct {
	*Datamap
}

func (r *UUIDDatamap) Delete(id uuid.UUID) error {
	return r.Datamap.Delete(id.String())
}

func (r *UUIDDatamap) FindById(id uuid.UUID) (interface{}, error) {
	return r.Datamap.FindById(id.String())
}

func (r *UUIDDatamap) Save(id uuid.UUID, row interface{}) error {
	return r.Datamap.Save(id.String(), row)
}

func NewUUID(notFoundError error) *UUIDDatamap {
	return &UUIDDatamap{
		Datamap: New(notFoundError),
	}
}
