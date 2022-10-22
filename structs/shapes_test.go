package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	testArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		testArea(t, rectangle, 100)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10}
		testArea(t, circle, 314.1592653589793)
	})
}
