package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a trigger that can activate an agent.
//
// Triggers can take any arguments and then return what the input should be for the agent.
// Triggers also can validate the arguments and return an error if the arguments are invalid.
type ITrigger interface {
	// Activates the trigger and returns the input for the agent.
	//
	// Returns: The input for the agent.
	Activate(args interface{}, invoker ITriggerInvoker)
	// binds the outcome to the trigger.
	//
	// Returns: The bound outcome.
	BindOutcome(outcome IOutcome) ITrigger
	// Validates the arguments for the trigger.
	//
	// Returns: The validation result.
	Validate(args interface{}) *map[string]interface{}
	// Sets the name of the trigger.
	//
	// Returns: The trigger instance for chaining.
	WithName(name *string) ITrigger
	// The unique name of this trigger instance.
	//
	// If not set, defaults to triggerType.
	Name() *string
	SetName(n *string)
	Outcome() IOutcome
	SetOutcome(o IOutcome)
	// The configuration for the trigger.
	TriggerConfig() *map[string]interface{}
	SetTriggerConfig(t *map[string]interface{})
	// The type of trigger.
	TriggerType() *string
	SetTriggerType(t *string)
}

// The jsii proxy for ITrigger
type jsiiProxy_ITrigger struct {
	_ byte // padding
}

func (i *jsiiProxy_ITrigger) Activate(args interface{}, invoker ITriggerInvoker) {
	if err := i.validateActivateParameters(args, invoker); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"activate",
		[]interface{}{args, invoker},
	)
}

func (i *jsiiProxy_ITrigger) BindOutcome(outcome IOutcome) ITrigger {
	if err := i.validateBindOutcomeParameters(outcome); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		i,
		"bindOutcome",
		[]interface{}{outcome},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITrigger) Validate(args interface{}) *map[string]interface{} {
	if err := i.validateValidateParameters(args); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"validate",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITrigger) WithName(name *string) ITrigger {
	if err := i.validateWithNameParameters(name); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		i,
		"withName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrigger)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ITrigger) Outcome() IOutcome {
	var returns IOutcome
	_jsii_.Get(
		j,
		"outcome",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrigger)SetOutcome(val IOutcome) {
	_jsii_.Set(
		j,
		"outcome",
		val,
	)
}

func (j *jsiiProxy_ITrigger) TriggerConfig() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrigger)SetTriggerConfig(val *map[string]interface{}) {
	if err := j.validateSetTriggerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfig",
		val,
	)
}

func (j *jsiiProxy_ITrigger) TriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrigger)SetTriggerType(val *string) {
	if err := j.validateSetTriggerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerType",
		val,
	)
}

