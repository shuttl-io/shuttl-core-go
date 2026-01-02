package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IOutcome interface {
	BindToRequest(request interface{})
	Send(messageStream IModelResponseStream)
}

// The jsii proxy for IOutcome
type jsiiProxy_IOutcome struct {
	_ byte // padding
}

func (i *jsiiProxy_IOutcome) BindToRequest(request interface{}) {
	if err := i.validateBindToRequestParameters(request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"bindToRequest",
		[]interface{}{request},
	)
}

func (i *jsiiProxy_IOutcome) Send(messageStream IModelResponseStream) {
	if err := i.validateSendParameters(messageStream); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"send",
		[]interface{}{messageStream},
	)
}

