package bot

type BotState struct {
	CurrentComponent BotComponent
	Data map[string]interface{}
}