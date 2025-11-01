package model_errors

import "fmt"

type MaterialDoesNotExistError struct {
	MaterialName string
	ModelName    string
}

func (err MaterialDoesNotExistError) Error() string {
	return fmt.Sprintf(
		"the material \"%v\" does not exist in the model \"%v\"",
		err.MaterialName,
		err.ModelName,
	)
}
