package orb

// A MultiPoint represents a set of points in the 2D Eucledian or Cartesian plane.
type MultiPointZM []PointZM

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPointZM) GeoJSONType() string {
	return "MultiPointZM"
}

// Dimensions returns 0 because a MultiPoint is a 0d object.
func (mp MultiPointZM) Dimensions() int {
	return 0
}

// Clone returns a new copy of the points.
func (mp MultiPointZM) Clone() MultiPointZM {
	if mp == nil {
		return nil
	}

	points := make([]PointZM, len(mp))
	copy(points, mp)

	return MultiPointZM(points)
}

// Bound returns a bound around the points. Uses rectangular coordinates.
func (mp MultiPointZM) Bound() Bound {
	if len(mp) == 0 {
		return emptyBound
	}

	b := Bound{mp[0].Point(), mp[0].Point()}
	for _, p := range mp {
		b = b.Extend(p.Point())
	}

	return b
}

// Equal compares two MultiPoint objects. Returns true if lengths are the same
// and all points are Equal, and in the same order.
func (mp MultiPointZM) Equal(multiPoint MultiPointZM) bool {
	if len(mp) != len(multiPoint) {
		return false
	}

	for i := range mp {
		if !mp[i].Equal(multiPoint[i]) {
			return false
		}
	}

	return true
}
