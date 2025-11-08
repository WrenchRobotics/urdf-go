package transmission_decoding_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/decoding/transmission_decoding"
)

/*
TestTransmissionElement_Unmarshal1
Description:

	Tests the unmarshaling of some example XML data from a Robotiq gripper URDF file.
		<transmission name="finger_joint_trans">
			<type>transmission_interface/SimpleTransmission</type>
			<joint name="finger_joint">
				<hardwareInterface>PositionJointInterface</hardwareInterface>
			</joint>
			<actuator name="finger_joint_motor">
				<mechanicalReduction>1</mechanicalReduction>
			</actuator>
		</transmission>
	We should expect all values to match what was given.
*/
func TestTransmissionElement_Unmarshal1(t *testing.T) {
	// Setup
	toDecode := `<transmission name="finger_joint_trans">
		<type>transmission_interface/SimpleTransmission</type>
		<joint name="finger_joint">
			<hardwareInterface>PositionJointInterface</hardwareInterface>
		</joint>
		<actuator name="finger_joint_motor">
			<mechanicalReduction>1</mechanicalReduction>
		</actuator>
	</transmission>`

	// Decode
	var transmissionElt transmission_decoding.TransmissionElement
	err := xml.Unmarshal([]byte(toDecode), &transmissionElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	if transmissionElt.Name != "finger_joint_trans" {
		t.Errorf("unexpected value for Name: got %q, want %q",
			transmissionElt.Name, "finger_joint_trans")
	}

	if transmissionElt.Type.Name != "transmission_interface/SimpleTransmission" {
		t.Errorf("unexpected value for Type.Name: got %q, want %q",
			transmissionElt.Type.Name, "transmission_interface/SimpleTransmission")
	}
	// - Joints
	// (There should only be one joint in this example)
	if len(transmissionElt.Joints) != 1 {
		t.Errorf("unexpected number of Joints: got %d, want %d",
			len(transmissionElt.Joints), 1)
	}

	if transmissionElt.Joints[0].Name != "finger_joint" {
		t.Errorf("unexpected value for Joint.Name: got %q, want %q",
			transmissionElt.Joints[0].Name, "finger_joint")
	}

	if len(transmissionElt.Joints[0].HardwareInterfaces) != 1 {
		t.Errorf("unexpected number of Joint.HardwareInterfaces: got %d, want %d",
			len(transmissionElt.Joints[0].HardwareInterfaces), 1)
	}
	for _, hwInterface := range transmissionElt.Joints[0].HardwareInterfaces {
		if hwInterface.Name != "PositionJointInterface" {
			t.Errorf("unexpected value for Joint.HardwareInterface.Name: got %q, want %q",
				hwInterface.Name, "PositionJointInterface")
		}
	}

	// - Actuators
	if len(transmissionElt.Actuators) != 1 {
		t.Errorf("unexpected number of Actuators: got %d, want %d",
			len(transmissionElt.Actuators), 1)
	}
	if transmissionElt.Actuators[0].Name != "finger_joint_motor" {
		t.Errorf("unexpected value for Actuator.Name: got %q, want %q",
			transmissionElt.Actuators[0].Name, "finger_joint_motor")
	}

	if transmissionElt.Actuators[0].MechanicalReduction.Value != 1.0 {
		t.Errorf("unexpected value for Actuator.MechanicalReduction.Value: got %v, want %v",
			transmissionElt.Actuators[0].MechanicalReduction.Value, 1.0)
	}
}
