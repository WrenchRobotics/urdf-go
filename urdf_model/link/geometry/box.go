package geometry

import (
	geometry_type "github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry/type"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Box struct {
	Dimensions pose.Vector3 `xml:"size,attr"`
}

func (b *Box) Type() geometry_type.GeometryType {
	return geometry_type.Box
}

func (b *Box) Clear() {
	b.Dimensions.Clear()
}
