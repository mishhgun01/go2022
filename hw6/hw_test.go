package hw6

import "testing"

func TestGeom_CalculateDistance(t *testing.T) {
	g1 := Geom{X1: -1, X2: 2, Y1: 0, Y2: 3}
	got := g1.CalculateDistance()
	want := -1.0
	if got != want {
		t.Fatalf("want: %+v, got: %+v", want, got)
	}
	g1 = Geom{X1: 1, X2: 0, Y1: 0, Y2: 0}
	got = g1.CalculateDistance()
	want = 1.0
	if got != want {
		t.Fatalf("want: %+v, got: %+v", want, got)
	}
	g1 = Geom{X1: 1, X2: 1, Y1: 1, Y2: 1}
	got = g1.CalculateDistance()
	want = 0.0
	if got != want {
		t.Fatalf("want: %+v, got: %+v", want, got)
	}
}
