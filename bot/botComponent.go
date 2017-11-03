package bot

import "../parser"

type BotComponent struct {
	Question func(state *BotState) string
	Name string
	Parser func(message string) parser.ParserResult
	Handler func(message string, state *BotState) string
	Connector BotComponentConnector
}