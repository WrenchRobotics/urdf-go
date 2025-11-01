package urdfmodel

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/decoding"
	urdfmodel "github.com/WrenchRobotics/urdf-go/urdf_model"
)

/*
TestModel_DeriveModelFrom1
Description:

Tests the DeriveModelFrom1 function to ensure it correctly
derives a new model from a successfully processed <robot> tag.
*/
func TestModel_DeriveModelFrom1(t *testing.T) {
	// Setup
	toDecode := `<robot name="test_robot">
		<link name="link1">
			<visual>
				<geometry>
					<box size="1 1 1"/>
				</geometry>
				<material name="test_material">
					<color rgba="0.5 0.5 0.5 1.0"/>
				</material>
			</visual>
		</link>
		<link name="link2"/>
		<joint name="joint1" type="fixed">
			<parent link="link1"/>
			<child link="link2"/>
		</joint>
	</robot>`

	// Decode
	var robotElt decoding.RobotElement
	err := xml.Unmarshal([]byte(toDecode), &robotElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Derive model
	model, err := urdfmodel.DeriveModelFrom(&robotElt)
	if err != nil {
		t.Errorf("there was an issue deriving the model: %v", err)
	}

	// Check values
	if model.Name != "test_robot" {
		t.Errorf("expected model name to be 'test_robot', got '%s'", model.Name)
	}
	if model.NumLinks() != 2 {
		t.Errorf("expected 2 links, got %d", model.NumLinks())
	}
	if model.NumJoints() != 1 {
		t.Errorf("expected 1 joint, got %d", model.NumJoints())
	}
	if model.NumMaterials() != 1 {
		t.Errorf("expected 1 material, got %d", model.NumMaterials())
	}
}
