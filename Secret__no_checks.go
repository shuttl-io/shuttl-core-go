//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func validateSecret_FromEnvParameters(envVarName *string) error {
	return nil
}

func (j *jsiiProxy_Secret) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Secret) validateSetSourceParameters(val *string) error {
	return nil
}

