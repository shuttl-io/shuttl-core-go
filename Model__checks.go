//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func validateModel_OpenAIParameters(identifier *string, apiKey ISecret) error {
	if identifier == nil {
		return fmt.Errorf("parameter identifier is required, but nil was provided")
	}

	if apiKey == nil {
		return fmt.Errorf("parameter apiKey is required, but nil was provided")
	}

	return nil
}

