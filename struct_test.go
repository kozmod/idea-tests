package idea_tests

import "testing"

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

func (s *Square) area() float64 {
	return s.side * s.side
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func TestShape(t *testing.T) {
	square := Square{5}
	if 25 != square.area() {
		t.Errorf("ex = %d; actual = %f", 25, square.area())
	}
}
