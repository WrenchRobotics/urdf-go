package inertial

import "github.com/WrenchRobotics/urdf-go/common/pose"

/*
Represents the inertial quantities for a link (moment of inertia,
mass, center of mass) according to the URDF specification.

It uses the xml struct tags to allow for easy decoding and encoding
using the xml.Unmarshal or xml.Marshal methods.

You should be able to parse any inertial element. For example, this:

	<inertial>
		<origin rpy="0 0 0" xyz="8.625E-08 -4.6583E-06 0.03145"/>
		<mass value="0.22652"/>
		<inertia ixx="0.00020005" ixy="-4.2442E-10" ixz="-2.9069E-10" iyy="0.00017832" iyz="-3.4402E-08" izz="0.00013478"/>
	</inertial>
*/
type Inertial struct {

	// Center of mass with which Moment of Inertia (`Inertia`) is defined.
	Origin pose.Pose `xml:"origin"`

	// Mass of the object
	Mass Mass `xml:"mass"`

	// Moment of inertia for the object
	Inertia Inertia `xml:"inertia"`
}

/*
Clears the internal variables for all of the internal variables
*/
func (i *Inertial) Clear() {
	i.Origin.Clear()
	i.Mass.Clear()
	i.Inertia.Clear()
}
