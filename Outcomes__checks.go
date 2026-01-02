//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (o *jsiiProxy_Outcomes) validateBindToRequestParameters(request interface{}) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_Outcomes) validateSendParameters(messageStream IModelResponseStream) error {
	if messageStream == nil {
		return fmt.Errorf("parameter messageStream is required, but nil was provided")
	}

	return nil
}

func validateNewOutcomesParameters(outcomes *[]IOutcome) error {
	if outcomes == nil {
		return fmt.Errorf("parameter outcomes is required, but nil was provided")
	}

	return nil
}

