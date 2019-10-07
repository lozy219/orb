package orb

// LineStringZM represents a set of points to be thought of as a polyline.
type LineStringZM []PointZM

// GeoJSONType returns the GeoJSON type for the object.
func (ls LineStringZM) GeoJSONType() string {
	return "LineStringZM"
}

// Dimensions returns 1 because a LineString is a 1d object.
func (ls LineStringZM) Dimensions() int {
	return 1
}

// Reverse will reverse the line string.
// This is done inplace, ie. it modifies the original data.
func (ls LineStringZM) Reverse() {
	l := len(ls) - 1
	for i := 0; i <= l/2; i++ {
		ls[i], ls[l-i] = ls[l-i], ls[i]
	}
}

// Bound returns a rect around the line string. Uses rectangular coordinates.
func (ls LineStringZM) Bound() Bound {
	return MultiPointZM(ls).Bound()
}

// Equal compares two line strings. Returns true if lengths are the same
// and all points are Equal.
func (ls LineStringZM) Equal(lineString LineStringZM) bool {
	return MultiPointZM(ls).Equal(MultiPointZM(lineString))
}

// Clone returns a new copy of the line string.
func (ls LineStringZM) Clone() LineStringZM {
	ps := MultiPointZM(ls)
	return LineStringZM(ps.Clone())
}
