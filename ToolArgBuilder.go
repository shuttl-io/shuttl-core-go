package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type ToolArgBuilder interface {
	ArgType() *string
	SetArgType(val *string)
	DefaultValue() interface{}
	SetDefaultValue(val interface{})
	Description() *string
	SetDescription(val *string)
	EnumValues() *[]*string
	SetEnumValues(val *[]*string)
	Required() *bool
	SetRequired(val *bool)
	DefaultTo(defaultValue interface{}) ToolArgBuilder
	IsRequired() ToolArgBuilder
}

// The jsii proxy struct for ToolArgBuilder
type jsiiProxy_ToolArgBuilder struct {
	_ byte // padding
}

func (j *jsiiProxy_ToolArgBuilder) ArgType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"argType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolArgBuilder) DefaultValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolArgBuilder) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolArgBuilder) EnumValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolArgBuilder) Required() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}


func NewToolArgBuilder(argType *string, description *string, required *bool, defaultValue interface{}, enumValues *[]*string) ToolArgBuilder {
	_init_.Initialize()

	if err := validateNewToolArgBuilderParameters(argType, description, required); err != nil {
		panic(err)
	}
	j := jsiiProxy_ToolArgBuilder{}

	_jsii_.Create(
		"@shuttl-io/core.ToolArgBuilder",
		[]interface{}{argType, description, required, defaultValue, enumValues},
		&j,
	)

	return &j
}

func NewToolArgBuilder_Override(t ToolArgBuilder, argType *string, description *string, required *bool, defaultValue interface{}, enumValues *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.ToolArgBuilder",
		[]interface{}{argType, description, required, defaultValue, enumValues},
		t,
	)
}

func (j *jsiiProxy_ToolArgBuilder)SetArgType(val *string) {
	if err := j.validateSetArgTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"argType",
		val,
	)
}

func (j *jsiiProxy_ToolArgBuilder)SetDefaultValue(val interface{}) {
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_ToolArgBuilder)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ToolArgBuilder)SetEnumValues(val *[]*string) {
	_jsii_.Set(
		j,
		"enumValues",
		val,
	)
}

func (j *jsiiProxy_ToolArgBuilder)SetRequired(val *bool) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (t *jsiiProxy_ToolArgBuilder) DefaultTo(defaultValue interface{}) ToolArgBuilder {
	if err := t.validateDefaultToParameters(defaultValue); err != nil {
		panic(err)
	}
	var returns ToolArgBuilder

	_jsii_.Invoke(
		t,
		"defaultTo",
		[]interface{}{defaultValue},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_ToolArgBuilder) IsRequired() ToolArgBuilder {
	var returns ToolArgBuilder

	_jsii_.Invoke(
		t,
		"isRequired",
		nil, // no parameters
		&returns,
	)

	return returns
}

