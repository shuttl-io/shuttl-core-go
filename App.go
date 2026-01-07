package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type App interface {
	Agents() *[]Agent
	Name() *string
	Server() IServer
	AddAgent(agent Agent)
	AddToolkit(toolkit Toolkit)
	Serve() *bool
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	_ byte // padding
}

func (j *jsiiProxy_App) Agents() *[]Agent {
	var returns *[]Agent
	_jsii_.Get(
		j,
		"agents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Server() IServer {
	var returns IServer
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}


func NewApp(name *string, server IServer) App {
	_init_.Initialize()

	if err := validateNewAppParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_App{}

	_jsii_.Create(
		"@shuttl-io/core.App",
		[]interface{}{name, server},
		&j,
	)

	return &j
}

func NewApp_Override(a App, name *string, server IServer) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.App",
		[]interface{}{name, server},
		a,
	)
}

func (a *jsiiProxy_App) AddAgent(agent Agent) {
	if err := a.validateAddAgentParameters(agent); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAgent",
		[]interface{}{agent},
	)
}

func (a *jsiiProxy_App) AddToolkit(toolkit Toolkit) {
	if err := a.validateAddToolkitParameters(toolkit); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addToolkit",
		[]interface{}{toolkit},
	)
}

func (a *jsiiProxy_App) Serve() *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"serve",
		nil, // no parameters
		&returns,
	)

	return returns
}

