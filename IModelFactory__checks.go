//go:build !no_runtime_type_checking

package core

import (
	"fmt"
)

func (i *jsiiProxy_IModelFactory) validateCreateParameters(props IModelFactoryProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

