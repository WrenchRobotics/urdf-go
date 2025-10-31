package link_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/WrenchRobotics/urdf-go/urdf_model/link"
	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
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
	t.Errorf(toDecode)

	// Decode
	var visualElt link.Visual
	err := xml.Unmarshal([]byte(toDecode), &visualElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	t.Errorf("%v", visualElt)

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
