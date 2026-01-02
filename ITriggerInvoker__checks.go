//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_ITriggerInvoker) validateDefaultOutcomeParameters(stream IModelResponseStream) error {
	if stream == nil {
		return fmt.Errorf("parameter stream is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITriggerInvoker) validateInvokeParameters(prompt *[]*InputContent) error {
	if prompt == nil {
		return fmt.Errorf("parameter prompt is required, but nil was provided")
	}
	for idx_cf0719, v := range *prompt {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter prompt[%#v]", idx_cf0719) }); err != nil {
			return err
		}
	}

	return nil
}

