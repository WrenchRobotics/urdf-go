package model_errors

import "fmt"

type LinkDoesNotExistError struct {
	LinkName  string
	ModelName string
}

func (err LinkDoesNotExistError) Error() string {
	return fmt.Sprintf(
		"The link \"%v\" does not exist in the model \"%v\".",
		err.LinkName,
		err.ModelName,
	)
}
