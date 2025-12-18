// This package is used to define a form of enum with the internal type JointType.
package joint_type

import "encoding/xml"

/*
A form of an "enum" that defines the many types of joints
that are possible aaccording to the URDF specification.
*/
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

/*
UnmarshalXMLAttr allows for JointType objects to be decoded from
the "attributes" of a tag.

For example, consider the following tag from a gripper URDF:

	<joint name="finger_joint" type="revolute">

It assigns the "revolute" joint type using the attribute `type`.

This should not normally be called by the user. It is used to
implement an interface from the xml package and do efficient decoding
when we use `xml.Unmarshal`.
*/
func (jt *JointType) UnmarshalXMLAttr(attr xml.Attr) error {
	*jt = JointType(attr.Value)
	return nil
}
