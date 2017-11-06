package stateCache

import "goproj/bot"

type InMemoryBotStateCache struct {
	cache map[string] *bot.BotState
}
func (this *InMemoryBotStateCache) GetState(id string) *bot.BotState {
	return this.cache[id]
}
func (this *InMemoryBotStateCache) SetState(id string, state *bot.BotState) {
	this.cache[id] = state
}
func (this *InMemoryBotStateCache) DeleteState(id string) {
	delete(this.cache, id)
}
func (this *InMemoryBotStateCache) Init() {
	this.cache = make(map[string] *bot.BotState)
}