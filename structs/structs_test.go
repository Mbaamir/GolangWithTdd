package structs

import (
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectange", shape: Rectangle{5, 2.5}, hasArea: 12.5},
		{name: "circle", shape: Circle{2.5}, hasArea: 19.634954084936208},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
