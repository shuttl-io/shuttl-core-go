package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type Model interface {
}

// The jsii proxy struct for Model
type jsiiProxy_Model struct {
	_ byte // padding
}

func NewModel() Model {
	_init_.Initialize()

	j := jsiiProxy_Model{}

	_jsii_.Create(
		"@shuttl-io/core.Model",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewModel_Override(m Model) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.Model",
		nil, // no parameters
		m,
	)
}

func Model_OpenAI(identifier *string, apiKey ISecret) IModelFactory {
	_init_.Initialize()

	if err := validateModel_OpenAIParameters(identifier, apiKey); err != nil {
		panic(err)
	}
	var returns IModelFactory

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Model",
		"openAI",
		[]interface{}{identifier, apiKey},
		&returns,
	)

	return returns
}

