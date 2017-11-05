package bot


func BuildSingleTransitionConnector(name string, component *BotComponent) BotComponentConnector{
	return BotComponentConnector{
		name,
		func(state *BotState,  bot *Bot) *BotComponent { return component}}
}
func BuildStateDependantStateTransitionConnector(name string, key string, componentMap map[interface{}]*BotComponent) BotComponentConnector{
	return BotComponentConnector{
		name,
		func(state *BotState,  bot *Bot) *BotComponent{
			val := state.Data[key]
			if value, ok := componentMap[val]; ok {
				return value
			} else {
				panic(stateDependentTransitionError{key, state.CurrentComponent.Name})
			}
		}}
}