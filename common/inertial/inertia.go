/*
inertial defines convenience structs and methods for defining inertial quantities (e.g., mass) from URDF files.
*/
package inertial

/*
Represents the compressed moment of inertia for a link according to the
URDF specification.

We use the xml struct tags so that you can parse `inertia` tags
in a URDF without adding extra code.

You should be able to decode any inertia of the following form:

	<inertia ixx="0.00020005" ixy="-4.2442E-10" ixz="-2.9069E-10" iyy="0.00017832" iyz="-3.4402E-08" izz="0.00013478"/>

using xml.Unmarshal().
*/
type Inertia struct {
	Ixx float64 `xml:"ixx,attr"`
	Ixy float64 `xml:"ixy,attr"`
	Ixz float64 `xml:"ixz,attr"`
	Iyy float64 `xml:"iyy,attr"`
	Iyz float64 `xml:"iyz,attr"`
	Izz float64 `xml:"izz,attr"`
}

/*
Sets all moment of inertia values to zero.
*/
func (in *Inertia) Clear() {
	in.Ixx = 0
	in.Ixy = 0
	in.Ixz = 0
	in.Iyy = 0
	in.Iyz = 0
	in.Izz = 0
}
