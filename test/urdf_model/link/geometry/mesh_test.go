package geometry_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
)

/*
TestMesh_Unmarshal1
Description:

	In this test, we provide the reasonable mesh tag:
		<mesh filename='500ml.STL' scale='0.001 0.001 0.001'/>
	and attempt to use the standard unmarshal to get its value.
	We should expect all values to match what we expect.
*/
func TestMesh_Unmarshal1(t *testing.T) {
	// Setup
	toDecode := `
	<mesh filename="500ml.STL" scale="0.001 0.001 0.001"/>
	`

	// Decode
	var meshGeometry geometry.Mesh
	err := xml.Unmarshal([]byte(toDecode), &meshGeometry)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
	if meshGeometry.Filename != "500ml.STL" {
		t.Errorf(
			"expected meshGeometry filename to be \"%v\" received \"%v\"",
			"500ml.STL",
			meshGeometry.Filename,
		)
	}

	for ii, scaleDimension := range meshGeometry.Scale {
		if scaleDimension != 0.001 {
			t.Errorf(
				"expected meshGeometry's scale to have %v at %v-th value; received %v",
				0.001,
				ii,
				scaleDimension,
			)
		}
	}
}
