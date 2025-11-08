package transmission_decoding

type TransmissionElement struct {
	Name      string                     `xml:"name,attr"`
	Type      TransmissionTypeElement    `xml:"type"`
	Joints    []TransmissionJointElement `xml:"joint"`
	Actuators []ActuatorElement          `xml:"actuator"`
}
