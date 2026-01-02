//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (i *jsiiProxy_IAgentStreamerWriter) validateWriteParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IAgentStreamerWriter) validateWriteObjectParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

