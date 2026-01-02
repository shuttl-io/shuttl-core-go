package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/shuttl-io/shuttl-core-go/core/jsii"
)

type Rate interface {
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
	WithOnTrigger(onTrigger IRateTriggerOnTrigger) Rate
}

// The jsii proxy struct for Rate
type jsiiProxy_Rate struct {
	jsiiProxy_BaseTrigger
}

func (j *jsiiProxy_Rate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rate) Outcome() IOutcome {
	var returns IOutcome
	_jsii_.Get(
		j,
		"outcome",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rate) TriggerConfig() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rate) TriggerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerType",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_Rate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Rate)SetOutcome(val IOutcome) {
	_jsii_.Set(
		j,
		"outcome",
		val,
	)
}

func (j *jsiiProxy_Rate)SetTriggerConfig(val *map[string]interface{}) {
	if err := j.validateSetTriggerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfig",
		val,
	)
}

func (j *jsiiProxy_Rate)SetTriggerType(val *string) {
	if err := j.validateSetTriggerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerType",
		val,
	)
}

func Rate_Cron(expression *string, timezone *string) Rate {
	_init_.Initialize()

	if err := validateRate_CronParameters(expression); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Rate",
		"cron",
		[]interface{}{expression, timezone},
		&returns,
	)

	return returns
}

func Rate_Days(value *float64) Rate {
	_init_.Initialize()

	if err := validateRate_DaysParameters(value); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Rate",
		"days",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Rate_Hours(value *float64) Rate {
	_init_.Initialize()

	if err := validateRate_HoursParameters(value); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Rate",
		"hours",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Rate_Milliseconds(value *float64) Rate {
	_init_.Initialize()

	if err := validateRate_MillisecondsParameters(value); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Rate",
		"milliseconds",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Rate_Minutes(value *float64) Rate {
	_init_.Initialize()

	if err := validateRate_MinutesParameters(value); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Rate",
		"minutes",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Rate_Months(value *float64) Rate {
	_init_.Initialize()

	if err := validateRate_MonthsParameters(value); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Rate",
		"months",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Rate_Seconds(value *float64) Rate {
	_init_.Initialize()

	if err := validateRate_SecondsParameters(value); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Rate",
		"seconds",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Rate_Weeks(value *float64) Rate {
	_init_.Initialize()

	if err := validateRate_WeeksParameters(value); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.StaticInvoke(
		"@shuttl-io/core.Rate",
		"weeks",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rate) Activate(args interface{}, invoker ITriggerInvoker) {
	if err := r.validateActivateParameters(args, invoker); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"activate",
		[]interface{}{args, invoker},
	)
}

func (r *jsiiProxy_Rate) BindOutcome(outcome IOutcome) ITrigger {
	if err := r.validateBindOutcomeParameters(outcome); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		r,
		"bindOutcome",
		[]interface{}{outcome},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rate) ParseArgs(args interface{}) *TriggerOutput {
	if err := r.validateParseArgsParameters(args); err != nil {
		panic(err)
	}
	var returns *TriggerOutput

	_jsii_.Invoke(
		r,
		"parseArgs",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rate) Validate(args interface{}) *map[string]interface{} {
	if err := r.validateValidateParameters(args); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"validate",
		[]interface{}{args},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rate) WithName(name *string) ITrigger {
	if err := r.validateWithNameParameters(name); err != nil {
		panic(err)
	}
	var returns ITrigger

	_jsii_.Invoke(
		r,
		"withName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rate) WithOnTrigger(onTrigger IRateTriggerOnTrigger) Rate {
	if err := r.validateWithOnTriggerParameters(onTrigger); err != nil {
		panic(err)
	}
	var returns Rate

	_jsii_.Invoke(
		r,
		"withOnTrigger",
		[]interface{}{onTrigger},
		&returns,
	)

	return returns
}

