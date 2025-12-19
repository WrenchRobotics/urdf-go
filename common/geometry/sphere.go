package geometry

import geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"

/*
Sphere is an object that exactly represents a sphere according
to the URDF specification.

It uses the xml struct tags so that you can parse URDF
elements without needing additional code/functions.

You should be able to parse any cylinder of the form:

	<sphere radius="0.1" />
*/
type Sphere struct {
	// Radius of the sphere
	Radius float64 `xml:"radius,attr"`
}

/*
Always returns `geometry_type.Sphere`.

This method is defined in order to implement (in other words,
satisfy the requirements of) the `GeometryImplementation` interface in this package.
*/
func (s *Sphere) Type() geometry_type.GeometryType {
	return geometry_type.Sphere
}

/*
Clears all internal variables for this type (only Radius).

This method is defined in order to implement (in other words,
satisfy the requirements of) the `GeometryImplementation` interface in this package.
*/
func (s *Sphere) Clear() {
	s.Radius = 0
}
