//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_Toolkit) validateAddToolParameters(tool ITool) error {
	if tool == nil {
		return fmt.Errorf("parameter tool is required, but nil was provided")
	}

	return nil
}

func validateNewToolkitParameters(props *ToolkitProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

