//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (j *jsiiProxy_IPCResponseError) validateSetCodeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IPCResponseError) validateSetMessageParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

