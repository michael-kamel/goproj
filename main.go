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
