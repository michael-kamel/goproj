package Processor

import (
	"../Parser"
	"math/rand"
	"fmt"
	//"time"
	//"strings"
	"../ScriptParserAndBuilder"
	"../SessionManagement"
	"../ResponseHandlers"
	"encoding/json"
)

//structs
type Bot struct {
	UnhandledMessages []string
}
type BotComponent struct {
	Reply []string
	Name string
	CustomResponse string
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
func Process(uuid string, message string) []byte {
	//fmt.Println(uuid)
	session := SessionManagement.UserSessions[uuid]
	if(session == nil) {
		return messageResponse("Session Error")
	}
	//fmt.Println(session)
	parserResult := Parser.DetermineTransition(ScriptParserAndBuilder.ConstructedBot[session.State].Transitions ,message)
	if !parserResult.Success {
		//fmt.Println(len(session.RejectMessages))
		return messageResponse(session.RejectMessages[rand.Intn(len(session.RejectMessages))])
	}


	var transition ScriptParserAndBuilder.Transition
	for i := 0; i < len(ScriptParserAndBuilder.ConstructedBot[session.State].Transitions); i++ { //can be improved, need to have maps of transitions inside the states
		if(ScriptParserAndBuilder.ConstructedBot[session.State].Transitions[i].NextState == parserResult.NextState) {
			transition = ScriptParserAndBuilder.ConstructedBot[session.State].Transitions[i]
		}
	}
	//direction-specific parsing and validation
	if(transition.CustomParser != "null") {
		if r := Parser.DSPV[transition.CustomParser](transition, message, session); r != "okay" {
			return messageResponse(r)
		}
	}

	session.State = parserResult.NextState;
	session.RejectMessages = transition.Rejects;
	fmt.Println(*session)

	//response handlers
	if transition.CustomResponse != "null" {
		return messageResponse(transition.Replies[rand.Intn(len(transition.Replies))] + " " + ResponseHandlers.ResponseHandlers[transition.CustomResponse](transition, message, session))
	}
	return messageResponse(transition.Replies[rand.Intn(len(transition.Replies))])




	//return parserResult.NextState
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
func messageResponse(s string) []byte {
	j, _ := json.Marshal(map[string]string{"message":s});
	return j
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