package joint_decoding

type JointLimitsElement struct {
	Lower    float64 `xml:"lower,attr"`
	Upper    float64 `xml:"upper,attr"`
	Effort   float64 `xml:"effort,attr"`
	Velocity float64 `xml:"velocity,attr"`
}
