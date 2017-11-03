package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../validators"
	"../parser"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	jData, err := json.Marshal(
		map[string]string{"message":"Hi! E7na sherket el mor3ebeen el ma7dooda, would you like to buy or sell?", "uuid":"generateSthRandomHere"});
	if err != nil {
		panic(err)
	}
	w.Write(jData)
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	s := validators.ValidateChat(w, r)
	if !s.Success {
		w.Write([]byte(s.Message))
		return
	}
	//fmt.Println(s.KeyValues["message"]);
	//testKeywords := new []
	//parserResult := parser.Parse([]string{"buy", "sell"}, s.KeyValues["message"])
	//fmt.Println(parserResult)

	//testing a parser function
	returnedParser := parser.GenerateParser([]string{"buy", "sell"}, parser.Parse)
	parserResult := returnedParser(s.KeyValues["message"])
	fmt.Println(parserResult)
	
}