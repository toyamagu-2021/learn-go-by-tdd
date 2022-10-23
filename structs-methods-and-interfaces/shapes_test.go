package shapes

import (
	"math"
	"testing"
)

// func TestPerimeter(t *testing.T) {
// 	areaTests := []struct {
// 		name         string
// 		shape        Shape
// 		hasPerimeter float64
// 	}{
// 		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, hasPerimeter: 40.0},
// 		{name: "Circle", shape: Circle{10}, hasPerimeter: 2 * math.Pi * 10},
// 		// {name: "Triangle", shape: Triangle{12, 6}, hasPerimeter: ?},
// 	}

// 	for _, tt := range areaTests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := tt.shape.Perimeter()
// 			if got != tt.hasPerimeter {
// 				t.Errorf("%#v got %.2f want %.2g", tt.shape, got, tt.hasPerimeter)
// 			}
// 		})

// 	}
// }

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
