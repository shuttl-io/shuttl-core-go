//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IOutcome) validateBindToRequestParameters(request interface{}) error {
	return nil
}

func (i *jsiiProxy_IOutcome) validateSendParameters(messageStream IModelResponseStream) error {
	return nil
}

