package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IModelStreamer interface {
	Recieve(model IModel, content *ModelResponse)
}

// The jsii proxy for IModelStreamer
type jsiiProxy_IModelStreamer struct {
	_ byte // padding
}

func (i *jsiiProxy_IModelStreamer) Recieve(model IModel, content *ModelResponse) {
	if err := i.validateRecieveParameters(model, content); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"recieve",
		[]interface{}{model, content},
	)
}

