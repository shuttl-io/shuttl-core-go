//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamingOutcome) validateBindToRequestParameters(request interface{}) error {
	return nil
}

func (s *jsiiProxy_StreamingOutcome) validateSendParameters(messageStream IModelResponseStream) error {
	return nil
}

