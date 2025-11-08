package transmission

import (
	"github.com/WrenchRobotics/urdf-go/decoding/transmission_decoding"
	"github.com/WrenchRobotics/urdf-go/urdf_model/transmission/hardware_interface"
	"github.com/WrenchRobotics/urdf-go/urdf_model/transmission/transmission_type"
)

type Transmission struct {
	Name      string
	Type      transmission_type.TransmissionType
	Joints    []TransmissionJointReference
	Actuators []TransmissionActuatorReference
}

func (t *Transmission) FromDecodingElement(e *transmission_decoding.TransmissionElement) error {
	t.Name = e.Name

	// Transmission Type
	err := t.Type.FromDecodingElement(&e.Type)
	if err != nil {
		return err
	}

	// Joint References
	for _, jointElement := range e.Joints {
		var jointRef TransmissionJointReference
		jointRef.Name = jointElement.Name
		for _, hwInterfaceElement := range jointElement.HardwareInterfaces {
			var hwInterface hardware_interface.HardwareInterface
			hwInterface.FromDecodingElement(&hwInterfaceElement)
			jointRef.HardwareInterfaces = append(jointRef.HardwareInterfaces, &hwInterface)
		}
		t.Joints = append(t.Joints, jointRef)
	}

	// Actuator References
	for _, actuatorElement := range e.Actuators {
		var actuatorRef TransmissionActuatorReference
		actuatorRef.Name = actuatorElement.Name
		if actuatorElement.MechanicalReduction != nil {
			actuatorRef.MechanicalReduction = &actuatorElement.MechanicalReduction.Value
		}
		t.Actuators = append(t.Actuators, actuatorRef)
	}

	return nil
}
