package Parser

import (
	// "fmt"
	"strings"
	"strconv"
	"../ScriptParserAndBuilder"
	"../SessionManagement"
	valid "github.com/asaskevich/govalidator"
)

type ParserResult struct {
	Success bool
	NextState string
}

type KeywordsToNextState struct {
	Keywords []string //list of keywords
	nextState string //the next state
}


//direction-specific parsing and validation (& handling)
//returns "okay" if input is valid, otherwise the error
var DSPV map[string]func(ScriptParserAndBuilder.Transition, string, *SessionManagement.UserSession) string = map[string]func(ScriptParserAndBuilder.Transition, string, *SessionManagement.UserSession) string {
	"phase2buy":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		price, err := strconv.Atoi(message);
		if err != nil {
			return "Please enter a valid price"
		}

		session.Data.Price = price
		return "okay"
	},
	"category":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		categories := []string{"Residential", "Adminstrative", "Commercial", "Industrial", "Other"};

		for i := 0; i < len(categories); i++ {
			if(strings.Contains(strings.ToLower(message), strings.ToLower(categories[i]))) {
				session.Data.Category = categories[i];
				return "okay"
			}
		}

		return "Invalid category. Try again! [Residential, Adminstrative, Commercial, Industrial, Other]."
	},
	"location":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		locations := []string{"6th of October", "Abbassiya", "Agouza", "Al Rehab", "Dokki", "El Sadat City", "El Salam City", "El Sayeda Zeinab", "El Shorouk City", "El Tagammoa El Khames", "Faisal", "Gesr El Suez", "Giza", "El Haram", "Heliopolis", "Helwan", "Imbaba", "Katameya", "Maadi", "Madinaty", "Manial", "Sheraton", "Mokattam", "Nasr City", "New Cairo", "Sheikh Zayed", "Shoubra", "Smart Village", "Zamalek"}
	
		for i := 0; i < len(locations); i++ {
			if(strings.Contains(strings.ToLower(message), strings.ToLower(locations[i]))) {
				session.Data.Location = locations[i];
				return "okay"
			}
		}

		return "Location not valid. Try again. [6th of October, Abbassiya, Agouza, Al Rehab, Dokki, El Sadat City, El Salam City, El Sayeda Zeinab, El Shorouk City, El Tagammoa El Khames, Faisal, Gesr El Suez, Giza, El Haram, Heliopolis, Helwan, Imbaba, Katameya, Maadi, Madinaty, Manial, Sheraton, Mokattam, Nasr City, New Cairo, Sheikh Zayed, Shoubra, Smart Village, Zamalek]"
	},
	"space":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		space, err := strconv.Atoi(message);
		if err != nil {
			return "Space invalid. Please enter a number."
		}

		if space > 2000 || space < 1000 {
			return "Space invalid. Just enter a number between 1000 and 2000."
		}

		session.Data.Space = space;

		return "okay"
	},
	"price":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		price, err := strconv.Atoi(message);
		if err != nil {
			return "Price invalid. Please enter a number."
		}

		if price < 1000000 || price > 500000000 {
			return "Price invalid. Just enter a number between 10 000 000 and 500 000 000."
		}

		session.Data.Price = price;

		return "okay"
	},
	"description":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		
		session.Data.Description = message;

		return "okay"
	},
	"name":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {

		session.Data.Name = message;

		return "okay"
	},
	"phone":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {

		if len(message) < 6 {
			return "Minimum 6 digits. Try again!"
		}

		if !valid.IsInt(message) {
			return "Well this is a phone number, therefore use numbers only!"
		}

		session.Data.Phone = message;
		return "okay"
		
	},
	"email":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		
		if valid.IsEmail(message) {
			session.Data.Email = message;
			return "okay"
		}

		return "Invalid email! Try again please"
	},
	"address":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		
		latLong := strings.Split(message, ",")


		if len(latLong) != 2 || !valid.IsFloat(latLong[0]) || !valid.IsFloat(latLong[1]) {
			return "Invalid! Try again! Example: -34.0,151.0"
		}

		session.Data.Address = message
		return "okay"
	},
	"chosenItem":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		
		itemNo, err := strconv.Atoi(message);
		if err != nil {
			return "Please input a number from the list. " + session.ReceivedItemsMessage
		}
		for i := 0; i < len(session.ReceivedItems); i++ {
			if session.ReceivedItems[i].Number == itemNo {
				session.ChosenItem = session.ReceivedItems[i].ID
				return "okay"
			}
		}

		return "Invalid Input Try Again. Choose a number. " + session.ReceivedItemsMessage
	},
}


func DetermineTransition(transitions []ScriptParserAndBuilder.Transition, message string) ParserResult {
	if(len(transitions) == 1) {
		return ParserResult{Success:true, NextState:transitions[0].NextState}
	}

	for i := 0; i < len(transitions); i++ {
		pr := parse(transitions[i].Keywords, message)
		if(pr) {
			return ParserResult{Success:true, NextState:transitions[i].NextState}
		}
	}
	return ParserResult{Success:false}
}
func parse(keywords []string, message string) bool {
	for i := 0; i < len(keywords); i++ {
		if strings.Contains(strings.ToLower(message), keywords[i]) {
			return true //ParserResult{Success:true, Message:keywords[i]}
		}
	}
	return false //ParserResult{Success:false}
}



//useless
func GenerateParser(keywords []string, parserFunction func([]string, string) ParserResult) func(string) ParserResult {
	return func(message string) ParserResult {
		return parserFunction(keywords, message);
	}
}