package bot

import "../parser"
type BotQuestion func(state *BotState) string
type BotParser func(message string) parser.ParserResult
type BotHandler func(message interface{}, state *BotState) string

type BotComponent struct {
	Name string
	Question BotQuestion
	Parser BotParser
	Handler BotHandler
}