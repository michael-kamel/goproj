package scriptBuilder

import (
	"strings"
	"goproj/bot"
	"fmt"
	"goproj/repositories"
	"goproj/models"
)

type FunctionalityProvider interface {
	GetQuestion(name string) bot.BotQuestion
	GetHandler(name string) bot.BotHandler
	GetTransition(name string) bot.BotTransition
}

type DefaultFunctionalityProvider struct {
	Questions map[string]bot.BotQuestion
	Handlers map[string]bot.BotHandler
	Transitions map[string]bot.BotTransition
	BuyerRepository repositories.BuyerRequestRepository
	ListingRepository repositories.ListingRepository
}

func (this *DefaultFunctionalityProvider) Init() {
	this.Questions = map[string]bot.BotQuestion{}
	this.Handlers = map[string]bot.BotHandler{
		"BuySellHandler":func(message interface{}, state *bot.BotState) string { 
			state.Data["type"] = message
			return fmt.Sprintf("So you are a %ser", message)
		},
		"BuyerRequestGetter":func(message interface{}, state *bot.BotState) string {
			spec := models.ListingSpecification{
				state.Data["category"].(string),
				state.Data["location"].(string),
				state.Data["space"].(int),
				message.(int),
			}
			listings, err := this.ListingRepository.GetListings(spec)
			if err != nil {
				state.Data["llen"] = 0
				return "Service is not available"
			}
			state.Data["llen"] = len(listings)

			if len(listings) == 0 {
				return "There are no listings that match your criteria"
			} else {
				
				listingIds := make([]string, len(listings))
				for i, list := range listings {
					listingIds[i] = list.Id
				} 
				state.Data["iListings"] = listingIds
				listing := listings[0]
				return fmt.Sprintf("%s/%s/%d/%d/%s/%s/%d", listing.Category, listing.Location, listing.Space, listing.Price, listing.Description, listing.Address, len(listings)-1)
				//return fmt.Sprintf("We have a listing that match your criteria with the following info /n/t Category:%s  /n/t Location:%s /n/t Space:%d /n/t Price:%d /n/t Description:%s /n/t Address:%s /n and %d more", listing.Category, listing.Location, listing.Space, listing.Price, listing.Description, listing.Address, len(listings)-1)
			}

		},
		"BuyerRecord":func(message interface{}, state *bot.BotState) string { 
			ownerInfo := models.OwnerInfo{
				state.Data["name"].(string),
				state.Data["phone"].(string),
				message.(string),
			}
			request := models.BuyerRequest{
				ownerInfo,
				state.Data["iListings"].([]string),
			}
			err := this.BuyerRepository.AddBuyerRequest(request)
			if err != nil {
				return "Service is not available"
			} else {
				return "Your response has been recorded"
			}
		},
		"SellerDataRecord":func(message interface{}, state *bot.BotState) string { 
			ownerInfo := models.OwnerInfo{
				state.Data["name"].(string),
				state.Data["phone"].(string),
				state.Data["email"].(string),
			}
			listing := models.Listing{
				"",
				ownerInfo,
				state.Data["category"].(string),
				state.Data["location"].(string),
				state.Data["space"].(int),
				state.Data["price"].(int),
				state.Data["address"].(string),
				message.(string),
			}
			err := this.ListingRepository.AddListing(listing)
			if err != nil {
				return "Service is not available"
			} else {
				return "Your response has been recorded"
			}
		},
		"ByeHandler":func(message interface{}, state *bot.BotState) string { 
			msg := strings.ToLower(message.(string))
			state.Data["restart"] = msg
			if msg == "yes" {
				return ""
			} else {
				return "Thank you for using our services"
			}
		},
	}
	this.Transitions = map[string]bot.BotTransition{
		"BuySellConnection":func(state *bot.BotState, bot *bot.Bot) *bot.BotComponent {
			customerType := strings.ToLower(state.Data["type"].(string))
			if customerType == "buy" {
				return bot.Components["CategoryComponent"]
			} else {
				return bot.Components["NameComponent"]
			}
		},
		"PriceConnector":func(state *bot.BotState, bot *bot.Bot) *bot.BotComponent {
			customerType := strings.ToLower(state.Data["type"].(string))
			if customerType == "buy" {
				return bot.Components["PriceBuyerComponent"]
			} else {
				return bot.Components["PriceSellerComponent"]
			}
		},
		"EmailConnector":func(state *bot.BotState, bot *bot.Bot) *bot.BotComponent {
			customerType := strings.ToLower(state.Data["type"].(string))
			if customerType == "buy" {
				return bot.Components["EmailBuyerComponent"]
			} else {
				return bot.Components["EmailSellerComponent"]
			}
		},
		"FromPriceBuyerTransition":func(state *bot.BotState, bot *bot.Bot) *bot.BotComponent {
			listingsCount := state.Data["llen"].(int)
			if listingsCount == 0 {
				return bot.Components["ByeComponent"]
			} else {
				return bot.Components["InterestedComponent"]
			}
		},
		"FromInterestedTransition":func(state *bot.BotState, bot *bot.Bot) *bot.BotComponent {
			interested := strings.ToLower(state.Data["interested"].(string))
			if interested == "yes" {
				return bot.Components["NameComponent"]
			} else {
				return bot.Components["ByeComponent"]
			}
		},
		"RestartTransition":func(state *bot.BotState, bot *bot.Bot) *bot.BotComponent {
			restart := strings.ToLower(state.Data["restart"].(string))
			if restart == "yes" {
				state.Data = make(map[string]interface{})
				return bot.Components["BuySellComponent"]
			} else {
				return bot.Components["ByeComponent"]
			}
		},
	}
	
}
func (this *DefaultFunctionalityProvider) GetQuestion(name string) bot.BotQuestion {
	return this.Questions[name]
}
func (this *DefaultFunctionalityProvider) GetHandler(name string) bot.BotHandler {
	return this.Handlers[name]
}
func (this *DefaultFunctionalityProvider) GetTransition(name string) bot.BotTransition {
	return this.Transitions[name]
}