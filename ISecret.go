package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ISecret interface {
	ResolveSecret() *string
}

// The jsii proxy for ISecret
type jsiiProxy_ISecret struct {
	_ byte // padding
}

func (i *jsiiProxy_ISecret) ResolveSecret() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"resolveSecret",
		nil, // no parameters
		&returns,
	)

	return returns
}

