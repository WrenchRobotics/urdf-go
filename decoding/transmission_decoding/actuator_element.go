package transmission_decoding

type ActuatorElement struct {
	Name                string                     `xml:"name,attr"`
	MechanicalReduction MechanicalReductionElement `xml:"mechanicalReduction"`
}
