package color

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

/*
Color is a representation of the Red, Green, Blue and
A (?, really the transparency) of a color in a URDF model.
*/
type Color [4]float64

/*
NewColor uses the four components of a color
(Red, Green, Blue, A?) to define a Color object.
*/
func NewColor(r, g, b, a float64) Color {
	return Color{r, g, b, a}
}

/*
Clear resets the color to have zero of each color
component and sets the color to full transparency (a=0).
*/
func (c *Color) Clear() {
	*c = Color{0, 0, 0, 0}
}

/*
UnmarshalXMLAttr is used to decode the `color` attribute in
an XML element's tag and produces a Color value.

This method is typically not called directly. It is used in
URDF decoding when calling xml.Unmarshal() on an object that contains
fields with type Color or *Color.
*/
func (c *Color) UnmarshalXMLAttr(attr xml.Attr) error {
	// Split value into three parts by the "space"
	valueWithoutBrackets := strings.ReplaceAll(attr.Value, "[", "")
	valueWithoutBrackets = strings.ReplaceAll(valueWithoutBrackets, "]", "")
	values := strings.Split(valueWithoutBrackets, " ")

	// Check that there are four elements
	if len(values) != 4 {
		return fmt.Errorf(
			"there was a problem parsing \"%v\" as a Color. Expected there to be 3 spaces, but there were not!",
			values,
		)
	}

	// Extract each element
	var err error
	for ii, valueII := range values {
		c[ii], err = strconv.ParseFloat(valueII, 64)
		// fmt.Println(ii, "has value", valueII)
		if err != nil {
			return fmt.Errorf(
				"there was a problem interpreting the %v-th component of Vector3 (%v): %v",
				ii,
				values[ii],
				err,
			)
		}
	}

	return nil

}
