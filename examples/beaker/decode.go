package main

import (
	"encoding/xml"
	"fmt"

	"github.com/WrenchRobotics/urdf-go/urdf_model/link/geometry"
)

func main() {
	// Setup
	toDecode := `
	<geometry>
		<mesh filename="500ml.STL" scale="0.001 0.001 0.001"/>
	</geometry>
	`

	// Decode
	var geom geometry.Mesh
	err := xml.Unmarshal([]byte(toDecode), geom)
	if err != nil {
		fmt.Println("there was an issue decoding the input toDecode: %v", err)
	}

	// Print output geom
	fmt.Println("the output geom object: %v", geom)

}
