package link

import (
	"fmt"

	"github.com/WrenchRobotics/urdf-go/common/inertial"
	"github.com/WrenchRobotics/urdf-go/decoding"
	"github.com/WrenchRobotics/urdf-go/urdf_model/joint"
)

type Link struct {
	Name           string
	Inertial       *inertial.Inertial
	VisualArray    []*Visual
	CollisionArray []*Collision
	ParentJoint    *joint.Joint
	ParentLink     *Link
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

func (l *Link) FromDecodingElement(le *decoding.LinkElement) error {
	if le == nil {
		return fmt.Errorf(
			"the provided link element pointer was nil",
		)
	}
	l.Name = le.Name
	if le.Inertial != nil {
		l.Inertial = le.Inertial
	}
	for _, ve := range le.Visual {
		visual := &Visual{}
		if err := visual.FromDecodingElement(ve); err != nil {
			return fmt.Errorf("failed to decode visual: %v", err)
		}
		l.VisualArray = append(l.VisualArray, visual)
	}
	for _, ce := range le.Collision {
		collision := &Collision{}
		if err := collision.FromDecodingElement(ce); err != nil {
			return fmt.Errorf("failed to decode collision: %v", err)
		}
		l.CollisionArray = append(l.CollisionArray, collision)
	}

	return nil
}
