package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	numbers := [7]int{5, 9, 4, 3, 7, 0, 5}

	got := Sum(numbers)
	expected := 33

	if got != expected {
		t.Errorf("expected %d got %d, given %v", expected, got, numbers)
	} else {
		t.Logf("expected %d got %d, given %v", expected, got, numbers)
	}
}