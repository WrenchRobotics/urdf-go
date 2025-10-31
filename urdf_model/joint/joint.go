package joint

import (
	joint_type "github.com/WrenchRobotics/urdf-go/urdf_model/joint/type"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Joint struct {
	Name                         string               `xml:"name,attr"`
	Type                         joint_type.JointType `xml:"type,attr"`
	ChildLinkRef                 JointLinkReference   `xml:"child"`
	ParentLinkRef                JointLinkReference   `xml:"parent"`
	ParentToJointOriginTransform pose.Pose            `xml:"origin"`
	Axis                         JointAxis            `xml:"axis"`
	Dynamics                     *JointDynamics       `xml:"dynamics"`
	Limits                       *JointLimits         `xml:"limit"`
	Mimic                        *JointMimic          `xml:"mimic"`
}
