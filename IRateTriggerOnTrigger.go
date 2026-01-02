package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IRateTriggerOnTrigger interface {
	OnTrigger() *[]*InputContent
}

// The jsii proxy for IRateTriggerOnTrigger
type jsiiProxy_IRateTriggerOnTrigger struct {
	_ byte // padding
}

func (i *jsiiProxy_IRateTriggerOnTrigger) OnTrigger() *[]*InputContent {
	var returns *[]*InputContent

	_jsii_.Invoke(
		i,
		"onTrigger",
		nil, // no parameters
		&returns,
	)

	return returns
}

