package link

import (
	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Collision struct {
	Name     string
	Origin   *pose.Pose
	Geometry *geometry.Geometry
}

func (c *Collision) Clear() {
	c.Name = ""
	if c.Origin != nil {
		c.Origin.Clear()
	}
	if c.Geometry != nil {
		(*c.Geometry).Clear()
	}
}
