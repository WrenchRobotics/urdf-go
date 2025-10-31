package joint

import "github.com/WrenchRobotics/urdf-go/urdf_model/pose"

type JointAxis struct {
	Direction pose.Vector3 `xml:"xyz,attr"`
}

func (ja *JointAxis) Clear() {
	ja.Direction.Clear()
}
