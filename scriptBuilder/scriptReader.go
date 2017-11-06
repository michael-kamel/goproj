package scriptBuilder

import (
	"encoding/json"
	"goproj/bot"
)

func JsonScriptToBot(functionalityProvider FunctionalityProvider, scriptBuilder *ScriptBuilder, jsonData []byte) map[string]*bot.Bot {
	description := BotDescription{}
	json.Unmarshal(jsonData, &description)
	scriptBuilder.Init()
	scriptBuilder.FunctionalityProvider = functionalityProvider
	for _, component := range description.Components {
		scriptBuilder.LoadComponent(component)
	}
	for _, connector := range description.Connectors {
		scriptBuilder.LoadConnector(connector)
	}
	botSlice := make(map[string]*bot.Bot)
	for _, bDesc := range description.Bots {
		bbot := scriptBuilder.BuildBot(bDesc)
		botSlice[bDesc.Name] = &bbot
	}

	return botSlice
}