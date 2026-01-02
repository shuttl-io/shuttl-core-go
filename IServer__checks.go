//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (i *jsiiProxy_IServer) validateAcceptParameters(app interface{}) error {
	if app == nil {
		return fmt.Errorf("parameter app is required, but nil was provided")
	}

	return nil
}

