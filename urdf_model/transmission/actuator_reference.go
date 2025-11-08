package transmission

import (
	"github.com/WrenchRobotics/urdf-go/common/mechanical_reduction"
)

type TransmissionActuatorReference struct {
	Name                string
	MechanicalReduction *mechanical_reduction.MechanicalReduction
}
