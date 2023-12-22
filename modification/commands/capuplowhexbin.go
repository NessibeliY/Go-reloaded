package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CapUpLowHexBin(action, firstChunkWithReg, firstChunk, reStringLow, reStringUp, reStringCap, reStringHex, reStringBin string) string {
	var funcName func(string) string

	reString := ""

	if action == "low" {
		funcName = strings.ToLower
		reString = reStringLow
	} else if action == "up" {
		funcName = strings.ToUpper
		reString = reStringUp
	} else if action == "cap" {
		funcName = capitalise
		reString = reStringCap
	} else if action == "hex" {
		funcName = hexToDecimal
		reString = reStringHex
	} else if action == "bin" {
		funcName = binToDecimal
		reString = reStringBin
	}

	re := regexp.MustCompile(reString)

	submatch := re.FindStringSubmatch(firstChunkWithReg)

	num := 1
	if submatch[1] != "" {
		num, _ = strconv.Atoi(submatch[1])
	}

	finalText := ""

	if num <= 1000 {
		endText := regexp.MustCompile(`(?:[\w-]+(?:[^\w-]+|\b)){0,` + strconv.Itoa(num) + `}$`)
		finalText += endText.ReplaceAllStringFunc(firstChunk, func(match string) string {
			return funcName(match)
		})

	} else {
		endText := regexp.MustCompile(`(?:[\w-]+(?:[^\w-]+|\b)){0,1000}$`)
		finalText += endText.ReplaceAllStringFunc(firstChunk, func(match string) string {
			return funcName(match)
		})

		for num < 1000 {
			intermediateText := endText.ReplaceAllString(firstChunk, "")
			endText = regexp.MustCompile(`(?:[\w-]+(?:[^\w-]+|\b)){0,` + strconv.Itoa(num-1000) + `}$`)
			finalText += endText.ReplaceAllStringFunc(intermediateText, func(match string) string {
				return funcName(match)
			})
			num = num - 1000
		}

	}

	return finalText
}

func capitalise(s string) string {
	s = strings.ToLower(s)
	s = strings.Title(s)
	return s
}

func hexToDecimal(hexValue string) string {
	re := regexp.MustCompile(`[ \n\t]+`)
	spacesAfter := re.FindString(hexValue)
	newValue := hexValue
	if spacesAfter != "" {
		newValue = strings.ReplaceAll(newValue, " ", "")
		newValue = strings.ReplaceAll(newValue, "\n", "")
		newValue = strings.ReplaceAll(newValue, "\t", "")
	}
	if !isHexadecimal(newValue) {
		fmt.Println("Please provide hexadecimal number")
		return hexValue
	}
	decimalValue, _ := strconv.ParseInt(newValue, 16, 64)
	strDecimalValue := strconv.Itoa(int(decimalValue))
	return strDecimalValue + spacesAfter
}

func binToDecimal(binValue string) string {
	re := regexp.MustCompile(`[ \n\t]+`)
	spacesAfter := re.FindString(binValue)
	newValue := binValue
	if spacesAfter != "" {
		newValue = strings.ReplaceAll(newValue, " ", "")
		newValue = strings.ReplaceAll(newValue, "\n", "")
		newValue = strings.ReplaceAll(newValue, "\t", "")
	}
	if !isBinary(newValue) {
		fmt.Println("Please provide binary number")
		return binValue
	}

	decimalValue, _ := strconv.ParseInt(newValue, 2, 64)
	strDecimalValue := strconv.Itoa(int(decimalValue))
	return strDecimalValue + spacesAfter
}

func isBinary(s string) bool {
	base := "01"
	for _, char := range s {
		if !strings.Contains(base, string(char)) {
			return false
		}
	}
	return true
}

func isHexadecimal(s string) bool {
	base := "0123456789abcdefABCDEF"
	for _, char := range s {
		if !strings.Contains(base, string(char)) {
			return false
		}
	}
	return true
}
