package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{
		width:  10.0,
		height: 10.0,
	}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAria(t *testing.T) {
	areaTables := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, table := range areaTables {
		got := table.shape.Area()
		if got != table.want {
			t.Errorf("got %g want %g", got, table.want)
		}
	}
}
