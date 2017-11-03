package bot

type BotComponentConnector struct {
	Transition func(state *BotState) BotComponent
}
