package color

type Color [4]float64

func NewColor(r, g, b, a float64) Color {
	return Color{r, g, b, a}
}

func (c *Color) Clear() {
	*c = Color{0, 0, 0, 0}
}
