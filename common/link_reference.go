/*
This package contains objects and functions that are useful for URDF decoding as well as internal manipulations of URDF models in Go.
*/
package common

/*
LinkReference is an object that references links in a URDF model.
It is primarily used in objects that reference links (for example,
in JointElement objects).
*/
type LinkReference struct {
	LinkName string `xml:"link,attr"`
}
