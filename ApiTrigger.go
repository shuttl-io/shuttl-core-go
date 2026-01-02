package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

// Represents a trigger that can activate an agent via an API call.
//
// This API trigger is the default trigger for agents.
type ApiTrigger interface {
	BaseTrigger
	// The unique name of this trigger instance.
	//
	// If not set, defaults to triggerType.
	Name() *string
	SetName(val *string)
	Outcome() IOutcome
	SetOutcome(val IOutcome)
	// The configuration for the trigger.
	TriggerConfig() *map[string]interface{}
	SetTriggerConfig(val *map[string]interface{})
	// The type of trigger.
	TriggerType() *string
	SetTriggerType(val *string)
	// Activates the trigger and returns the input for the agent.
	Activate(args interface{}, invoker ITriggerInvoker)
	// binds the outcome to the trigger.
	BindOutcome(outcome IOutcome) ITrigger
	ParseArgs(args interface{}) *TriggerOutput
	// Validates the arguments for the trigger.
	Validate(args interface{}) *map[string]interface{}
	// Sets the name of the trigger.
	WithName(name *string) ITrigger
}

// The jsii proxy struct for ApiTrigger
type jsiiProxy_ApiTrigger struct {
	jsiiProxy_BaseTrigger
}

func (j *jsiiProxy_ApiTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiTrigger) Outcome() IOutcome {
	var returns IOutcome
	_jsii_.Get(
		j,
		"outcome",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiTrigger) TriggerConfig() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiTrigger) TriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerType",
		&returns,
	)
	return returns
}


func NewApiTrigger(config *ApiTriggerConfig) ApiTrigger {
	_init_.Initialize()

	if err := validateNewApiTriggerParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiTrigger{}

	_jsii_.Create(
		"@shuttl-io/core.ApiTrigger",
		[]interface{}{config},
		&j,
	)

	return &j
}

func NewApiTrigger_Override(a ApiTrigger, config *ApiTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.ApiTrigger",
		[]interface{}{config},
		a,
	)
}

func (j *jsiiProxy_ApiTrigger)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiTrigger)SetOutcome(val IOutcome) {
	_jsii_.Set(
		j,
		"outcome",
		val,
	)
}

func (j *jsiiProxy_ApiTrigger)SetTriggerConfig(val *map[string]interface{}) {
	if err := j.validateSetTriggerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfig",
		val,
	)
}

func (j *jsiiProxy_ApiTrigger)SetTriggerType(val *string) {
	if err := j.validateSetTriggerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerType",
		val,
	)
}

func (a *jsiiProxy_ApiTrigger) Activate(args interface{}, invoker ITriggerInvoker) {
	if err := a.validateActivateParameters(args, invoker); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"activate",
		[]interface{}{args, invoker},
	)
}

func (a *jsiiProxy_ApiTrigger) BindOutcome(outcome IOutcome) ITrigger {
	if err := a.validateBindOutcomeParameters(outcome); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		a,
		"bindOutcome",
		[]interface{}{outcome},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiTrigger) ParseArgs(args interface{}) *TriggerOutput {
	if err := a.validateParseArgsParameters(args); err != nil {
		panic(err)
	}
	var returns *TriggerOutput

	_jsii_.Invoke(
		a,
		"parseArgs",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiTrigger) Validate(args interface{}) *map[string]interface{} {
	if err := a.validateValidateParameters(args); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"validate",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiTrigger) WithName(name *string) ITrigger {
	if err := a.validateWithNameParameters(name); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		a,
		"withName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

