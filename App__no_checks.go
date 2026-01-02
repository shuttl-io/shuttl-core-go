//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_App) validateAddAgentParameters(agent Agent) error {
	return nil
}

func (a *jsiiProxy_App) validateAddToolkitParameters(toolkit Toolkit) error {
	return nil
}

func validateNewAppParameters(name *string, server IServer) error {
	return nil
}

