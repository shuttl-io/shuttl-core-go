//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (i *jsiiProxy_IOutcome) validateBindToRequestParameters(request interface{}) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IOutcome) validateSendParameters(messageStream IModelResponseStream) error {
	if messageStream == nil {
		return fmt.Errorf("parameter messageStream is required, but nil was provided")
	}

	return nil
}

