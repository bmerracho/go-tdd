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
	areaTests := []struct {
		name     string
		shape    Shape
		hasValue float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasValue: 72.0},
		{name: "Circle", shape: Circle{10}, hasValue: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasValue: 36.0},
	}

	for _, tableTest := range areaTests {
		got := tableTest.shape.Area()
		want := tableTest.hasValue

		t.Run("check the area of shape", func(t *testing.T) {
			if got != want {
				t.Errorf("%#v got %g want %g", tableTest.shape, got, want)
			}
		})

	}
}
