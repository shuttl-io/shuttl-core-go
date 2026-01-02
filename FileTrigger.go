package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type FileTrigger interface {
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

// The jsii proxy struct for FileTrigger
type jsiiProxy_FileTrigger struct {
	jsiiProxy_BaseTrigger
}

func (j *jsiiProxy_FileTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileTrigger) Outcome() IOutcome {
	var returns IOutcome
	_jsii_.Get(
		j,
		"outcome",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileTrigger) TriggerConfig() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileTrigger) TriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerType",
		&returns,
	)
	return returns
}


func NewFileTrigger(config *FileTriggerConfig) FileTrigger {
	_init_.Initialize()

	if err := validateNewFileTriggerParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FileTrigger{}

	_jsii_.Create(
		"@shuttl-io/core.FileTrigger",
		[]interface{}{config},
		&j,
	)

	return &j
}

func NewFileTrigger_Override(f FileTrigger, config *FileTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@shuttl-io/core.FileTrigger",
		[]interface{}{config},
		f,
	)
}

func (j *jsiiProxy_FileTrigger)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FileTrigger)SetOutcome(val IOutcome) {
	_jsii_.Set(
		j,
		"outcome",
		val,
	)
}

func (j *jsiiProxy_FileTrigger)SetTriggerConfig(val *map[string]interface{}) {
	if err := j.validateSetTriggerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfig",
		val,
	)
}

func (j *jsiiProxy_FileTrigger)SetTriggerType(val *string) {
	if err := j.validateSetTriggerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerType",
		val,
	)
}

func (f *jsiiProxy_FileTrigger) Activate(args interface{}, invoker ITriggerInvoker) {
	if err := f.validateActivateParameters(args, invoker); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"activate",
		[]interface{}{args, invoker},
	)
}

func (f *jsiiProxy_FileTrigger) BindOutcome(outcome IOutcome) ITrigger {
	if err := f.validateBindOutcomeParameters(outcome); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		f,
		"bindOutcome",
		[]interface{}{outcome},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileTrigger) ParseArgs(args interface{}) *TriggerOutput {
	if err := f.validateParseArgsParameters(args); err != nil {
		panic(err)
	}
	var returns *TriggerOutput

	_jsii_.Invoke(
		f,
		"parseArgs",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileTrigger) Validate(args interface{}) *map[string]interface{} {
	if err := f.validateValidateParameters(args); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"validate",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileTrigger) WithName(name *string) ITrigger {
	if err := f.validateWithNameParameters(name); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		f,
		"withName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

