package link

import (
	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Visual struct {
	Origin   pose.Pose         `xml:"origin"`
	Geometry geometry.Geometry `xml:"geometry"`
	Name     string            `xml:"name,attr"`
	Material Material          `xml:"material"`
}

func (v *Visual) Clear() {
	v.Name = ""
	v.Origin.Clear()
	v.Geometry.Clear()
	v.Material.Clear()
}
