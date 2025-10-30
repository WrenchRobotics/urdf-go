package geometry

import (
	"fmt"

	model_errors "github.com/WrenchRobotics/urdf-go/urdf_model/errors"
	geometry_type "github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry/type"
)

type Geometry struct {
	// Only one of these should be non-nil at a time
	Box      *Box      `xml:"box"`
	Cylinder *Cylinder `xml:"cylinder"`
	Mesh     *Mesh     `xml:"mesh"`
	Sphere   *Sphere   `xml:"sphere"`
}

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

func (g *Geometry) GetImplementationMap() map[geometry_type.GeometryType]GeometryImplementation {
	return map[geometry_type.GeometryType]GeometryImplementation{
		geometry_type.Box:      g.Box,
		geometry_type.Cylinder: g.Cylinder,
		geometry_type.Mesh:     g.Mesh,
		geometry_type.Sphere:   g.Sphere,
	}
}

func (g *Geometry) Clear() {
	for _, geomImpl := range g.GetImplementationMap() {
		if geomImpl != nil {
			geomImpl.Clear()
		}
	}
}

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
