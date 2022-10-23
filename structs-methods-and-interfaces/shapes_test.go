package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{Width: 10.0, Height: 10.0}
		got := r.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{Radius: 10.0}
		got := c.Perimeter()
		want := 2 * math.Pi * c.Radius

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, hasArea: 100},
		{name: "Circle", shape: Circle{10}, hasArea: math.Pi * 10 * 10},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
