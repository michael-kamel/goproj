package Parser

import (
	//"fmt"
	"strings"
	"strconv"
	"../ScriptParserAndBuilder"
	"../SessionManagement"
)

type ParserResult struct {
	Success bool
	NextState string
}

type KeywordsToNextState struct {
	Keywords []string //list of keywords
	nextState string //the next state
}


//direction-specific parsing and validation
//returns "okay" if input is valid, otherwise the error
var DSPV map[string]func(ScriptParserAndBuilder.Transition, string, *SessionManagement.UserSession) string = map[string]func(ScriptParserAndBuilder.Transition, string, *SessionManagement.UserSession) string {
	"phase2buy":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		price, err := strconv.Atoi(message);
		if err != nil {
			return "Please enter a valid price"
		}

		session.Data.ItemPrice = price
		return "okay"
	},
}


func DetermineTransition(transitions []ScriptParserAndBuilder.Transition, message string) ParserResult {
	if(len(transitions) == 1) {
		return ParserResult{Success:true, NextState:transitions[0].NextState}
	}

	for i := 0; i < len(transitions); i++ {
		pr := parse(transitions[i].Keywords, message)
		if(pr) {
			return ParserResult{Success:true, NextState:transitions[i].NextState}
		}
	}
	return ParserResult{Success:false}
}
func parse(keywords []string, message string) bool {
	for i := 0; i < len(keywords); i++ {
		if strings.Contains(strings.ToLower(message), keywords[i]) {
			return true //ParserResult{Success:true, Message:keywords[i]}
		}
	}
	return false //ParserResult{Success:false}
}



//useless
func GenerateParser(keywords []string, parserFunction func([]string, string) ParserResult) func(string) ParserResult {
	return func(message string) ParserResult {
		return parserFunction(keywords, message);
	}
}