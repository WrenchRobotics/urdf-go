package decoding_test

import (
	"encoding/xml"
	"testing"

	joint_type "github.com/WrenchRobotics/urdf-go/common/joint/type"
	"github.com/WrenchRobotics/urdf-go/decoding/joint_decoding"
)

/*
TestJoint_Unmarshal1
Description:

Tests the unmarshaling of a Joint struct from XML data, ensuring that all fields are correctly populated.

This test uses the following XML snippet as input:

	<joint name="finger_joint" type="revolute">
		<origin rpy="0 0 3.141592653589793" xyz="0 -0.0306011 0.054904"/>
		<parent link="robotiq_arg2f_base_link"/>
		<child link="left_outer_knuckle"/>
		<axis xyz="1 0 0"/>
		<limit effort="1000" lower="0" upper="0.8" velocity="2.0"/>
	</joint>
*/
func TestJoint_Unmarshal1(t *testing.T) {
	// Setup

	// Define the XML input
	toDecode := `<joint name="finger_joint" type="revolute">
		<origin rpy="0 0 3.141592653589793" xyz="0 -0.0306011 0.054904"/>
		<parent link="robotiq_arg2f_base_link"/>
		<child link="left_outer_knuckle"/>
		<axis xyz="1 0 0"/>
		<limit effort="1000" lower="0" upper="0.8" velocity="2.0"/>
	</joint>`

	// Decode
	var jointElt joint_decoding.JointElement
	err := xml.Unmarshal([]byte(toDecode), &jointElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	// - Name
	if jointElt.Name != "finger_joint" {
		t.Errorf(
			"expected jointElt.Name to be 'finger_joint'; received %v",
			jointElt.Name,
		)
	}

	// - Type
	if jointElt.Type != joint_type.RevoluteJoint {
		t.Errorf(
			"expected jointElt.Type to be 'revolute'; received %v",
			jointElt.Type,
		)
	}

	// - ChildLinkRef
	if jointElt.ChildLinkRef.LinkName != "left_outer_knuckle" {
		t.Errorf(
			"expected jointElt.ChildLinkRef.LinkName to be 'left_outer_knuckle'; received %v",
			jointElt.ChildLinkRef.LinkName,
		)
	}

	// - ParentLinkRef
	if jointElt.ParentLinkRef.LinkName != "robotiq_arg2f_base_link" {
		t.Errorf(
			"expected jointElt.ParentLinkRef.LinkName to be 'robotiq_arg2f_base_link'; received %v",
			jointElt.ParentLinkRef.LinkName,
		)
	}

	// - Axis
	if jointElt.Axis.Direction.X() != 1 || jointElt.Axis.Direction.Y() != 0 || jointElt.Axis.Direction.Z() != 0 {
		t.Errorf(
			"expected jointElt.Axis.Direction to be (1,0,0); received (%v,%v,%v)",
			jointElt.Axis.Direction.X(),
			jointElt.Axis.Direction.Y(),
			jointElt.Axis.Direction.Z(),
		)
	}

	// - Limits
	if jointElt.Limits == nil {
		t.Errorf("expected jointElt.Limits to be non-nil")
	} else {
		if jointElt.Limits.Effort != 1000 {
			t.Errorf(
				"expected jointElt.Limits.Effort to be 1000; received %v",
				jointElt.Limits.Effort,
			)
		}
		if jointElt.Limits.Lower != 0 {
			t.Errorf(
				"expected jointElt.Limits.Lower to be 0; received %v",
				jointElt.Limits.Lower,
			)
		}
		if jointElt.Limits.Upper != 0.8 {
			t.Errorf(
				"expected jointElt.Limits.Upper to be 0.8; received %v",
				jointElt.Limits.Upper,
			)
		}
		if jointElt.Limits.Velocity != 2.0 {
			t.Errorf(
				"expected jointElt.Limits.Velocity to be 2.0; received %v",
				jointElt.Limits.Velocity,
			)
		}
	}

}

/*
TestJoint_Unmarshal2
Description:

Tests the unmarshaling of a Joint struct from XML data, specifically for a fixed joint type.

This test uses the following XML snippet as input:

	<joint name="left_inner_finger_joint" type="revolute">
		<origin rpy="0 0 0" xyz="0 0.0061 0.0471"/>
		<parent link="left_outer_finger"/>
		<child link="left_inner_finger"/>
		<axis xyz="1 0 0"/>
		<limit effort="1000" lower="-0.8757" upper="0" velocity="2.0"/>
		<mimic joint="finger_joint" multiplier="-1" offset="0"/>
	</joint>
*/
func TestJoint_Unmarshal2(t *testing.T) {
	// Setup

	// Define the XML input
	toDecode := `<joint name="left_inner_finger_joint" type="revolute">
		<origin rpy="0 0 0" xyz="0 0.0061 0.0471"/>
		<parent link="left_outer_finger"/>
		<child link="left_inner_finger"/>
		<axis xyz="1 0 0"/>
		<limit effort="1000" lower="-0.8757" upper="0" velocity="2.0"/>
		<mimic joint="finger_joint" multiplier="-1" offset="0"/>
	</joint>`

	// Decode
	var jointElt joint_decoding.JointElement
	err := xml.Unmarshal([]byte(toDecode), &jointElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	// - Name
	if jointElt.Name != "left_inner_finger_joint" {
		t.Errorf(
			"expected jointElt.Name to be 'left_inner_finger_joint'; received %v",
			jointElt.Name,
		)
	}

	// - Type
	if jointElt.Type != joint_type.RevoluteJoint {
		t.Errorf(
			"expected jointElt.Type to be 'revolute'; received %v",
			jointElt.Type,
		)
	}

	// - Mimic
	if jointElt.Mimic == nil {
		t.Errorf("expected jointElt.Mimic to be non-nil")
	} else {
		if jointElt.Mimic.JointName != "finger_joint" {
			t.Errorf(
				"expected jointElt.Mimic.JointName to be 'finger_joint'; received %v",
				jointElt.Mimic.JointName,
			)
		}
		if jointElt.Mimic.Multiplier != -1 {
			t.Errorf(
				"expected jointElt.Mimic.Multiplier to be -1; received %v",
				jointElt.Mimic.Multiplier,
			)
		}
		if jointElt.Mimic.Offset != 0 {
			t.Errorf(
				"expected jointElt.Mimic.Offset to be 0; received %v",
				jointElt.Mimic.Offset,
			)
		}
	}

}

/*
TestJoint_Unmarshal3
Description:

Tests the unmarshaling of a Joint struct from XML data, specifically for a fixed joint type.

This test uses the following XML snippet as input:

	<joint name="arm_joint" type="prismatic">
		<parent link="link_before_joint"/>
		<child link="link_after_joint"/>
		<origin xyz="0 0 0" rpy="0 0 0"/>
		<axis xyz="0 0 1"/>
		<limit lower="-3.14" upper="3.14" effort="10" velocity="1.0"/>
		<dynamics damping="0.1" friction="0.01"/>
	</joint>
*/
func TestJoint_Unmarshal3(t *testing.T) {
	// Setup

	// Define the XML input
	toDecode := `<joint name="arm_joint" type="prismatic">
		<parent link="link_before_joint"/>
		<child link="link_after_joint"/>
		<origin xyz="0 0 0" rpy="0 0 0"/>
		<axis xyz="0 0 1"/>
		<limit lower="-3.14" upper="3.14" effort="10" velocity="1.0"/>
		<dynamics damping="0.1" friction="0.01"/>
	</joint>`

	// Decode
	var jointElt joint_decoding.JointElement
	err := xml.Unmarshal([]byte(toDecode), &jointElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	// - Name
	if jointElt.Name != "arm_joint" {
		t.Errorf(
			"expected jointElt.Name to be 'arm_joint'; received %v",
			jointElt.Name,
		)
	}

	// - Type
	if jointElt.Type != joint_type.PrismaticJoint {
		t.Errorf(
			"expected jointElt.Type to be 'prismatic'; received %v",
			jointElt.Type,
		)
	}

	// - Dynamics
	if jointElt.Dynamics == nil {
		t.Errorf("expected jointElt.Dynamics to be non-nil")
	} else {
		if jointElt.Dynamics.Damping != 0.1 {
			t.Errorf(
				"expected jointElt.Dynamics.Damping to be 0.1; received %v",
				jointElt.Dynamics.Damping,
			)
		}
		if jointElt.Dynamics.Friction != 0.01 {
			t.Errorf(
				"expected jointElt.Dynamics.Friction to be 0.01; received %v",
				jointElt.Dynamics.Friction,
			)
		}
	}

}
