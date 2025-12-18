package pose

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

/*
An array of three elements which can be used in the decoding or encoding
of 3D vectors in a URDF.
*/
type Vector3 [3]float64

// Defines a new vector using three input values.
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{x, y, z}
}

// Returns the first value in the vector.
func (v *Vector3) X() float64 {
	return v[0]
}

// Returns the second value in the vector.
func (v *Vector3) Y() float64 {
	return v[1]
}

// Returns the third value in the vector.
func (v Vector3) Z() float64 {
	return v[2]
}

// Returns the addition of the two vectors.
func (v *Vector3) Plus(other Vector3) Vector3 {
	return Vector3{v[0] + other[0], v[1] + other[1], v[2] + other[2]}
}

// Sets the current vector to all zeros.
func (v *Vector3) Clear() {
	v[0] = 0
	v[1] = 0
	v[2] = 0
}

/*
MarshalXMLAttr returns an xml attribute that can be easily encoded into
a new URDF file.

This method should not normally be called by users. It is used
to implement some important interfaces in encoding/xml, and thus is
mainly used by encoding/xml when you call those methods (e.g., when you call xml.Marshal()).
*/
func (v Vector3) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: fmt.Sprintf("%v %v %v", v.X(), v.Y(), v.Z()),
	}, nil
}

/*
UnmarshalXMLAttr updates the current vector using the values
obtained from the attributes of an xml tag.

For example, the following tag contains a quantity that would be
decoded using this method:

	<origin rpy="0 0 3.141592653589793" xyz="0 -0.0306011 0.054904"/>

Note: This method should not be normally called by users.
It is defined so that this type implements an interface from encoding/xml.
Usually, it is called by the methods from encoding/xml that need it (e.g., xml.Unmarshal).
*/
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
