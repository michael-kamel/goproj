package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../validators"
	"../SessionManagement"
	//"math/rand"
	"time"
	//"strconv"
	//"../parser"
	"../scriptParserAndBuilder"
	"../bot"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	generatedUUID := time.Now().UnixNano() / int64(time.Millisecond)
	SessionManagement.GenerateNewUserSession(fmt.Sprintf("%v", generatedUUID))

	//response := bot.Process("", )

	//respond with welcome message
	w.Header().Set("Content-Type", "application/json");
	jData, _ := json.Marshal(
		map[string]string{"message":"Hi! E7na sherket el mor3ebeen el ma7dooda, would you like to buy or sell?", "uuid":"generateSthRandomHere"});
	w.Write(jData)
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	s := validators.ValidateChat(w, r)
	if !s.Success {
		w.Write([]byte(s.Message))
		return
	}

	userSession = SessionManagement.GetUserSession(s.KeyValues["UUID"])

	//fmt.Println(s.KeyValues["message"]);
	//testKeywords := new []
	//parserResult := parser.Parse([]string{"buy", "sell"}, s.KeyValues["message"])
	//fmt.Println(parserResult)

	//testing a parser function
	//returnedParser := parser.GenerateParser([]string{"buy", "sell"}, parser.Parse)
	//parserResult := returnedParser(s.KeyValues["message"])
	//fmt.Println(parserResult)
	
}