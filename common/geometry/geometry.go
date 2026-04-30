/*
This package contains convenience methods for handling the different types of geometries that are allowed in a URDF.
*/
package geometry

import (
	"fmt"

	geometry_type "github.com/WrenchRobotics/urdf-go/common/geometry/type"
	model_errors "github.com/WrenchRobotics/urdf-go/errors"
)

// The Geometry object is an object that is used to conveniently parse any
// `geometry` tag in a URDF.
//
// Typically, a geometry tag encloses a specific "implementation" of geometry, like in this case:
//
//	<geometry>
//		<mesh filename="../meshes/visual/robotiq_arg2f_85_inner_finger.obj" scale="0.001 0.001 0.001"/>
//	</geometry>
//
// Or in this example:
//
//	<geometry>
//		<box size='1.2 2.3 7'/>
//	</geometry>
//
// Because of the uncertainty in "what the geometry tag contains", we make this struct
// capable of containing _pointers_ any possible geometry implementation. During decoding,
// only ONE of the pointers will be "not nil". You can retrieve the active one by calling
// Geomtry.GetActiveImplementation() method.
type Geometry struct {
	// If this geometry is of Box type, then this pointer will be non-nil.
	Box *Box `xml:"box"`

	// If this geometry is of Cylinder type, then this pointer will be non-nil.
	Cylinder *Cylinder `xml:"cylinder"`

	// If this geometry is of Mesh type, then this pointer will be non-nil.
	Mesh *Mesh `xml:"mesh"`

	// If this geometry is of Sphere type, then this pointer will be non-nil.
	Sphere *Sphere `xml:"sphere"`
}

/*
Check validates whether or not the Geometry object is valid.

The object is valid if one AND ONLY one geometry implementation is active.
*/
func (g Geometry) Check() error {
	// Return an error if:
	// - No geometry is not nil
	// - More than one geometry is not nil

	count := 0
	activeTypes := []geometry_type.GeometryType{}
	for geomType, geomImpl := range g.GetImplementationMap() {
		switch geomImpl := geomImpl.(type) {
		case *Box:
			var nilBox *Box = nil
			if geomImpl != nilBox {
				count++
				activeTypes = append(activeTypes, geomType)
			}
		case *Cylinder:
			var nilCyl *Cylinder = nil
			if geomImpl != nilCyl {
				count++
				activeTypes = append(activeTypes, geomType)
			}
		case *Mesh:
			var nilMesh *Mesh = nil
			if geomImpl != nilMesh {
				count++
				activeTypes = append(activeTypes, geomType)
			}
		case *Sphere:
			var nilSphere *Sphere = nil
			if geomImpl != nilSphere {
				count++
				activeTypes = append(activeTypes, geomType)
			}
		default:
			return &model_errors.UnknownGeometryImplementationError{
				ImplementationName: fmt.Sprintf("%T", geomImpl),
				MethodName:         "Check",
			}
		}
	}

	// No geometry is set
	if count == 0 {
		return fmt.Errorf("no geometry is set")
	}

	// More than one geometry is set
	if count > 1 {
		return fmt.Errorf("more than one geometry is set: %v", activeTypes)
	}

	// Otherwsie, all good
	return nil
}

/*
Defines a map between each of the possible geometry types (see geometry_type package)
to the value that is stored in this object (if the type is not represented
in this object, then you should receive an empty object I think).

TODO(Kwesi): Investigate if we should delete this.
*/
func (g *Geometry) GetImplementationMap() map[geometry_type.GeometryType]GeometryImplementation {
	return map[geometry_type.GeometryType]GeometryImplementation{
		geometry_type.Box:      g.Box,
		geometry_type.Cylinder: g.Cylinder,
		geometry_type.Mesh:     g.Mesh,
		geometry_type.Sphere:   g.Sphere,
	}
}

/*
Clears the internal variables of the active geometry, if possible.
*/
func (g *Geometry) Clear() {
	if g.Box != nil {
		g.Box.Clear()
	}
	if g.Cylinder != nil {
		g.Cylinder.Clear()
	}
	if g.Mesh != nil {
		g.Mesh.Clear()
	}
	if g.Sphere != nil {
		g.Sphere.Clear()
	}
}

/*
Returns the active implementation, if there is one.

If the geometry is not well-defined (i.e., fails the tests in Check())
then this will return nil.
*/
func (g *Geometry) GetActiveImplementation() GeometryImplementation {
	if g.Check() != nil {
		return nil
	}

	for _, geomImpl := range g.GetImplementationMap() {
		switch geomImpl := geomImpl.(type) {
		case *Box:
			var nilBox *Box = nil
			if geomImpl != nilBox {
				return geomImpl
			}
		case *Cylinder:
			var nilCyl *Cylinder = nil
			if geomImpl != nilCyl {
				return geomImpl
			}
		case *Mesh:
			var nilMesh *Mesh = nil
			if geomImpl != nilMesh {
				return geomImpl
			}
		case *Sphere:
			var nilSphere *Sphere = nil
			if geomImpl != nilSphere {
				return geomImpl
			}
		}
	}
	return nil
}
