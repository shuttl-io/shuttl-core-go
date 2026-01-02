//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAgentStreamerWriter) validateWriteParameters(value *string) error {
	return nil
}

func (i *jsiiProxy_IAgentStreamerWriter) validateWriteObjectParameters(value interface{}) error {
	return nil
}

