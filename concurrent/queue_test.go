package concurrent

import
(
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	q := NewQueue()
	q.Insert(1)
	q.Insert(2)
	q.Insert(3)
	q.Insert(4)
	if q.Size() != 4 {
		t.Errorf(fmt.Sprintf("TestInsert. Expected %d, found %d", 4, q.Size()))
	}
}

func TestPop(t *testing.T) {
	q := NewQueue()
	q.Insert(1)
	q.Insert(2)
	q.Insert(3)
	q.Insert(4)
	item, err := q.Pop()
	if err != nil {
		t.Errorf("TestPop. Found error")
	}
	if item.(int) != 1 {
		t.Errorf(fmt.Sprintf("TestPop. Expected %d, found %d", 1, item))
	}

}