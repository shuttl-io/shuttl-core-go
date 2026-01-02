//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITool) validateExecuteParameters(args *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_ITool) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ITool) validateSetNameParameters(val *string) error {
	return nil
}

