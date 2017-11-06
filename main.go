package main

import (
	"os"
	"fmt"
	"goproj/handlers"
	"goproj/bot"
	"goproj/parser"
	"goproj/repositories"
	"goproj/api"
	"goproj/models"
	"goproj/scriptBuilder"
	"io/ioutil"
	"goproj/stateCache"
	"goproj/uidGenerator"
	"github.com/gin-gonic/gin"
)
type jsonResponse struct {
	Message string
	UUID string
}
var (
	bots map[string]*bot.Bot
	botStateCache stateCache.BotStateCache
)
func main() {
	botStateCache := &stateCache.InMemoryBotStateCache{}
	botStateCache.Init()
	bots := loadBots([]string{"./descriptions/bots.json"})
	idGenerator := uidGenerator.DefaultUIDGenerator{}
	idGenerator.Init()
	handlerFuncs := handlers.Handlers(bots, botStateCache, &idGenerator)
	/*sstate := &bot.BotState{
		bots["Bolt"].RootComponent,
		make(map[string]interface{}),
	}
	fmt.Println(bots["Bolt"].Process("a", sstate))
	fmt.Println(bots["Bolt"].Process("buyer", sstate))
	fmt.Println(bots["Bolt"].Process("Industrial", sstate))
	fmt.Println(bots["Bolt"].Process("Nasr City", sstate))
	fmt.Println(bots["Bolt"].Process("1500", sstate))
	fmt.Println(bots["Bolt"].Process("47000000", sstate))
	fmt.Println(bots["Bolt"].Process("Yes", sstate))
	fmt.Println(bots["Bolt"].Process("Mike", sstate))
	fmt.Println(bots["Bolt"].Process("12412142", sstate))
	fmt.Println(bots["Bolt"].Process("mike@mike.com", sstate))*/
	router := gin.Default()
	router.GET("/welcome", handlerFuncs["Welcome"])
	router.POST("/chat", handlerFuncs["Chat"])
	router.OPTIONS("/chat", handlerFuncs["BS"])
	router.Run(os.Getenv("PORT"))
}
func testBot() {
	welcomeComponent := bot.BotComponent {
		"WelcomeComponent",
		bot.BuildSimpleQuestion(""),
		func(message string) parser.ParserResult{return parser.ParserResult{Success:true,Message:""}},
		bot.BuildSimpleHandler("Hello and welcome to the shit service"),
	}
	buySellComponent := bot.BotComponent {
		"BuySellComponent",
		bot.BuildSimpleQuestion("Are you a buyer or seller?"),
		func(message string) parser.ParserResult {
			if message == "buy" { 
				return parser.ParserResult{Success:true,Message:"buyer"}
			} else  if message == "sell"{
				return parser.ParserResult{Success:true,Message:"seller"}
			} else {
				return parser.ParserResult{Success:false,Message:""}
			}
		},
		func(message interface{}, state *bot.BotState) string { 
			state.Data["type"] = message
			return fmt.Sprintf("So you are a %s", message)
		},
	}
	byeComponent := bot.BotComponent{
		"ByeComponent",
		bot.BuildSimpleQuestion("Your response has been recorded"),
		func(message string) parser.ParserResult { return parser.ParserResult{Success:true,Message:""}},
		bot.BuildSimpleHandler(""),
	}
	welcomeConnector := bot.BuildSingleTransitionConnector("WelcomeConnector",&buySellComponent)
	buySellConnector := bot.BuildSingleTransitionConnector("BuySellConnector", &byeComponent)
	byeConnector := bot.BuildSingleTransitionConnector("ByeConnector", &byeComponent)

	botDesc := bot.Bot{
		"BoltBot",
		map[string]*bot.BotComponent{
			welcomeComponent.Name:&welcomeComponent,
			buySellComponent.Name:&buySellComponent,
			byeComponent.Name:&byeComponent,
		},
		map[string]*bot.BotComponentConnector{
			welcomeConnector.Name:&welcomeConnector,
			buySellConnector.Name:&buySellConnector,
			byeConnector.Name:&byeConnector,
		},
		make(map[string]string),
		[]string{"I don't understand", "No Comprendo", "what the hell"},
		&welcomeComponent,
	}
	
	botDesc.Connect(&welcomeComponent, &welcomeConnector)
	botDesc.Connect(&buySellComponent, &buySellConnector)
	botDesc.Connect(&byeComponent, &byeConnector)
	fmt.Println(botDesc.Transitions)
	botState := bot.BotState{
		&welcomeComponent,
		make(map[string]interface{}),
	}
	fmt.Println(botDesc.Process("hi", &botState))
	fmt.Println(botDesc.Process("ana gy ahazar", &botState))
	fmt.Println(botDesc.Process("ana gy ahazar", &botState))
	fmt.Println(botDesc.Process("buy", &botState))
	fmt.Println(botDesc.Process(".", &botState))
	fmt.Println(botDesc.Process("..", &botState))

	botState2 := bot.BotState{
		&welcomeComponent,
		make(map[string]interface{}),
	}
	fmt.Println("wat")
	fmt.Println(botDesc.Process("hi", &botState2))
	fmt.Println(botDesc.Process("ana gy ahazar", &botState2))
	fmt.Println(botDesc.Process("ana gy ahazar", &botState2))
	fmt.Println(botDesc.Process("ana gy ahazar", &botState2))
	fmt.Println(botDesc.Process("buy", &botState2))
	fmt.Println(botDesc.Process(".", &botState2))
	fmt.Println(botDesc.Process("..", &botState2))
}
func testBuyerRequestServices() {
	caller := api.MinJsonApiService {Headers:map[string]string{}}
	caller.Init()
	buyerRequestRepo := repositories.BuyerRequestAPIService {
		&caller,
		"http://127.0.0.1:8080/addbuyerrequest",
	}
	buyerRequestRepo.AddBuyerRequest(models.BuyerRequest{
		models.OwnerInfo{
			Name:"testbot",
			Phone:"0122333",
			Email:"bot@bot.com",
		},
		[]string{"59fa611cf1173f24a8b2cd3e", "59fcfe2032373e1995b908fa"},
	})
}
func testBuyerEnquireServices() {
	caller := api.MinJsonApiService {Headers:map[string]string{}}
	caller.Init()
	listingRepo := repositories.ListingAPIService {
		&caller,
		"http://127.0.0.1:8080/getlistingbyspec",
		"http://127.0.0.1:8080/addlisting",
	}
	listings, _ := listingRepo.GetListings(models.ListingSpecification{
		"Industrial",
		"Nasr City",
		1900,
		44000000,
	})
	fmt.Println(listings)

	listing := models.Listing{
		"",
		models.OwnerInfo{
			Name:"testbot",
			Phone:"0122333",
			Email:"bot@bot.com",
		},
		"Industrial",
		"Nasr City",
		1860,
		50000000,
		"Bot place",
	}
	listingRepo.AddListing(listing)
}
func testScript(){
	fProvider := scriptBuilder.DefaultFunctionalityProvider{}
	fProvider.Init()
	sBuilder := scriptBuilder.ScriptBuilder{}
	sBuilder.Init()
	dat, err := ioutil.ReadFile("goproj/descriptions/bots.json")
	if err != nil {
		panic(err)
	}
	bots := scriptBuilder.JsonScriptToBot(&fProvider, &sBuilder, dat)
	var bolt = bots["Noob"]
	botState := bot.BotState{
		bolt.RootComponent,
		make(map[string]interface{}),
	}
	fmt.Println(bolt.Process("hi", &botState))
	fmt.Println(bolt.Process("ana gy ahazar", &botState))
	fmt.Println(bolt.Process("ana gy ahazar", &botState))
	fmt.Println(bolt.Process("ana gy ahazar", &botState))
	fmt.Println(bolt.Process("buy", &botState))
	fmt.Println(bolt.Process(".", &botState))
	fmt.Println(bolt.Process("..", &botState))
	
}
func loadBots(directories []string) map[string]*bot.Bot {
	apiService := api.MinJsonApiService{}
	apiService.Init()
	listingRepository := repositories.ListingAPIService{
		&apiService,
		os.Getenv("LISTING_GET"),
		os.Getenv("LISTING_POST"),
	}
	buyerRequestRepository := repositories.BuyerRequestAPIService{
		&apiService,
		os.Getenv("BUYER_REQUEST_POST"),
	}
	fProvider := scriptBuilder.DefaultFunctionalityProvider{
		BuyerRepository:&buyerRequestRepository,
		ListingRepository:&listingRepository,
	}
	fProvider.Init()
	sBuilder := scriptBuilder.ScriptBuilder{}
	sBuilder.Init()
	gBots := make(map[string]*bot.Bot)
	for _, dir := range directories {
		dat, err := ioutil.ReadFile(dir)
		if err != nil {
			panic(err)
		}
		bots := scriptBuilder.JsonScriptToBot(&fProvider, &sBuilder, dat)
		for key, value := range bots {
			gBots[key] = value
		}
	}
	return gBots
}
