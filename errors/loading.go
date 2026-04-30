package model_errors

import "fmt"

// Deprecated: use NoRobotsFoundInFileError instead.
// This error is not specific enough to be useful, so the name is misleading.
type NoRobotsFoundError struct {
	FilePath string
}

func (err NoRobotsFoundError) Error() string {
	return fmt.Sprintf(
		"no <robot> elements were found in the URDF file at path \"%v\"",
		err.FilePath,
	)
}

// NoRobotsFoundInFileError is an error which indicates that no
// <robot> elements were found in a URDF file at a given path.
type NoRobotsFoundInFileError struct {
	FilePath string
}

func (err NoRobotsFoundInFileError) Error() string {
	return fmt.Sprintf(
		"no <robot> elements were found in the URDF file at path \"%v\"",
		err.FilePath,
	)
}

type NoRobotsFoundInContentsError struct{}

func (err NoRobotsFoundInContentsError) Error() string {
	return "no <robot> elements were found in the URDF contents"
}
