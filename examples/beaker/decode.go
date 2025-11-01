package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/WrenchRobotics/urdf-go/decoding"
)

func main() {
	// Setup
	urdfPath := "500ml.urdf"

	// Extract the contents of the urdf
	content, err := os.ReadFile(urdfPath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Decode
	var robot decoding.RobotElement
	err = xml.Unmarshal([]byte(content), &robot)
	if err != nil {
		fmt.Println("there was an issue decoding the input toDecode:", err)
	}

	// Print output geom
	fmt.Printf("the output geom object: %+v\n", robot)

}
