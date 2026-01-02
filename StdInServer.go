package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

// A server that communicates via STDIN/STDOUT using JSON messages.
//
// Uses newline-delimited JSON (NDJSON) protocol where each message
// is a single line of JSON followed by a newline character.
//
// Host CLI writes requests to the process's STDIN.
// This server writes responses to STDOUT.
// Debug/log messages go to STDERR to avoid polluting the protocol.
type StdInServer interface {
	IServer
	Accept(app interface{})
	Start()
	Stop()
}

// The jsii proxy struct for StdInServer
type jsiiProxy_StdInServer struct {
	jsiiProxy_IServer
}

func NewStdInServer() StdInServer {
	_init_.Initialize()

	j := jsiiProxy_StdInServer{}

	_jsii_.Create(
		"@shuttl-io/core.StdInServer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStdInServer_Override(s StdInServer) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.StdInServer",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_StdInServer) Accept(app interface{}) {
	if err := s.validateAcceptParameters(app); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"accept",
		[]interface{}{app},
	)
}

func (s *jsiiProxy_StdInServer) Start() {
	_jsii_.InvokeVoid(
		s,
		"start",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StdInServer) Stop() {
	_jsii_.InvokeVoid(
		s,
		"stop",
		nil, // no parameters
	)
}

