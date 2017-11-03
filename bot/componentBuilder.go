package bot

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