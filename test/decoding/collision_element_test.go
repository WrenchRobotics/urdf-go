package decoding_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/decoding"
)

/*
TestVisual_Unmarshal1
Description:

	In this test, we provide the reasonable geometry element with internal box tag:
			<collision>
				<geometry>
					<box size='1.2 2.3 7'/>
				</geometry>
				<origin xyz="1 2 3" rpy="0.1 0.2 0.3" />
			</collision>
	and attempt to use the standard unmarshal to get its value.
	We should expect all values to match what we expect.
*/
func TestCollision_Unmarshal1(t *testing.T) {
	// Setup
	toDecode := `<collision>	
		<geometry>
			<cylinder length="2.5" radius="1.0"/>
		</geometry>
		<origin xyz="1 2 3" rpy="0.1 0.2 0.3" />
	</collision>`

	// Decode
	var collisionElt decoding.CollisionElement
	err := xml.Unmarshal([]byte(toDecode), &collisionElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	err = collisionElt.Geometry.Check()
	if err != nil {
		t.Errorf(
			"unable to decode Geometry element of collision element: %v",
			err,
		)
		return
	}

	// Check values
	cylinderGeometry, ok := collisionElt.Geometry.GetActiveImplementation().(*geometry.Cylinder)
	if !ok {
		t.Errorf("The decoded geometry is not a cylinder, but is of type %T", collisionElt.Geometry.GetActiveImplementation())
	}

	if cylinderGeometry.Length != 2.5 {
		t.Errorf(
			"expected cylinderGeometry's length to be 2.5; received %v",
			cylinderGeometry.Length,
		)
	}

	if cylinderGeometry.Radius != 1.0 {
		t.Errorf(
			"expected cylinderGeometry's radius to be 1.0; received %v",
			cylinderGeometry.Radius,
		)
	}
}

/*
TestCollisionElement_Clear
Description:

	Tests that the Clear method properly resets all fields of a CollisionElement
	to their zero values. This includes:
	- Name string set to empty string
	- Origin.Position and Origin.Rotation cleared
	- Geometry implementations cleared
*/
func TestCollisionElement_Clear(t *testing.T) {
	// Setup - Create a collision element with data
	toDecode := `<collision name="test_collision">	
		<geometry>
			<box size="1.0 2.0 3.0"/>
		</geometry>
		<origin xyz="1 2 3" rpy="0.1 0.2 0.3" />
	</collision>`

	var collisionElt decoding.CollisionElement
	err := xml.Unmarshal([]byte(toDecode), &collisionElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Verify the element has data before clearing
	if collisionElt.Name != "test_collision" {
		t.Errorf("expected Name to be 'test_collision' before Clear; received %v", collisionElt.Name)
	}

	if collisionElt.Origin.Position[0] != 1.0 || collisionElt.Origin.Position[1] != 2.0 || collisionElt.Origin.Position[2] != 3.0 {
		t.Errorf("expected Origin.Position to be [1, 2, 3] before Clear; received %v", collisionElt.Origin.Position)
	}

	boxGeometry, ok := collisionElt.Geometry.GetActiveImplementation().(*geometry.Box)
	if !ok {
		t.Errorf("The decoded geometry is not a box, but is of type %T", collisionElt.Geometry.GetActiveImplementation())
	}

	if boxGeometry.Dimensions[0] != 1.0 || boxGeometry.Dimensions[1] != 2.0 || boxGeometry.Dimensions[2] != 3.0 {
		t.Errorf("expected Box dimensions to be [1.0, 2.0, 3.0] before Clear; received %v", boxGeometry.Dimensions)
	}

	// Call Clear
	collisionElt.Clear()

	// Verify all fields are cleared
	if collisionElt.Name != "" {
		t.Errorf("expected Name to be empty string after Clear; received %v", collisionElt.Name)
	}

	// Verify Origin is cleared
	if collisionElt.Origin.Position[0] != 0.0 || collisionElt.Origin.Position[1] != 0.0 || collisionElt.Origin.Position[2] != 0.0 {
		t.Errorf("expected Origin.Position to be [0, 0, 0] after Clear; received %v", collisionElt.Origin.Position)
	}

	if collisionElt.Origin.Rotation[0] != 0.0 || collisionElt.Origin.Rotation[1] != 0.0 || collisionElt.Origin.Rotation[2] != 0.0 {
		t.Errorf("expected Origin.Rotation to be [0, 0, 0] after Clear; received %v", collisionElt.Origin.Rotation)
	}

	// Verify Geometry is cleared (Box dimensions should be zero)
	if boxGeometry.Dimensions[0] != 0.0 || boxGeometry.Dimensions[1] != 0.0 || boxGeometry.Dimensions[2] != 0.0 {
		t.Errorf("expected Box dimensions to be [0.0, 0.0, 0.0] after Clear; received %v", boxGeometry.Dimensions)
	}
}
