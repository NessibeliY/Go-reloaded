package apostrophes

import (
	"regexp"
	"strings"
)

func Apostrophes(finalText string) string {
	reWords := `(?i)\b(((is|do(es)?)n\W*'\W*t)|(i\W*'\W*m)|(you\W*'\W*re)|(s?he\W*'\W*s)|(it\W*'\W*s)|(he\W*'\W*d)|(let\W*'\W*s)|(would\W*'\W*ve)|(who\W*'\W*s)|(can\W*'\W*t)|(couldn\W*'\W*t)|(won\W*'\W*t)|(wouldn\W*'\W*t)|(they\W*'\W*d))\b`
	re := regexp.MustCompile(reWords)
	finalText = re.ReplaceAllStringFunc(finalText, ApostrophesAdditional)
	return finalText
}

func FinalApostrophes(finalText string) string {
	reWords := `(?i)\b(((is|do(es)?)n_t)|(i_m)|(you_re)|(s?he_s)|(it_s)|(he_d)|(let_s)|(would_ve)|(who_s)|(can_t)|(couldn_t)|(won_t)|(wouldn_t)|(they_d))\b`
	re := regexp.MustCompile(reWords)
	finalText = re.ReplaceAllStringFunc(finalText, FinalApostrophesAdditional)
	return finalText
}

func FinalApostrophesAdditional(can string) string {
	can = strings.ReplaceAll(can, "_", "'")
	return can
}

func ApostrophesAdditional(can string) string {
	can = strings.ReplaceAll(can, " ", "")
	can = strings.ReplaceAll(can, "\n", "")
	can = strings.ReplaceAll(can, "\t", "")
	can = strings.ReplaceAll(can, "'", "_")
	return can
}
