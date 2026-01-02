//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITriggerInvoker) validateDefaultOutcomeParameters(stream IModelResponseStream) error {
	return nil
}

func (i *jsiiProxy_ITriggerInvoker) validateInvokeParameters(prompt *[]*InputContent) error {
	return nil
}

