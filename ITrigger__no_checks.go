//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	return nil
}

func (i *jsiiProxy_ITrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	return nil
}

func (i *jsiiProxy_ITrigger) validateValidateParameters(args interface{}) error {
	return nil
}

func (i *jsiiProxy_ITrigger) validateWithNameParameters(name *string) error {
	return nil
}

func (j *jsiiProxy_ITrigger) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ITrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_ITrigger) validateSetTriggerTypeParameters(val *string) error {
	return nil
}

