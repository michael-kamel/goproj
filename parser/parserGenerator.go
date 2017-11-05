package parser

import (
	"regexp"
	"strconv"
)

func GenerateKeywordParser(keywords []string) func(string) ParserResult {
	return func(message string) ParserResult {
		return Parse(keywords, message);
	}
}
func GenerateRegexpParser(regExp string) func(string) ParserResult {
	var validMsg = regexp.MustCompile(regExp)
	return func(message string) ParserResult {
		found := validMsg.MatchString(message)
		if found {
			return ParserResult{Success:true, Message:message}
		} else {
			return ParserResult{Success:false}
		}
	}
}
func GenerateIdentityParser() func(string) ParserResult {
	return func(message string) ParserResult {
		return ParserResult{Success:true, Message:message}
	}
}
func GenerateNumericParser(min int, max int) func(string) ParserResult {
	return func(message string) ParserResult {
		i, err := strconv.Atoi(message)
		if err != nil {
			return ParserResult{Success:false}
		}
		if i >= min && i <= max {
			return ParserResult{Success:true, Message:i}
		} else {
			return ParserResult{Success:false}
		}
	}
}