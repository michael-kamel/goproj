package bot

import "../parser"

type BotComponent struct {
	Reply []string
	Name string
	CustomFunction string
	Handler func(message interface{}, state *BotState) string
	Parser func(message string) parser.ParserResult
	Connector BotComponentConnector
}