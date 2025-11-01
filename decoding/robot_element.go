package decoding

import "github.com/WrenchRobotics/urdf-go/decoding/joint_decoding"

type RobotElement struct {
	Name   string                         `xml:"name,attr"`
	Links  []*LinkElement                 `xml:"link"`
	Joints []*joint_decoding.JointElement `xml:"joint"`
}
