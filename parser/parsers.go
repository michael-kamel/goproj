package parser

import (
	"strings"
)

func Parse(keywords []string, message string) ParserResult {
	for i := 0; i < len(keywords); i++ {
		if strings.Contains(strings.ToLower(message), strings.ToLower(keywords[i])) {
			return ParserResult{Success:true, Message:keywords[i]}
		}
	}
	return ParserResult{Success:false}
}