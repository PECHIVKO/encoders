package encoder

import (
	"strings"
	"unicode"

	"github.com/PECHIVKO/encoders/pig-latin/format"
)

var vowels []rune = []rune{'a', 'e', 'i', 'o', 'u'}
var vowelSounds []string = []string{"a", "e", "i", "o", "u", "hon"}

// Encode encodes input string
func Encode(input string) (output string) {
	input = format.RemoveUnexpectedSymbols(input)

	var inputSlice []string
	var outputSlice []string

	inputSlice = strings.Split(input, " ")

	for _, wordInput := range inputSlice {
		var wordOutput string
		if startsWithVowelSound(wordInput) {
			wordOutput = vowelsPigLatin(wordInput)
		} else {
			wordOutput = consonantsPigLatin(wordInput)
		}
		outputSlice = append(outputSlice, wordOutput)
	}
	output = strings.Join(outputSlice, " ")
	return
}

// encode word as it starts from vowel sound
func vowelsPigLatin(input string) (output string) {
	outputSlice := []string{input, "yay"}
	output = strings.Join(outputSlice, "")
	output = format.FormatOutput(input, output)
	return
}

// encode word as it starts from consonant sound
func consonantsPigLatin(input string) (output string) {
	consostantsClusterCount := 0
	for _, character := range input {
		if isConsostant(character) {
			consostantsClusterCount++
		} else {
			break
		}
	}
	outputSlice := []string{input[consostantsClusterCount:], input[:consostantsClusterCount], "ay"}
	output = strings.Join(outputSlice, "")
	output = format.FormatOutput(input, output)
	return
}

//  isConsostant checks if character is consostant
func isConsostant(character rune) bool {
	character = unicode.ToLower(character)

	for _, r := range vowels {
		if character == r {
			return false
		}
	}
	return true
}

// startsWithVowelSound checks if string starts from vowel sound
func startsWithVowelSound(s string) bool {

	s = strings.ToLower(s)
	for _, sound := range vowelSounds {
		if strings.HasPrefix(s, sound) {
			return true
		}
	}
	return false
}
