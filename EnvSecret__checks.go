//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func validateNewEnvSecretParameters(envVarName *string) error {
	if envVarName == nil {
		return fmt.Errorf("parameter envVarName is required, but nil was provided")
	}

	return nil
}

