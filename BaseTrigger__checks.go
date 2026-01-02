//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (b *jsiiProxy_BaseTrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	if invoker == nil {
		return fmt.Errorf("parameter invoker is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseTrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	if outcome == nil {
		return fmt.Errorf("parameter outcome is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseTrigger) validateParseArgsParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseTrigger) validateValidateParameters(args interface{}) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseTrigger) validateWithNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseTrigger) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseTrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseTrigger) validateSetTriggerTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBaseTriggerParameters(triggerType *string, config *map[string]interface{}) error {
	if triggerType == nil {
		return fmt.Errorf("parameter triggerType is required, but nil was provided")
	}

	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}

	return nil
}

