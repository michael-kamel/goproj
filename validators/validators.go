package validators

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type ValidationReturn struct {
	Success bool
	Message string
}

func ValidateChat(w http.ResponseWriter, r *http.Request) ValidationReturn {
	//check HTTP verb
	if(r.Method != "POST") {
		return ValidationReturn{Success:false, Message:"Invalid HTTP Verb"}
	}

	//decode body as json
	var JSONData map[string]string
	decoder := json.NewDecoder(r.Body);
	e := decoder.Decode(&JSONData);
	if e != nil {
		//panic(e)
		return ValidationReturn{Success:false, Message:e.Error()}
	}
	fmt.Println(JSONData);

	//check for the key "message"
	if _, exists := JSONData["message"]; !exists { //_ is the message
		return ValidationReturn{Success:false, Message:"JSON received doesn't have key \"Message\""}
	}

	// // a string slice to hold the keys
	// k := make([]string, len(c))

	// // iteration counter
	// i := 0

	// // copy c's keys into k
	// for s, _ := range c {
	// 	k[i] = s
	// 	i++
	// }

	return ValidationReturn{}
}