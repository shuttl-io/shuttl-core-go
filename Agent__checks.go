//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_Agent) validateGetToolParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Agent) validateGetToolCallResultParameters(callID *string, result interface{}) error {
	if callID == nil {
		return fmt.Errorf("parameter callID is required, but nil was provided")
	}

	if result == nil {
		return fmt.Errorf("parameter result is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Agent) validateInvokeParameters(prompt interface{}, attachments *[]*FileAttachment) error {
	if prompt == nil {
		return fmt.Errorf("parameter prompt is required, but nil was provided")
	}
	switch prompt.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]interface{}:
		prompt := prompt.(*[]interface{})
		for idx_cf0719, v := range *prompt {
			switch v.(type) {
			case *ModelContent:
				v := v.(*ModelContent)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter prompt[%#v]", idx_cf0719) }); err != nil {
					return err
				}
			case ModelContent:
				v_ := v.(ModelContent)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter prompt[%#v]", idx_cf0719) }); err != nil {
					return err
				}
			case *map[string]interface{}:
				// ok
			case map[string]interface{}:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter prompt[%#v] must be one of the allowed types: *ModelContent, *map[string]interface{}; received %#v (a %T)", idx_cf0719, v, v)
				}
			}
		}
	case []interface{}:
		prompt_ := prompt.([]interface{})
		prompt := &prompt_
		for idx_cf0719, v := range *prompt {
			switch v.(type) {
			case *ModelContent:
				v := v.(*ModelContent)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter prompt[%#v]", idx_cf0719) }); err != nil {
					return err
				}
			case ModelContent:
				v_ := v.(ModelContent)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter prompt[%#v]", idx_cf0719) }); err != nil {
					return err
				}
			case *map[string]interface{}:
				// ok
			case map[string]interface{}:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter prompt[%#v] must be one of the allowed types: *ModelContent, *map[string]interface{}; received %#v (a %T)", idx_cf0719, v, v)
				}
			}
		}
	default:
		return fmt.Errorf("parameter prompt must be one of the allowed types: *string, *[]interface{}; received %#v (a %T)", prompt, prompt)
	}

	for idx_3930e6, v := range *attachments {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter attachments[%#v]", idx_3930e6) }); err != nil {
			return err
		}
	}

	return nil
}

func (a *jsiiProxy_Agent) validateRespondWithToolCallParameters(modelInstance IModel, callID *string, result interface{}, streamer IModelStreamer) error {
	if modelInstance == nil {
		return fmt.Errorf("parameter modelInstance is required, but nil was provided")
	}

	if callID == nil {
		return fmt.Errorf("parameter callID is required, but nil was provided")
	}

	if result == nil {
		return fmt.Errorf("parameter result is required, but nil was provided")
	}

	if streamer == nil {
		return fmt.Errorf("parameter streamer is required, but nil was provided")
	}

	return nil
}

func validateNewAgentParameters(props *AgentProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

