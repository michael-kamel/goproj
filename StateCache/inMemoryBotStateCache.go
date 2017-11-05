package StateCache

import "../bot"

type InMemoryBotStateCache struct {
	cache map[string] *bot.BotState
}
func (this *InMemoryBotStateCache) GetState(id string) *bot.BotState {
	return this.cache[id]
}

