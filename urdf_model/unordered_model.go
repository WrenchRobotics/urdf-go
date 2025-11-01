package urdfmodel

import (
	"github.com/WrenchRobotics/urdf-go/decoding"
	"github.com/WrenchRobotics/urdf-go/urdf_model/joint"
	"github.com/WrenchRobotics/urdf-go/urdf_model/link"
)

type UnorderedModel struct {
	Name      string                      `xml:"name,attr"`
	Links     []*link.Link                `xml:"link"`     // Complete list of links, organized by name
	Joints    []*joint.Joint              `xml:"joint"`    // Complete list of joints, organized by name
	Materials []*decoding.MaterialElement `xml:"material"` // Complete list of materials, organized by name
}
