package scriptBuilder

type ComponentDescriptor struct {
	Name string
	Question QuestionDescriptor
	Parser ParserDescriptor
	Handler HandlerDescriptor
}

type QuestionDescriptor struct {
	Type string
	Data string
}
type ParserDescriptor struct {
	Type string
	Data string
}
type HandlerDescriptor struct {
	Type string
	Data string
}

type ConnectorDescriptor struct {
	Name string
	Type string
	Data string
}

type ConnectionDescriptor struct {
	Component string
	Connector string
}
type BotDescriptor struct {
	Name string
	Componenets []string
	Connectors []string
	Connections []ConnectionDescriptor
	UnhandledMessages []string
	RootComponent string
}

type BotDescription struct {
	Components []ComponentDescriptor
	Connectors []ConnectorDescriptor
	Bots []BotDescriptor
}