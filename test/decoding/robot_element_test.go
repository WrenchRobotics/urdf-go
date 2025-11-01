package decoding_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/decoding"
)

/*
TestUnorderedModel_Unmarshal1
Description:

	In this test, we provide a reasonable URDF model with unordered links and joints:
	<robot name="flask">
		<link name="flask_base_link">
			<inertial>
				<origin
					xyz="0 0 0"
					rpy="0 0 0" />
				<mass
					value="5.0" />
				<inertia
					ixx="1"
					ixy="0.0"
					ixz="0.0"
					iyy="1"
					iyz="0.0"
					izz="1" />
			</inertial>
			<visual>
				<geometry>
					<mesh filename="500ml.STL" scale="0.001 0.001 0.001"/>
				</geometry>
				<material name="flask_glass">
					<color rgba="1.0 1.0 1.0 0.4"/>
				</material>
			</visual>
			<collision>
				<geometry>
					<mesh filename="500ml.STL" scale="0.001 0.001 0.001"/>
				</geometry>
			</collision>
		</link>
	</robot>
	and attempt to use the standard unmarshal to get its value.
	We should expect all values to match what we expect.
*/
func TestUnorderedModel_Unmarshal1(t *testing.T) {
	// Setup
	toDecode := `
<robot name="flask">
	<link name="flask_base_link">
		<inertial>
			<origin
				xyz="0 0 0"
				rpy="0 0 0" />
			<mass
				value="5.0" />
			<inertia
				ixx="1"
				ixy="0.0"
				ixz="0.0"
				iyy="1"
				iyz="0.0"
				izz="1" />
		</inertial>
		<visual>
			<geometry>
				<mesh filename="500ml.STL" scale="0.001 0.001 0.001"/>
			</geometry>
			<material name="flask_glass">
				<color rgba="1.0 1.0 1.0 0.4"/>
			</material>
		</visual>
		<collision>
			<geometry>
				<mesh filename="500ml.STL" scale="0.001 0.001 0.001"/>
			</geometry>
		</collision>
	</link>
</robot>
`

	// Decode
	var robot decoding.RobotElement
	err := xml.Unmarshal([]byte(toDecode), &robot)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	// - Name
	if robot.Name != "flask" {
		t.Errorf(
			"expected unorderedModel.Name to be 'flask'; received %v",
			robot.Name,
		)
	}

	// - Links
	if len(robot.Links) != 1 {
		t.Errorf(
			"expected unorderedModel to have %v links; received %v",
			1,
			len(robot.Links),
		)
	} else {
		// Verify link name
		expectedLinkName := "flask_base_link"
		if robot.Links[0].Name != expectedLinkName {
			t.Errorf(
				"expected unorderedModel link name to be \"%v\"; received \"%v\"",
				expectedLinkName,
				robot.Links[0].Name,
			)
		}
	}

	// - Joints
	if len(robot.Joints) != 0 {
		t.Errorf(
			"expected unorderedModel to have %v joints; received %v",
			0,
			len(robot.Joints),
		)
	}
}
