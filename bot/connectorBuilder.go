package bot


func BuildSingleTransitionConnector(component *BotComponent) BotComponentConnector{
	return BotComponentConnector{func(state *BotState) BotComponent { 
		return *component
	}}
}
func BuildStateDependantStateTransitionConnector(key string, componentMap map[interface{}]*BotComponent) BotComponentConnector{
	return BotComponentConnector{func(state *BotState) BotComponent{
		val := state.Data[key]
		if value, ok := componentMap[val]; ok {
			return *value
		} else {
			panic(stateDependentTransitionError{key, state.CurrentComponent.Name})
		}
	}}
}