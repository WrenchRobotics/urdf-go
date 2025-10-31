package inertial

type Mass struct {
	Value float64 `xml:"value,attr"`
}

func (m *Mass) Clear() {
	m.Value = 0
}
