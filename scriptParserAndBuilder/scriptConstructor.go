package scriptParserAndBuilder

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"reflect"
	//"../bot"
	//"../parser"
)

type ParsedBot struct {
	BotScript []struct {
		Name string
		Transitions []struct {
			Keywords []string
			NextPhase string
			Replies []string
			CustomFunction string
		}
	}
}

var ConstructedBot ParsedBot = ConstructBot()

func ConstructBot() ParsedBot {
	wd, _ := os.Getwd();
    file, e := ioutil.ReadFile(wd + "/scriptConstructor/BotScript1.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
	}
	//fmt.Printf(string(file))
	fmt.Println(reflect.TypeOf(file))

	var constructedBot ParsedBot
	err := json.Unmarshal(file, &constructedBot)
	if err != nil {
		fmt.Println("Cannot unmarshal the json ", err)
		os.Exit(1)
	}

	return constructedBot

	//fmt.Printf("%+v\n", testObj)
	
	//validate struct
	// b, err := json.Marshal(testObj)
    // if err != nil {
    //     fmt.Printf("Error: %s", err)
    //     return;
    // }
	// fmt.Println(string(b))
	
	//components := make([]bot.BotComponent, len(importedBot.BotScript))
	//connectors := []bot.BotComponentConnector{}
	/*components := map[string]bot.BotComponent{}
	for i := 0; i < len(importedBot.BotScript); i++ {
		c := bot.BotComponent{};
		c.Name = importedBot.BotScript[i].Name;
		c.Reply = importedBot.BotScript[i].Reply;
		c.CustomFunction = importedBot.BotScript[i].CustomFunction
		c.Parser = parser.GenerateParser(importedBot.BotScript[i].CustomFunction)
		components[importedBot.BotScript[i].Name] = c;

		//components = append(components, c)
	}

	for i := 0; i < len(importedBot.BotScript); i++ {
		components[importedBot.BotScript[i].Name] = bot.BotComponent{}
		//components = append(components, c)
	}*/


	//connectors := map[string]bot.BotComponentConnector{}
	// connectors := []bot.BotComponentConnector{}
	// for k, v := range components {
	// 	for i := 0; i < len(v.Transitions); i++ {
	// 		fmt.Printf(v.Transitions[i].NextPhase)
	// 	}
	// 	bot.BuildSingleTransitionConnector(&byeComponent)
	// 	connectors = append(connectors, )
	// }
}