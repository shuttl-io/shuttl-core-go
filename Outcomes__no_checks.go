//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_Outcomes) validateBindToRequestParameters(request interface{}) error {
	return nil
}

func (o *jsiiProxy_Outcomes) validateSendParameters(messageStream IModelResponseStream) error {
	return nil
}

func validateNewOutcomesParameters(outcomes *[]IOutcome) error {
	return nil
}

