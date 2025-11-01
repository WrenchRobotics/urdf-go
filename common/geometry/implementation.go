package geometry

import geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"

type GeometryImplementation interface {
	Type() geometry_type.GeometryType
	Clear()
}
