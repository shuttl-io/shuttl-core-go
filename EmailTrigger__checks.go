//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (e *jsiiProxy_EmailTrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	if invoker == nil {
		return fmt.Errorf("parameter invoker is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EmailTrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	if outcome == nil {
		return fmt.Errorf("parameter outcome is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EmailTrigger) validateParseArgsParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EmailTrigger) validateValidateParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EmailTrigger) validateWithNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EmailTrigger) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EmailTrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_EmailTrigger) validateSetTriggerTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEmailTriggerParameters(config *EmailTriggerConfig) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

