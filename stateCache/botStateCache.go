package stateCache

import "../bot"

type BotStateCache interface{
	GetState(id string) bot.BotState 
}