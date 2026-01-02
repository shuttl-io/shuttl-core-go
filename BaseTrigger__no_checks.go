//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseTrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	return nil
}

func (b *jsiiProxy_BaseTrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	return nil
}

func (b *jsiiProxy_BaseTrigger) validateParseArgsParameters(args interface{}) error {
	return nil
}

func (b *jsiiProxy_BaseTrigger) validateValidateParameters(args interface{}) error {
	return nil
}

func (b *jsiiProxy_BaseTrigger) validateWithNameParameters(name *string) error {
	return nil
}

func (j *jsiiProxy_BaseTrigger) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BaseTrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_BaseTrigger) validateSetTriggerTypeParameters(val *string) error {
	return nil
}

func validateNewBaseTriggerParameters(triggerType *string, config *map[string]interface{}) error {
	return nil
}

