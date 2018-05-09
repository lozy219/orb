package orb

import (
	"testing"
)

func TestPointZM(t *testing.T) {
	p := PointZM{1, 2, 3, 4}
	if v := p.Lon(); v != 1 {
		t.Errorf("incorrect lon: %v != 1", v)
	}

	if v := p.Lat(); v != 2 {
		t.Errorf("incorrect lat: %v != 2", v)
	}

	if v := p.M(); v != 4 {
		t.Errorf("incorrect M: %v != 4", v)
	}
}
func TestPointZMEqual(t *testing.T) {
	p1 := PointZM{1, 0, 2, 0}
	p2 := PointZM{1, 0, 2, 0}

	p3 := PointZM{2, 3, 4, 5}
	p4 := PointZM{2, 4, 3, 5}

	if !p1.Equal(p2) {
		t.Errorf("expected: %v == %v", p1, p2)
	}

	if p2.Equal(p3) {
		t.Errorf("expected: %v != %v", p2, p3)
	}

	if p3.Equal(p4) {
		t.Errorf("expected: %v != %v", p3, p4)
	}
}
func TestPointZMPoint(t *testing.T) {
	p1 := PointZM{1, 2, 3, 4}
	p2 := Point{1, 2}

	if !p2.Equal(p1.Point()) {
		t.Errorf("incorrect conversion: %v != %v", p1, p2)
	}
}
