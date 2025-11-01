package link

import (
	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/common/pose"
)

type Collision struct {
	Name     string
	Origin   *pose.Pose
	Geometry geometry.GeometryImplementation
}

func (c *Collision) Clear() {
	c.Name = ""
	c.Origin.Clear()
	c.Geometry.Clear()
}
