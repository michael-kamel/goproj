package scriptConstructor

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"reflect"
	"../bot"
)

type BotScript struct {
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

func ConstructBot() {
	wd, _ := os.Getwd();
    file, e := ioutil.ReadFile(wd + "/scriptConstructor/BotScript1.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
	}
	//fmt.Printf(string(file))
	fmt.Println(reflect.TypeOf(file))

	var importedBot BotScript
	err := json.Unmarshal(file, &importedBot)
	if err != nil {
		fmt.Println("Cannot unmarshal the json ", err)
		os.Exit(1)
	}

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
	components := map[string]bot.BotComponent{}
	for i := 0; i < len(importedBot.BotScript); i++ {
		components[importedBot.BotScript[i].Name] = bot.BotComponent{}
		//components = append(components, c)
	}
	//connectors := map[string]bot.BotComponentConnector{}
	
}