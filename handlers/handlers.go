package handlers

import (
	"os"
	"net/http"
	"goproj/bot"
	"goproj/stateCache"
	"goproj/uidGenerator"
	"github.com/gin-gonic/gin"
)
type Message struct {
	Message string `json:"message" binding:"required"`
}
func Handlers(bots map[string]*bot.Bot, botStateCache stateCache.BotStateCache, idGenerator uidGenerator.UIDGenerator) map[string]gin.HandlerFunc{
	handlers := map[string]gin.HandlerFunc{
		"Welcome":func (c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin","*")
			id := idGenerator.GenerateUID()
			bolt := bots[os.Getenv("ACTIVE_BOT")]
			state := bot.BotState{
				bolt.RootComponent,
				make(map[string]interface{}),
			}
			botStateCache.SetState(id, &state)
			message := bolt.Process("", &state)
			c.JSON(200, gin.H{
				"message":message,
				"uuid":id,
			})
		},
		"Chat":func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin","*")
			var inMsg Message
			err := c.ShouldBindJSON(&inMsg)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Please send me a message."})
				return
			}
			id := c.GetHeader("Authorization")
			bolt := bots[os.Getenv("ACTIVE_BOT")]
			outMsg := bolt.Process(inMsg.Message, botStateCache.GetState(id))
			c.JSON(200, gin.H{
				"message":outMsg,
			})
		},
		"BS":func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin","*")
			c.Writer.Header().Set("Access-Control-Allow-Headers","authorization,Content-type")
		},
	}
	return handlers
}


