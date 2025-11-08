package transmission_decoding

type TransmissionElement struct {
	Name     string                   `xml:"name,attr"`
	Type     TransmissionTypeElement  `xml:"type"`
	Joint    TransmissionJointElement `xml:"joint"`
	Actuator ActuatorElement          `xml:"actuator"`
}
