package transmission

type TransmissionActuatorReference struct {
	Name                string
	MechanicalReduction *float64 // nil if not specified
}
