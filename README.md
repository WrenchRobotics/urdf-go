# urdf-go
A Golang wrapper for interpreting URDF files (a version of xml).

## Installation

You can use this library in your own Go projects by running:
```
go get github.com/WrenchRobotics/urdf-go
```

## Usage

This library is designed so that you only need a single loader to get a complete model of the URDF.

```
package main

import (
	"fmt"

	"github.com/WrenchRobotics/urdf-go/loaders"
)

func main() {
	// Setup
	urdfPath := "500ml.urdf"

	// Load using our loading library
	urdfModel, err := loaders.FromURDFFile(urdfPath)
	if err != nil {
		panic(fmt.Errorf("there was an issue loading the URDF file: %v", err))
	}

	// Print information about the loaded model
	fmt.Println("the number of links in the model:", urdfModel.NumLinks())
	fmt.Println("the number of joints in the model:", urdfModel.NumJoints())
	fmt.Println("the number of materials in the model:", urdfModel.NumMaterials())

}
```

More details to come soon.
