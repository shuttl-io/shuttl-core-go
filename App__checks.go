//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (a *jsiiProxy_App) validateAddAgentParameters(agent Agent) error {
	if agent == nil {
		return fmt.Errorf("parameter agent is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_App) validateAddToolkitParameters(toolkit Toolkit) error {
	if toolkit == nil {
		return fmt.Errorf("parameter toolkit is required, but nil was provided")
	}

	return nil
}

func validateNewAppParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

