package main

import (
	"auth/domain/repository"
	"github.com/satori/go.uuid"
	"github.com/mbict/go-dry/datamap"
	"fmt"
)

type TestUUIDRepository interface {
	Delete(id uuid.UUID) error
	Find() ([]*TestUUIDModel, error)
	FindById(id uuid.UUID) (*TestUUIDModel, error)
	FindByName(string) ([]*TestUUIDModel, error)
	FindFirstByName(string) (*TestUUIDModel, error)
	Save(client *TestUUIDModel) error
}

type TestUUIDModel struct {
	Id   uuid.UUID
	Name string
}

type testUUIDModelRepository struct {
	data *datamap.UUIDDatamap
}

// delete example
func (r *testUUIDModelRepository) Delete(id uuid.UUID) error {
	return r.data.Delete(id)
}

// find all items example no filtering
func (r *testUUIDModelRepository) Find() ([]*TestUUIDModel, error) {
	//we return all items
	matchFunc := func(interface{}) bool {
		return true
	}
	return r.data.Find([]*TestUUIDModel{}, matchFunc).([]*TestUUIDModel), nil
}

// default get by primary key example
func (r *testUUIDModelRepository) FindById(id uuid.UUID) (*TestUUIDModel, error) {
	row, err := r.data.FindById(id)
	if row != nil {
		return row.(*TestUUIDModel), err
	}
	return nil, err
}

// find all that match the name
func (r *testUUIDModelRepository) FindByName(name string) ([]*TestUUIDModel, error) {
	//match by name
	matchFunc := func(v interface{}) bool {
		return v.(*TestUUIDModel).Name == name
	}
	return r.data.Find([]*TestUUIDModel{}, matchFunc).([]*TestUUIDModel), nil
}

// find the first that matches the name
func (r *testUUIDModelRepository) FindFirstByName(name string) (*TestUUIDModel, error) {
	//match by name
	matchFunc := func(v interface{}) bool {
		return v.(*TestUUIDModel).Name == name
	}
	res, err := r.data.FindOne(matchFunc)
	if err != nil {
		return nil, err
	}
	return res.(*TestUUIDModel), nil
}

// save a record to the map
func (r *testUUIDModelRepository) Save(model *TestUUIDModel) error {
	return r.data.Save(model.Id, model)
}

func NewTestUUIDRepository() TestUUIDRepository {
	return &testUUIDModelRepository{
		data: datamap.NewUUID(repository.ErrNoResult),
	}
}

func main() {
	repo := NewTestUUIDRepository()
	id := uuid.NewV4()
	repo.Save(&TestUUIDModel{Id: id, Name: "test"})
	repo.Save(&TestUUIDModel{Id: uuid.NewV4(), Name: "test 2"})

	findResult, err := repo.Find()
	fmt.Printf("result Find:(%v) error:(%s)\n", findResult, err)

	findByIdResult, err := repo.FindById(id)
	fmt.Printf("result FindById:(%v) error:(%s)\n", findByIdResult, err)

	findByNameResult, err := repo.FindByName("test")
	fmt.Printf("result FindByName:(%v) error:(%s)\n", findByNameResult, err)

	findFirstByNameResult, err := repo.FindFirstByName("test 2")
	fmt.Printf("result FindFirstByName:(%v) error:(%s)\n", findFirstByNameResult, err)
}
