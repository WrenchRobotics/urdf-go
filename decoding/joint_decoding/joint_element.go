package joint_decoding

import (
	"github.com/WrenchRobotics/urdf-go/common"
	joint_type "github.com/WrenchRobotics/urdf-go/common/joint/type"
	"github.com/WrenchRobotics/urdf-go/common/pose"
)

type JointElement struct {
	Name          string               `xml:"name,attr"`
	Type          joint_type.JointType `xml:"type,attr"`
	ChildLinkRef  common.LinkReference `xml:"child"`
	ParentLinkRef common.LinkReference `xml:"parent"`
	// Optional elements
	Origin   *pose.Pose            `xml:"origin"`
	Axis     *JointAxisElement     `xml:"axis"`
	Dynamics *JointDynamicsElement `xml:"dynamics"`
	Limits   *JointLimitsElement   `xml:"limit"`
	Mimic    *JointMimicElement    `xml:"mimic"`
}
