package StateCache

import "../bot"

type BotStateCache interface{
	GetState(id string) bot.BotState 
}