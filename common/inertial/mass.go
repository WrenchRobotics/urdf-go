package inertial

/*
Represents the mass of a link with the attributes required
by the URDF specification.

It uses the xml struct tags so that you can decode or encode
a mass using methods like, xml.Unmarshal or xml.Marhsal.

For example, this object should be decode-able using xml.Unmarshal:

	<mass value="0.22652"/>
*/
type Mass struct {
	Value float64 `xml:"value,attr"`
}

/*
Sets the value of the mass to zero.
*/
func (m *Mass) Clear() {
	m.Value = 0
}
