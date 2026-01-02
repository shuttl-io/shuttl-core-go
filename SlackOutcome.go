package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type SlackOutcome interface {
	IOutcome
	BindToRequest(request interface{})
	Send(messageStream IModelResponseStream)
}

// The jsii proxy struct for SlackOutcome
type jsiiProxy_SlackOutcome struct {
	jsiiProxy_IOutcome
}

func NewSlackOutcome() SlackOutcome {
	_init_.Initialize()

	j := jsiiProxy_SlackOutcome{}

	_jsii_.Create(
		"@shuttl-io/core.SlackOutcome",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewSlackOutcome_Override(s SlackOutcome) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.SlackOutcome",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_SlackOutcome) BindToRequest(request interface{}) {
	if err := s.validateBindToRequestParameters(request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bindToRequest",
		[]interface{}{request},
	)
}

func (s *jsiiProxy_SlackOutcome) Send(messageStream IModelResponseStream) {
	if err := s.validateSendParameters(messageStream); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"send",
		[]interface{}{messageStream},
	)
}

