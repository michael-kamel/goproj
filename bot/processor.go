package bot

import (
	"math/rand"
	//"time"
	"strings"
)

func Process(message string, state *BotState, bot Bot) string {
	parseResult := state.CurrentComponent.Parser(message)
	if parseResult.Success {
		answer := state.CurrentComponent.Handler(parseResult.Message, state)
		state.CurrentComponent = state.CurrentComponent.Connector.Transition(state)
		return composeResponse(answer, state.CurrentComponent.Question(state))
	} else {
		message := bot.UnhandledMessages[rand.Intn(len(bot.UnhandledMessages))]
		return message
	}
}
func composeResponse(question string, answer string) string {
	if question != "" {
		return strings.Join([]string{question, answer}, ". ")
	} else {
		return answer
	}
}