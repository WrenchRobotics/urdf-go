package geometry

import geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"

type Cylinder struct {
	Radius float64 `xml:"radius,attr"`
	Length float64 `xml:"length,attr"`
}

func (c *Cylinder) Type() geometry_type.GeometryType {
	return geometry_type.Cylinder
}

func (c *Cylinder) Clear() {
	c.Radius = 0
	c.Length = 0
}
