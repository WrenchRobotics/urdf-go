package link

import (
	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Visual struct {
	Origin   *pose.Pose
	Geometry *geometry.Geometry
	Name     string
	Material *Material
}

func (v *Visual) Clear() {
	v.Name = ""
	if v.Origin != nil {
		v.Origin.Clear()
	}
	if v.Geometry != nil {
		(*v.Geometry).Clear()
	}
	if v.Material != nil {
		v.Material.Name = ""
		v.Material.TextureFilename = ""
		if v.Material.Color != nil {
			colorPtr := v.Material.Color
			colorPtr.Clear()
		}
	}
}
