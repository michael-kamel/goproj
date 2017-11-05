package repositories

import "fmt"

type unFulfilledRequest  struct {
    name string
    reason string
}
func (e *unFulfilledRequest) Error() string {
    return fmt.Sprintf("Request %s was not fulfilled reason: %s", e.name, e.reason)
}