//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (t *jsiiProxy_ToolArgBuilder) validateDefaultToParameters(defaultValue interface{}) error {
	if defaultValue == nil {
		return fmt.Errorf("parameter defaultValue is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ToolArgBuilder) validateSetArgTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ToolArgBuilder) validateSetDescriptionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ToolArgBuilder) validateSetRequiredParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewToolArgBuilderParameters(argType *string, description *string, required *bool) error {
	if argType == nil {
		return fmt.Errorf("parameter argType is required, but nil was provided")
	}

	if description == nil {
		return fmt.Errorf("parameter description is required, but nil was provided")
	}

	if required == nil {
		return fmt.Errorf("parameter required is required, but nil was provided")
	}

	return nil
}

