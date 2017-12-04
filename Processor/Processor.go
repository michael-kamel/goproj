package Processor

import (
	// "fmt"
	"../Parser"
	"math/rand"
	// "fmt"
	//"time"
	//"strings"
	"../ScriptParserAndBuilder"
	"../SessionManagement"
	"../ResponseHandlers"
	"encoding/json"
	//"github.com/davecgh/go-spew/spew"
)


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
	//fmt.Println(*session)


	//response handlers
	if transition.CustomResponse != "null" {
		s := messageResponse(transition.Replies[rand.Intn(len(transition.Replies))] + " " + ResponseHandlers.ResponseHandlers[transition.CustomResponse](transition, message, session));
		//spew.Dump(session)
		return s
	}
	//spew.Dump(session)


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


