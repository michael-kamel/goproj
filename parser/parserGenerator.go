package parser

func GenerateParser(keywords []string, parserFunction func([]string, string) ParserResult) func(string) ParserResult {
	return func(message string) ParserResult {
		return parserFunction(keywords, message);
	}
}