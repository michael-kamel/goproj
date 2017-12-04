package ResponseHandlers

import (
	"fmt"
	"strings"
	"strconv"
	"../ScriptParserAndBuilder"
	"../SessionManagement"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	// "github.com/davecgh/go-spew/spew"
)

const endPoint string = "https://go-proj-node-backend.herokuapp.com/";

// type listing struct {
// 	description:derrrrrrrr id:5a12c6e3c7244b0012b97554 ownerInfo:map[email:der@der.der name:der der phone:123456789] category:Residential location:6th of October space:200 price:1e+07
// }

var ResponseHandlers map[string]func(ScriptParserAndBuilder.Transition, string, *SessionManagement.UserSession) string = map[string]func(ScriptParserAndBuilder.Transition, string, *SessionManagement.UserSession) string {
	"test":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		return "item, item, item"
	},
	"setStateBuy":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		session.Data.BuyerOrSeller = "Buyer";
		return ""
	},
	"setStateSell":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		session.Data.BuyerOrSeller = "Buyer";
		return ""
	},
	"submit":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {

		url := endPoint + "addlisting"
		fmt.Println("URL:>", url)
	
		bodyString := `{
			"ownerInfo": {
				"name":"` + session.Data.Name + `",
				"phone":"` + session.Data.Phone + `",
				"email":"` + session.Data.Email + `"
			},
			"category":"` + session.Data.Category + `",
			"location":"` + session.Data.Location + `",
			"address":"` + session.Data.Address + `",
			"space":"` + strconv.Itoa(session.Data.Space) + `",
			"price":"` + strconv.Itoa(session.Data.Price) + `",
			"description":"` + session.Data.Description + `"
		}`;

		var jsonStr = []byte(bodyString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
	
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "Couldn't parse data received from server. Please inform the developers! " + err.Error();
		}
		defer resp.Body.Close()
	
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))



		//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err != nil {
			panic(err)
		}
		fmt.Println("response from backend")
		fmt.Println(dat)

		if dat["success"] == true {
			return "Your data was recorded and hopefully we'll find you a client. Thanks! To start over type buy or sell!"
		}

		return "Sorry your data wasn't recorded. You can try again by typing buy or sell. Please send this error to the developers: " + string(body)
	},
	"query":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {

		url := endPoint + "getlistingbyspec"
		url += "?category=" + session.Data.Category
		url += "&location=" + session.Data.Location
		url += "&space=" + strconv.Itoa(session.Data.Space)
		url += "&price=" + strconv.Itoa(session.Data.Price)

		url = strings.Replace(url, " ", "%20", -1)

		fmt.Println("URL:>", url)
			
		req, err := http.NewRequest("GET", url, nil)
		//req.Header.Set("Content-Type", "application/json")
	
		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			return "Error sending request to server. Please inform the developers. " + err.Error();
		}
	
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))



		//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err != nil {
			return "Couldn't parse json from server. Please inform the developers. " + err.Error();
		}
		fmt.Println("response from backend")
		fmt.Println(dat)

		if dat["success"] == true {
			list, ok := dat["listings"].([]interface{})
			if ok {
				if(len(list) == 0) {
					return "No results were found for your search criteria. You can start over by typing buy or sell!"
				}
				var r string = "";
				for i := 0; i < len(list); i++ {
					item, ok := list[i].(map[string]interface{});
					if ok {
						//r += strconv.Itoa(i) + ") " + string(item) + "\n";

						session.ReceivedItems = append(session.ReceivedItems, SessionManagement.ReceivedItem{Number:i+1, ID:item["id"].(string)})
					
						parsedSpace := strconv.FormatFloat(item["space"].(float64), 'f', 2, 64)
						parsedPrice := strconv.FormatFloat(item["price"].(float64), 'f', 2, 64)
						r += strconv.Itoa(i+1) + ") Description: " + item["description"].(string) + " Location: " + item["location"].(string) + " Space: " + parsedSpace + " Price: " + parsedPrice + "."
					//category:Residential location:6th of October    space:200 price:1e+07 description:derrrrrrrr
					}
				}

				fmt.Println(session.ReceivedItems)
				session.ReceivedItemsMessage = r
				rr := "Okay relevant results are displayed below. Choose one by typing its number. " + r + " If you see something you want to buy type its number";

				return rr;
			}

			return "Couldn't parse server response. Please inform the developers."
		}

		return "Couldn't get response from server. Send this error to the developers: " + string(body)
	},
	"request":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		
		url := endPoint + "addbuyerrequest"
		fmt.Println("URL:>", url)
	
		bodyString := `{
			"listings": ["` + session.ChosenItem + `"],
			"buyerInfo": {
				"name":"` + session.Data.Name + `",
				"phone":"` + session.Data.Phone + `",
				"email":"` + session.Data.Email + `"
			}
		}`;

		var jsonStr = []byte(bodyString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
	
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "Couldn't parse data received from server. Please inform the developers! " + err.Error();
		}
		defer resp.Body.Close()
	
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))



		//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err != nil {
			panic(err)
		}
		fmt.Println("response from backend")
		fmt.Println(dat)

		if dat["success"] == true {
			return "Your buy request was recorded. Thanks! To start over type buy or sell!"
		}

		return "Sorry your data wasn't recorded. You can try again by typing buy or sell. Please send this error to the developers: " + string(body)
	},
}














/*

func main() {

  fmt.Println("enter link")

  var link string //= "https://www.google.com"

  fmt.Scanf("%s", &link)

  

  resp, _ := http.Get(link)

  

  //fmt.Println(link)

  //defer next1(resp.Body)

  defer resp.Body.Close()

  //body, _ := ioutil.ReadAll(resp.Body)

  //bodyString := string(body)

  tokenised := html.NewTokenizer(resp.Body)

  for {

    sth := tokenised.Next()

    switch {

      case sth == html.ErrorToken:

        return

      case sth == html.StartTagToken:

        sth2 := tokenised.Token()

        isAnchor := sth2.Data == "a"

        if isAnchor {

          for _, a := range sth2.Attr {

            if(a.Key == "href") {

              fmt.Println()

              fmt.Println()

              fmt.Println(a.Val)

              break

            }

          }

          for {

            sth = tokenised.Next()

            if sth == html.EndTagToken && tokenised.Token().Data == "a" {

              break

            }

            if sth == html.TextToken {

              sth2 = tokenised.Token()

              fmt.Println(sth2)

              break

            }

          }

        }

     }

  }

}



func next1(body string) {

  fmt.Println(body)

}

#Advanced_Lab*/