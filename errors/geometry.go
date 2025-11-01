package model_errors

type UnknownGeometryImplementationError struct {
	ImplementationName string
	MethodName         string
}

func (e *UnknownGeometryImplementationError) Error() string {
	return "Unknown geometry implementation '" + e.ImplementationName + "' for method '" + e.MethodName + "'"
}
