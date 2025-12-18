package geometry

import (
	geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"
	"github.com/WrenchRobotics/urdf-go/common/pose"
)

/*
Represents a mesh with the fields (required and not) from the URDF
specification.

A mesh geometry can be defined as follows:

	<mesh filename="../meshes/visual/robotiq_arg2f_85_inner_finger.obj" scale="0.001 0.001 0.001"/>

The object uses the xml decorators on each field so that they are automatically
retrieved when decoding a URDF file.
*/
type Mesh struct {

	// Path to the file that represents the mesh
	Filename string `xml:"filename,attr"`

	// Vector of the three values that defines the
	// scale to apply to each dimension of the mesh (X,Y, and Z)
	// By default, this should be assumed to be a vector of all ones.
	Scale *pose.Vector3 `xml:"scale,attr"`
}

/*
Always returns `geometry_type.Mesh`.

This method is defined in order to implement (in other words,
satisfy the requirements of) the `GeometryImplementation` interface in this package.
*/
func (m *Mesh) Type() geometry_type.GeometryType {
	return geometry_type.Mesh
}

/*
Clears the properties of the current mesh, if they exist.

This method is defined in order to implement (in other words,
satisfy the requirements of) the `GeometryImplementation` interface in this package.
*/
func (m *Mesh) Clear() {
	m.Filename = ""

	if m.Scale != nil {
		m.Scale.Clear()
	}
}
