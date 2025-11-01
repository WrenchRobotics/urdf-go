package decoding

import "github.com/WrenchRobotics/urdf-go/common/geometry"

type Geometry struct {
	// Only one of these should be non-nil at a time
	Box      *geometry.Box      `xml:"box"`
	Cylinder *geometry.Cylinder `xml:"cylinder"`
	Mesh     *geometry.Mesh     `xml:"mesh"`
	Sphere   *geometry.Sphere   `xml:"sphere"`
}
