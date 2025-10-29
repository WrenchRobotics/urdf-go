package link

import "github.com/WrenchRobotics/urdf-go/urdf_model/joint"

type Link struct {
	Name           string
	Inertial       *Inertial
	VisualArray    []*Visual
	CollisionArray []*Collision
	ParentJoint    *joint.Joint
	ChildJoints    []*joint.Joint
	ChildLinks     []*Link
}

func (l *Link) Clear() {
	l.Name = ""
	if l.Inertial != nil {
		l.Inertial.Clear()
	}
	for _, v := range l.VisualArray {
		if v == nil {
			continue
		}
		v.Name = ""
		v.Origin.Clear()
		if v.Geometry != nil {
			v.Geometry.Clear()
		}
		if v.Material != nil {
			v.Material.Name = ""
			v.Material.TextureFilename = ""
			if v.Material.Color != nil {
				v.Material.Color.Clear()
			}
		}
	}
	l.VisualArray = nil
	for _, c := range l.CollisionArray {
		if c == nil {
			continue
		}
		c.Clear()
	}
	l.CollisionArray = nil
	l.ParentJoint = nil
	l.ChildJoints = nil
	l.ChildLinks = nil
}
