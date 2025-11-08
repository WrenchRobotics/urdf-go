package main

import (
	"fmt"

	"github.com/WrenchRobotics/urdf-go/loaders"
)

func main() {
	// Setup
	urdfPath := "robotiq_2f_85.urdf"

	// Load using our loading library
	urdfModel, err := loaders.FromURDFFile(urdfPath)
	if err != nil {
		panic(fmt.Errorf("there was an issue loading the URDF file: %v", err))
	}

	// Print information about the loaded model
	fmt.Println("the number of links in the model:", urdfModel.NumLinks())
	fmt.Println("the number of joints in the model:", urdfModel.NumJoints())
	fmt.Println("the number of materials in the model:", urdfModel.NumMaterials())
	fmt.Println("the number of transmissions in the model:", urdfModel.NumTransmissions())

	// Print details about each transmission
	for i, transmissionName := range urdfModel.GetAllTransmissionNames() {
		transmission, err := urdfModel.GetTransmission(transmissionName)
		if err != nil {
			fmt.Printf("Error retrieving transmission %d: %v\n", i+1, err)
			continue
		}
		fmt.Printf("Transmission %d: Name=%s, Type=%s\n", i+1, transmission.Name, transmission.Type)
	}
}
