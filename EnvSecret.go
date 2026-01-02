package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type EnvSecret interface {
	ISecret
	EnvVarName() *string
	ResolveSecret() *string
}

// The jsii proxy struct for EnvSecret
type jsiiProxy_EnvSecret struct {
	jsiiProxy_ISecret
}

func (j *jsiiProxy_EnvSecret) EnvVarName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envVarName",
		&returns,
	)
	return returns
}


func NewEnvSecret(envVarName *string) EnvSecret {
	_init_.Initialize()

	if err := validateNewEnvSecretParameters(envVarName); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnvSecret{}

	_jsii_.Create(
		"@shuttl-io/core.EnvSecret",
		[]interface{}{envVarName},
		&j,
	)

	return &j
}

func NewEnvSecret_Override(e EnvSecret, envVarName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.EnvSecret",
		[]interface{}{envVarName},
		e,
	)
}

func (e *jsiiProxy_EnvSecret) ResolveSecret() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"resolveSecret",
		nil, // no parameters
		&returns,
	)

	return returns
}

