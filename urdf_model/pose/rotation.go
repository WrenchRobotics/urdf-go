package pose

import (
	"encoding/xml"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Rotation [4]float64

func NewRotation(x, y, z, w float64) Rotation {
	return Rotation{x, y, z, w}
}

/*
FromRollPitchYaw
Description:

	Assume roll (x), pitch (y) and yaw (z) are angles in radians.
*/
func (r *Rotation) FromRollPitchYaw(roll, pitch, yaw float64) {
	// Setup
	cr, sr := math.Cos(roll*0.5), math.Sin(roll*0.5)
	cp, sp := math.Cos(pitch*0.5), math.Sin(pitch*0.5)
	cy, sy := math.Cos(yaw*0.5), math.Sin(yaw*0.5)

	// Assemble Quaternion
	r[3] = cr*cp*cy + sr*sp*sy
	r[0] = sr*cp*cy - cr*sp*sy
	r[1] = cr*sp*cy + sr*cp*sy
	r[2] = cr*cp*sy - sr*sp*cy
}

func (r *Rotation) X() float64 {
	return r[0]
}

func (r *Rotation) Y() float64 {
	return r[1]
}

func (r *Rotation) Z() float64 {
	return r[2]
}

func (r *Rotation) W() float64 {
	return r[3]
}

func (r *Rotation) ToRollPitchYaw() (roll, pitch, yaw float64) {
	// Setup
	sqx := r.X() * r.X()
	sqy := r.Y() * r.Y()
	sqz := r.Z() * r.Z()
	sqw := r.W() * r.W()

	sarg := -2.0 * (r.X()*r.Z() - r.W()*r.Y())
	piOver2 := math.Pi / 2.0
	switch {
	case sarg <= -0.99999:
		roll = 0
		pitch = -piOver2
		yaw = -2.0 * math.Atan2(r.X(), -r.Y())
	case sarg >= 0.99999:
		roll = 0
		pitch = piOver2
		yaw = 2.0 * math.Atan2(-r.X(), r.Y())
	default:
		roll = math.Atan2(2.0*(r.Y()*r.Z()+r.W()*r.X()), sqw-sqx-sqy+sqz)
		pitch = math.Asin(sarg)
		yaw = math.Atan2(2.0*(r.X()*r.Y()+r.W()*r.Z()), sqw+sqx-sqy-sqz)
	}

	return
}

func (r *Rotation) Normalize() {
	norm := math.Sqrt(r.X()*r.X() + r.Y()*r.Y() + r.Z()*r.Z() + r.W()*r.W())
	if norm <= 0 {
		r[0] = 0
		r[1] = 0
		r[2] = 0
		r[3] = 1
		return
	}
	r[0] /= norm
	r[1] /= norm
	r[2] /= norm
	r[3] /= norm
}

func (r Rotation) Multiply(other Rotation) Rotation {
	return Rotation{
		r.W()*other.X() + r.X()*other.W() + r.Y()*other.Z() - r.Z()*other.Y(),
		r.W()*other.Y() - r.X()*other.Z() + r.Y()*other.W() + r.Z()*other.X(),
		r.W()*other.Z() + r.X()*other.Y() - r.Y()*other.X() + r.Z()*other.W(),
		r.W()*other.W() - r.X()*other.X() - r.Y()*other.Y() - r.Z()*other.Z(),
	}
}

func (r *Rotation) Rotate(vec Vector3) Vector3 {
	vecRot := Rotation{vec.X(), vec.Y(), vec.Z(), 0}
	tmp := r.Multiply(vecRot).Multiply(r.Inverse())
	return Vector3{tmp.X(), tmp.Y(), tmp.Z()}
}

func (r Rotation) Inverse() Rotation {
	norm := r.X()*r.X() + r.Y()*r.Y() + r.Z()*r.Z() + r.W()*r.W()
	if norm > 0 {
		invNorm := 1.0 / norm
		return Rotation{
			-r.X() * invNorm,
			-r.Y() * invNorm,
			-r.Z() * invNorm,
			r.W() * invNorm,
		}
	}
	return r
}

func (r *Rotation) Clear() {
	r[0] = 0
	r[1] = 0
	r[2] = 0
	r[3] = 1
}

func (r *Rotation) UnmarshalXMLAttr(attr xml.Attr) error {
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
	var rpy [3]float64
	for ii, valueII := range values {
		rpy[ii], err = strconv.ParseFloat(valueII, 64)
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

	r.FromRollPitchYaw(rpy[0], rpy[1], rpy[2])

	return nil

}
