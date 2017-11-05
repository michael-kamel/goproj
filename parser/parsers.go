package parser

import (
	//"fmt"
	"strings"
)

type KeywordsToNextState struct {
	Keywords []string //list of keywords
	nextState string //the next state
}

func MultiParseState(IO []KeywordsToNextState, message string) ParserResult {
	for i := 0; i < len(IO); i++ {
		pr := Parse(IO[i].Keywords, message)
		if(pr) {
			return ParserResult{Success:true, NextState:IO[i].nextState}
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
