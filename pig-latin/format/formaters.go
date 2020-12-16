package format

import (
	"log"
	"regexp"
	"strings"
)

// removeUnexpectedSymbols removes non alphabetic symbols and redundant spaces from input string
func RemoveUnexpectedSymbols(input string) (output string) {
	words, err := regexp.Compile(`[^A-Za-z\s']`)
	if err != nil {
		log.Fatal(err)
	}
	doubleSpaces, err := regexp.Compile(`\s+`)
	if err != nil {
		log.Fatal(err)
	}

	output = words.ReplaceAllString(input, "")
	output = doubleSpaces.ReplaceAllString(output, " ")
	output = strings.TrimSpace(output)

	return
}

// FormatOutput formats output string due to input case
func FormatOutput(reference string, input string) (output string) {
	if isUpperCase(reference) {
		output = strings.ToUpper(input)
	} else if startsWithCapital(reference) {
		output = strings.Title(strings.ToLower(input))
	} else {
		output = strings.ToLower(input)
	}
	return
}

// isUpperCase check if string is UPPERCASE
func isUpperCase(s string) bool {
	if strings.ToUpper(s) == s {
		return true
	}
	return false
}

// startsWithCapital check if string starts from capital
func startsWithCapital(s string) bool {
	if strings.Title(strings.ToLower(s)) == s {
		return true
	}
	return false
}
