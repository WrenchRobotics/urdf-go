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

type NoRootLinkError struct {
	ModelName string
}

func (err NoRootLinkError) Error() string {
	return fmt.Sprintf(
		"The model \"%v\" does not have a root link.",
		err.ModelName,
	)
}

type MultipleRootLinksError struct {
	ModelName string
	NumRoots  int
}

func (err MultipleRootLinksError) Error() string {
	return fmt.Sprintf(
		"The model \"%v\" has %v root links; expected exactly one.",
		err.ModelName,
		err.NumRoots,
	)
}
