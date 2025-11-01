package joint_decoding

import "github.com/WrenchRobotics/urdf-go/common/pose"

type JointAxisElement struct {
	Direction pose.Vector3 `xml:"xyz,attr"`
}

func (ja *JointAxisElement) Clear() {
	ja.Direction.Clear()
}
