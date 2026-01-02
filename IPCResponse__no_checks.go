//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IPCResponse) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IPCResponse) validateSetSuccessParameters(val *bool) error {
	return nil
}

