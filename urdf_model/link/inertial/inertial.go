package inertial

import "github.com/WrenchRobotics/urdf-go/urdf_model/pose"

type Inertial struct {
	Origin  pose.Pose `xml:"origin"`
	Mass    Mass      `xml:"mass"`
	Inertia Inertia   `xml:"inertia"`
}

func (i *Inertial) Clear() {
	i.Origin.Clear()
	i.Mass.Clear()
	i.Inertia.Clear()
}
