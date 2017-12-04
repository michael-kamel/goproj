package InputHandlers

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type ValidationReturn struct {
	Success bool
	Message string
	KeyValues map[string]string
}


var InputHandlers map[string]func(http.ResponseWriter, *http.Request) ValidationReturn = map[string]func(http.ResponseWriter, *http.Request) ValidationReturn {
	"default":func(w http.ResponseWriter, r *http.Request) ValidationReturn {
		if(r.Method != "POST") {
			//w.Write([]byte("Invalid HTTP Verb!"))
			return ValidationReturn{Success:false, Message:"Invalid HTTP Verb"}
		}
		var JSONData map[string]string
		decoder := json.NewDecoder(r.Body);
		e := decoder.Decode(&JSONData);
		if e != nil {
			return ValidationReturn{Success:false, Message:e.Error()}
		}
		//fmt.Println(JSONData);

		if _, exists := JSONData["message"]; !exists { //_ is the message
			return ValidationReturn{Success:false, Message:"JSON received doesn't have key \"Message\""}
		}

		if _, exists := JSONData["UUID"]; !exists {
			return ValidationReturn{Success:false, Message:"JSON received doesn't have key \"UUID\""}
		}

		return ValidationReturn{Success:true, KeyValues:JSONData}
	},
	/*"phase2buy":func(w http.ResponseWriter, r *http.Request) ValidationReturn {
		return ValidationReturn{}
	},
	"phase2sell":func(w http.ResponseWriter, r *http.Request) ValidationReturn {
		return ValidationReturn{}
	},
	"bye":func(w http.ResponseWriter, r *http.Request) ValidationReturn {
		return ValidationReturn{}
	},*/
}




func ValidateChat(w http.ResponseWriter, r *http.Request) ValidationReturn {
	//check HTTP verb
	if(r.Method != "POST") {
		return ValidationReturn{Success:false, Message:"Invalid HTTP Verb"}
	}

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

	//check for the key "UUID"
	// if _, exists := JSONData["UUID"]; !exists {
	// 	return ValidationReturn{Success:false, Message:"JSON received doesn't have key \"UUID\""}
	// }

	auth := r.Header.Get("Authorization")

	println(auth)

	JSONData["authorization"] = auth;

	// // a string slice to hold the keys
	// k := make([]string, len(c))

	// // iteration counter
	// i := 0

	// // copy c's keys into k
	// for s, _ := range c {
	// 	k[i] = s
	// 	i++
	// }

	return ValidationReturn{Success:true, KeyValues:JSONData}
}