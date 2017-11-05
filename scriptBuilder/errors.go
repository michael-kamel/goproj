package scriptBuilder

import "fmt"

type unRecognizedQuestionType  struct {
    name string
}
func (e *unRecognizedQuestionType) Error() string {
    return fmt.Sprintf("Question type %s is not recongnized", e.name)
}

type unRecongnizedParserType struct {
    name string
}
func (e *unRecongnizedParserType) Error() string {
    return fmt.Sprintf("Parser type %s is not recongnized", e.name)
}

type unRecongnizedHandlerType struct {
    name string
}
func (e *unRecongnizedHandlerType) Error() string {
    return fmt.Sprintf("Handler type %s is not recongnized", e.name)
}

type unRecongnizedTransitionType struct {
    name string
}
func (e *unRecongnizedTransitionType) Error() string {
    return fmt.Sprintf("Transition type %s is not recongnized", e.name)
}

