package decoding

type MaterialElement struct {
	Name            string        `xml:"name,attr"`
	TextureFilename *string       `xml:"texture"`
	Color           *ColorElement `xml:"color"`
}

func (m *MaterialElement) Clear() {
	m.Name = ""
	m.TextureFilename = nil
	if m.Color != nil {
		m.Color.Clear()
	}
}
