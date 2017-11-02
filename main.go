package main

import (
	"fmt"
	"net/http"
	//"encoding/json"
	"./handlers"
)

type jsonResponse struct {
	Message string
	UUID string
}

func main() {
	http.HandleFunc("/welcome", handlers.WelcomeHandler);
	http.HandleFunc("/chat", handlers.ChatHandler);
	http.ListenAndServe(":9000", nil);
	fmt.Println("listening on 9000");
}