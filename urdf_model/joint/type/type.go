package joint_type

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
