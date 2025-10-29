package pose

import (
	"encoding/xml"
	"fmt"
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

func (v Vector3) Unmarshal
