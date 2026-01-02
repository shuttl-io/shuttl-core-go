//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Rate) validateActivateParameters(args interface{}, invoker ITriggerInvoker) error {
	return nil
}

func (r *jsiiProxy_Rate) validateBindOutcomeParameters(outcome IOutcome) error {
	return nil
}

func (r *jsiiProxy_Rate) validateParseArgsParameters(args interface{}) error {
	return nil
}

func (r *jsiiProxy_Rate) validateValidateParameters(args interface{}) error {
	return nil
}

func (r *jsiiProxy_Rate) validateWithNameParameters(name *string) error {
	return nil
}

func (r *jsiiProxy_Rate) validateWithOnTriggerParameters(onTrigger IRateTriggerOnTrigger) error {
	return nil
}

func validateRate_CronParameters(expression *string) error {
	return nil
}

func validateRate_DaysParameters(value *float64) error {
	return nil
}

func validateRate_HoursParameters(value *float64) error {
	return nil
}

func validateRate_MillisecondsParameters(value *float64) error {
	return nil
}

func validateRate_MinutesParameters(value *float64) error {
	return nil
}

func validateRate_MonthsParameters(value *float64) error {
	return nil
}

func validateRate_SecondsParameters(value *float64) error {
	return nil
}

func validateRate_WeeksParameters(value *float64) error {
	return nil
}

func (j *jsiiProxy_Rate) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Rate) validateSetTriggerConfigParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_Rate) validateSetTriggerTypeParameters(val *string) error {
	return nil
}

