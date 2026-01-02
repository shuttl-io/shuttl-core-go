//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_ApiTrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	if invoker == nil {
		return fmt.Errorf("parameter invoker is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiTrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	if outcome == nil {
		return fmt.Errorf("parameter outcome is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiTrigger) validateParseArgsParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiTrigger) validateValidateParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiTrigger) validateWithNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApiTrigger) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApiTrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ApiTrigger) validateSetTriggerTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewApiTriggerParameters(config *ApiTriggerConfig) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

