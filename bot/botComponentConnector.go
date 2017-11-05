package bot

type BotTransition func(state *BotState, bot *Bot) *BotComponent
type BotComponentConnector struct {
	Name string
	Transition BotTransition
}
