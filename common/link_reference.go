package common

/*
LinkReference is an object that references links in a URDF model.
It is primarily used in objects that reference links (for example,
in JointElement objects).
*/
type LinkReference struct {
	LinkName string `xml:"link,attr"`
}
