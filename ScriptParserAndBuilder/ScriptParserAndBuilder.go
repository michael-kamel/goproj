package ScriptParserAndBuilder

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	//"reflect"
	//"../bot"
	//"../parser"
)

type ParsedBot struct {
	BotScript []State 
}
type State struct {
	Name string
	Transitions []Transition
}
type Transition struct {
	Keywords []string
	NextState string
	Replies []string
	Rejects []string
	CustomFunction string
}


var ConstructedBot map[string]State = ConstructBot()

func ConstructBot() map[string]State {
	wd, _ := os.Getwd();
    file, e := ioutil.ReadFile(wd + "/ScriptParserAndBuilder/BotScript1.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
	}
	//fmt.Printf(string(file))
	//fmt.Println(reflect.TypeOf(file))

	var parsedBot ParsedBot
	err := json.Unmarshal(file, &parsedBot)
	if err != nil {
		fmt.Println("Cannot unmarshal the json ", err)
		os.Exit(1)
	}

	//fmt.Println(parsedBot)

	StatesMap := map[string]State{}

	for i := 0; i < len(parsedBot.BotScript); i++ {
		StatesMap[parsedBot.BotScript[i].Name]  = parsedBot.BotScript[i]
	}

	//fmt.Println(StatesMap)

	return StatesMap //constructedBot

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