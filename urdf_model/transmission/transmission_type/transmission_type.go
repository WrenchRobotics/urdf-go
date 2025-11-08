package transmission_type

import "github.com/WrenchRobotics/urdf-go/decoding/transmission_decoding"

type TransmissionType string

const (
	Simple TransmissionType = "transmission_interface/SimpleTransmission"
)

func (t *TransmissionType) FromDecodingElement(e *transmission_decoding.TransmissionTypeElement) error {
	*t = TransmissionType(e.Name)
	return nil
}
