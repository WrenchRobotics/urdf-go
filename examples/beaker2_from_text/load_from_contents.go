package main

import (
	"fmt"

	"github.com/WrenchRobotics/urdf-go/loaders"
)

func main() {
	// Setup
	urdfContents := []byte(`
	<robot name="erlenmeyer_flask_500ml">
    <link name="flask_base_link">
        <inertial>
            <origin
                xyz="0 0 0"
                rpy="0 0 0" />
            <mass
                value="5.0" />
            <inertia
                ixx="1"
                ixy="0.0"
                ixz="0.0"
                iyy="1"
                iyz="0.0"
                izz="1" />
        </inertial>
        <visual>
            <geometry>
                <mesh filename="500ml.STL" scale="0.001 0.001 0.001"/>
            </geometry>
            <material name="flask_glass">
                <color rgba="1.0 1.0 1.0 0.4"/>
            </material>
        </visual>
        <collision>
            <geometry>
                <mesh filename="500ml.STL" scale="0.001 0.001 0.001"/>
            </geometry>
        </collision>
    </link>

</robot>
	`)

	// Load using our loading library
	urdfModel, err := loaders.FromURDFContents(urdfContents)
	if err != nil {
		panic(fmt.Errorf("there was an issue loading the URDF contents: %v", err))
	}

	// Print information about the loaded model
	fmt.Println("the number of links in the model:", urdfModel.NumLinks())
	fmt.Println("the number of joints in the model:", urdfModel.NumJoints())
	fmt.Println("the number of materials in the model:", urdfModel.NumMaterials())

}
