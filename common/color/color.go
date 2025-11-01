package color

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Color [4]float64

func NewColor(r, g, b, a float64) Color {
	return Color{r, g, b, a}
}

func (c *Color) Clear() {
	*c = Color{0, 0, 0, 0}
}

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
