package decoding

import (
	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/common/pose"
)

type CollisionElement struct {
	Name     string            `xml:"name,attr"`
	Origin   *pose.Pose        `xml:"origin"`
	Geometry geometry.Geometry `xml:"geometry"`
}

func (c *CollisionElement) Clear() {
	c.Name = ""
	c.Origin.Clear()
	c.Geometry.Clear()
}
