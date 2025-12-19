// This package defines a number of helper structs and methods for managing poses and 3D math.
package pose

/*
Pose is an object which exactly represents a pose according to the
URDF specification.

It uses the xml struct tags so that you can take advantage
of the xml encoding/decoding features (e.g., xml.Marhsal, xml.Unmarshal)
to easily read or write URDFs. No need to write  any additional
functions.
*/
type Pose struct {

	// 3D position of a quantity (e.g., a link) in the URDF
	Position Vector3 `xml:"xyz,attr"`

	// Rotation of the quantity (e.g., a link) in the URDF
	Rotation Rotation `xml:"rpy,attr"`
}

/*
Sets the position to the origin and sets the rotation to the identity rotation.
*/
func (p *Pose) Clear() {
	p.Position.Clear()
	p.Rotation.Clear()
}
