//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (s *jsiiProxy_StreamingOutcome) validateBindToRequestParameters(request interface{}) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StreamingOutcome) validateSendParameters(messageStream IModelResponseStream) error {
	if messageStream == nil {
		return fmt.Errorf("parameter messageStream is required, but nil was provided")
	}

	return nil
}

