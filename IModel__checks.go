//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IModel) validateInvokeParameters(prompt *[]interface{}, streamer IModelStreamer) error {
	if prompt == nil {
		return fmt.Errorf("parameter prompt is required, but nil was provided")
	}
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

	if streamer == nil {
		return fmt.Errorf("parameter streamer is required, but nil was provided")
	}

	return nil
}

