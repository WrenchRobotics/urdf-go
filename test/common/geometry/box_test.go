package geometry_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/WrenchRobotics/urdf-go/common/geometry"
)

/*
TestBox_Unmarshal1
Description:

	In this test, we provide the reasonable box tag:
		<box size='1.2 2.3 7'/>
	and attempt to use the standard unmarshal to get its value.
	We should expect all values to match what we expect.
*/
func TestBox_Unmarshal1(t *testing.T) {
	// Setup
	dims := []float64{1.2, 2.3, 7}
	toDecode := `
	<box size="` + fmt.Sprintf("%v %v %v", dims[0], dims[1], dims[2]) + `"/>
	`

	// Decode
	var boxGeometry geometry.Box
	err := xml.Unmarshal([]byte(toDecode), &boxGeometry)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Check values
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
