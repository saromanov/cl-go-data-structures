package dict

import
(
	"testing"
)

func constructDict()*Dict {
	d := New()
	d.Add("A", "B")
	d.Add("C", "A")
	d.Add("V", "E")
	d.Add("Q", "C")
	d.Add("Z", "P")
	return d

}

func TestGet(t *testing.T) {
	d := constructDict()
	value, err := d.Get("A")
	if err != nil {
		t.Errorf("%v", err)
	}
	if value == nil {
		t.Errorf("Expected value is %s", "B")
	}
}

func TestRemove(t *testing.T) {
	d := constructDict()
	d.Remove("C")
	_, err := d.Get("C")
	if err == nil {
		t.Errorf("Element %s must be deleted", "C")
	}
}

func TestSize(t *testing.T) {
	d := constructDict()
	if d.Size() != 5 {
		t.Errorf("Expected size of the dice is %d", d.Size())
	}
}