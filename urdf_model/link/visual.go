package link

import (
	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Visual struct {
	Origin   pose.Pose         `xml:"origin"`
	Geometry geometry.Geometry `xml:"geometry"`
	Name     string
	Material *Material
}

func (v *Visual) Clear() {
	v.Name = ""
	v.Origin.Clear()
	v.Geometry.Clear()
	if v.Material != nil {
		v.Material.Name = ""
		v.Material.TextureFilename = ""
		if v.Material.Color != nil {
			colorPtr := v.Material.Color
			colorPtr.Clear()
		}
	}
}
