package stateCache

import "goproj/bot"

type BotStateCache interface{
	GetState(id string) *bot.BotState
	SetState(id string, state *bot.BotState)
	DeleteState(id string)
}