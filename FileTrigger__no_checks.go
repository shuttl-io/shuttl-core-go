//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileTrigger) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	return nil
}

func (f *jsiiProxy_FileTrigger) validateBindOutcomeParameters(outcome IOutcome) error {
	return nil
}

func (f *jsiiProxy_FileTrigger) validateParseArgsParameters(args interface{}) error {
	return nil
}

func (f *jsiiProxy_FileTrigger) validateValidateParameters(args interface{}) error {
	return nil
}

func (f *jsiiProxy_FileTrigger) validateWithNameParameters(name *string) error {
	return nil
}

func (j *jsiiProxy_FileTrigger) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FileTrigger) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_FileTrigger) validateSetTriggerTypeParameters(val *string) error {
	return nil
}

func validateNewFileTriggerParameters(config *FileTriggerConfig) error {
	return nil
}

