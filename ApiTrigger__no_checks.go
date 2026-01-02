//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiTrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	return nil
}

func (a *jsiiProxy_ApiTrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	return nil
}

func (a *jsiiProxy_ApiTrigger) validateParseArgsParameters(args interface{}) error {
	return nil
}

func (a *jsiiProxy_ApiTrigger) validateValidateParameters(args interface{}) error {
	return nil
}

func (a *jsiiProxy_ApiTrigger) validateWithNameParameters(name *string) error {
	return nil
}

func (j *jsiiProxy_ApiTrigger) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiTrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiTrigger) validateSetTriggerTypeParameters(val *string) error {
	return nil
}

func validateNewApiTriggerParameters(config *ApiTriggerConfig) error {
	return nil
}

