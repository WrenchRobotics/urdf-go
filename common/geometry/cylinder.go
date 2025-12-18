package geometry

import geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"

/*
Cylinder is an object which exactly represents a Cylinder object
according to the URDF specification.

It uses the xml encoding decorators so that you can parse xml snippets without adding more code.

You should be able to parse any cylinder. For example, this:

	<cylinder radius="0.1" length="1.0" />
*/
type Cylinder struct {
	Radius float64 `xml:"radius,attr"`
	Length float64 `xml:"length,attr"`
}

/*
Always returns `geometry_type.Cylinder`.

This method is defined in order to implement (in other words,
satisfy the requirements of) the `GeometryImplementation` interface in this package.
*/
func (c *Cylinder) Type() geometry_type.GeometryType {
	return geometry_type.Cylinder
}

/*
Clears the internal variables of a Cylinder object (Radius and Length),

This method is defined in order to implement (in other words,
satisfy the requirements of) the `GeometryImplementation` interface in this package.
*/
func (c *Cylinder) Clear() {
	c.Radius = 0
	c.Length = 0
}
