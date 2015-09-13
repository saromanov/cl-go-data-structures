package dict

import
(
	"sync"
	"errors"
)

type Dict struct {
	mutex sync.RWMutex
	data map[interface{}]interface{}
}

func NewDict()*Dict {
	dict := new(Dict)
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
