//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Toolkit) validateAddToolParameters(tool ITool) error {
	return nil
}

func validateNewToolkitParameters(props *ToolkitProps) error {
	return nil
}

