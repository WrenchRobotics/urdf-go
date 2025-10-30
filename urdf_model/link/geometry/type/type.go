package geometry_type

type GeometryType string

const (
	Box      GeometryType = "box"
	Cylinder GeometryType = "cylinder"
	Sphere   GeometryType = "sphere"
	Mesh     GeometryType = "mesh"
)
