package concurrent

import (
	"container/list"
	"sync"
)

type List struct {
	l     *list.List
	mutex sync.RWMutex
}

func NewList() *List {
	l := new(List)
	l.l = list.New()
	l.mutex = sync.RWMutex{}
	return l
}

func (l *List) PushBack(v interface{}) *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.PushBack(v)
}

func (l *List) PushFront(v interface{}) *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.PushFront(v)
}

func (l *List) Len() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.Len()
}
