//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (f *jsiiProxy_FileTrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	if invoker == nil {
		return fmt.Errorf("parameter invoker is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FileTrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	if outcome == nil {
		return fmt.Errorf("parameter outcome is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FileTrigger) validateParseArgsParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FileTrigger) validateValidateParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (f *jsiiProxy_FileTrigger) validateWithNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FileTrigger) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FileTrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FileTrigger) validateSetTriggerTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewFileTriggerParameters(config *FileTriggerConfig) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

