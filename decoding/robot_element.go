package decoding

import (
	"encoding/xml"

	"github.com/WrenchRobotics/urdf-go/decoding/joint_decoding"
	"github.com/WrenchRobotics/urdf-go/decoding/transmission_decoding"
)

type RobotElement struct {
	XMLName       xml.Name                                     `xml:""`
	Name          string                                       `xml:"name,attr"`
	Links         []*LinkElement                               `xml:"link"`
	Joints        []*joint_decoding.JointElement               `xml:"joint"`
	Transmissions []*transmission_decoding.TransmissionElement `xml:"transmission"`
}
