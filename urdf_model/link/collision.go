package link

import (
	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Collision struct {
	Name     string            `xml:"name,attr"`
	Origin   pose.Pose         `xml:"origin"`
	Geometry geometry.Geometry `xml:"geometry"`
}

func (c *Collision) Clear() {
	c.Name = ""
	c.Origin.Clear()
	c.Geometry.Clear()
}
