package link

import "github.com/WrenchRobotics/urdf-go/urdf_model/pose"

type Inertial struct {
	Origin *pose.Pose
	Mass   float64
	Ixx    float64
	Ixy    float64
	Ixz    float64
	Iyy    float64
	Iyz    float64
	Izz    float64
}

func (i *Inertial) Clear() {
	i.Origin.Clear()
	i.Mass = 0
	i.Ixx = 0
	i.Ixy = 0
	i.Ixz = 0
	i.Iyy = 0
	i.Iyz = 0
	i.Izz = 0
}
