package structmethodinterface

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	cheackArea := func(tb testing.TB, got, want float64) {
		t.Helper()

		if got != want {
			tb.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Get area of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		cheackArea(t, rectangle.Area(), 72.0)
	})

	t.Run("Get area of Circle", func(t *testing.T) {
		circle := Circle{10}
		cheackArea(t, circle.Area(), 314.1592653589793)
	})
}
