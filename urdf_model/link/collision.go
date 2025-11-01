package link

import (
	"fmt"

	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/common/pose"
	"github.com/WrenchRobotics/urdf-go/decoding"
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

func (c *Collision) FromDecodingElement(ce *decoding.CollisionElement) error {
	if ce == nil {
		return fmt.Errorf(
			"the provided collision element pointer was nil",
		)
	}
	c.Name = ce.Name
	if ce.Origin != nil {
		if c.Origin == nil {
			c.Origin = &pose.Pose{}
		}
		c.Origin = ce.Origin
	}

	c.Geometry = ce.Geometry.GetActiveImplementation()

	return nil
}
