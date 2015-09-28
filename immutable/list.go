package immutable

import (
	"container/list"
	"sync"
)

type List struct {
	l     *list.List
	mutex *sync.RWMutex
}

//NewList provides creates list object
func NewList() *List {
	l := new(List)
	l.l = list.New()
	l.mutex = &sync.RWMutex{}
	return l
}

//PushBack provides append new element to back and returns new list
func (l *List) PushBack(v interface{}) *List {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	newlist := NewList()
	newlist = l
	newlist.l.PushBack(v)
	return newlist
}

//PushBack provides append new element to back and returns new list
func (l *List) PushFront(v interface{}) *List {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	newlist := NewList()
	newlist = l
	newlist.l.PushFront(v)
	return newlist
}

//Front provides skipping element from front
func (l *List) Front() *List {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	newlist := NewList()
	newlist = l
	newlist.l.Front()
	return newlist
}

//Front provides skipping element from front
func (l *List) Back() *List {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	newlist := NewList()
	newlist = l
	newlist.l.Back()
	return newlist
}

//Front provides skipping element from front
func (l *List) HeadFront() *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.Back()
}

//Front provides skipping element from front
func (l *List) HeadBack() *list.Element {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return l.l.Front()
}
