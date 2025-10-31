package joint

type JointDynamics struct {
	Damping  float64 `xml:"damping,attr"`
	Friction float64 `xml:"friction,attr"`
}

func NewJointDynamics(damping, friction float64) JointDynamics {
	return JointDynamics{
		Damping:  damping,
		Friction: friction,
	}
}

func MakeZeroJointDynamics() JointDynamics {
	return JointDynamics{
		Damping:  0.0,
		Friction: 0.0,
	}
}
