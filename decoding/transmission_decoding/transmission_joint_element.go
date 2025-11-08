package transmission_decoding

type TransmissionJointElement struct {
	Name              string                   `xml:"name,attr"`
	HardwareInterface HardwareInterfaceElement `xml:"hardwareInterface"`
}
