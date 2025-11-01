package inertial

type Inertia struct {
	Ixx float64 `xml:"ixx,attr"`
	Ixy float64 `xml:"ixy,attr"`
	Ixz float64 `xml:"ixz,attr"`
	Iyy float64 `xml:"iyy,attr"`
	Iyz float64 `xml:"iyz,attr"`
	Izz float64 `xml:"izz,attr"`
}

func (in *Inertia) Clear() {
	in.Ixx = 0
	in.Ixy = 0
	in.Ixz = 0
	in.Iyy = 0
	in.Iyz = 0
	in.Izz = 0
}
