package punctuation

import (
	"regexp"
	"strings"
)

func Punctuation(originalText string) string {
	reString := `[ \n\t]*([.,!?:;]+)[ ]*`
	re := regexp.MustCompile(reString)
	submatches := re.FindAllStringSubmatch(originalText, -1)
	if len(submatches) == 0 {
		return originalText
	}
	finalText := ""
	finalText = re.ReplaceAllString(originalText, "$1 ")

	reString = `([ \n\t]*[.,!?:;]+[ ]*){2,}`
	re = regexp.MustCompile(reString)
	signsRepeated := re.FindAllString(finalText, -1)
	signsRepeatedReworked := []string{}
	textSplited := re.Split(finalText, -1)

	for _, groupOfSigns := range signsRepeated {
		groupOfSigns = strings.ReplaceAll(groupOfSigns, " ", "")
		groupOfSigns = strings.ReplaceAll(groupOfSigns, "\n", "")
		groupOfSigns = strings.ReplaceAll(groupOfSigns, "\t", "")
		signsRepeatedReworked = append(signsRepeatedReworked, groupOfSigns)
	}
	finalText = ""

	for i := 0; i < len(signsRepeatedReworked); i++ {
		finalText += textSplited[i] + signsRepeatedReworked[i] + " "
	}
	finalText += textSplited[len(textSplited)-1]

	return finalText
}
