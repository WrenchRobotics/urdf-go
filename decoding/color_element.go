package decoding

import "github.com/WrenchRobotics/urdf-go/common/color"

type ColorElement struct {
	Color color.Color `xml:"rgba,attr"`
}

func (c *ColorElement) Clear() {
	c.Color.Clear()
}
