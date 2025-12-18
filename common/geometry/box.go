package geometry

import (
	geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"
	"github.com/WrenchRobotics/urdf-go/common/pose"
)

/*
Box is an object that specifies a box with three lengths
(length, width, height) as is assumed in the URDF specification.

You should be able to parse any box in a URDF, like this one:

	<box size="0.022 0.00635 0.0375"/>
*/
type Box struct {
	Dimensions pose.Vector3 `xml:"size,attr"`
}

/*
Type always returns `geometry_type.Box`.

This method is defined in order to implement (in other words,
satisfy the requirements of) the `GeometryImplementation` interface in this package.
*/
func (b *Box) Type() geometry_type.GeometryType {
	return geometry_type.Box
}

/*
Clears all internal variables for Box (only the Dimensions).

This method is defined in order to implement (in other words,
satisfy the requirements of) the `GeometryImplementation` interface in this package.
*/
func (b *Box) Clear() {
	b.Dimensions.Clear()
}
