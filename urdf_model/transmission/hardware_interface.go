package transmission

import "github.com/WrenchRobotics/urdf-go/decoding/transmission_decoding"

type HardwareInterface string

const (
	PositionHardwareInterface HardwareInterface = "position"
	VelocityHardwareInterface HardwareInterface = "velocity"
	EffortHardwareInterface   HardwareInterface = "effort"
)

func (hi *HardwareInterface) FromDecodingElement(e *transmission_decoding.HardwareInterfaceElement) error {
	*hi = HardwareInterface(e.Name)
	return nil
}
