package decoding

import "github.com/WrenchRobotics/urdf-go/common/inertial"

type LinkElement struct {
	Name      string              `xml:"name,attr"`
	Visual    []*VisualElement    `xml:"visual"`
	Collision []*CollisionElement `xml:"collision"`
	Inertial  *inertial.Inertial  `xml:"inertial"`
}
