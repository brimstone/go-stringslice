package stringslice

import "testing"

func TestCompare(t *testing.T) {
	fuji := New("Apple")
	mcintosh := New("Apple")
	banana := New("Banana")
	basket := New("Apple")
	basket.Set("Banana")

	if !fuji.Compare(mcintosh) {
		t.Errorf("%v is not %v", fuji, mcintosh)
	}

	if fuji.Compare(banana) {
		t.Errorf("%v is %v", fuji, banana)
	}

	if fuji.String() != "[Apple]" {
		t.Errorf("%v is not %v", fuji.String(), "[Apple]")
	}

	if basket.String() != "[Apple Banana]" {
		t.Errorf("%v is not %v", basket.String(), "[Apple Banana]")
	}

	if fuji.Compare(basket) {
		t.Errorf("%v is %v", fuji, basket)
	}
}
