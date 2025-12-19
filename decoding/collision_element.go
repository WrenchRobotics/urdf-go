// Package for decoding all elements from URDF files (a specific format of XML files used in robotics applications).
package decoding

import (
	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/common/pose"
)

/*
An object which a Collision element from the
URDF specification, with its optional and required fields.

It uses the struct tags so that you can parse or marshal xml snippets without adding more code.

For example, this:

	<collision name="my_collision">
	  <origin xyz="0 0 0" rpy="0 0 0"/>
	  <geometry>
	    <box size="1 1 1"/>
	  </geometry>
	</collision>

Can be parsed into a CollisionElement object.

Note that the Origin field is a pointer to a Pose object.
This is because the Origin element is optional in the URDF specification.
If the Origin element is not present in the XML being parsed,
the Origin field will be nil.

Also note that the Geometry field is a struct with all optional fields.
See the comments for the geometry.Geometry struct for more information.
*/
type CollisionElement struct {
	Name     string            `xml:"name,attr"`
	Origin   *pose.Pose        `xml:"origin"`
	Geometry geometry.Geometry `xml:"geometry"`
}

/*
Clears the internal variables of a CollisionElement object (Name, Origin, and Geometry).
*/
func (c *CollisionElement) Clear() {
	c.Name = ""
	c.Origin.Clear()
	c.Geometry.Clear()
}
