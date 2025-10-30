package link

import (
	"github.com/WrenchRobotics/urdf-go/urdf_model/color"
)

type Material struct {
	Name            string
	TextureFilename string
	Color           *color.Color
}

func (m *Material) Clear() {
	m.Name = ""
	m.TextureFilename = ""
	if m.Color != nil {
		m.Color.Clear()
	}
}
