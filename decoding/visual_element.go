package decoding

import (
	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/common/pose"
)

type VisualElement struct {
	Origin   *pose.Pose         `xml:"origin"`
	Geometry *geometry.Geometry `xml:"geometry"`
	Name     string             `xml:"name,attr"`
	Material *MaterialElement   `xml:"material"`
}

func (ve *VisualElement) Clear() {
	if ve.Origin != nil {
		ve.Origin.Clear()
	}
	if ve.Geometry != nil {
		ve.Geometry.Clear()
	}
	if ve.Material != nil {
		ve.Material.Clear()
	}
	ve.Name = ""
}
