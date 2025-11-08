package transmission

import hardwareinterface "github.com/WrenchRobotics/urdf-go/common/hardware_interface"

type TransmissionJointReference struct {
	Name               string
	HardwareInterfaces []*hardwareinterface.HardwareInterface
}
