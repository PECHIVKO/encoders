package main

import (
	"bufio"
	"os"
	"strings"
)

var encodetable = map[string]string{
	"a": "1", "A": "1",
	"e": "2", "E": "2",
	"i": "3", "I": "3",
	"o": "4", "O": "4",
	"u": "5", "U": "5",
}

var decodetable = map[string]string{
	"1": "a",
	"2": "e",
	"3": "i",
	"4": "o",
	"5": "u",
}

func encode(inputPhrase string) (outputPhrase string) {
	outputPhrase = inputPhrase

	for oldChar, newChar := range encodetable {
		outputPhrase = strings.ReplaceAll(outputPhrase, oldChar, newChar)
	}
	return
}

func decode(inputPhrase string) (outputPhrase string) {
	outputPhrase = inputPhrase

	for oldChar, newChar := range decodetable {
		outputPhrase = strings.ReplaceAll(outputPhrase, oldChar, newChar)
	}
	return
}

// getPhrase gets phrase from command
func getPhrase(command string) (phrase string) {
	return command[7 : len(command)-1]
}

// isValidCommand checks if command is valid
func isValidCommand(command string) bool {
	if (strings.HasPrefix(command, "encode(") || strings.HasPrefix(command, "decode")) && strings.HasSuffix(command, ")") {
		return true
	}
	return false
}

// inputHandler handle user input
func inputHandler(input string) (output string) {
	if isValidCommand(input) {
		if strings.HasPrefix(input, "e") {
			output = encode(getPhrase(input))
			return
		}
		output = decode(getPhrase(input))
		return
	}
	output = "Invalid command"
	return
}

func main() {
	var input string

	reader := bufio.NewReader(os.Stdin)
	println("Use command 'encode(your phrase)' to encode OR 'decode(y45r phr1s2)' to decode:")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	println(inputHandler(input))
}
