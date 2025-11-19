package link_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/common/geometry"
	"github.com/WrenchRobotics/urdf-go/decoding"
	"github.com/WrenchRobotics/urdf-go/urdf_model/link"
)

/*
TestLink_FromDecodingElement1
Description:

	Tests the FromDecodingElement method with a basic link that has only a name.
	This verifies that the method correctly handles a minimal LinkElement.
*/
func TestLink_FromDecodingElement1(t *testing.T) {
	// Setup - Create a LinkElement with just a name
	toDecode := `<link name="test_link"/>`

	var linkElt decoding.LinkElement
	err := xml.Unmarshal([]byte(toDecode), &linkElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
		return
	}

	// Convert to Link
	testLink := &link.Link{}
	err = testLink.FromDecodingElement(&linkElt)
	if err != nil {
		t.Errorf("FromDecodingElement failed: %v", err)
		return
	}

	// Check values
	if testLink.Name != "test_link" {
		t.Errorf("expected link name to be 'test_link', got '%s'", testLink.Name)
	}
	if testLink.Inertial != nil {
		t.Errorf("expected inertial to be nil, got %v", testLink.Inertial)
	}
	if len(testLink.VisualArray) != 0 {
		t.Errorf("expected no visual elements, got %d", len(testLink.VisualArray))
	}
	if len(testLink.CollisionArray) != 0 {
		t.Errorf("expected no collision elements, got %d", len(testLink.CollisionArray))
	}
}

/*
TestLink_FromDecodingElement2_with_visual
Description:

	Tests the FromDecodingElement method with a link that contains a visual element.
	This verifies that visual elements are correctly converted.
*/
func TestLink_FromDecodingElement2_with_visual(t *testing.T) {
	// Setup - Create a LinkElement with a visual
	toDecode := `<link name="test_link">
		<visual>
			<geometry>
				<box size="1.0 2.0 3.0"/>
			</geometry>
			<origin xyz="0.1 0.2 0.3" rpy="0 0 0"/>
		</visual>
	</link>`

	var linkElt decoding.LinkElement
	err := xml.Unmarshal([]byte(toDecode), &linkElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
		return
	}

	// Convert to Link
	testLink := &link.Link{}
	err = testLink.FromDecodingElement(&linkElt)
	if err != nil {
		t.Errorf("FromDecodingElement failed: %v", err)
		return
	}

	// Check values
	if testLink.Name != "test_link" {
		t.Errorf("expected link name to be 'test_link', got '%s'", testLink.Name)
	}
	if len(testLink.VisualArray) != 1 {
		t.Errorf("expected 1 visual element, got %d", len(testLink.VisualArray))
		return
	}

	// Check visual geometry
	visual := testLink.VisualArray[0]
	if visual == nil {
		t.Errorf("visual element is nil")
		return
	}

	boxGeometry, ok := visual.Geometry.(*geometry.Box)
	if !ok {
		t.Errorf("expected box geometry, got %T", visual.Geometry)
		return
	}

	expectedDims := []float64{1.0, 2.0, 3.0}
	for i, dim := range boxGeometry.Dimensions {
		if dim != expectedDims[i] {
			t.Errorf("expected dimension[%d] to be %v, got %v", i, expectedDims[i], dim)
		}
	}
}

