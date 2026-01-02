//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IModelStreamer) validateRecieveParameters(model IModel, content *ModelResponse) error {
	return nil
}

