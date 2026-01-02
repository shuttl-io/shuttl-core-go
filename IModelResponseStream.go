package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IModelResponseStream interface {
	Next() *ModelResponseStreamValue
}

// The jsii proxy for IModelResponseStream
type jsiiProxy_IModelResponseStream struct {
	_ byte // padding
}

func (i *jsiiProxy_IModelResponseStream) Next() *ModelResponseStreamValue {
	var returns *ModelResponseStreamValue

	_jsii_.Invoke(
		i,
		"next",
		nil, // no parameters
		&returns,
	)

	return returns
}

