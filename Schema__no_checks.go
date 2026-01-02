//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func validateSchema_BooleanValueParameters(description *string) error {
	return nil
}

func validateSchema_EnumValueParameters(description *string, enumValues *[]*string) error {
	return nil
}

func validateSchema_NumberValueParameters(description *string) error {
	return nil
}

func validateSchema_ObjectValueParameters(properties *map[string]ToolArgBuilder) error {
	return nil
}

func validateSchema_StringValueParameters(description *string) error {
	return nil
}

