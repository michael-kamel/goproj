package main

import (
	"fmt"
	"net/http"
	//"encoding/json"
	"./Handlers"
	//"./bot"
	//"./parser"
	//"./ScriptParserAndBuilder"
	"math/rand"
	"time"
)

type jsonResponse struct {
	Message string
	UUID string
}

func main() {
	//fmt.Println(InputHandlers["test"](1,2))
	rand.Seed(time.Now().UTC().UnixNano());
	//fmt.Println(ScriptParserAndBuilder.ConstructedBot);

	http.HandleFunc("/welcome", Handlers.WelcomeHandler);
	http.HandleFunc("/chat", Handlers.ChatHandler);
	fmt.Println("will listen on 9000");
	http.ListenAndServe(":9000", nil);
}
/*
func testBot() {
	welcomeComponent := bot.BotComponent{}
	buySellComponent := bot.BotComponent{}
	byeComponent := bot.BotComponent{}
	welcomeConnector := bot.BuildSingleTransitionConnector(&buySellComponent)
	buySellConnector := bot.BuildSingleTransitionConnector(&byeComponent)
	byeConnector := bot.BuildSingleTransitionConnector(&byeComponent)

	welcomeComponent.Reply = "wewerw"
	welcomeComponent.Name = "WelcomeComponent"
	welcomeComponent.Parser = func(message string) parser.ParserResult { return parser.ParserResult{Success:true,Message:""}}
	welcomeComponent.Handler = bot.BuildSimpleHandler("Hello and welcome to the shit service")
	welcomeComponent.Connector = welcomeConnector

	buySellComponent.Question = bot.BuildSimpleQuestion("Are you a buyer or seller?")
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
	buySellComponent.Handler = func(message interface{}, state *bot.BotState) string { 
		state.Data["type"] = message
		return fmt.Sprintf("So you are a %s", message)
	}
	buySellComponent.Connector = buySellConnector

	byeComponent.Question = bot.BuildSimpleQuestion("Your response has been recorded")
	byeComponent.Name = "ByeComponent"
	byeComponent.Parser = func(message string) parser.ParserResult { return parser.ParserResult{Success:true,Message:""}}
	byeComponent.Handler = bot.BuildSimpleHandler("")
	byeComponent.Connector = byeConnector


	botDesc := bot.Bot{[]string{"I don't understand", "No Comprendo", "what the hell"}}
	
	
	botState := bot.BotState{
		welcomeComponent,
		map[string]interface{}{}}
	fmt.Println(bot.Process("test", &botState, botDesc))
	fmt.Println(bot.Process("ana gy ahazar", &botState, botDesc))
	fmt.Println(bot.Process("ana gy ahazar", &botState, botDesc))
	fmt.Println(bot.Process("ana gy ahazar", &botState, botDesc))
	fmt.Println(bot.Process("buy", &botState, botDesc))
	fmt.Println(bot.Process(".", &botState, botDesc))
	fmt.Println(bot.Process("bye", &botState, botDesc))


	botState2 := bot.BotState{
		welcomeComponent,
		map[string]interface{}{}}
	fmt.Println(bot.Process("test", &botState2, botDesc))
	fmt.Println(bot.Process("bye", &botState2, botDesc))

	botState3 := bot.BotState{
		welcomeComponent,
		map[string]interface{}{}}
	fmt.Println(bot.Process("test", &botState3, botDesc))
	fmt.Println(bot.Process("buy", &botState3, botDesc))
	fmt.Println(bot.Process("bye", &botState3, botDesc))
}*/