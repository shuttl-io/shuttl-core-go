package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type Outcomes interface {
	IOutcome
	BindToRequest(request interface{})
	Send(messageStream IModelResponseStream)
}

// The jsii proxy struct for Outcomes
type jsiiProxy_Outcomes struct {
	jsiiProxy_IOutcome
}

func NewOutcomes(outcomes *[]IOutcome) Outcomes {
	_init_.Initialize()

	if err := validateNewOutcomesParameters(outcomes); err != nil {
		panic(err)
	}
	j := jsiiProxy_Outcomes{}

	_jsii_.Create(
		"@shuttl-io/core.Outcomes",
		[]interface{}{outcomes},
		&j,
	)

	return &j
}

func NewOutcomes_Override(o Outcomes, outcomes *[]IOutcome) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.Outcomes",
		[]interface{}{outcomes},
		o,
	)
}

func Outcomes_Combine(outcomes ...IOutcome) Outcomes {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range outcomes {
		args = append(args, a)
	}

	var returns Outcomes

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Outcomes",
		"combine",
		args,
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Outcomes) BindToRequest(request interface{}) {
	if err := o.validateBindToRequestParameters(request); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"bindToRequest",
		[]interface{}{request},
	)
}

func (o *jsiiProxy_Outcomes) Send(messageStream IModelResponseStream) {
	if err := o.validateSendParameters(messageStream); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"send",
		[]interface{}{messageStream},
	)
}