/*
TestLink_FromDecodingElement3_with_collision
Description:

	Tests the FromDecodingElement method with a link that contains a collision element.
	This verifies that collision elements are correctly converted.
*/
func TestLink_FromDecodingElement3_with_collision(t *testing.T) {
	// Setup - Create a LinkElement with a collision
	toDecode := `<link name="test_link">
		<collision>
			<geometry>
				<cylinder radius="0.5" length="2.0"/>
			</geometry>
			<origin xyz="1 2 3" rpy="0.1 0.2 0.3"/>
		</collision>
	</link>`

	var linkElt decoding.LinkElement
	err := xml.Unmarshal([]byte(toDecode), &linkElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
		return
	}

	// Convert to Link
	testLink := &link.Link{}
	err = testLink.FromDecodingElement(&linkElt)
	if err != nil {
		t.Errorf("FromDecodingElement failed: %v", err)
		return
	}

	// Check values
	if testLink.Name != "test_link" {
		t.Errorf("expected link name to be 'test_link', got '%s'", testLink.Name)
	}
	if len(testLink.CollisionArray) != 1 {
		t.Errorf("expected 1 collision element, got %d", len(testLink.CollisionArray))
		return
	}

	// Check collision geometry
	collision := testLink.CollisionArray[0]
	if collision == nil {
		t.Errorf("collision element is nil")
		return
	}

	cylinderGeometry, ok := collision.Geometry.(*geometry.Cylinder)
	if !ok {
		t.Errorf("expected cylinder geometry, got %T", collision.Geometry)
		return
	}

	if cylinderGeometry.Radius != 0.5 {
		t.Errorf("expected cylinder radius to be 0.5, got %v", cylinderGeometry.Radius)
	}
	if cylinderGeometry.Length != 2.0 {
		t.Errorf("expected cylinder length to be 2.0, got %v", cylinderGeometry.Length)
	}
}

/*
TestLink_FromDecodingElement4_with_inertial
Description:

	Tests the FromDecodingElement method with a link that contains an inertial element.
	This verifies that inertial elements are correctly converted.
*/
func TestLink_FromDecodingElement4_with_inertial(t *testing.T) {
	// Setup - Create a LinkElement with inertial
	toDecode := `<link name="test_link">
		<inertial>
			<mass value="1.5"/>
			<origin xyz="0 0 0" rpy="0 0 0"/>
			<inertia ixx="0.1" ixy="0" ixz="0" iyy="0.1" iyz="0" izz="0.1"/>
		</inertial>
	</link>`

	var linkElt decoding.LinkElement
	err := xml.Unmarshal([]byte(toDecode), &linkElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
		return
	}

	// Convert to Link
	testLink := &link.Link{}
	err = testLink.FromDecodingElement(&linkElt)
	if err != nil {
		t.Errorf("FromDecodingElement failed: %v", err)
		return
	}

	// Check values
	if testLink.Name != "test_link" {
		t.Errorf("expected link name to be 'test_link', got '%s'", testLink.Name)
	}
	if testLink.Inertial == nil {
		t.Errorf("expected inertial to be non-nil")
		return
	}

	if testLink.Inertial.Mass.Value != 1.5 {
		t.Errorf("expected mass value to be 1.5, got %v", testLink.Inertial.Mass.Value)
	}
}

/*
TestLink_FromDecodingElement5_with_multiple_elements
Description:

	Tests the FromDecodingElement method with a link that contains multiple visual
	and collision elements. This verifies that arrays are properly populated.
*/
func TestLink_FromDecodingElement5_with_multiple_elements(t *testing.T) {
	// Setup - Create a LinkElement with multiple visuals and collisions
	toDecode := `<link name="test_link">
		<visual>
			<geometry>
				<box size="1.0 1.0 1.0"/>
			</geometry>
		</visual>
		<visual>
			<geometry>
				<sphere radius="0.5"/>
			</geometry>
		</visual>
		<collision>
			<geometry>
				<box size="1.0 1.0 1.0"/>
			</geometry>
		</collision>
		<collision>
			<geometry>
				<sphere radius="0.5"/>
			</geometry>
		</collision>
	</link>`

	var linkElt decoding.LinkElement
	err := xml.Unmarshal([]byte(toDecode), &linkElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
		return
	}

	// Convert to Link
	testLink := &link.Link{}
	err = testLink.FromDecodingElement(&linkElt)
	if err != nil {
		t.Errorf("FromDecodingElement failed: %v", err)
		return
	}

	// Check values
	if testLink.Name != "test_link" {
		t.Errorf("expected link name to be 'test_link', got '%s'", testLink.Name)
	}
	if len(testLink.VisualArray) != 2 {
		t.Errorf("expected 2 visual elements, got %d", len(testLink.VisualArray))
	}
	if len(testLink.CollisionArray) != 2 {
		t.Errorf("expected 2 collision elements, got %d", len(testLink.CollisionArray))
	}

	// Verify first visual is a box
	if testLink.VisualArray[0] != nil {
		if _, ok := testLink.VisualArray[0].Geometry.(*geometry.Box); !ok {
			t.Errorf("expected first visual to be a box, got %T", testLink.VisualArray[0].Geometry)
		}
	}

	// Verify second visual is a sphere
	if testLink.VisualArray[1] != nil {
		if _, ok := testLink.VisualArray[1].Geometry.(*geometry.Sphere); !ok {
			t.Errorf("expected second visual to be a sphere, got %T", testLink.VisualArray[1].Geometry)
		}
	}
}

