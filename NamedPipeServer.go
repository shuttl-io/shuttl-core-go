package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

// A server that communicates via named pipes (FIFOs).
//
// This avoids conflicts with JSII's stdin/stdout usage.
// Supports both Unix FIFOs and Windows named pipes.
//
// The CLI creates named pipes and passes paths via environment variables:
// - _SHUTTL_REQUEST_PIPE: Path to the request pipe (CLI writes, server reads)
// - _SHUTTL_RESPONSE_PIPE: Path to the response pipe (server writes, CLI reads).
type NamedPipeServer interface {
	IServer
	Accept(app interface{})
	IsRunning() *bool
	Start()
	Stop()
}

// The jsii proxy struct for NamedPipeServer
type jsiiProxy_NamedPipeServer struct {
	jsiiProxy_IServer
}

func NewNamedPipeServer() NamedPipeServer {
	_init_.Initialize()

	j := jsiiProxy_NamedPipeServer{}

	_jsii_.Create(
		"@shuttl-io/core.NamedPipeServer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNamedPipeServer_Override(n NamedPipeServer) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.NamedPipeServer",
		nil, // no parameters
		n,
	)
}

func (n *jsiiProxy_NamedPipeServer) Accept(app interface{}) {
	if err := n.validateAcceptParameters(app); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"accept",
		[]interface{}{app},
	)
}

func (n *jsiiProxy_NamedPipeServer) IsRunning() *bool {
	var returns *bool

	_jsii_.Invoke(
		n,
		"isRunning",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NamedPipeServer) Start() {
	_jsii_.InvokeVoid(
		n,
		"start",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NamedPipeServer) Stop() {
	_jsii_.InvokeVoid(
		n,
		"stop",
		nil, // no parameters
	)
}

