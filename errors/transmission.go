package model_errors

import "fmt"

type TransmissionDoesNotExistError struct {
	TransmissionName string
	ModelName        string
}

func (err TransmissionDoesNotExistError) Error() string {
	return fmt.Sprintf(
		"the transmission \"%v\" does not exist in the model \"%v\"",
		err.TransmissionName,
		err.ModelName,
	)
}
