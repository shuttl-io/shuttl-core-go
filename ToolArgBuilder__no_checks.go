//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_ToolArgBuilder) validateDefaultToParameters(defaultValue interface{}) error {
	return nil
}

func (j *jsiiProxy_ToolArgBuilder) validateSetArgTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ToolArgBuilder) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ToolArgBuilder) validateSetRequiredParameters(val *bool) error {
	return nil
}

func validateNewToolArgBuilderParameters(argType *string, description *string, required *bool) error {
	return nil
}

