package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IModel interface {
	Invoke(prompt *[]interface{}, streamer IModelStreamer)
	ThreadId() *string
}

// The jsii proxy for IModel
type jsiiProxy_IModel struct {
	_ byte // padding
}

func (i *jsiiProxy_IModel) Invoke(prompt *[]interface{}, streamer IModelStreamer) {
	if err := i.validateInvokeParameters(prompt, streamer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"invoke",
		[]interface{}{prompt, streamer},
	)
}

func (j *jsiiProxy_IModel) ThreadId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threadId",
		&returns,
	)
	return returns
}

