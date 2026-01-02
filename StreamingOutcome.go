package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type StreamingOutcome interface {
	IOutcome
	BindToRequest(request interface{})
	Send(messageStream IModelResponseStream)
}

// The jsii proxy struct for StreamingOutcome
type jsiiProxy_StreamingOutcome struct {
	jsiiProxy_IOutcome
}

func NewStreamingOutcome() StreamingOutcome {
	_init_.Initialize()

	j := jsiiProxy_StreamingOutcome{}

	_jsii_.Create(
		"@shuttl-io/core.StreamingOutcome",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStreamingOutcome_Override(s StreamingOutcome) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.StreamingOutcome",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_StreamingOutcome) BindToRequest(request interface{}) {
	if err := s.validateBindToRequestParameters(request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bindToRequest",
		[]interface{}{request},
	)
}

func (s *jsiiProxy_StreamingOutcome) Send(messageStream IModelResponseStream) {
	if err := s.validateSendParameters(messageStream); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"send",
		[]interface{}{messageStream},
	)
}

