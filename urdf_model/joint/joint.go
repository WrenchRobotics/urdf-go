package joint

import (
	"fmt"

	joint_type "github.com/WrenchRobotics/urdf-go/common/joint/type"
	"github.com/WrenchRobotics/urdf-go/common/pose"
	"github.com/WrenchRobotics/urdf-go/decoding/joint_decoding"
)

type Joint struct {
	Name                         string
	Type                         joint_type.JointType
	ChildLinkName                string
	ParentLinkName               string
	ParentToJointOriginTransform *pose.Pose
	Axis                         *pose.Vector3
	Dynamics                     *joint_decoding.JointDynamicsElement
	Limits                       *joint_decoding.JointLimitsElement
	Mimic                        *joint_decoding.JointMimicElement
}

func (j *Joint) FromDecodingElement(je *joint_decoding.JointElement) error {
	if je == nil {
		return fmt.Errorf(
			"the provided joint element pointer was nil",
		)
	}
	j.Name = je.Name
	j.Type = joint_type.JointType(je.Type)

	j.ChildLinkName = je.ChildLinkRef.LinkName
	j.ParentLinkName = je.ParentLinkRef.LinkName
	if je.Origin != nil {
		if j.ParentToJointOriginTransform == nil {
			j.ParentToJointOriginTransform = &pose.Pose{}
		}
		j.ParentToJointOriginTransform = je.Origin
	}
	if je.Axis != nil {
		if j.Axis == nil {
			j.Axis = &pose.Vector3{}
		}
		j.Axis = &je.Axis.Direction
	}
	j.Dynamics = je.Dynamics
	j.Limits = je.Limits
	j.Mimic = je.Mimic

	return nil
}
