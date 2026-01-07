package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IServer interface {
	Accept(app interface{})
	IsRunning() *bool
	Start()
	Stop()
}

// The jsii proxy for IServer
type jsiiProxy_IServer struct {
	_ byte // padding
}

func (i *jsiiProxy_IServer) Accept(app interface{}) {
	if err := i.validateAcceptParameters(app); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"accept",
		[]interface{}{app},
	)
}

func (i *jsiiProxy_IServer) IsRunning() *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"isRunning",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServer) Start() {
	_jsii_.InvokeVoid(
		i,
		"start",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IServer) Stop() {
	_jsii_.InvokeVoid(
		i,
		"stop",
		nil, // no parameters
	)
}

