package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Request message format from the host CLI.
type IPCRequest interface {
	// Optional parameters for the method.
	Body() *map[string]interface{}
	SetBody(b *map[string]interface{})
	// Unique request ID for correlation.
	Id() *string
	SetId(i *string)
	// The method/action to invoke.
	Method() *string
	SetMethod(m *string)
}

// The jsii proxy for IPCRequest
type jsiiProxy_IPCRequest struct {
	_ byte // padding
}

func (j *jsiiProxy_IPCRequest) Body() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCRequest)SetBody(val *map[string]interface{}) {
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_IPCRequest) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCRequest)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IPCRequest) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCRequest)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

