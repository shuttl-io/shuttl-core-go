//go:build !no_runtime_type_checking

package core

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IApiAuthenticator) validateAuthenticateParameters(args *ApiTriggerArgs) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(args, func() string { return "parameter args" }); err != nil {
		return err
	}

	return nil
}

