package geometry

import geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"

/*
Defines the minimum requirements for each "type" of geometry
that is allowed to define a geometry tag in a URDF.
*/
type GeometryImplementation interface {
	Type() geometry_type.GeometryType
	// Clears the internal variables inside of the current implementation.
	Clear()
}
