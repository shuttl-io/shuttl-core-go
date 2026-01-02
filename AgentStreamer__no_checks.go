//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AgentStreamer) validateRecieveParameters(model IModel, content *ModelResponse) error {
	return nil
}

func validateNewAgentStreamerParameters(agent Agent, controlID *string) error {
	return nil
}

