package link_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/urdf_model/link"
	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
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
	var collisionElt link.Collision
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
