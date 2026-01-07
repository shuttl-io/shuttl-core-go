//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NamedPipeServer) validateAcceptParameters(app interface{}) error {
	return nil
}

