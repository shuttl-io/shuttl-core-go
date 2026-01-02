package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type AgentStreamer interface {
	IModelStreamer
	Recieve(model IModel, content *ModelResponse)
}

// The jsii proxy struct for AgentStreamer
type jsiiProxy_AgentStreamer struct {
	jsiiProxy_IModelStreamer
}

func NewAgentStreamer(agent Agent, controlID *string, writer IAgentStreamerWriter) AgentStreamer {
	_init_.Initialize()

	if err := validateNewAgentStreamerParameters(agent, controlID); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentStreamer{}

	_jsii_.Create(
		"@shuttl-io/core.AgentStreamer",
		[]interface{}{agent, controlID, writer},
		&j,
	)

	return &j
}

func NewAgentStreamer_Override(a AgentStreamer, agent Agent, controlID *string, writer IAgentStreamerWriter) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.AgentStreamer",
		[]interface{}{agent, controlID, writer},
		a,
	)
}

func (a *jsiiProxy_AgentStreamer) Recieve(model IModel, content *ModelResponse) {
	if err := a.validateRecieveParameters(model, content); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"recieve",
		[]interface{}{model, content},
	)
}

