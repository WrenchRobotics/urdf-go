package geometry

import geometry_type "github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry/type"

type Cylinder struct {
	Radius float64
	Length float64
}

func (c *Cylinder) Type() geometry_type.GeometryType {
	return geometry_type.Cylinder
}

func (c *Cylinder) Clear() {
	c.Radius = 0
	c.Length = 0
}
