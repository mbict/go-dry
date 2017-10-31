package datamap

import (
	"github.com/jinzhu/copier"
	"reflect"
	"sync"
)

type Datamap struct {
	data          map[string]interface{}
	notFoundError error
	mu            sync.RWMutex
}

func (r *Datamap) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return r.notFoundError
	}

	delete(r.data, id)
	return nil
}

func (r *Datamap) FindById(id string) (interface{}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if u, ok := r.data[id]; ok {
		rowCopy := reflect.New(reflect.TypeOf(u).Elem()).Interface()
		err := copier.Copy(rowCopy, u)
		return rowCopy, err
	}
	return nil, r.notFoundError
}

func (r *Datamap) Save(id string, row interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[id] = row

	return nil
}

func (r *Datamap) FindOne(fn func(row interface{}) bool) (interface{}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, v := range r.data {
		if fn(v) == true {
			rowCopy := reflect.New(reflect.TypeOf(v).Elem()).Interface()
			err := copier.Copy(rowCopy, v)
			return rowCopy, err
		}
	}
	return nil, r.notFoundError
}

func (r *Datamap) Find(dst interface{}, fn func(row interface{}) bool) interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Slice {
		panic("not a slice")
	}

	for _, v := range r.data {
		if fn(v) == true {
			rowCopy := reflect.New(reflect.TypeOf(v).Elem()).Interface()
			copier.Copy(rowCopy, v)
			rv = reflect.Append(rv, reflect.ValueOf(rowCopy))
		}
	}
	return rv.Interface()
}

func New(notFoundError error) *Datamap {
	return &Datamap{
		data:          make(map[string]interface{}),
		notFoundError: notFoundError,
	}
}
