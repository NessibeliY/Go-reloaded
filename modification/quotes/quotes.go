package quotes

import (
	"regexp"
	"strings"
)

func Quotes(originalText string) string {
	finalText := originalText

	reString := `'([\w\W]*?)'`
	re := regexp.MustCompile(reString)
	finalText = re.ReplaceAllStringFunc(finalText, TrimSingleQuotes)

	reString = `"([\w\W]*?)"`
	re = regexp.MustCompile(reString)
	finalText = re.ReplaceAllStringFunc(finalText, TrimDoubleQuotes)

	return finalText
}

func TrimSingleQuotes(originalText string) string {
	trimmedText := strings.Trim(originalText, " \n\t'")
	return "'" + trimmedText + "'"
}

func TrimDoubleQuotes(originalText string) string {
	trimmedText := strings.Trim(originalText, " \n\t\"")
	return "\"" + trimmedText + "\""
}
