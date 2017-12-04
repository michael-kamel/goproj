package Handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../InputHandlers"
	"../SessionManagement"
	//"math/rand"
	"time"
	//"strconv"
	//"../parser"
	//"../scriptParserAndBuilder"
	"../Processor"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if(r.Method != "GET") {
		w.Write([]byte("Invalid HTTP Verb!"))
		return
	}
	generatedUUID := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(generatedUUID);
	SessionManagement.GenerateNewUserSession(fmt.Sprintf("%v", generatedUUID))

	//response := bot.Process("", )

	//respond with welcome message
	w.Header().Set("Content-Type", "application/json");
	w.Header().Set("Access-Control-Allow-Origin", "*");

	jData, _ := json.Marshal(map[string]string{"message":"Hi! E7na sherket el mor3ebeen el ma7dooda, would you like to buy or sell?", "uuid":fmt.Sprintf("%v", generatedUUID)});
	w.Write(jData)
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Headers", "authorization, Content-Type")
	
	s := InputHandlers.ValidateChat(w, r)
	if !s.Success {
		w.Write([]byte(s.Message))
		return
	}


	//userSession := *SessionManagement.UserSessions[s.KeyValues["UUID"]]
	//fmt.Println(userSession);



	w.Write(Processor.Process(s.KeyValues["authorization"], s.KeyValues["message"]))




	//fmt.Println(s.KeyValues["message"]);
	//testKeywords := new []
	//parserResult := parser.Parse([]string{"buy", "sell"}, s.KeyValues["message"])
	//fmt.Println(parserResult)

	//testing a parser function
	//returnedParser := parser.GenerateParser([]string{"buy", "sell"}, parser.Parse)
	//parserResult := returnedParser(s.KeyValues["message"])
	//fmt.Println(parserResult)
	
}