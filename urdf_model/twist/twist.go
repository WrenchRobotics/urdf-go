package twist

import "github.com/WrenchRobotics/urdf-go/common/pose"

type Twist struct {
	Linear  pose.Vector3
	Angular pose.Vector3
}

func (t *Twist) Clear() {
	t.Linear.Clear()
	t.Angular.Clear()
}
