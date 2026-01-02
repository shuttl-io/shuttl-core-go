package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type EmailTrigger interface {
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

// The jsii proxy struct for EmailTrigger
type jsiiProxy_EmailTrigger struct {
	jsiiProxy_BaseTrigger
}

func (j *jsiiProxy_EmailTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailTrigger) Outcome() IOutcome {
	var returns IOutcome
	_jsii_.Get(
		j,
		"outcome",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailTrigger) TriggerConfig() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailTrigger) TriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerType",
		&returns,
	)
	return returns
}


func NewEmailTrigger(config *EmailTriggerConfig) EmailTrigger {
	_init_.Initialize()

	if err := validateNewEmailTriggerParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmailTrigger{}

	_jsii_.Create(
		"@shuttl-io/core.EmailTrigger",
		[]interface{}{config},
		&j,
	)

	return &j
}

func NewEmailTrigger_Override(e EmailTrigger, config *EmailTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.EmailTrigger",
		[]interface{}{config},
		e,
	)
}

func (j *jsiiProxy_EmailTrigger)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmailTrigger)SetOutcome(val IOutcome) {
	_jsii_.Set(
		j,
		"outcome",
		val,
	)
}

func (j *jsiiProxy_EmailTrigger)SetTriggerConfig(val *map[string]interface{}) {
	if err := j.validateSetTriggerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfig",
		val,
	)
}

func (j *jsiiProxy_EmailTrigger)SetTriggerType(val *string) {
	if err := j.validateSetTriggerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerType",
		val,
	)
}

func (e *jsiiProxy_EmailTrigger) Activate(args interface{}, invoker ITriggerInvoker) {
	if err := e.validateActivateParameters(args, invoker); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"activate",
		[]interface{}{args, invoker},
	)
}

func (e *jsiiProxy_EmailTrigger) BindOutcome(outcome IOutcome) ITrigger {
	if err := e.validateBindOutcomeParameters(outcome); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		e,
		"bindOutcome",
		[]interface{}{outcome},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailTrigger) ParseArgs(args interface{}) *TriggerOutput {
	if err := e.validateParseArgsParameters(args); err != nil {
		panic(err)
	}
	var returns *TriggerOutput

	_jsii_.Invoke(
		e,
		"parseArgs",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailTrigger) Validate(args interface{}) *map[string]interface{} {
	if err := e.validateValidateParameters(args); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"validate",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailTrigger) WithName(name *string) ITrigger {
	if err := e.validateWithNameParameters(name); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		e,
		"withName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

