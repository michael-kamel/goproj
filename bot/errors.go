package bot

import "fmt"

type stateDependentTransitionError  struct {
    key  string
    name string
}
func (e *stateDependentTransitionError) Error() string {
    return fmt.Sprintf("Could not make a transition from component %s using key %s", e.name, e.key)
}

type unRegisteredComponentError struct {
    name string
}
func (e *unRegisteredComponentError) Error() string {
    return fmt.Sprintf("Component %s is not registered", e.name)
}

type unRegisteredConnectorError struct {
    name string
}
func (e *unRegisteredConnectorError) Error() string {
    return fmt.Sprintf("Connector %s is not registered", e.name)
}