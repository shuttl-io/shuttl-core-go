package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type Secret interface {
	Name() *string
	SetName(val *string)
	Source() *string
	SetSource(val *string)
}

// The jsii proxy struct for Secret
type jsiiProxy_Secret struct {
	_ byte // padding
}

func (j *jsiiProxy_Secret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_Secret)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Secret)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

// Create a Secret from an environment variable.
//
// Returns: A new Secret.
func Secret_FromEnv(envVarName *string) ISecret {
	_init_.Initialize()

	if err := validateSecret_FromEnvParameters(envVarName); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Secret",
		"fromEnv",
		[]interface{}{envVarName},
		&returns,
	)

	return returns
}

