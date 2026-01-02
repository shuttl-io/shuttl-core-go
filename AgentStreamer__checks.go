//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_AgentStreamer) validateRecieveParameters(model IModel, content *ModelResponse) error {
	if model == nil {
		return fmt.Errorf("parameter model is required, but nil was provided")
	}

	if content == nil {
		return fmt.Errorf("parameter content is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(content, func() string { return "parameter content" }); err != nil {
		return err
	}

	return nil
}

func validateNewAgentStreamerParameters(agent Agent, controlID *string) error {
	if agent == nil {
		return fmt.Errorf("parameter agent is required, but nil was provided")
	}

	if controlID == nil {
		return fmt.Errorf("parameter controlID is required, but nil was provided")
	}

	return nil
}

