package scriptBuilder

import (
	"goproj/bot"
	"goproj/parser"
	"strings"
	"strconv"
)

type ScriptBuilder struct {
	Components map[string]*bot.BotComponent
	Connectors map[string]*bot.BotComponentConnector
	FunctionalityProvider FunctionalityProvider
}

func (this *ScriptBuilder) Init(){
	this.Components = make(map[string]*bot.BotComponent)
	this.Connectors = make(map[string]*bot.BotComponentConnector)
}
func (this *ScriptBuilder) LoadComponent(componentDescriptor ComponentDescriptor) {
	component := bot.BotComponent{}
	component.Name = componentDescriptor.Name
	switch componentDescriptor.Question.Type {
		case "Simple":
			component.Question = bot.BuildSimpleQuestion(componentDescriptor.Question.Data)
		case "Custom":
			component.Question = this.FunctionalityProvider.GetQuestion(componentDescriptor.Question.Data)
		default:
			panic(unRecognizedQuestionType{componentDescriptor.Question.Type})
	}
	switch componentDescriptor.Parser.Type {
		case "KeywordParser": {
			keywords := strings.Split(componentDescriptor.Parser.Data, ",")
			component.Parser = parser.GenerateKeywordParser(keywords)
		}
		case "RegexpParser":
			component.Parser = parser.GenerateRegexpParser(componentDescriptor.Parser.Data)
		case "IdentityParser":
			component.Parser = parser.GenerateIdentityParser()
		case "NumericParser": {
			keywords := strings.Split(componentDescriptor.Parser.Data, ",")
			min, _ := strconv.Atoi(keywords[0])
			max, _ := strconv.Atoi(keywords[1])
			component.Parser = parser.GenerateNumericParser(min, max)
		}
		default:
			panic(unRecongnizedParserType{componentDescriptor.Parser.Type})
	}
	switch componentDescriptor.Handler.Type {
		case "Simple":
			component.Handler = bot.BuildSimpleHandler(componentDescriptor.Handler.Data)
		case "Custom": 
			component.Handler = this.FunctionalityProvider.GetHandler(componentDescriptor.Handler.Data)
		case "DataFill": {
			keywords := strings.Split(componentDescriptor.Handler.Data, ",")
			component.Handler = bot.BuildDataFillHandler(keywords[0], keywords[1])
		}
		default:
			panic(unRecongnizedHandlerType{componentDescriptor.Handler.Type})
	}
	this.Components[componentDescriptor.Name] = &component
}
func (this *ScriptBuilder) LoadConnector(connectorDescriptor ConnectorDescriptor) {
	switch connectorDescriptor.Type {
		case "Simple": {
			connector := bot.BuildSingleTransitionConnector(connectorDescriptor.Name, this.Components[connectorDescriptor.Data])
			this.Connectors[connectorDescriptor.Name] = &connector
		}
		case "Custom": {
			this.Connectors[connectorDescriptor.Name] = &bot.BotComponentConnector{connectorDescriptor.Name,this.FunctionalityProvider.GetTransition(connectorDescriptor.Data)}
		}
		default:
			panic(unRecongnizedTransitionType{connectorDescriptor.Type})
	}
}
func (this *ScriptBuilder) BuildBot(botDescriptor BotDescriptor) bot.Bot {
	bbot := bot.Bot{}
	bbot.Name = botDescriptor.Name
	bbot.UnhandledMessages = botDescriptor.UnhandledMessages
	components := make(map[string]*bot.BotComponent)
	connectors := make(map[string]*bot.BotComponentConnector)
	transitions := make(map[string]string)
	for key, value := range this.Components {
		components[key] = value
	}
	for key, value := range this.Connectors {
		connectors[key] = value
	}
	for _, value := range botDescriptor.Connections {
		transitions[value.Component] = value.Connector
	}
	bbot.Components = components
	bbot.Connectors = connectors
	bbot.Transitions = transitions
	bbot.RootComponent = components[botDescriptor.RootComponent]
	return bbot
}