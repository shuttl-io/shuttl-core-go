package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ITriggerInvoker interface {
	DefaultOutcome(stream IModelResponseStream)
	Invoke(prompt *[]*InputContent) IModelResponseStream
}

// The jsii proxy for ITriggerInvoker
type jsiiProxy_ITriggerInvoker struct {
	_ byte // padding
}

func (i *jsiiProxy_ITriggerInvoker) DefaultOutcome(stream IModelResponseStream) {
	if err := i.validateDefaultOutcomeParameters(stream); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"defaultOutcome",
		[]interface{}{stream},
	)
}

func (i *jsiiProxy_ITriggerInvoker) Invoke(prompt *[]*InputContent) IModelResponseStream {
	if err := i.validateInvokeParameters(prompt); err != nil {
		panic(err)
	}
	var returns IModelResponseStream

	_jsii_.Invoke(
		i,
		"invoke",
		[]interface{}{prompt},
		&returns,
	)

	return returns
}

