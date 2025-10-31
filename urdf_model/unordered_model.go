package urdfmodel

import (
	"github.com/WrenchRobotics/urdf-go/urdf_model/joint"
	"github.com/WrenchRobotics/urdf-go/urdf_model/link"
)

type UnorderedModel struct {
	Name      string           `xml:"name,attr"`
	Links     []*link.Link     `xml:"link"`     // Complete list of links, organized by name
	Joints    []*joint.Joint   `xml:"joint"`    // Complete list of joints, organized by name
	Materials []*link.Material `xml:"material"` // Complete list of materials, organized by name
}
