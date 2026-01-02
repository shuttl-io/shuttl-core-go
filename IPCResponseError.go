package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPCResponseError interface {
	Code() *string
	SetCode(c *string)
	Message() *string
	SetMessage(m *string)
}

// The jsii proxy for IPCResponseError
type jsiiProxy_IPCResponseError struct {
	_ byte // padding
}

func (j *jsiiProxy_IPCResponseError) Code() *string {
	var returns *string
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCResponseError)SetCode(val *string) {
	if err := j.validateSetCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_IPCResponseError) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPCResponseError)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

