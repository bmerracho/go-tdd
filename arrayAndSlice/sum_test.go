package arrayAndSlice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{5, 3, 4, 2, 1}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 3})
	want := []int{3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestSumAllTrail(t *testing.T) {

	checkSum := func(tb testing.TB, got, want []int) {
		tb.Helper()

		if !reflect.DeepEqual(got, want) {
			tb.Errorf("want %v got %v", want, got)
		}
	}

	t.Run("make the sums of trail", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2}, []int{5, 3})
		want := []int{2, 3}

		checkSum(t, got, want)

	})
	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{5, 3})
		want := []int{0, 3}

		checkSum(t, got, want)
	})
}
