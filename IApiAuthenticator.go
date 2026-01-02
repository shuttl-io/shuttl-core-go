package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IApiAuthenticator interface {
	Authenticate(args *ApiTriggerArgs) *bool
}

// The jsii proxy for IApiAuthenticator
type jsiiProxy_IApiAuthenticator struct {
	_ byte // padding
}

func (i *jsiiProxy_IApiAuthenticator) Authenticate(args *ApiTriggerArgs) *bool {
	if err := i.validateAuthenticateParameters(args); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"authenticate",
		[]interface{}{args},
		&returns,
	)

	return returns
}

