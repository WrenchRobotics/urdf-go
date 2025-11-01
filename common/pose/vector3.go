package pose

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Vector3 [3]float64

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{x, y, z}
}

func (v *Vector3) X() float64 {
	return v[0]
}

func (v *Vector3) Y() float64 {
	return v[1]
}

func (v Vector3) Z() float64 {
	return v[2]
}

func (v *Vector3) Plus(other Vector3) Vector3 {
	return Vector3{v[0] + other[0], v[1] + other[1], v[2] + other[2]}
}

func (v *Vector3) Clear() {
	v[0] = 0
	v[1] = 0
	v[2] = 0
}

func (v Vector3) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: fmt.Sprintf("%v %v %v", v.X(), v.Y(), v.Z()),
	}, nil
}

func (v *Vector3) UnmarshalXMLAttr(attr xml.Attr) error {
	// Split value into three parts by the "space"
	valueWithoutBrackets := strings.ReplaceAll(attr.Value, "[", "")
	valueWithoutBrackets = strings.ReplaceAll(valueWithoutBrackets, "]", "")
	values := strings.Split(valueWithoutBrackets, " ")

	// Check that there are three elements
	if len(values) != 3 {
		return fmt.Errorf(
			"there was a problem parsing \"%v\" as a Vector3. Expected there to be 2 spaces, but there were not!",
			values,
		)
	}

	// Extract each element
	var err error
	for ii, valueII := range values {
		v[ii], err = strconv.ParseFloat(valueII, 64)
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
