package model_errors

import "fmt"

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
