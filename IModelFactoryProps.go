package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IModelFactoryProps interface {
	SystemPrompt() *string
	Tools() *[]ITool
}

// The jsii proxy for IModelFactoryProps
type jsiiProxy_IModelFactoryProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IModelFactoryProps) SystemPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelFactoryProps) Tools() *[]ITool {
	var returns *[]ITool
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}

