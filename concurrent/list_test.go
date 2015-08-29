package concurrent

import
(
	"testing"
	"fmt"
)

func TestPush(t *testing.T) {
	ls := NewList()
	ls.PushBack("A")
	ls.PushBack("B")
	ls.PushBack("C")
	l := ls.Len()
	if l != 3 {
		t.Errorf(fmt.Sprintf("TestPush. Expected length %d, found %d", 3, l))
	}
}

func TestFront(t *testing.T) {
	ls := NewList()
	ls.PushBack("A")
	ls.PushBack("B")
	ls.PushBack("C")
	fr := ls.Front().Value
	if fr != "A" {
		t.Errorf(fmt.Sprintf("TestFront. Expected value %s, found %s", "A", fr))
	}
}