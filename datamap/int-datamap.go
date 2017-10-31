package datamap

import (
	"strconv"
)

type IntDatamap struct {
	*Datamap
}

func (r *IntDatamap) Delete(id int) error {
	return r.Datamap.Delete(strconv.Itoa(id))
}

func (r *IntDatamap) FindById(id int) (interface{}, error) {
	return r.Datamap.FindById(strconv.Itoa(id))
}

func (r *IntDatamap) Save(id int, row interface{}) error {
	return r.Datamap.Save(strconv.Itoa(id), row)
}

func NewInt(notFoundError error) *IntDatamap {
	return &IntDatamap{
		Datamap: New(notFoundError),
	}
}
