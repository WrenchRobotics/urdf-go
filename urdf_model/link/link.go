package link

import (
	"github.com/WrenchRobotics/urdf-go/urdf_model/joint"
	"github.com/WrenchRobotics/urdf-go/urdf_model/link/inertial"
)

type Link struct {
	Name           string             `xml:"name,attr"`
	Inertial       *inertial.Inertial `xml:"inertial"`
	VisualArray    []*Visual          `xml:"visual"`
	CollisionArray []*Collision       `xml:"collision"`
	ParentJoint    *joint.Joint       `xml:"parent_joint"`
	ChildJoints    []*joint.Joint     `xml:"child_joints"`
	ChildLinks     []*Link            `xml:"child_links"`
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
		v.Clear()
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
