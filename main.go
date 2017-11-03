package main

import (
	"fmt"
	"net/http"
	//"encoding/json"
	"./handlers"
	"./bot"
	"./parser"
)

type jsonResponse struct {
	Message string
	UUID string
}

func main() {
	testBot()
	http.HandleFunc("/welcome", handlers.WelcomeHandler);
	http.HandleFunc("/chat", handlers.ChatHandler);
	http.ListenAndServe(":9000", nil);
	fmt.Println("listening on 9000");
}
func testBot() {
	welcomeComponent := bot.BotComponent{}
	buySellComponent := bot.BotComponent{}
	byeComponent := bot.BotComponent{}
	welcomeConnector := bot.BuildSingleTransitionConnector(&buySellComponent)
	buySellConnector := bot.BuildSingleTransitionConnector(&byeComponent)
	byeConnector := bot.BuildSingleTransitionConnector(&byeComponent)

	welcomeComponent.Question = func(state *bot.BotState) string { return ""}
	welcomeComponent.Name = "WelcomeComponent"
	welcomeComponent.Parser = func(message string) parser.ParserResult { return parser.ParserResult{Success:true,Message:""}}
	welcomeComponent.Handler = func(message string, state *bot.BotState) string { return "Hello and welcome to the shit service"}
	welcomeComponent.Connector = welcomeConnector

	buySellComponent.Question = func(state *bot.BotState) string { return "Are you a buyer or seller"}
	buySellComponent.Name = "BuySellComponent"
	buySellComponent.Parser = func(message string) parser.ParserResult {
		if message == "buy" { 
			return parser.ParserResult{Success:true,Message:"buyer"}
		} else  if message == "sell"{
			return parser.ParserResult{Success:true,Message:"seller"}
		} else {
			return parser.ParserResult{Success:false,Message:""}
		}
	}
	buySellComponent.Handler = func(message string, state *bot.BotState) string { 
		state.Data["type"] = message
		return fmt.Sprintf("So you are a %s", message)
	}

	buySellComponent.Connector = buySellConnector

	byeComponent.Question = func(state *bot.BotState) string { return "Your response has been recorded"}
	byeComponent.Name = "ByeComponent"
	byeComponent.Parser = func(message string) parser.ParserResult { return parser.ParserResult{Success:true,Message:""}}
	byeComponent.Handler = func(message string, state *bot.BotState) string { return ""}
	byeComponent.Connector = byeConnector


	botDesc := bot.Bot{[]string{"I don't understand", "No Comprendo", "what the hell"}}
	botState := bot.BotState{
		welcomeComponent,
		map[string]interface{}{}}
	fmt.Println(bot.Process("test", &botState, botDesc))
	fmt.Println(bot.Process("ana gy ahazar", &botState, botDesc))
	fmt.Println(bot.Process("buy", &botState, botDesc))
	fmt.Println(bot.Process(".", &botState, botDesc))
	fmt.Println(bot.Process("..", &botState, botDesc))
}