package concurrent

import (
	"container/list"
	"sync"
)

type List struct {
	l     *list.List
	mutex sync.RWMutex
}


//NewList provides creates list object
func NewList() *List {
	l := new(List)
	l.l = list.New()
	l.mutex = sync.RWMutex{}
	return l
}


//PushBack provides append new element to back
func (l *List) PushBack(v interface{}) *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.PushBack(v)
}

//PushBack provides append new element to front
func (l *List) PushFront(v interface{}) *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.PushFront(v)
}

//Len returns length of this list
func (l *List) Len() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.Len()
}


//Front provides getting element from front
func (l *List) Front() *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.Front()
}


//back provides getting element from back
func (l *List) Back() *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.Back()
}
