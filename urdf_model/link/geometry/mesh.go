package geometry

import (
	geometry_type "github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry/type"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Mesh struct {
	Filename string       `xml:"filename,attr"`
	Scale    pose.Vector3 `xml:"scale,attr"`
}

func (m *Mesh) Type() geometry_type.GeometryType {
	return geometry_type.Mesh
}

func (m *Mesh) Clear() {
	m.Filename = ""
	m.Scale.Clear()
}
