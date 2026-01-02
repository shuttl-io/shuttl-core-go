//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Agent) validateGetToolParameters(name *string) error {
	return nil
}

func (a *jsiiProxy_Agent) validateGetToolCallResultParameters(callID *string, result interface{}) error {
	return nil
}

func (a *jsiiProxy_Agent) validateInvokeParameters(prompt interface{}, attachments *[]*FileAttachment) error {
	return nil
}

func (a *jsiiProxy_Agent) validateRespondWithToolCallParameters(modelInstance IModel, callID *string, result interface{}, streamer IModelStreamer) error {
	return nil
}

func validateNewAgentParameters(props *AgentProps) error {
	return nil
}

