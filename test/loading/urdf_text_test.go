package loading_test

import (
	"strings"
	"testing"

	model_errors "github.com/WrenchRobotics/urdf-go/errors"
	"github.com/WrenchRobotics/urdf-go/loaders"
)

/*
TestFromURDFContents_loads_urdf_and_detects_link_and_joint_count
Description:

	Tests that FromURDFContents successfully loads a URDF from inline
	contents and reports the expected number of links and joints.
*/
func TestFromURDFContents_loads_urdf_and_detects_link_and_joint_count(t *testing.T) {
	// Setup
	urdfContents := []byte(`
<robot name="simple_robot">
  <link name="base_link"/>
  <link name="child_link"/>
  <joint name="base_to_child" type="fixed">
    <parent link="base_link"/>
    <child link="child_link"/>
    <origin xyz="0 0 1" rpy="0 0 0"/>
  </joint>
</robot>`)

	// Execute
	model, err := loaders.FromURDFContents(urdfContents)

	// Verify
	if err != nil {
		t.Fatalf("expected URDF contents to load successfully, got error: %v", err)
	}
	if model == nil {
		t.Fatalf("expected model to be non-nil after loading URDF contents")
	}
	if model.NumLinks() != 2 {
		t.Fatalf("expected 2 links, got %d", model.NumLinks())
	}
	if model.NumJoints() != 1 {
		t.Fatalf("expected 1 joint, got %d", model.NumJoints())
	}
}

/*
TestFromURDFContents_invalid_xml_does_not_contain_robot_element
Description:

	Tests that FromURDFContents returns an error when the provided URDF
	contents contain invalid XML that does not include a robot element.
*/
func TestFromURDFContents_invalid_xml_does_not_contain_robot_element(t *testing.T) {
	// Setup
	invalidURDFContents := []byte(`
<not_a_robot>
  <link name="base_link"/>
</not_a_robot>`)

	// Execute
	model, err := loaders.FromURDFContents(invalidURDFContents)

	// Verify
	if err == nil {
		t.Errorf(
			"model = %#v, err = %v; expected an error for URDF contents without a robot element, got nil",
			model,
			err,
		)
		t.Fatalf("expected an error for URDF contents without a robot element, got nil")
	}
	if model != nil {
		t.Fatalf("expected model to be nil when loading URDF contents without a robot element, got %#v", model)
	}
	if !strings.Contains(err.Error(), model_errors.NoRobotsFoundInContentsError{}.Error()) {
		t.Fatalf(
			"expected error to contain '%v', got %q",
			model_errors.NoRobotsFoundInContentsError{}.Error(),
			err.Error(),
		)
	}
}
