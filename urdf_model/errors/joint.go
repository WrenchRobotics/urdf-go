package model_errors

import "fmt"

type JointDoesNotExistError struct {
	JointName string
	ModelName string
}

func (err JointDoesNotExistError) Error() string {
	return fmt.Sprintf(
		"the joint named \"%v\" does not exist in the model \"%v\".",
		err.JointName,
		err.ModelName,
	)
}
