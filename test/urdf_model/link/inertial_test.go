package link_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/urdf_model/link/inertial"
)

/*
TestInertial_Unmarshal1
Description:

	This test demonstrates that the Inertial element can be properly unmarshaled
	from the following correct XML snippet:
			<inertial>
				<origin xyz="0 1 2" rpy="0.1 0.2 0.3" />
				<mass value="1e-9" />
				<inertia ixx="0.7" ixy="0" ixz="0" iyy="0.8" iyz="0" izz="0.9" />
			</inertial>
	We should expect all values to match what we expect.

	This example comes from:
		https://github.com/rimim/AWD/blob/main/awd/data/assets/go_bdx/go_bdx.urdf
*/
func TestInertial_Unmarshal1(t *testing.T) {
	// Setup
	toDecode := `<inertial>
		<origin xyz="0 1 2" rpy="0.1 0.2 0.3" />
		<mass value="1e-9" />
		<inertia ixx="0.7" ixy="0" ixz="0" iyy="0.8" iyz="0" izz="0.9" />
	</inertial>`

	// Decode
	var inertialElt inertial.Inertial
	err := xml.Unmarshal([]byte(toDecode), &inertialElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	if inertialElt.Origin.Position.X() != 0 {
		t.Errorf(
			"expected inertialElt.Origin.Position.X() to be 0; received %v",
			inertialElt.Origin.Position.X(),
		)
	}

	if inertialElt.Origin.Position.Y() != 1 {
		t.Errorf(
			"expected inertialElt.Origin.Position.Y() to be 1; received %v",
			inertialElt.Origin.Position.Y(),
		)
	}

	if inertialElt.Origin.Position.Z() != 2 {
		t.Errorf(
			"expected inertialElt.Origin.Position.Z() to be 2; received %v",
			inertialElt.Origin.Position.Z(),
		)
	}

	// Rotation checks
	roll, pitch, yaw := inertialElt.Origin.Rotation.ToRollPitchYaw()
	if roll != 0.1 {
		t.Errorf(
			"expected inertialElt.Origin.Rotation.Roll() to be 0.1; received %v",
			roll,
		)
	}
	if pitch != 0.2 {
		t.Errorf(
			"expected inertialElt.Origin.Rotation.Pitch() to be 0.2; received %v",
			pitch,
		)
	}
	if yaw != 0.3 {
		t.Errorf(
			"expected inertialElt.Origin.Rotation.Yaw() to be 0.3; received %v",
			yaw,
		)
	}

	// Mass check
	if inertialElt.Mass.Value != 1e-9 {
		t.Errorf(
			"expected inertialElt.Mass to be 1e-9; received %v",
			inertialElt.Mass,
		)
	}

	// Moment of Inertia checks
	if inertialElt.Inertia.Ixx != 0.7 {
		t.Errorf(
			"expected inertialElt.Inertia.Ixx to be 0.7; received %v",
			inertialElt.Inertia.Ixx,
		)
	}

	if inertialElt.Inertia.Iyy != 0.8 {
		t.Errorf(
			"expected inertialElt.Inertia.Iyy to be 0.8; received %v",
			inertialElt.Inertia.Iyy,
		)
	}

	if inertialElt.Inertia.Izz != 0.9 {
		t.Errorf(
			"expected inertialElt.Inertia.Izz to be 0.9; received %v",
			inertialElt.Inertia.Izz,
		)
	}
}
