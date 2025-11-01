package model_errors

import "fmt"

type NoRobotsFoundError struct {
	FilePath string
}

func (err NoRobotsFoundError) Error() string {
	return fmt.Sprintf(
		"no <robot> elements were found in the URDF file at path \"%v\"",
		err.FilePath,
	)
}
