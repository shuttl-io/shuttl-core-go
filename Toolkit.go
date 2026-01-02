package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type Toolkit interface {
	Description() *string
	Name() *string
	Tools() *[]ITool
	AddTool(tool ITool)
}

// The jsii proxy struct for Toolkit
type jsiiProxy_Toolkit struct {
	_ byte // padding
}

func (j *jsiiProxy_Toolkit) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Toolkit) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Toolkit) Tools() *[]ITool {
	var returns *[]ITool
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}


func NewToolkit(props *ToolkitProps) Toolkit {
	_init_.Initialize()

	if err := validateNewToolkitParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Toolkit{}

	_jsii_.Create(
		"@shuttl-io/core.Toolkit",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewToolkit_Override(t Toolkit, props *ToolkitProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.Toolkit",
		[]interface{}{props},
		t,
	)
}

func (t *jsiiProxy_Toolkit) AddTool(tool ITool) {
	if err := t.validateAddToolParameters(tool); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addTool",
		[]interface{}{tool},
	)
}