/*
TestLink_FromDecodingElement6_nil_pointer
Description:

	Tests the FromDecodingElement method with a nil pointer.
	This verifies that the method correctly handles error cases.
*/
func TestLink_FromDecodingElement6_nil_pointer(t *testing.T) {
	// Setup
	testLink := &link.Link{}

	// Try to convert from nil LinkElement
	err := testLink.FromDecodingElement(nil)
	if err == nil {
		t.Errorf("expected an error when passing nil pointer, got none")
	}
}

/*
TestLink_FromDecodingElement7_complete_link
Description:

	Tests the FromDecodingElement method with a complete link that has
	name, inertial, visual, and collision elements.
*/
func TestLink_FromDecodingElement7_complete_link(t *testing.T) {
	// Setup - Create a complete LinkElement
	toDecode := `<link name="complete_link">
		<inertial>
			<mass value="2.0"/>
			<origin xyz="0 0 0" rpy="0 0 0"/>
			<inertia ixx="0.2" ixy="0" ixz="0" iyy="0.2" iyz="0" izz="0.2"/>
		</inertial>
		<visual>
			<geometry>
				<box size="1.0 2.0 3.0"/>
			</geometry>
			<origin xyz="0 0 0" rpy="0 0 0"/>
			<material name="test_material">
				<color rgba="1.0 0.0 0.0 1.0"/>
			</material>
		</visual>
		<collision>
			<geometry>
				<box size="1.0 2.0 3.0"/>
			</geometry>
			<origin xyz="0 0 0" rpy="0 0 0"/>
		</collision>
	</link>`

	var linkElt decoding.LinkElement
	err := xml.Unmarshal([]byte(toDecode), &linkElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
		return
	}

	// Convert to Link
	testLink := &link.Link{}
	err = testLink.FromDecodingElement(&linkElt)
	if err != nil {
		t.Errorf("FromDecodingElement failed: %v", err)
		return
	}

	// Check all values
	if testLink.Name != "complete_link" {
		t.Errorf("expected link name to be 'complete_link', got '%s'", testLink.Name)
	}
	if testLink.Inertial == nil {
		t.Errorf("expected inertial to be non-nil")
	}
	if len(testLink.VisualArray) != 1 {
		t.Errorf("expected 1 visual element, got %d", len(testLink.VisualArray))
	}
	if len(testLink.CollisionArray) != 1 {
		t.Errorf("expected 1 collision element, got %d", len(testLink.CollisionArray))
	}

	// Verify visual has material
	if len(testLink.VisualArray) > 0 && testLink.VisualArray[0] != nil {
		if testLink.VisualArray[0].Material == nil {
			t.Errorf("expected visual to have material, got nil")
		} else if testLink.VisualArray[0].Material.Name != "test_material" {
			t.Errorf("expected material name to be 'test_material', got '%s'", testLink.VisualArray[0].Material.Name)
		}
	}
}
