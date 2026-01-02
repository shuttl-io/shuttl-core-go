//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmailTrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	return nil
}

func (e *jsiiProxy_EmailTrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	return nil
}

func (e *jsiiProxy_EmailTrigger) validateParseArgsParameters(args interface{}) error {
	return nil
}

func (e *jsiiProxy_EmailTrigger) validateValidateParameters(args interface{}) error {
	return nil
}

func (e *jsiiProxy_EmailTrigger) validateWithNameParameters(name *string) error {
	return nil
}

func (j *jsiiProxy_EmailTrigger) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmailTrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_EmailTrigger) validateSetTriggerTypeParameters(val *string) error {
	return nil
}

func validateNewEmailTriggerParameters(config *EmailTriggerConfig) error {
	return nil
}

