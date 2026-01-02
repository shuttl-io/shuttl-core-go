//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func validateSecret_FromEnvParameters(envVarName *string) error {
	if envVarName == nil {
		return fmt.Errorf("parameter envVarName is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Secret) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Secret) validateSetSourceParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

