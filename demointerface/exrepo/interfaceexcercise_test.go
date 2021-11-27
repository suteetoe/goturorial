package exrepo

import "testing"

func TestCubeVolumes(t *testing.T) {
	var want, given float64

	given = 3
	cube := &cube{
		edge: given,
	}
	want = 27
	if get := VolumeOf(cube); want != get {
		t.Errorf("when give egde %f, want volumes %f but has %f",
			given, want, get)
	}
}

func TestTriangularPrismVolumes(t *testing.T) {
	var want, given float64

	triangularPrism := &triangularPrism{
		base:     10,
		attitude: 10,
		height:   10,
	}
	want = 500
	if get := VolumeOf(triangularPrism); want != get {
		t.Errorf("when give egde %f, want volumes %f but has %f",
			given, want, get)
	}
}
