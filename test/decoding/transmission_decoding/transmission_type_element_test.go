package transmission_decoding_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/decoding/transmission_decoding"
)

/*
TestTransmissionTypeElement_Unmarshal1
Description:

	In this test, we provide the following transmission definition from a Robotiq
	gripper URDF file:
		<type>transmission_interface/SimpleTransmission</type>
	and attempt to use the standard unmarshal to get its value.
	We should expect all values to match what was given.
*/
func TestTransmissionTypeElement_Unmarshal1(t *testing.T) {
	// Setup
	toDecode := `<type>transmission_interface/SimpleTransmission</type>`

	// Decode
	var transmissionTypeElt transmission_decoding.TransmissionTypeElement
	err := xml.Unmarshal([]byte(toDecode), &transmissionTypeElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	if transmissionTypeElt.Name != "transmission_interface/SimpleTransmission" {
		t.Errorf("unexpected value for Name: got %q, want %q",
			transmissionTypeElt.Name, "transmission_interface/SimpleTransmission")
	}
}
