package concurrent

import (
	"container/heap"
	"sync"
)

type QueueItem struct {
	value interface{}
	priority int
	index int
}

type Queue []*QueueItem

