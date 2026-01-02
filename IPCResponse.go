package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Response message format sent back to the host CLI.
type IPCResponse interface {
	// Error information (on failure).
	ErrorObj() IPCResponseError
	SetErrorObj(e IPCResponseError)
	// Correlates to the request ID.
	Id() *string
	SetId(i *string)
	// The result data (on success).
	Result() interface{}
	SetResult(r interface{})
	// Whether the request was successful.
	Success() *bool
	SetSuccess(s *bool)
}

// The jsii proxy for IPCResponse
type jsiiProxy_IPCResponse struct {
	_ byte // padding
}

func (j *jsiiProxy_IPCResponse) ErrorObj() IPCResponseError {
	var returns IPCResponseError
	_jsii_.Get(
		j,
		"errorObj",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCResponse)SetErrorObj(val IPCResponseError) {
	_jsii_.Set(
		j,
		"errorObj",
		val,
	)
}

func (j *jsiiProxy_IPCResponse) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCResponse)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IPCResponse) Result() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCResponse)SetResult(val interface{}) {
	_jsii_.Set(
		j,
		"result",
		val,
	)
}

func (j *jsiiProxy_IPCResponse) Success() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCResponse)SetSuccess(val *bool) {
	if err := j.validateSetSuccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"success",
		val,
	)
}

