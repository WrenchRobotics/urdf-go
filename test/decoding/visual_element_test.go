package decoding_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/decoding"
)

/*
TestVisual_Unmarshal1
Description:

	In this test, we provide the reasonable geometry element with internal box tag:
			<visual>
				<geometry>
					<box size='1.2 2.3 7'/>
				</geometry>
				<material name="flask_glass">
					<color rgba="1.0 1.0 1.0 0.4"/>
				</material>
			</visual>
	and attempt to use the standard unmarshal to get its value.
	We should expect all values to match what we expect.
*/
func TestVisual_Unmarshal1(t *testing.T) {
	// Setup
	dims := []float64{1.2, 2.3, 7}
	toDecode := `<visual>	
		<geometry>
			<box size="` + fmt.Sprintf("%v %v %v", dims[0], dims[1], dims[2]) + `"/>
		</geometry>
		<material name="flask_glass">
			<color rgba="1.0 1.0 1.0 0.4"/>
		</material>
		<origin xyz="1 2 3" rpy="0.1 0.2 0.3" />
	</visual>`

	// Decode
	var visualElt decoding.VisualElement
	err := xml.Unmarshal([]byte(toDecode), &visualElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	err = visualElt.Geometry.Check()
	if err != nil {
		t.Errorf(
			"unable to decode Geometry element of visual element: %v",
			err,
		)
		return
	}

	// Check values
	boxGeometry, ok := visualElt.Geometry.GetActiveImplementation().(*geometry.Box)
	if !ok {
		t.Errorf("The decoded geometry is not a box, but is of type %T", visualElt.Geometry.GetActiveImplementation())
	}

	for ii, scaleDimension := range boxGeometry.Dimensions {
		if scaleDimension != dims[ii] {
			t.Errorf(
				"expected boxGeometry's size to have %v at %v-th value; received %v",
				dims[ii],
				ii,
				scaleDimension,
			)
		}
	}
}

/*
TestVisualElement_Clear
Description:

	Tests that the Clear method properly resets all fields of a VisualElement
	to their zero values. This includes:
	- Name string set to empty string
	- Origin.Position and Origin.Rotation cleared (if Origin is not nil)
	- Geometry implementations cleared (if Geometry is not nil)
	- Material cleared (if Material is not nil)
*/
func TestVisualElement_Clear(t *testing.T) {
	// Setup - Create a visual element with all fields populated
	toDecode := `<visual name="test_visual">	
		<geometry>
			<cylinder radius="1.5" length="2.5"/>
		</geometry>
		<material name="test_material">
			<color rgba="0.5 0.6 0.7 0.8"/>
		</material>
		<origin xyz="1 2 3" rpy="0.1 0.2 0.3" />
	</visual>`

	var visualElt decoding.VisualElement
	err := xml.Unmarshal([]byte(toDecode), &visualElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Verify the element has data before clearing
	if visualElt.Name != "test_visual" {
		t.Errorf("expected Name to be 'test_visual' before Clear; received %v", visualElt.Name)
	}

	if visualElt.Origin == nil {
		t.Errorf("expected Origin to be non-nil before Clear")
	} else {
		if visualElt.Origin.Position[0] != 1.0 || visualElt.Origin.Position[1] != 2.0 || visualElt.Origin.Position[2] != 3.0 {
			t.Errorf("expected Origin.Position to be [1, 2, 3] before Clear; received %v", visualElt.Origin.Position)
		}
	}

	if visualElt.Geometry == nil {
		t.Errorf("expected Geometry to be non-nil before Clear")
	} else {
		cylinderGeometry, ok := visualElt.Geometry.GetActiveImplementation().(*geometry.Cylinder)
		if !ok {
			t.Errorf("The decoded geometry is not a cylinder, but is of type %T", visualElt.Geometry.GetActiveImplementation())
		} else {
			if cylinderGeometry.Radius != 1.5 {
				t.Errorf("expected Cylinder radius to be 1.5 before Clear; received %v", cylinderGeometry.Radius)
			}
			if cylinderGeometry.Length != 2.5 {
				t.Errorf("expected Cylinder length to be 2.5 before Clear; received %v", cylinderGeometry.Length)
			}
		}
	}

	if visualElt.Material == nil {
		t.Errorf("expected Material to be non-nil before Clear")
	} else {
		if visualElt.Material.Name != "test_material" {
			t.Errorf("expected Material.Name to be 'test_material' before Clear; received %v", visualElt.Material.Name)
		}
	}

	// Call Clear
	visualElt.Clear()

	// Verify all fields are cleared
	if visualElt.Name != "" {
		t.Errorf("expected Name to be empty string after Clear; received %v", visualElt.Name)
	}

	// Verify Origin is cleared
	if visualElt.Origin != nil {
		if visualElt.Origin.Position[0] != 0.0 || visualElt.Origin.Position[1] != 0.0 || visualElt.Origin.Position[2] != 0.0 {
			t.Errorf("expected Origin.Position to be [0, 0, 0] after Clear; received %v", visualElt.Origin.Position)
		}

		if visualElt.Origin.Rotation[0] != 0.0 || visualElt.Origin.Rotation[1] != 0.0 || visualElt.Origin.Rotation[2] != 0.0 {
			t.Errorf("expected Origin.Rotation to be [0, 0, 0] after Clear; received %v", visualElt.Origin.Rotation)
		}
	}

	// Verify Geometry is cleared (Cylinder dimensions should be zero)
	if visualElt.Geometry != nil {
		cylinderGeometry, ok := visualElt.Geometry.GetActiveImplementation().(*geometry.Cylinder)
		if ok {
			if cylinderGeometry.Radius != 0.0 {
				t.Errorf("expected Cylinder radius to be 0.0 after Clear; received %v", cylinderGeometry.Radius)
			}
			if cylinderGeometry.Length != 0.0 {
				t.Errorf("expected Cylinder length to be 0.0 after Clear; received %v", cylinderGeometry.Length)
			}
		}
	}

	// Verify Material is cleared
	if visualElt.Material != nil {
		if visualElt.Material.Name != "" {
			t.Errorf("expected Material.Name to be empty string after Clear; received %v", visualElt.Material.Name)
		}

		if visualElt.Material.Color != nil {
			if visualElt.Material.Color.Color[0] != 0.0 || visualElt.Material.Color.Color[1] != 0.0 ||
				visualElt.Material.Color.Color[2] != 0.0 || visualElt.Material.Color.Color[3] != 0.0 {
				t.Errorf("expected Material.Color to be [0.0, 0.0, 0.0, 0.0] after Clear; received %v", visualElt.Material.Color.Color)
			}
		}
	}
}

/*
TestVisualElement_Clear_WithNilFields
Description:

	Tests that the Clear method handles nil fields gracefully without panicking.
	This verifies the nil-safe implementation of the Clear method.
*/
func TestVisualElement_Clear_WithNilFields(t *testing.T) {
	// Setup - Create a visual element with minimal fields (some nil)
	toDecode := `<visual name="minimal_visual"/>`

	var visualElt decoding.VisualElement
	err := xml.Unmarshal([]byte(toDecode), &visualElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Verify the element has only the name before clearing
	if visualElt.Name != "minimal_visual" {
		t.Errorf("expected Name to be 'minimal_visual' before Clear; received %v", visualElt.Name)
	}

	// Call Clear - should not panic even with nil fields
	visualElt.Clear()

	// Verify Name is cleared
	if visualElt.Name != "" {
		t.Errorf("expected Name to be empty string after Clear; received %v", visualElt.Name)
	}
}
