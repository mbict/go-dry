package main

import (
	"errors"
	"fmt"
	"github.com/mbict/go-dry/datamap"
)

type TestIntRepository interface {
	Delete(id int) error
	Find() ([]*TestIntModel, error)
	FindById(id int) (*TestIntModel, error)
	FindByName(string) ([]*TestIntModel, error)
	FindFirstByName(string) (*TestIntModel, error)
	Save(client *TestIntModel) error
}

type TestIntModel struct {
	Id   int
	Name string
}

type testIntModelRepository struct {
	data *datamap.IntDatamap
}

// delete example
func (r *testIntModelRepository) Delete(id int) error {
	return r.data.Delete(id)
}

// find all items example no filtering
func (r *testIntModelRepository) Find() ([]*TestIntModel, error) {
	//we return all items
	matchFunc := func(interface{}) bool {
		return true
	}
	return r.data.Find([]*TestIntModel{}, matchFunc).([]*TestIntModel), nil
}

// default get by primary key example
func (r *testIntModelRepository) FindById(id int) (*TestIntModel, error) {
	row, err := r.data.FindById(id)
	if row != nil {
		return row.(*TestIntModel), err
	}
	return nil, err
}

// find all that match the name
func (r *testIntModelRepository) FindByName(name string) ([]*TestIntModel, error) {
	//match by name
	matchFunc := func(v interface{}) bool {
		return v.(*TestIntModel).Name == name
	}
	return r.data.Find([]*TestIntModel{}, matchFunc).([]*TestIntModel), nil
}

// find the first that matches the name
func (r *testIntModelRepository) FindFirstByName(name string) (*TestIntModel, error) {
	//match by name
	matchFunc := func(v interface{}) bool {
		return v.(*TestIntModel).Name == name
	}
	res, err := r.data.FindOne(matchFunc)
	if err != nil {
		return nil, err
	}
	return res.(*TestIntModel), nil
}

// save a record to the map
func (r *testIntModelRepository) Save(model *TestIntModel) error {
	return r.data.Save(model.Id, model)
}

func NewTestIntRepository() TestIntRepository {
	return &testIntModelRepository{
		data: datamap.NewInt(errors.New("no result")),
	}
}

func main() {
	repo := NewTestIntRepository()
	repo.Save(&TestIntModel{Id: 1, Name: "test"})
	repo.Save(&TestIntModel{Id: 2, Name: "test 2"})

	findResult, err := repo.Find()
	fmt.Printf("result Find:(%v) error:(%s)\n", findResult, err)

	findByIdResult, err := repo.FindById(2)
	fmt.Printf("result FindById:(%v) error:(%s)\n", findByIdResult, err)

	findByNameResult, err := repo.FindByName("test")
	fmt.Printf("result FindByName:(%v) error:(%s)\n", findByNameResult, err)

	findFirstByNameResult, err := repo.FindFirstByName("test 2")
	fmt.Printf("result FindFirstByName:(%v) error:(%s)\n", findFirstByNameResult, err)
}
