package transmission_decoding_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/decoding/transmission_decoding"
)

/*
TestTransmissionJointElement_Unmarshal1
Description:

	Tests the unmarshalling of a TransmissionJointElement from XML data.
	In this case, we provide the following example from Robotiq gripper URDF file:
		<joint name="right_outer_knuckle_joint">
			<hardwareInterface>PositionJointInterface</hardwareInterface>
		</joint>
	We should expect all values to match what was given.
*/
func TestTransmissionJointElement_Unmarshal1(t *testing.T) {
	// Setup
	toDecode := `<joint name="right_outer_knuckle_joint">
		<hardwareInterface>PositionJointInterface</hardwareInterface>
	</joint>`

	// Decode
	var transmissionJointElt transmission_decoding.TransmissionJointElement
	err := xml.Unmarshal([]byte(toDecode), &transmissionJointElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	// - Name
	if transmissionJointElt.Name != "right_outer_knuckle_joint" {
		t.Errorf("unexpected value for Name: got %q, want %q",
			transmissionJointElt.Name, "right_outer_knuckle_joint")
	}

	// - Hardware Interfaces
	if len(transmissionJointElt.HardwareInterfaces) != 1 {
		t.Errorf("unexpected number of HardwareInterfaces: got %d, want %d",
			len(transmissionJointElt.HardwareInterfaces), 1)
	}
	for _, hwInterface := range transmissionJointElt.HardwareInterfaces {
		if hwInterface.Name != "PositionJointInterface" {
			t.Errorf("unexpected value for HardwareInterface.Name: got %q, want %q",
				hwInterface.Name, "PositionJointInterface")
		}
	}
}
