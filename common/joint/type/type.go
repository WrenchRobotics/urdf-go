package joint_type

import "encoding/xml"

type JointType string

const (
	RevoluteJoint   JointType = "revolute"
	ContinuousJoint JointType = "continuous"
	PrismaticJoint  JointType = "prismatic"
	FixedJoint      JointType = "fixed"
	FloatingJoint   JointType = "floating"
	PlanarJoint     JointType = "planar"
	UnknownJoint    JointType = "unknown"
)

func (jt *JointType) UnmarshalXMLAttr(attr xml.Attr) error {
	*jt = JointType(attr.Value)
	return nil
}
