package orb

// A PointZM is a 4-tuple containing lon/lat/altitude/extra information
type PointZM [4]float64

// GeoJSONType returns the GeoJSON type for the object.
func (p PointZM) GeoJSONType() string {
	return "PointZM"
}

// Dimensions returns 0 because a point is a 0d object.
func (p PointZM) Dimensions() int {
	return 0
}

// Bound returns a single point bound of the point.
func (p PointZM) Bound() Bound {
	return Bound{p.Point(), p.Point()}
}

// Point returns itself so it implements the Pointer interface.
func (p PointZM) Point() Point {
	return Point{p.X(), p.Y()}
}

// Y returns the vertical coordinate of the point.
func (p PointZM) Y() float64 {
	return p[1]
}

// X returns the horizontal coordinate of the point.
func (p PointZM) X() float64 {
	return p[0]
}

// Z returns the altitude of the point.
func (p PointZM) Z() float64 {
	return p[2]
}

// M returns the extra data of the point.
func (p PointZM) M() float64 {
	return p[3]
}

// Lat returns the vertical, latitude coordinate of the point.
func (p PointZM) Lat() float64 {
	return p[1]
}

// Lon returns the horizontal, longitude coordinate of the point.
func (p PointZM) Lon() float64 {
	return p[0]
}

// Equal checks if the point represents the same point or vector.
func (p PointZM) Equal(point PointZM) bool {
	return p[0] == point[0] && p[1] == point[1] && p[2] == point[2] && p[3] == point[3]
}
