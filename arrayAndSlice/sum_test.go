package arrayAndSlice

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{5, 3, 4, 2, 1}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %q want %q, given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{5, 4, 2, 1}

		got := Sum(numbers)
		want := 12

		if got != want {
			t.Errorf("got %q, want %q given %v", got, want, numbers)
		}
	})
}
