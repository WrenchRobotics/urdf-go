package transmission_decoding_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/decoding/transmission_decoding"
)

/*
TestActuatorElement_Unmarshal1
Description:

	Tests the unmarshalling of an ActuatorElement from XML data.
		<actuator name="right_outer_knuckle_motor">
			<mechanicalReduction>1</mechanicalReduction>
		</actuator>
	We should expect all values to match what was given.
*/
func TestActuatorElement_Unmarshal1(t *testing.T) {
	// Setup
	toDecode := `<actuator name="right_outer_knuckle_motor">
		<mechanicalReduction>1</mechanicalReduction>
	</actuator>`

	// Decode
	var actuatorElt transmission_decoding.ActuatorElement
	err := xml.Unmarshal([]byte(toDecode), &actuatorElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	if actuatorElt.Name != "right_outer_knuckle_motor" {
		t.Errorf("unexpected value for Name: got %q, want %q",
			actuatorElt.Name, "right_outer_knuckle_motor")
	}

	if actuatorElt.MechanicalReduction.Value != 1.0 {
		t.Errorf("unexpected value for MechanicalReduction.Value: got %v, want %v",
			actuatorElt.MechanicalReduction.Value, 1.0)
	}
}
