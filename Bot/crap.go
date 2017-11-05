package Bot

import (
	"../Parser"
	"math/rand"
	"fmt"
	//"time"
	//"strings"
	"../ScriptParserAndBuilder"
	"../SessionManagement"
)

//structs
type Bot struct {
	UnhandledMessages []string
}
type BotComponent struct {
	Reply []string
	Name string
	CustomFunction string
	Handler func(message interface{}, state *BotState) string
	Parser func(message string) Parser.ParserResult
	Connector BotComponentConnector
}
type BotComponentConnector struct { //useless
	Transition func(state *BotState) BotComponent
}
type BotState struct {
	CurrentComponent BotComponent
	Data map[string]interface{}
}
type stateDependentTransitionError  struct {
    key  string
    name string
}



//processor
func Process(session SessionManagement.UserSession, message string) string {
	parserResult := Parser.MultiParseState(ScriptParserAndBuilder.ConstructedBot[session.State].Transitions ,message)
	return parserResult.NextState
	/*
	if parseResult.Success {
		state.CurrentComponent = state.CurrentComponent.Connector.Transition(state)
		if state.CurrentComponent.CustomFunction == "null" {
			return composeResponse(state.CurrentComponent.Reply)
		} else {
			return "TODO" //should execute custom handler //state.CurrentComponent.Handler(parseResult.NextState, state)
		}
	} else {
		message := bot.UnhandledMessages[rand.Intn(len(bot.UnhandledMessages))]
		return message
	}*/
}



//builders
func BuildSingleTransitionConnector(component *BotComponent) BotComponentConnector{
	return BotComponentConnector{func(state *BotState) BotComponent { 
		return *component
	}}
}
func BuildStateDependantStateTransitionConnector(key string, componentMap map[interface{}]*BotComponent) BotComponentConnector{
	return BotComponentConnector{func(state *BotState) BotComponent{
		val := state.Data[key]
		if value, ok := componentMap[val]; ok {
			return *value
		} else {
			panic(stateDependentTransitionError{key, state.CurrentComponent.Name})
		}
	}}
}

func BuildSimpleQuestion(message string) func(*BotState)string{
	return func(state *BotState) string{
		return message
	}
}
func BuildSimpleHandler(response string) func(interface{}, *BotState) string {
	return func(message interface{}, state *BotState) string {
		return response
	}
}


func composeResponse(possibleAnswers []string) string {
	/*if question != "" {
		return strings.Join([]string{question, answer}, ". ")
	} else {
		return answer
	}*/
	return possibleAnswers[rand.Intn(len(possibleAnswers))]
}



func (e *stateDependentTransitionError) Error() string {
    return fmt.Sprintf("Could not make a transition from component %s using key %s", e.name, e.key)
}