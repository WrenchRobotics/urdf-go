package pose

type Pose struct {
	Position Vector3
	Rotation Rotation
}

func (p *Pose) Clear() {
	p.Position.Clear()
	p.Rotation.Clear()
}
