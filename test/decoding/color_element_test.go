package decoding_test

import (
	"encoding/xml"
	"testing"

	"github.com/WrenchRobotics/urdf-go/decoding"
)

/*
TestColorElement_Clear
Description:

	Tests that the Clear method properly resets all fields of a ColorElement
	to their zero values. The Color field should be set to [0, 0, 0, 0].
*/
func TestColorElement_Clear(t *testing.T) {
	// Setup - Create a color element with data
	toDecode := `<color rgba="0.5 0.6 0.7 0.8"/>`

	var colorElt decoding.ColorElement
	err := xml.Unmarshal([]byte(toDecode), &colorElt)
	if err != nil {
		t.Errorf("there was an issue decoding the input toDecode: %v", err)
	}

	// Verify the element has data before clearing
	if colorElt.Color[0] != 0.5 || colorElt.Color[1] != 0.6 || colorElt.Color[2] != 0.7 || colorElt.Color[3] != 0.8 {
		t.Errorf("expected Color to be [0.5, 0.6, 0.7, 0.8] before Clear; received %v", colorElt.Color)
	}

	// Call Clear
	colorElt.Clear()

	// Verify all fields are cleared
	if colorElt.Color[0] != 0.0 || colorElt.Color[1] != 0.0 || colorElt.Color[2] != 0.0 || colorElt.Color[3] != 0.0 {
		t.Errorf("expected Color to be [0.0, 0.0, 0.0, 0.0] after Clear; received %v", colorElt.Color)
	}
}
