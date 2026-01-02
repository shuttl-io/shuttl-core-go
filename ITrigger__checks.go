//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (i *jsiiProxy_ITrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	if invoker == nil {
		return fmt.Errorf("parameter invoker is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	if outcome == nil {
		return fmt.Errorf("parameter outcome is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITrigger) validateValidateParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITrigger) validateWithNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ITrigger) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ITrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ITrigger) validateSetTriggerTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

