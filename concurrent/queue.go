package concurrent

import (
	"sync"
	"sync/atomic"
	"errors"
)

type QueueItem struct {
	value interface{}
	priority int
	index int
}

type Queue struct {
	item []*QueueItem
	count int32
	mutex* sync.RWMutex
}

func NewQueue()*Queue {
	q := new(Queue)
	q.item = []*QueueItem{}
	q.count = 0
	q.mutex = &sync.RWMutex{}
	return q
}

func (q* Queue) Insert(elem interface{}) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	q.item = append(q.item, &QueueItem{value:elem})
	atomic.AddInt32(&q.count, 1)
}

func (q *Queue) Pop()(interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("Queue is empty")
	}
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	defer atomic.AddInt32(&q.count, -1)
	elem := q.item[0].value
	q.item = append(q.item[:0], q.item[1:]...)
	return elem, nil

}

func (q * Queue) Size() int32 {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	value := atomic.LoadInt32(&q.count)
	return value
}

