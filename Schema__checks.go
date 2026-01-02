//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func validateSchema_BooleanValueParameters(description *string) error {
	if description == nil {
		return fmt.Errorf("parameter description is required, but nil was provided")
	}

	return nil
}

func validateSchema_EnumValueParameters(description *string, enumValues *[]*string) error {
	if description == nil {
		return fmt.Errorf("parameter description is required, but nil was provided")
	}

	if enumValues == nil {
		return fmt.Errorf("parameter enumValues is required, but nil was provided")
	}

	return nil
}

func validateSchema_NumberValueParameters(description *string) error {
	if description == nil {
		return fmt.Errorf("parameter description is required, but nil was provided")
	}

	return nil
}

func validateSchema_ObjectValueParameters(properties *map[string]ToolArgBuilder) error {
	if properties == nil {
		return fmt.Errorf("parameter properties is required, but nil was provided")
	}

	return nil
}

func validateSchema_StringValueParameters(description *string) error {
	if description == nil {
		return fmt.Errorf("parameter description is required, but nil was provided")
	}

	return nil
}

