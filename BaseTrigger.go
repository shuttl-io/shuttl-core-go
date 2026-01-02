package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type BaseTrigger interface {
	ITrigger
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

// The jsii proxy struct for BaseTrigger
type jsiiProxy_BaseTrigger struct {
	jsiiProxy_ITrigger
}

func (j *jsiiProxy_BaseTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseTrigger) Outcome() IOutcome {
	var returns IOutcome
	_jsii_.Get(
		j,
		"outcome",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseTrigger) TriggerConfig() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseTrigger) TriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerType",
		&returns,
	)
	return returns
}


func NewBaseTrigger_Override(b BaseTrigger, triggerType *string, config *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.BaseTrigger",
		[]interface{}{triggerType, config},
		b,
	)
}

func (j *jsiiProxy_BaseTrigger)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BaseTrigger)SetOutcome(val IOutcome) {
	_jsii_.Set(
		j,
		"outcome",
		val,
	)
}

func (j *jsiiProxy_BaseTrigger)SetTriggerConfig(val *map[string]interface{}) {
	if err := j.validateSetTriggerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfig",
		val,
	)
}

func (j *jsiiProxy_BaseTrigger)SetTriggerType(val *string) {
	if err := j.validateSetTriggerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerType",
		val,
	)
}

func (b *jsiiProxy_BaseTrigger) Activate(args interface{}, invoker ITriggerInvoker) {
	if err := b.validateActivateParameters(args, invoker); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"activate",
		[]interface{}{args, invoker},
	)
}

func (b *jsiiProxy_BaseTrigger) BindOutcome(outcome IOutcome) ITrigger {
	if err := b.validateBindOutcomeParameters(outcome); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		b,
		"bindOutcome",
		[]interface{}{outcome},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseTrigger) ParseArgs(args interface{}) *TriggerOutput {
	if err := b.validateParseArgsParameters(args); err != nil {
		panic(err)
	}
	var returns *TriggerOutput

	_jsii_.Invoke(
		b,
		"parseArgs",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseTrigger) Validate(args interface{}) *map[string]interface{} {
	if err := b.validateValidateParameters(args); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"validate",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseTrigger) WithName(name *string) ITrigger {
	if err := b.validateWithNameParameters(name); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		b,
		"withName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

