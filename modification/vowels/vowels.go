package vowels

import (
	"regexp"
	"unicode/utf8"
)

func AAns(originalText string) string {
	// Vowels and h-s
	reString := `(?i)\ba\b\W+\n*[aeiouh]`
	re := regexp.MustCompile(reString)
	findAllAs := re.FindAllString(originalText, -1)
	splittedText := re.Split(originalText, -1)
	newSliceOfText := []string{}
	for _, v := range findAllAs {
		newSliceOfText = append(newSliceOfText, v[:1]+"n"+v[1:])
	}
	finalText := ""

	for i, textChunk := range splittedText {
		if i != len(splittedText)-1 {
			finalText += textChunk + newSliceOfText[i]
		} else {
			finalText += textChunk
		}
	}

	finalText = re.ReplaceAllStringFunc(finalText, ReplaceAs)

	// Consonants
	reString = `(?i)\ban\b\W+\n*[bcdfghjklmnpqrstvwxz]`
	re = regexp.MustCompile(reString)
	finalText = re.ReplaceAllStringFunc(finalText, ReplaceAns)

	return finalText
}

func ReplaceAs(originalText string) string {
	firstLetter := ExtractFirstRune(originalText)
	remainingText := TrimFirstRune(originalText)
	if firstLetter == "A" {
		return "An" + remainingText
	}
	return "an" + remainingText
}

func ReplaceAns(originalText string) string {
	firstLetter := ExtractFirstRune(originalText)
	remainingText := TrimFirstRune(originalText)

	remainingText = TrimFirstRune(remainingText)

	return firstLetter + remainingText
}

func ExtractFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[:i]
}

func TrimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
