package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type Schema interface {
	Properties() *map[string]ToolArgBuilder
}

// The jsii proxy struct for Schema
type jsiiProxy_Schema struct {
	_ byte // padding
}

func (j *jsiiProxy_Schema) Properties() *map[string]ToolArgBuilder {
	var returns *map[string]ToolArgBuilder
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}


func Schema_BooleanValue(description *string) ToolArgBuilder {
	_init_.Initialize()

	if err := validateSchema_BooleanValueParameters(description); err != nil {
		panic(err)
	}
	var returns ToolArgBuilder

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Schema",
		"booleanValue",
		[]interface{}{description},
		&returns,
	)

	return returns
}

func Schema_EnumValue(description *string, enumValues *[]*string) ToolArgBuilder {
	_init_.Initialize()

	if err := validateSchema_EnumValueParameters(description, enumValues); err != nil {
		panic(err)
	}
	var returns ToolArgBuilder

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Schema",
		"enumValue",
		[]interface{}{description, enumValues},
		&returns,
	)

	return returns
}

func Schema_NumberValue(description *string) ToolArgBuilder {
	_init_.Initialize()

	if err := validateSchema_NumberValueParameters(description); err != nil {
		panic(err)
	}
	var returns ToolArgBuilder

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Schema",
		"numberValue",
		[]interface{}{description},
		&returns,
	)

	return returns
}

func Schema_ObjectValue(properties *map[string]ToolArgBuilder) Schema {
	_init_.Initialize()

	if err := validateSchema_ObjectValueParameters(properties); err != nil {
		panic(err)
	}
	var returns Schema

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Schema",
		"objectValue",
		[]interface{}{properties},
		&returns,
	)

	return returns
}

func Schema_StringValue(description *string) ToolArgBuilder {
	_init_.Initialize()

	if err := validateSchema_StringValueParameters(description); err != nil {
		panic(err)
	}
	var returns ToolArgBuilder

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Schema",
		"stringValue",
		[]interface{}{description},
		&returns,
	)

	return returns
}

