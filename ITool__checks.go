//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (i *jsiiProxy_ITool) validateExecuteParameters(args *map[string]interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ITool) validateSetDescriptionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ITool) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

