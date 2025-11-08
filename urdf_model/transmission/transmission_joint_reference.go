package transmission

import "github.com/WrenchRobotics/urdf-go/urdf_model/transmission/hardware_interface"

type TransmissionJointReference struct {
	Name               string
	HardwareInterfaces []*hardware_interface.HardwareInterface
}
