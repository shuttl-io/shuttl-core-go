package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAgentStreamerWriter interface {
	Write(value *string)
	WriteObject(value interface{})
}

// The jsii proxy for IAgentStreamerWriter
type jsiiProxy_IAgentStreamerWriter struct {
	_ byte // padding
}

func (i *jsiiProxy_IAgentStreamerWriter) Write(value *string) {
	if err := i.validateWriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"write",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IAgentStreamerWriter) WriteObject(value interface{}) {
	if err := i.validateWriteObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"writeObject",
		[]interface{}{value},
	)
}

