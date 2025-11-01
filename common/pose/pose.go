package pose

type Pose struct {
	Position Vector3  `xml:"xyz,attr"`
	Rotation Rotation `xml:"rpy,attr"`
}

func (p *Pose) Clear() {
	p.Position.Clear()
	p.Rotation.Clear()
}
