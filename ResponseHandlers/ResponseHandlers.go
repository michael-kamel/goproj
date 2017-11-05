package ResponseHandlers

import (
	//"fmt"
	//"strings"
	//"strconv"
	"../ScriptParserAndBuilder"
	"../SessionManagement"
)

var ResponseHandlers map[string]func(ScriptParserAndBuilder.Transition, string, *SessionManagement.UserSession) string = map[string]func(ScriptParserAndBuilder.Transition, string, *SessionManagement.UserSession) string {
	"test":func(transition ScriptParserAndBuilder.Transition, message string, session *SessionManagement.UserSession) string {
		return "item, item, item"
	},
}