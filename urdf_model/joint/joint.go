package joint

import (
	"github.com/WrenchRobotics/urdf-go/common"
	joint_type "github.com/WrenchRobotics/urdf-go/common/joint/type"
	"github.com/WrenchRobotics/urdf-go/common/pose"
	"github.com/WrenchRobotics/urdf-go/decoding/joint_decoding"
)

type Joint struct {
	Name                         string
	Type                         joint_type.JointType
	ChildLinkRef                 *common.LinkReference
	ParentLinkRef                *common.LinkReference
	ParentToJointOriginTransform *pose.Pose
	Axis                         *pose.Vector3
	Dynamics                     *joint_decoding.JointDynamicsElement
	Limits                       *joint_decoding.JointLimitsElement
	Mimic                        *joint_decoding.JointMimicElement
}
