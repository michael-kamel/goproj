package Parser

import (
	//"fmt"
	"strings"
	"../ScriptParserAndBuilder"
)

type ParserResult struct {
	Success bool
	NextState string
}

type KeywordsToNextState struct {
	Keywords []string //list of keywords
	nextState string //the next state
}

func MultiParseState(transitions []ScriptParserAndBuilder.Transition, message string) ParserResult {
	for i := 0; i < len(transitions); i++ {
		pr := Parse(transitions[i].Keywords, message)
		if(pr) {
			return ParserResult{Success:true, NextState:transitions[i].NextState}
		}
	}
	return ParserResult{Success:false}
}

func Parse(keywords []string, message string) bool {
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