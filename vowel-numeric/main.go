package main

import (
	"bufio"
	"os"
	"strings"
)

func encode(inputPhrase string) (outputPhrase string) {
	var encodedSlice []rune

	for _, r := range inputPhrase {
		switch r {
		case 'a', 'A':
			encodedSlice = append(encodedSlice, '1')
		case 'e', 'E':
			encodedSlice = append(encodedSlice, '2')
		case 'i', 'I':
			encodedSlice = append(encodedSlice, '3')
		case 'o', 'O':
			encodedSlice = append(encodedSlice, '4')
		case 'u', 'U':
			encodedSlice = append(encodedSlice, '5')
		default:
			encodedSlice = append(encodedSlice, r)
		}
	}
	outputPhrase = string(encodedSlice)
	return
}

func decode(inputPhrase string) (outputPhrase string) {
	var decodedSlice []rune

	for _, r := range inputPhrase {
		switch r {
		case '1':
			decodedSlice = append(decodedSlice, 'a')
		case '2':
			decodedSlice = append(decodedSlice, 'e')
		case '3':
			decodedSlice = append(decodedSlice, 'i')
		case '4':
			decodedSlice = append(decodedSlice, 'o')
		case '5':
			decodedSlice = append(decodedSlice, 'u')
		default:
			decodedSlice = append(decodedSlice, r)
		}
	}
	outputPhrase = string(decodedSlice)
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
