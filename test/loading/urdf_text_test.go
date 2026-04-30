package loading_test

import (
	"testing"

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
