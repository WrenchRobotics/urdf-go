package joint

type JointMimic struct {
	Offset     float64 `xml:"offset,attr"`
	Multiplier float64 `xml:"multiplier,attr"`
	JointName  string  `xml:"joint,attr"`
}
