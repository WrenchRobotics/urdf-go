package transmission_decoding

type TransmissionJointElement struct {
	Name               string                     `xml:"name,attr"`
	HardwareInterfaces []HardwareInterfaceElement `xml:"hardwareInterface"`
}
