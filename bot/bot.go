package bot

import ( 
	"math/rand"
	"time"
	"strings"
)

type Bot struct {
	Name string
	Components map[string]*BotComponent
	Connectors map[string]*BotComponentConnector
	Transitions map[string]string
	UnhandledMessages []string
	RootComponent *BotComponent
}

func (this *Bot) Connect(component *BotComponent, connector *BotComponentConnector) {
	this.ConnectNames(component.Name, connector.Name)
}
func (this *Bot) ConnectNames(componentName string, connectorName string) {
	if _, ok := this.Components[componentName]; ok {
		if _, ok := this.Connectors[connectorName]; ok {
			this.Transitions[componentName] = connectorName
			return
		} else {
			panic(unRegisteredConnectorError{connectorName})
		}
	} else {
		panic(unRegisteredComponentError{componentName})
	}
}
func (this *Bot) Process(message string, state *BotState) string {
	parseResult := state.CurrentComponent.Parser(message)
	if parseResult.Success {
		answer := state.CurrentComponent.Handler(parseResult.Message, state)
		transitionConnectorName := this.Transitions[state.CurrentComponent.Name]
		transitionConnector := this.Connectors[transitionConnectorName]
		state.CurrentComponent = transitionConnector.Transition(state, this)
		return composeResponse(answer, state.CurrentComponent.Question(state))
	} else {
		rand.Seed(time.Now().Unix())
		return this.UnhandledMessages[rand.Intn(len(this.UnhandledMessages))]
	}
}
func composeResponse(question string, answer string) string {
	if question != "" {
		return strings.Join([]string{question, answer}, ". ")
	} else {
		return answer
	}
}