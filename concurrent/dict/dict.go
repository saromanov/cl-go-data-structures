package dict

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type Dict struct {
	mutex sync.RWMutex
	data  map[interface{}]interface{}
}

func New() *Dict {
	dict := new(Dict)
	dict.data = map[interface{}]interface{}{}
	dict.mutex = sync.RWMutex{}
	return dict
}

//Add provides append data to the dict
func (d *Dict) Add(key, value interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.data[key] = value
}

//Remove provides removing data from the dict
func (d *Dict) Remove(key interface{}) error {
	d.mutex.Lock()
	item, ok := d.data[key]
	if !ok {
		return errors.New("Element by key is not found")
	}

	delete(d.data, item)
	d.mutex.Unlock()
	return nil
}

//Exists returns true if dict contains key and false in otherwise
func (d *Dict) Exists(key interface{}) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, ok := d.data[key]
	return ok
}

//Get provides getting element by the key
func (d *Dict) Get(key interface{}) (interface{}, error) {
	if d.Exists(key) {
		return d.data[key], nil
	}

	return nil, errors.New("Element is not found")
}

//RandomGet provides getting value by random key
func (d *Dict) RandomGet() interface{} {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	rand.Seed(time.Now().UTC().UnixNano())
	item := rand.Intn(d.Size())
	counter := 0
	for _, value := range d.data {
		if counter == item {
			return value
		}
		counter++
	}

	return nil
}

//Size return current size of the dict
func (d *Dict) Size() int {
	return len(d.data)
}
