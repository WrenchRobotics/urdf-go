package geometry

import geometry_type "github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry/type"

type GeometryImplementation interface {
	Type() geometry_type.GeometryType
	Clear()
}
