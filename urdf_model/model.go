package urdfmodel

import (
	"github.com/WrenchRobotics/urdf-go/decoding"
	model_errors "github.com/WrenchRobotics/urdf-go/errors"
	"github.com/WrenchRobotics/urdf-go/urdf_model/joint"
	"github.com/WrenchRobotics/urdf-go/urdf_model/link"
	"github.com/WrenchRobotics/urdf-go/urdf_model/transmission"
)

type Model struct {
	Name string `xml:"name,attr"`
	// Unexported fields
	// (as a reminder, all fields with first character in lower case will not be exported outside of a Go package)
	rootLink      *link.Link
	links         map[string]*link.Link                 // Complete list of links, organized by name
	joints        map[string]*joint.Joint               // Complete list of joints, organized by name
	materials     map[string]*decoding.MaterialElement  // Complete list of materials, organized by name
	transmissions map[string]*transmission.Transmission // Complete list of transmissions, organized by name
}

/*
CheckTransmissionElement
Description:

	Checks that the specified transmission element is valid within the context of the model
*/
func (m *Model) CheckTransmissionElement(transmissionName string) error {
	// Extract transmission pointer
	transmissionPtr, err := m.GetTransmission(transmissionName)
	if err != nil {
		return err
	}

	// Check that all referenced joints exist
	for _, jointRef := range transmissionPtr.Joints {
		_, err := m.GetJoint(jointRef.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Model) DefineTreeRelationships() error {
	// Establish parent-child relationships between links and joints
	for _, joint := range m.joints {
		// Collect child and parent link names and check that they are non-empty
		childLinkName := joint.ChildLinkName
		parentLinkName := joint.ParentLinkName

		if childLinkName == "" {
			return model_errors.InvalidJointError{
				Reason: "the joint's child link name was empty",
			}
		}
		if parentLinkName == "" {
			return model_errors.InvalidJointError{
				Reason: "the joint's parent link name was empty",
			}
		}

		// Find the child and parent link pointers
		childLink, err := m.GetLink(childLinkName)
		if err != nil {
			return err
		}
		parentLink, err := m.GetLink(parentLinkName)
		if err != nil {
			return err
		}

		// For the child link:
		// - Set parent link as the parent defined in the joint
		childLink.ParentLink = parentLink

		// - Set Parent Joint as the joint for this iteration of the loop
		childLink.ParentJoint = joint

		// For the parent link:
		// - Set Child Joint as the joint which called this method
		parentLink.ChildJoints = append(
			parentLink.ChildJoints,
			joint,
		)

		// - Set Child Link as the child defined in the joint
		parentLink.ChildLinks = append(
			parentLink.ChildLinks,
			childLink,
		)

		// All changes should propagate because we are using pointers
	}

	return nil
}

/*
DefineTreeRoot
Description:

	Defines the root link of the model tree structure.

Assumptions:

	Assumes that the tree relationships have already been defined.
	In other words, assumes that DefineTreeRelationships() has already been called.
*/
func (m *Model) DefineTreeRoot() error {
	// Reset the root link pointer
	m.rootLink = nil

	// Find the links that have no parent in the tree
	var rootCandidates []*link.Link
	for _, link := range m.links {
		if link.ParentLink == nil {
			rootCandidates = append(rootCandidates, link)
		}
	}

	// Check that there is exactly one root candidate
	switch len(rootCandidates) {
	case 1:
		m.rootLink = rootCandidates[0]
		return nil
	case 0:
		return model_errors.NoRootLinkError{
			ModelName: m.Name,
		}
	default:
		return model_errors.MultipleRootLinksError{
			ModelName: m.Name,
			NumRoots:  len(rootCandidates),
		}
	}

}

func DeriveModelFrom(robotElement *decoding.RobotElement) (*Model, error) {
	// Setup
	model := &Model{
		Name:          robotElement.Name,
		links:         make(map[string]*link.Link),
		joints:        make(map[string]*joint.Joint),
		materials:     make(map[string]*decoding.MaterialElement),
		transmissions: make(map[string]*transmission.Transmission),
	}

	// Populate model with data from robotElement
	var foundMaterials []*decoding.MaterialElement
	// - Link information (basics only; no relationships yet)
	for _, linkElement := range robotElement.Links {
		model.links[linkElement.Name] = &link.Link{}
		err := model.links[linkElement.Name].FromDecodingElement(linkElement)
		if err != nil {
			return nil, err
		}
		// Collect material elements for later processing, if they exist
		for _, visualElement := range linkElement.Visual {
			if visualElement.Material != nil {
				foundMaterials = append(
					foundMaterials,
					visualElement.Material,
				)
			}
		}
	}
	// - Joint information
	for _, jointElement := range robotElement.Joints {
		model.joints[jointElement.Name] = &joint.Joint{}
		err := model.joints[jointElement.Name].FromDecodingElement(jointElement)
		if err != nil {
			return nil, err
		}
	}
	// - Material information
	for _, material := range foundMaterials {
		model.materials[material.Name] = &decoding.MaterialElement{}
		*model.materials[material.Name] = *material
	}
	// - Transmission information
	for _, transmissionElement := range robotElement.Transmissions {
		model.transmissions[transmissionElement.Name] = &transmission.Transmission{}
		err := model.transmissions[transmissionElement.Name].FromDecodingElement(transmissionElement)
		if err != nil {
			return nil, err
		}

		err = model.CheckTransmissionElement(transmissionElement.Name)
		if err != nil {
			return nil, err
		}
	}

	// - Establish link-joint relationships
	err := model.DefineTreeRelationships()
	if err != nil {
		return nil, err
	}

	// - Define root link
	err = model.DefineTreeRoot()
	if err != nil {
		return nil, err
	}

	// Return successfully constructed model
	return model, nil
}

func (m Model) GetAllJointNames() []string {
	var out []string
	for jointName := range m.joints {
		out = append(out, jointName)
	}
	return out
}

func (m Model) GetAllLinkNames() []string {
	var out []string
	for linkName := range m.links {
		out = append(out, linkName)
	}
	return out
}

func (m Model) GetAllMaterialNames() []string {
	var out []string
	for materialName := range m.materials {
		out = append(out, materialName)
	}
	return out
}

func (m Model) GetAllTransmissionNames() []string {
	var out []string
	for transmissionName := range m.transmissions {
		out = append(out, transmissionName)
	}
	return out
}

func (m Model) GetJoint(jointName string) (*joint.Joint, error) {
	// Extract pointer to joint, if it exists
	out, ok := m.joints[jointName]
	if !ok {
		return nil, model_errors.JointDoesNotExistError{
			JointName: jointName,
			ModelName: m.Name,
		}
	}

	// Return successfully retrieved value
	return out, nil
}

func (m Model) GetLink(linkName string) (*link.Link, error) {
	// Extract link pointer, if it exists
	out, ok := m.links[linkName]
	if !ok {
		return nil, model_errors.LinkDoesNotExistError{
			LinkName:  linkName,
			ModelName: m.Name,
		}
	}

	// Return successfully retrieved value
	return out, nil
}

func (m Model) GetMaterial(materialName string) (*decoding.MaterialElement, error) {
	// Extract pointer to material, if it exists
	out, ok := m.materials[materialName]
	if !ok {
		return nil, model_errors.MaterialDoesNotExistError{
			MaterialName: materialName,
			ModelName:    m.Name,
		}
	}

	// Return successfully retrieved material pointer
	return out, nil
}

func (m Model) GetTransmission(transmissionName string) (*transmission.Transmission, error) {
	// Extract pointer to transmission, if it exists
	out, ok := m.transmissions[transmissionName]
	if !ok {
		return nil, model_errors.TransmissionDoesNotExistError{
			TransmissionName: transmissionName,
			ModelName:        m.Name,
		}
	}

	// Return successfully retrieved transmission pointer
	return out, nil
}

func (m Model) NumJoints() int {
	return len(m.GetAllJointNames())
}

func (m Model) NumLinks() int {
	return len(m.GetAllLinkNames())
}

func (m Model) NumMaterials() int {
	return len(m.GetAllMaterialNames())
}

func (m Model) NumTransmissions() int {
	return len(m.GetAllTransmissionNames())
}
