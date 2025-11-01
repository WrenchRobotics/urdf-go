package link

import (
	"fmt"

	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/common/pose"
	"github.com/WrenchRobotics/urdf-go/decoding"
)

type Visual struct {
	Origin   *pose.Pose
	Geometry geometry.GeometryImplementation
	Name     string
	Material *decoding.MaterialElement
}

func (v *Visual) Clear() {
	if v.Origin != nil {
		v.Origin.Clear()
	}
	if v.Geometry != nil {
		v.Geometry.Clear()
	}
	if v.Material != nil {
		v.Material.Clear()
	}
	v.Name = ""
}

func (v *Visual) FromDecodingElement(ve *decoding.VisualElement) error {
	if ve == nil {
		return fmt.Errorf(
			"the provided visual element pointer was nil",
		)
	}
	v.Name = ve.Name
	if ve.Origin != nil {
		v.Origin = ve.Origin
	}
	if ve.Geometry != nil {
		v.Geometry = ve.Geometry.GetActiveImplementation()
	}
	if ve.Material != nil {
		v.Material = ve.Material
	}

	return nil
}
