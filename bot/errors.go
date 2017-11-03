package bot

import "fmt"

type stateDependentTransitionError  struct {
    key  string
    name string
}
func (e *stateDependentTransitionError) Error() string {
    return fmt.Sprintf("Could not make a transition from component %s using key %s", e.name, e.key)
}