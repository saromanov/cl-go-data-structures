package immutable

import
(
	"testing"
)

func TestHeadFront(t *testing.T) {
	lst := NewList()
	elem := lst.PushBack("A")
	elem2 := elem.PushBack("B")
	value := elem2.HeadFront().(string)
	if value != "B" {
		t.Errorf("Found %s, expected %s", value, "B")
	}
}

func TestHeadBack(t *testing.T) {
	lst := NewList()
	elem := lst.PushBack("A")
	elem2 := elem.PushBack("B")
	value := elem2.HeadFront().(string)
	if value != "A" {
		t.Errorf("Found %s, expected %s", value, "A")
	}
}