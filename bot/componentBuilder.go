package bot

func BuildSimpleQuestion(message string) BotQuestion{
	return func(state *BotState) string{
		return message
	}
}
func BuildSimpleHandler(response string) BotHandler {
	return func(message interface{}, state *BotState) string {
		return response
	}
}
func BuildDataFillHandler(key string, response string) BotHandler {
	return func(message interface{}, state *BotState) string {
		state.Data[key] = message
		return response
	}
}