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
