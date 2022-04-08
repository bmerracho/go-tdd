package arrayAndSlice

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{5, 3, 4, 2, 1}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
