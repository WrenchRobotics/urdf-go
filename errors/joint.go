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

type InvalidJointError struct {
	Reason string
}

func (err InvalidJointError) Error() string {
	return fmt.Sprintf(
		"the joint is invalid: %v",
		err.Reason,
	)
}
