package joint

import (
	joint_type "github.com/WrenchRobotics/urdf-go/urdf_model/joint/type"
	"github.com/WrenchRobotics/urdf-go/urdf_model/pose"
)

type Joint struct {
	Name                         string
	Type                         joint_type.JointType
	ChildLinkName                string
	ParentLinkName               string
	ParentToJointOriginTransform pose.Pose
	Dynamics                     *JointDynamics
	Limits                       *JointLimits
	Mimic                        *JointMimic
}
