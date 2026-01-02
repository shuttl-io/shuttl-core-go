package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type Agent interface {
	Model() IModelFactory
	Name() *string
	Outcomes() *[]IOutcome
	SystemPrompt() *string
	Toolkits() *[]Toolkit
	Tools() *[]ITool
	Triggers() *[]ITrigger
	GetTool(name *string) ITool
	GetToolCallResult(callID *string, result interface{}) *map[string]interface{}
	Invoke(prompt interface{}, threadId *string, streamer IModelStreamer, attachments *[]*FileAttachment) IModel
	RespondWithToolCall(modelInstance IModel, callID *string, result interface{}, streamer IModelStreamer)
}

// The jsii proxy struct for Agent
type jsiiProxy_Agent struct {
	_ byte // padding
}

func (j *jsiiProxy_Agent) Model() IModelFactory {
	var returns IModelFactory
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Outcomes() *[]IOutcome {
	var returns *[]IOutcome
	_jsii_.Get(
		j,
		"outcomes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SystemPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Toolkits() *[]Toolkit {
	var returns *[]Toolkit
	_jsii_.Get(
		j,
		"toolkits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Tools() *[]ITool {
	var returns *[]ITool
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Triggers() *[]ITrigger {
	var returns *[]ITrigger
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}


func NewAgent(props *AgentProps) Agent {
	_init_.Initialize()

	if err := validateNewAgentParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Agent{}

	_jsii_.Create(
		"@shuttl-io/core.Agent",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAgent_Override(a Agent, props *AgentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.Agent",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_Agent) GetTool(name *string) ITool {
	if err := a.validateGetToolParameters(name); err != nil {
		panic(err)
	}
	var returns ITool

	_jsii_.Invoke(
		a,
		"getTool",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetToolCallResult(callID *string, result interface{}) *map[string]interface{} {
	if err := a.validateGetToolCallResultParameters(callID, result); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getToolCallResult",
		[]interface{}{callID, result},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) Invoke(prompt interface{}, threadId *string, streamer IModelStreamer, attachments *[]*FileAttachment) IModel {
	if err := a.validateInvokeParameters(prompt, attachments); err != nil {
		panic(err)
	}
	var returns IModel

	_jsii_.Invoke(
		a,
		"invoke",
		[]interface{}{prompt, threadId, streamer, attachments},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) RespondWithToolCall(modelInstance IModel, callID *string, result interface{}, streamer IModelStreamer) {
	if err := a.validateRespondWithToolCallParameters(modelInstance, callID, result, streamer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"respondWithToolCall",
		[]interface{}{modelInstance, callID, result, streamer},
	)
}

