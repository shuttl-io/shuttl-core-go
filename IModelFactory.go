package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IModelFactory interface {
	Create(props IModelFactoryProps) IModel
}

// The jsii proxy for IModelFactory
type jsiiProxy_IModelFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_IModelFactory) Create(props IModelFactoryProps) IModel {
	if err := i.validateCreateParameters(props); err != nil {
		panic(err)
	}
	var returns IModel

	_jsii_.Invoke(
		i,
		"create",
		[]interface{}{props},
		&returns,
	)

	return returns
}

