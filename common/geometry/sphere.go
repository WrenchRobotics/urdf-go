package geometry

import geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"

type Sphere struct {
	Radius float64
}

func (s *Sphere) Type() geometry_type.GeometryType {
	return geometry_type.Sphere
}

func (s *Sphere) Clear() {
	s.Radius = 0
}
