package pose

import "math"

type Rotation [4]float64

func NewRotation(x, y, z, w float64) Rotation {
	return Rotation{x, y, z, w}
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

func (r *Rotation) Multiply(other Rotation) Rotation {
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

func (r *Rotation) Inverse() Rotation {
	norm := r.X()*r.X() + r.Y()*r.Y() + r.Z()*r.Z() + r.W()*r.W()
	if norm > 0 {
		invNorm := 1.0 / norm
		return Rotation{
			-r.X() * invNorm, 
			-r.Y() * invNorm, 
			-r.Z() * invNorm, 
			r.W() * invNorm
		}
	}
	return r
}