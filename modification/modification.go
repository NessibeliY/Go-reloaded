package modification

import (
	"go-reloaded/modification/apostrophes"
	"go-reloaded/modification/commands"
	"go-reloaded/modification/punctuation"
	"go-reloaded/modification/quotes"
	"go-reloaded/modification/vowels"
	"regexp"
	"strings"
)

func ModifyTheWholeText(finalText string) string {
	if finalText != "" {

		// Modification of original text

		// Cap, Up, Low, Hex, Bin
		reStringLow := `(?i) *\([ \n\t]*(?:` + `low` + `)[ \n\t]*,?[ \n\t]*(\d*)[ \n\t]*\)`

		reStringUp := `(?i) *\([ \n\t]*(?:` + `up` + `)[ \n\t]*,?[ \n\t]*(\d*)[ \n\t]*\)`

		reStringCap := `(?i) *\([ \n\t]*(?:` + `cap` + `)[ \n\t]*,?[ \n\t]*(\d*)[ \n\t]*\)`

		reStringHex := `(?i) *\([ \n\t]*(?:` + `hex` + `)[ \n\t]*,?[ \n\t]*(\d*)[ \n\t]*\)`

		reStringBin := `(?i) *\([ \n\t]*(?:` + `bin` + `)[ \n\t]*,?[ \n\t]*(\d*)[ \n\t]*\)`

		lowCapUpHexBin := `((?i) ?\([ \n\t]*(?:low|cap|up|hex|bin)[ \n\t]*,?[ \n\t]*(\d*)[ \n\t]*\))`

		pattern := regexp.MustCompile(lowCapUpHexBin)
		remainingChunk := ""
		// Modify can'ts
		finalText = apostrophes.Apostrophes(finalText)

		for len(pattern.FindAllString(finalText, -1)) > 0 {

			splittedText := pattern.Split(finalText, -1)
			firstChunk := splittedText[0]

			firstReg := pattern.FindString(finalText)
			firstChunkWithReg := firstChunk + firstReg

			lengthOfFirstChunk := len(firstChunkWithReg)
			remainingChunk = finalText[lengthOfFirstChunk:]

			if strings.Contains(firstReg, "cap") {
				finalText = commands.CapUpLowHexBin("cap", firstChunkWithReg, firstChunk, reStringLow, reStringUp, reStringCap, reStringHex, reStringBin)
			} else if strings.Contains(firstReg, "low") {
				finalText = commands.CapUpLowHexBin("low", firstChunkWithReg, firstChunk, reStringLow, reStringUp, reStringCap, reStringHex, reStringBin)
			} else if strings.Contains(firstReg, "up") {
				finalText = commands.CapUpLowHexBin("up", firstChunkWithReg, firstChunk, reStringLow, reStringUp, reStringCap, reStringHex, reStringBin)
			} else if strings.Contains(firstReg, "hex") {
				finalText = commands.CapUpLowHexBin("hex", firstChunkWithReg, firstChunk, reStringLow, reStringUp, reStringCap, reStringHex, reStringBin)
			} else if strings.Contains(firstReg, "bin") {
				finalText = commands.CapUpLowHexBin("bin", firstChunkWithReg, firstChunk, reStringLow, reStringUp, reStringCap, reStringHex, reStringBin)
			}
			finalText = finalText + remainingChunk

		}

		// Punctuation
		finalText = punctuation.Punctuation(finalText)

		// quotes
		finalText = quotes.Quotes(finalText)

		// A's and an's
		finalText = vowels.AAns(finalText)

		// Replace modified can'ts
		finalText = apostrophes.FinalApostrophes(finalText)
	}
	return finalText
}
