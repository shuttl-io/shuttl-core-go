package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ITool interface {
	Execute(args *map[string]interface{}) interface{}
	Description() *string
	SetDescription(d *string)
	Name() *string
	SetName(n *string)
	Schema() Schema
	SetSchema(s Schema)
}

// The jsii proxy for ITool
type jsiiProxy_ITool struct {
	_ byte // padding
}

func (i *jsiiProxy_ITool) Execute(args *map[string]interface{}) interface{} {
	if err := i.validateExecuteParameters(args); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"execute",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ITool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ITool) Schema() Schema {
	var returns Schema
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITool)SetSchema(val Schema) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

