package joint_decoding

type JointDynamicsElement struct {
	Damping  float64 `xml:"damping,attr"`
	Friction float64 `xml:"friction,attr"`
}

func NewJointDynamics(damping, friction float64) JointDynamicsElement {
	return JointDynamicsElement{
		Damping:  damping,
		Friction: friction,
	}
}

func MakeZeroJointDynamics() JointDynamicsElement {
	return JointDynamicsElement{
		Damping:  0.0,
		Friction: 0.0,
	}
}
