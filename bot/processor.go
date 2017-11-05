package bot

import (
	"math/rand"
	//"time"
	//"strings"
	"../scriptParserAndBuilder"
)

func Process(message string, state string) string {
	parseResult := state.CurrentComponent.Parser(message)
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