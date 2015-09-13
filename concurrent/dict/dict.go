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

func (d *Dict) Add(key, value interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.data[key] = value
}


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
