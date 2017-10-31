package main

import (
	"auth/domain/repository"
	"github.com/mbict/go-dry/datamap"
	"fmt"
)

type TestRepository interface {
	Delete(id string) error
	Find() ([]*TestModel, error)
	FindById(id string) (*TestModel, error)
	FindByName(string) ([]*TestModel, error)
	FindFirstByName(string) (*TestModel, error)
	Save(client *TestModel) error
}

type TestModel struct {
	Id   string
	Name string
}

type testModelRepository struct {
	data *datamap.Datamap
}

// delete example
func (r *testModelRepository) Delete(id string) error {
	return r.data.Delete(id)
}

// find all items example no filtering
func (r *testModelRepository) Find() ([]*TestModel, error) {
	//we return all items
	matchFunc := func(interface{}) bool {
		return true
	}
	return r.data.Find([]*TestModel{}, matchFunc).([]*TestModel), nil
}

// default get by primary key example
func (r *testModelRepository) FindById(id string) (*TestModel, error) {
	row, err := r.data.FindById(id)
	if row != nil {
		return row.(*TestModel), err
	}
	return nil, err
}

// find all that match the name
func (r *testModelRepository) FindByName(name string) ([]*TestModel, error) {
	//match by name
	matchFunc := func(v interface{}) bool {
		return v.(*TestModel).Name == name
	}
	return r.data.Find([]*TestModel{}, matchFunc).([]*TestModel), nil
}

// find the first that matches the name
func (r *testModelRepository) FindFirstByName(name string) (*TestModel, error) {
	//match by name
	matchFunc := func(v interface{}) bool {
		return v.(*TestModel).Name == name
	}
	res, err := r.data.FindOne(matchFunc)
	if err != nil {
		return nil, err
	}
	return res.(*TestModel), nil
}

// save a record to the map
func (r *testModelRepository) Save(model *TestModel) error {
	return r.data.Save(model.Id, model)
}

func NewTestRepository() TestRepository {
	return &testModelRepository{
		data: datamap.New(repository.ErrNoResult),
	}
}

func main() {
	repo := NewTestRepository()

	repo.Save(&TestModel{Id: "abcxxx", Name: "test"})
	repo.Save(&TestModel{Id: "efghij", Name: "test 2"})

	findResult, err := repo.Find()
	fmt.Printf("result Find:(%v) error:(%s)\n", findResult, err)

	findByIdResult, err := repo.FindById("abcxxx")
	fmt.Printf("result FindById:(%v) error:(%s)\n", findByIdResult, err)

	findByNameResult, err := repo.FindByName("test")
	fmt.Printf("result FindByName:(%v) error:(%s)\n", findByNameResult, err)

	findFirstByNameResult, err := repo.FindFirstByName("test 2")
	fmt.Printf("result FindFirstByName:(%v) error:(%s)\n", findFirstByNameResult, err)
}
